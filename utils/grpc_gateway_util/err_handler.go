package grpc_gateway_util

import (
	"context"
	grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	json "github.com/helloteemo/utils/fast_json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Result struct {
	Ret int32  `json:"ret"`
	Msg string `json:"msg,omitempty"`
}

// ErrHandler Grpc响应处理,会写流,如果发生错误就不需要再重复写流了
func ErrHandler(ctx context.Context, mux *grpc_gateway_runtime.ServeMux,
	marshaller grpc_gateway_runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {

	writer.Header().Set("Content-Type", `application/json`)
	writer.Header().Set(`X-Request-ID`, request.Header.Get(`X-Request-ID`))

	request.Header.Add(errHandlerFlagBit, errHandlerFlagBitVal)

	fromError, ok := status.FromError(err)
	if !ok {
		fromError = status.New(codes.Unknown, err.Error())
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		`req`:        request.URL.Path,
		`clientIp`:   request.Header.Get("X-Forwarded-For"),
		`ret`:        fromError.Proto().GetCode(),
		`err_msg`:    fromError.Message(),
		`request_id`: request.Header.Get("X-Request-ID"),
	}).Errorf("err: [%+v]", err)

	_, _ = writer.Write(json.MarshalNoError(&Result{
		Ret: -2000 - fromError.Proto().GetCode(), // 不要直接返回rpc异常给客户端或者服务端
		Msg: fromError.Message(),
	}))
}
