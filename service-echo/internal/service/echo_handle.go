package service

import (
	"context"
	"github.com/helloteemo/pb/echo"
	"github.com/helloteemo/pb/user"
	user_proxy "github.com/helloteemo/service-echo/internal/proxy/user-proxy"
	"github.com/sirupsen/logrus"
)

func (AppService) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	resp := new(echo.EchoResponse)

	logrus.WithContext(ctx).Infoln("echo request", req.GetMessage())

	client := user_proxy.GetGrpcClient()

	response, err := client.Echo(ctx, &user.UserRequest{UserId: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	resp.Message = "hello " + response.Message + " And " + req.GetMessage()
	return resp, nil
}
