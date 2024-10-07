package main

import (
	"log"
	"net"

	"github.com/FadyGamilM/usermanagement/internal/user"
	"github.com/FadyGamilM/usermanagement/proto"
	"google.golang.org/grpc"
)

func main() {
	// define a new grpc server
	grpcServer := grpc.NewServer()
	// define a new instant of our service (microservice which implements the rpcs)
	userService := user.UserService{}
	// register the microservice impl into the grpc server
	proto.RegisterUserServiceServer(grpcServer, &userService)
	// listen to a port
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("error listining on port | ERROR : %v", err)
	}

	log.Println("grpc server is up and running on port 50053")

	// blocking function that starts the server and serve requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error serving requests and starting the server | ERROR : %v", err)
	}

}
