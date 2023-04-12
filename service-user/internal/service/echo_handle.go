package service

import (
	"context"
	"fmt"
	"github.com/helloteemo/pb/user"
	"github.com/helloteemo/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

func (AppService) Echo(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	logrus.WithContext(ctx).Infoln("echo request")

	if req.GetUserId() < 0 {
		return nil, utils.GrpcError(codes.InvalidArgument, "user id is invalid")
	}
	return &user.UserResponse{
		Message: fmt.Sprintf("user id is %d", req.GetUserId()),
	}, nil
}
