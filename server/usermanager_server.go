package main

import (
	"context"
	"log"

	pb "github.com/FadyGamilM/usermanagement/usermanager"
)

const (
	server_port = ":50050"
)

type UserManagerServer struct {
	// must be composed here to register this type into gRPC
	pb.UnimplementedUserManagerServer
}

func (ums *UserManagerServer) CreateUser(c context.Context, req pb.NewUser) (*pb.User, error) {
	log.Printf("➜ new user to create : %v \n", req.UserInfo.Username)
	return nil, nil
}
