package service

import (
	"github.com/helloteemo/pb/echo"
	"github.com/helloteemo/service-echo/internal/config"
	user_proxy "github.com/helloteemo/service-echo/internal/proxy/user-proxy"
)

type AppService struct {
	echo.UnimplementedEchoServiceServer
}

func NewAppService() *AppService {
	imp := &AppService{}
	options := []option{
		withConfig(),
		withDB(),
	}
	for _, o := range options {
		o(imp)
	}
	return imp
}

type option func(imp *AppService)

func withDB() option {
	return func(imp *AppService) {
		// do something
	}
}

func withConfig() option {
	return func(imp *AppService) {
		config.Init()
		user_proxy.Init()
	}
}
