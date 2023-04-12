package config

import "github.com/helloteemo/utils/consul"

var (
	ProjectName = "user"
	GrpcPort    = 9001
)

func Init() {
	consul.Init("consul.dev:8500", "")
}
