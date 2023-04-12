package consul

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"strings"
)

// DefaultHealthImpl 默认的Grpc health 实现。各个服务可以根据实际情况来自定义
type DefaultHealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *DefaultHealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	logrus.Debugf("health checking.req:%+v", req)
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch 监听
func (h *DefaultHealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

type RegisterConfig struct {
	//Stage       string `json:"env" validate:"oneof=local dev test prod"`
	ServiceName                    string `json:"service_name" validate:"required"`
	MetricsPath                    string `json:"metrics_path" validate:"required"`
	Interval                       string `json:"interval"`                          //检查间隔，默认5s
	Timeout                        string `json:"timeout"`                           //超时时间，默认3s
	FailedFatal                    bool   `json:"failed_fatal"`                      //true:注册失败则fatal
	DeregisterCriticalServiceAfter string `json:"deregister_critical_service_after"` //多久之后注销
}

// GrpcRegisterByEth0IpPort 往consul服务中注册grpc服务
func GrpcRegisterByEth0IpPort(config RegisterConfig, port int) {
	ip, err := getIp()
	if err != nil {
		logrus.WithError(err).Errorf("注册失败, 获取ip失败")
		return
	}
	GrpcRegisterByIpPort(config, ip, port)
}

// getIp 进程启动时候使用, 故不考虑并发问题
func getIp() (string, error) {
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		return "", fmt.Errorf("获取主机ip失败1: %v", err)
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return "", fmt.Errorf("获取主机ip失败2: %v", err)
	}

	for _, a := range addrs {
		name := a.String()
		if idx := strings.Index(name, "/"); idx > 0 {
			name = name[:idx]
		}
		return name, nil
	}
	err = fmt.Errorf("无法获取主机ip")
	fmt.Println(err)
	return "", err
}

var registration *api.AgentServiceRegistration

// GrpcRegisterByIpPort 往consul服务中注册grpc服务
func GrpcRegisterByIpPort(config RegisterConfig, ip string, port int) {
	entry := logrus.WithField("config", config)
	var err error
	defer func() {
		if err != nil {
			entry = entry.WithError(err)
			if config.FailedFatal {
				entry.Fatalf("服务注册失败，终止服务")
			} else {
				entry.Info("服务注册失败，继续服务")
			}
		}
	}()
	if err = validator.New().Struct(config); err != nil {
		return
	}

	entry = entry.WithFields(logrus.Fields{
		"ip":   ip,
		"port": port,
	})

	if config.Interval == "" {
		config.Interval = "5s"
	}
	if config.Timeout == "" {
		config.Timeout = "3s"
	}
	address := fmt.Sprintf("%s:%d", ip, port)
	registration = &api.AgentServiceRegistration{
		Name:    config.ServiceName,
		ID:      address,
		Port:    port,
		Address: ip,
		Check: &api.AgentServiceCheck{
			Interval:                       config.Interval,
			Timeout:                        config.Timeout,
			GRPC:                           fmt.Sprintf("%s/%s", address, config.MetricsPath),
			DeregisterCriticalServiceAfter: config.DeregisterCriticalServiceAfter,
			Status:                         api.HealthPassing,
		},
	}
	err = Client.Agent().ServiceRegister(registration)
	if err != nil {
		return
	}
	serviceIdCh <- address
	entry.Infof("服务注册成功")
}

var serviceIdCh = make(chan string, 1)

// GrpcDRegister grpc deregister
func GrpcDRegister() {
	serviceId := <-serviceIdCh
	entry := logrus.WithField("serviceId", serviceId)
	if err := Client.Agent().ServiceDeregister(serviceId); err != nil {
		entry.WithError(err).Error("服务注销失败")
		return
	}
	entry.Info("服务注销成功")
}
