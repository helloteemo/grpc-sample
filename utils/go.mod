module github.com/helloteemo/utils

go 1.15

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/gosexy/to v0.0.0-20141221203644-c20e083e3123
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/hashicorp/consul/api v1.20.0
	github.com/json-iterator/go v1.1.12
	github.com/mercari/go-grpc-interceptor v0.0.0-20180110035004-b8ad3827e82a
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.10.0 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.51.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)
