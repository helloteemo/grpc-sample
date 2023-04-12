package log

import (
	"context"
	"github.com/helloteemo/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"testing"
)

func TestGrpcErrorMiddleware(t *testing.T) {
	middleware := GrpcErrorMiddleware(codes.InvalidArgument)
	ctx := ContextWithTraceId(context.Background())

	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:       true,
		DisableHTMLEscape: true,
	})

	_, _ = middleware(ctx, nil, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-none",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		// 没有异常
		return nil, nil
	})

	_, _ = middleware(ctx, nil, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-bad params",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, status.New(codes.InvalidArgument, "bad params").Err()
	})

	_, _ = middleware(ctx, map[string]interface{}{"h_m": 938}, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-status",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		// status 异常
		return nil, status.New(codes.Internal, "panic").Err()
	})

	_, _ = middleware(ctx, nil, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-pkg/errors",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, errors.New("pkg/errors异常")
	})

	_, _ = middleware(ctx, nil, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-utils.GrpcErrorWrap-pkg/errors",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, utils.GrpcWrapError(errors.New("utils.GrpcErrorWrap异常"),
			codes.Internal, "utils.GrpcErrorWrap异常1")
	})

	_, _ = middleware(ctx, nil, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-utils.GrpcErrorWrap-pkg/errors",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, utils.GrpcWrapErrorf(errors.New("utils.GrpcErrorWrap异常"),
			codes.Internal, "utils.GrpcErrorWrap异常2. %s", "abc")
	})

	_, _ = middleware(ctx, nil, &grpc.UnaryServerInfo{
		Server:     nil,
		FullMethod: "FullMethod-utils.GrpcErrorWrap-errors",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, utils.GrpcWrapError(os.ErrExist, codes.Internal, "utils.GrpcErrorWrap异常3")
	})
}
