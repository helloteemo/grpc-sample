package user_grpc

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/helloteemo/pb/user"
	"github.com/helloteemo/utils/grpc_gateway_util"
	"github.com/sirupsen/logrus"
)

func Register(ctx context.Context, grpcGatewayMux *grpc_gateway_runtime.ServeMux, engine *gin.Engine) {
	Init()

	grpcConn := GetGrpcConn()

	err := user.RegisterUserServiceHandler(ctx, grpcGatewayMux, grpcConn)
	if err != nil {
		logrus.WithError(err).Panic("failed to register app handler")
		return
	}

	// 只暴露app端的接口
	engine.Any(
		"/grpc-sample/user/*uri",
		// 重写resp流的中间件一定要最后使用
		grpc_gateway_util.GinHandler,
		gin.WrapH(grpcGatewayMux),
	)

}
