# proto

所有proto文件都定义在这里

# 目录结构

如果需要新增一个服务, 则直接在根目录中新增一个以服务名命名的文件夹, 然后在该文件夹中新增一个以服务名命名的proto文件, 例如新增一个名为`user`的服务, 则在根目录中新增一个`user`文件夹, 然后在`user`文件夹中新增一个`user.proto`文件, 该文件中定义的所有proto都属于`user`服务

# 包名

包名为应用的表示(APPID). 用于生成gRPC的请求路径, 或者在 Proto 之间进行引用Message

例如:

```protobuf
// RequestURL: /grpc_sample.user/${serviceName}/${methodName}
package grpc_sample.user;
```

其中 `grpc_sample` 为固定写法, `user` 为APPID

> 目前有两种写法: `grpc_sample.${APPID}` 和 `apis.${APPID}`, 两种写法都可以, 但是为了统一, 建议使用 `grpc_sample.${APPID}`

## go_package

```protobuf
// 其中 ${APPID} echo
option go_package = "github.com/helloteemo/pb/echo";
```

# Import

- 业务 proto 依赖，以根目录进行引入对应依赖的 proto。
- third_party，主要为依赖的第三方 proto，比如 protobuf、google rpc、google apis、gogo 定义。

