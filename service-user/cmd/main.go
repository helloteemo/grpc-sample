package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcOpentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/helloteemo/pb/user"
	"github.com/helloteemo/service-user/internal/config"
	"github.com/helloteemo/service-user/internal/service"
	"github.com/helloteemo/utils"
	"github.com/helloteemo/utils/consul"
	"github.com/helloteemo/utils/jaeger_tracer"
	"github.com/helloteemo/utils/log"
	"github.com/helloteemo/utils/stringx"
	"github.com/helloteemo/utils/validate"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func getGrpcPort() int {
	return config.GrpcPort
}

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

	// 解析配置文件
	log.InitLog()

	grpcPort := getGrpcPort()

	grpcListenAddr := fmt.Sprintf("0.0.0.0:%d", grpcPort)
	logrus.Infof("grpc %d is ready, listen: %s", grpcPort, grpcListenAddr)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logrus.Panicf("failed to listen: %v", err)
	}

	jaeger_tracer.InitJaeger(config.ProjectName)
	var tracer = jaeger_tracer.GetGlobalTracer()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandler(func(p interface{}) (err error) {
				stack := debug.Stack()
				logrus.WithError(err).WithField(`stack`, stringx.ZeroCopyBytes2String(stack)).
					Errorf("grpc server has panic.")
				return err
			})),
			grpcOpentracing.UnaryServerInterceptor(
				grpcOpentracing.WithTracer(tracer),
			),
			grpcPrometheus.UnaryServerInterceptor,
			log.GrpcTraceIdMiddleware(),
			log.GrpcErrorMiddleware(codes.InvalidArgument),
			validate.GrpcValidateParams(),
		)))

	imp := service.NewAppService()
	// 注册默认的grpc health
	grpc_health_v1.RegisterHealthServer(grpcServer, &consul.DefaultHealthImpl{})
	// 注册grpc服务
	user.RegisterUserServiceServer(grpcServer, imp)

	register2Consul()

	go func() {
		c := make(chan os.Signal, 4)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
		<-c
		consul.GrpcDRegister()
		logrus.Infof("graceful shutdown")
		grpcServer.GracefulStop()
		os.Exit(0)
	}()

	if err := grpcServer.Serve(listen); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func register2Consul() {

	grpcPort := getGrpcPort()

	if os.Getenv("stage") == "dev" {
		clientIp, err := utils.GetLocalIp()
		if err != nil {
			logrus.WithError(err).Errorln("failed to get client ip")
			clientIp = "127.0.0.1"
		}
		logrus.Infoln("clientIp", clientIp, "grpcPort", grpcPort)
		consul.GrpcRegisterByIpPort(consul.RegisterConfig{
			ServiceName: fmt.Sprintf("%s", config.ProjectName),
			MetricsPath: "grpc.health.v1.Health.Check",
			FailedFatal: true,
		}, clientIp, grpcPort)
		return
	}

	consul.GrpcRegisterByEth0IpPort(consul.RegisterConfig{
		ServiceName: fmt.Sprintf("%s", config.ProjectName),
		MetricsPath: "grpc.health.v1.Health.Check",
		FailedFatal: true,
	}, grpcPort)

	return
}
