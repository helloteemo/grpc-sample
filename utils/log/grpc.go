package log

import (
	"context"
	"github.com/mercari/go-grpc-interceptor/multiinterceptor"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// GrpcTraceIdMiddleware 生成trade_id
func GrpcTraceIdMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		incomingContext, ok := metadata.FromIncomingContext(ctx)
		if ok && len(incomingContext) > 0 && len(incomingContext.Get(TraceIdKey)) > 0 {
			ctx = context.WithValue(ctx, TraceIdKey, incomingContext.Get(TraceIdKey)[0])
		} else {
			ctx = ContextWithTraceId(ctx)
		}
		return handler(ctx, req)
	}
}

// GrpcErrorMiddleware 错误日志处理
// 可搭配 common.GrpcWrapError 使用
func GrpcErrorMiddleware(ignoreCode ...codes.Code) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}
		convert := status.Convert(err)
		if convert == nil {
			logrus.WithContext(ctx).WithFields(map[string]interface{}{`info`: info, `func`: `GrpcErrorMiddleware`, `req`: req}).
				Errorf("%+v", err)
			return resp, err
		}

		// 忽略错误码
		for _, code := range ignoreCode {
			if convert.Code() == code {
				return resp, err
			}
		}

		protoErr := convert.Proto()
		if protoErr != nil && protoErr.Code != 0 {
			logrus.WithContext(ctx).WithFields(map[string]interface{}{
				`req`:     req,
				`info`:    info,
				`func`:    `GrpcErrorMiddleware`,
				`code`:    protoErr.GetCode(),
				`details`: protoErr.Details,
			}).Error(protoErr.GetMessage())
		} else {
			logrus.WithContext(ctx).WithFields(map[string]interface{}{
				`info`: info,
				`req`:  req,
				`func`: `GrpcErrorMiddleware`,
				`code`: protoErr.GetCode(),
			}).Error(protoErr.GetMessage())
		}
		return resp, err
	}
}

// StreamServerInterceptor stream服务中间件
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		ctx := stream.Context()
		incomingContext, ok := metadata.FromIncomingContext(ctx)
		if ok && len(incomingContext) > 0 && len(incomingContext.Get(TraceIdKey)) > 0 {
			ctx = context.WithValue(ctx, TraceIdKey, incomingContext.Get(TraceIdKey)[0])
		} else {
			ctx = ContextWithTraceId(ctx)
		}
		stream = multiinterceptor.NewServerStreamWithContext(stream, ctx)
		return handler(srv, stream)
	}
}
