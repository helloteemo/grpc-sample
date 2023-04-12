package utils

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/helloteemo/utils/stringx"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// GrpcErrBadParams 参数异常
	GrpcErrBadParams = status.New(codes.InvalidArgument, "参数错误").Err()
	// GrpcErrInternal 内部错误
	GrpcErrInternal = status.New(codes.Internal, "请稍后再试吧，么么哒~~").Err()
)

// GrpcWrapInternal 其中err会在内部打印出来,但是err不会返回给客户端
func GrpcWrapInternal(err error) error {
	return GrpcWrapError(err, codes.Internal, "请稍后再试吧，么么哒~~")
}

// GrpcWrapInternalError 其中err会在内部打印出来,但是err不会返回给客户端
func GrpcWrapInternalError(err error, message string) error {
	return GrpcWrapError(err, codes.Internal, message)
}

// GrpcWrapInternalErrorf 其中err会在内部打印出来,但是err不会返回给客户端
func GrpcWrapInternalErrorf(err error, message string, args ...interface{}) error {
	return GrpcWrapErrorf(err, codes.Internal, message, args...)
}

// GrpcWrapError 包装grpc错误
func GrpcWrapError(err error, code codes.Code, message string) error {
	return status.FromProto(&spb.Status{
		Code:    int32(code),
		Message: message,
		Details: []*any.Any{
			{
				TypeUrl: "stack",
				Value:   stringx.ZeroCopyString2Bytes(fmt.Sprintf("%+v", err)),
			},
		},
	}).Err()
}

// GrpcWrapErrorf 包装grpc错误
func GrpcWrapErrorf(err error, code codes.Code, message string, args ...interface{}) error {
	return status.FromProto(&spb.Status{
		Code:    int32(code),
		Message: fmt.Sprintf(message, args...),
		Details: []*any.Any{
			{
				TypeUrl: "stack",
				Value:   stringx.ZeroCopyString2Bytes(fmt.Sprintf("%+v", err)),
			},
		},
	}).Err()
}

// GrpcErrorf 包装grpc错误
func GrpcErrorf(code codes.Code, message string, args ...interface{}) error {
	return status.Newf(code, message, args...).Err()
}

// GrpcError 包装grpc错误
func GrpcError(code codes.Code, message string) error {
	return status.New(code, message).Err()
}
