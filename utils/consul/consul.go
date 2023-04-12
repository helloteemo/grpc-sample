package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

var (
	Address string
	Client  *api.Client
)

func Init(address string, prefix string, options ...Option) {
	entry := logrus.WithFields(logrus.Fields{
		"address": address,
		"prefix":  prefix,
	})
	Address = address

	cfg := &api.Config{
		Address: address,
	}

	for _, option := range options {
		option(cfg)
	}

	consulClient, err := api.NewClient(cfg)
	if err != nil {
		entry.WithError(err).Panic("consul connect failed")
	}
	Client = consulClient
	entry.Info("consul connect ok")
}
