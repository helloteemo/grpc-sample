# grpc_gateway_util

## 功能点 

1. 异常处理
2. 异常日志输出,集成 `logrus`
3. response 格式封装

## Getting started

主函数需要添加 `ErrHandler` 函数和 `GinHandler` 函数

`ErrHandler` 加在 grpc gateway 的中间件中

`GinHandler` 加载 gin 的中间件中

```go
grpcGatewayMux := grpc_gateway_runtime.NewServeMux(
    // 这里处理异常
    grpc_gateway_runtime.WithErrorHandler(grpc_gateway_util.ErrHandler),
)

engine := gin.New()
// 使用 GinHandler 自动封装 data 
engine.Use(grpc_gateway_util.GinHandler)
engine.Any(
    "*",
    gin.WrapH(grpcGatewayMux),
)

if err = engine.Run(httpListenAddr); err != nil {
    logrus.WithError(err).Errorln("failed to listen http server")
}
```

这样 Grpc 函数在返回 err 后会在 ErrHandler 中被拦截然后返回固定格式的resp

```go
func (s *Server) GetPKStatus(ctx context.Context, req *v1.PkStatusReq) (reply *v1.PkStatusResp, err error) {
    return reply, status2.New(codes.Internal, "failed to get pk status").Err()
}
// 会得到
/*
{
    "ret": -2013,
    "msg": "failed to get pk status"
}
*/
```

如果返回的是data的话那么会在 grpc gateway 的 gin 中间件中被加到data字段中

```json
{
    "ret": 1,
    "data": {}
}
```