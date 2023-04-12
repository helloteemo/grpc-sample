package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	grpcGatewayRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	echo_grpc "github.com/helloteemo/gateway/internal/echo"
	user_grpc "github.com/helloteemo/gateway/internal/user"
	"github.com/helloteemo/utils"
	"github.com/helloteemo/utils/fast_json/grpc_gateway_runtime_json"
	"github.com/helloteemo/utils/grpc_gateway_util"
	"github.com/helloteemo/utils/jaeger_tracer"
	"github.com/helloteemo/utils/log"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   "2006-01-02 15:03:04",
		PrettyPrint:       false,
	})
}

func main() {
	// 获取启动参数
	var configFile string
	flag.StringVar(&configFile, "f", "", "Configuration file.")
	flag.StringVar(&configFile, "c", "", "Configuration file.")
	flag.Parse()

	jaeger_tracer.InitJaeger("gateway")

	grpcGatewayMux := grpcGatewayRuntime.NewServeMux(
		grpcGatewayRuntime.WithMarshalerOption(
			grpcGatewayRuntime.MIMEWildcard,
			&grpc_gateway_runtime_json.FastJsonBuiltin{},
		),
		grpcGatewayRuntime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			traceId := utils.UUID()
			request.Header.Add("X-Request-ID", traceId)
			return metadata.Pairs(log.TraceIdKey, traceId)
		}),
		grpcGatewayRuntime.WithErrorHandler(grpc_gateway_util.ErrHandler),
	)

	todo := context.TODO()
	engine := gin.New()
	engine.Any(`ping`, func(c *gin.Context) {
		c.String(200, "pong")
		return
	})

	echo_grpc.Register(todo, grpcGatewayMux, engine)
	user_grpc.Register(todo, grpcGatewayMux, engine)

	go func() {
		c := make(chan os.Signal, 4)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
		<-c
		logrus.Infof("graceful shutdown")
		os.Exit(0)
	}()

	var httpListenAddr = ":9000"
	logrus.Infof("httpSrv %d is ready, listen: %s", 9000, httpListenAddr)
	if err := engine.Run(httpListenAddr); err != nil {
		logrus.WithError(err).Errorln("failed to listen http server")
	}
}
