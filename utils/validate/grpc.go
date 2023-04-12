package validate

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcValidator grpc 参数校验器
type GrpcValidator interface {
	Validate() error
}

// GrpcValidateParams grpc参数校验器
func GrpcValidateParams() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {
		// 校验
		if p, ok := req.(GrpcValidator); ok {
			if err := p.Validate(); err != nil {
				return nil, status.New(codes.InvalidArgument, "参数异常,请稍后再试~").Err()
			}
		}
		// 通过验证
		return handler(ctx, req)
	}
}
