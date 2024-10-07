package user

import (
	"context"

	"github.com/FadyGamilM/usermanagement/proto"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

func (us *UserService) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return nil, nil
}
