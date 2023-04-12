module github.com/helloteemo/service-user

go 1.15

replace (
	github.com/helloteemo/pb => ../pb
	github.com/helloteemo/utils => ../utils
	golang.org/x/net => golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
)

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/helloteemo/pb v0.0.0-00010101000000-000000000000
	github.com/helloteemo/utils v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	google.golang.org/grpc v1.51.0
)
