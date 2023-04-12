package service

import (
	"github.com/helloteemo/pb/user"
	"github.com/helloteemo/service-user/internal/config"
)

type AppService struct {
	user.UnimplementedUserServiceServer
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
	}
}
