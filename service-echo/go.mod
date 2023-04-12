module github.com/helloteemo/service-echo

go 1.15

replace (
	github.com/helloteemo/pb => ../pb
	github.com/helloteemo/utils => ../utils
	github.com/mbobakov/grpc-consul-resolver => github.com/haozzzzzzzz/grpc-consul-resolver v0.0.0-20220801024211-c4d530953621
	golang.org/x/net => golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
)

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/helloteemo/pb v0.0.0-00010101000000-000000000000
	github.com/helloteemo/utils v0.0.0-00010101000000-000000000000
	github.com/mbobakov/grpc-consul-resolver v1.5.2
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.54.0
)
