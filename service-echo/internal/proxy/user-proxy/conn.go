package user_proxy

import (
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/helloteemo/pb/user"
	"github.com/helloteemo/utils/jaeger_tracer"
	"github.com/helloteemo/utils/log"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	grpcUrl string

	grpcConn   *grpc.ClientConn
	grpcClient user.UserServiceClient
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
		grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			ctx = log.GPPCClientContext(ctx)
			return invoker(ctx, method, req, reply, cc, opts...)
		}),
	)

	if err != nil {
		logrus.WithError(err).WithField(`urlPrefix`, grpcUrl).Panic("did not connect.")
	}
}

func Init() {
	grpcUrl = "consul://consul.dev:8500/user"
	initGrpcConn(grpcUrl)
	grpcConn = GetGrpcConn()
	grpcClient = user.NewUserServiceClient(grpcConn)
}

func GetGrpcConn() *grpc.ClientConn {
	return grpcConn
}

func GetGrpcClient() user.UserServiceClient {
	return grpcClient
}
