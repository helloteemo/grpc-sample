package jaeger_tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
)

// GetGlobalTracer 获取全局tracer
func GetGlobalTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

// InitJaeger 初始化Jaeger
func InitJaeger(service string) {
	cfg, _ := jaegerCfg.FromEnv()
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegerCfg.Logger(jaeger.StdLogger))
	if err != nil {
		logrus.WithError(err).Fatalf("ERROR: cannot init Jaeger")
	}
	opentracing.InitGlobalTracer(tracer)
	return
}
