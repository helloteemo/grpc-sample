package echo_grpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/helloteemo/utils/jaeger_tracer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	grpcUrl string

	grpcConn *grpc.ClientConn
)

func initGrpcConn(grpcUrl string) {
	var err error
	var ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	tracer := jaeger_tracer.GetGlobalTracer()
	grpcConn, err = grpc.DialContext(ctx, grpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer),
		),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer),
		),
	)

	if err != nil {
		logrus.WithError(err).WithField(`urlPrefix`, grpcUrl).Panic("did not connect.")
	}
}

func Init() {
	grpcUrl = "consul://consul.dev:8500/echo"
	initGrpcConn(grpcUrl)
	grpcConn = GetGrpcConn()
}

func GetGrpcConn() *grpc.ClientConn {
	return grpcConn
}
