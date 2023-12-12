package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/FadyGamilM/usermanagement/usermanager"
	"google.golang.org/grpc"
)

const (
	server_port = ":50050"
)

type UserManagerServer struct {
	// must be composed here to register this type into gRPC
	pb.UnimplementedUserManagerServer
}

func (ums *UserManagerServer) CreateUser(c context.Context, req *pb.NewUser) (*pb.User, error) {
	log.Printf("[info] ➜ new user to create : %v \n", req.UserInfo.Username)
	// define the user response and persist it in db and return the response
	rand.Seed(time.Now().Unix())
	userID := rand.Intn(1000000)
	res := &pb.User{
		UserId: int64(userID),
		UserInfo: &pb.UserDetails{
			Username: req.UserInfo.Username,
			Age:      req.UserInfo.Age,
		},
	}
	return res, nil
}

func main() {
	// define a listener
	listener, err := net.Listen("tcp", server_port)
	if err != nil {
		log.Fatalf("[X] ➜ failed to create a listener : %v \n", err)
	}

	server := grpc.NewServer()

	// now register the server as a new grpc service and register our struct
	pb.RegisterUserManagerServer(server, &UserManagerServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("[X] ➜ failed to start the server : %v \n", err)
	}
}
