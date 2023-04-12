package config

import "github.com/helloteemo/utils/consul"

var (
	ProjectName = "echo"
	GrpcPort    = 9002
)

func Init() {
	consul.Init("consul.dev:8500", "")
}
