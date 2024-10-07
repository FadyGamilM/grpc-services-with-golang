package user

import (
	"context"

	"github.com/FadyGamilM/usermanagement/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

func (us *UserService) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	if in.GetId() < 0 || in.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid id or username")
	}
	return &proto.CreateUserResponse{
		Id:        in.GetId(),
		Name:      in.GetName(),
		Activated: true,
	}, nil
}
