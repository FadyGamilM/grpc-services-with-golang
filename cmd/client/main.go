package main

import (
	"context"
	"log"

	"github.com/FadyGamilM/usermanagement/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	// intiilize grpc connection
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to define new grpc client | ERROR : %v", err)
	}
	defer conn.Close()

	// initialize the grpc client
	client := proto.NewUserServiceClient(conn)
	res, err := client.CreateUser(context.Background(), &proto.CreateUserRequest{Id: 0, Name: ""})
	if err != nil {
		// we need to check the error status before handling it as an internal error
		status, ok := status.FromError(err)
		if ok {
			log.Printf("error creating user and handled by the grpc microservice | error status : %s , error msg : %s\n", status.Code().String(), status.Message())
			panic(status.Err())
		}

		log.Println("error from createUser rpc | ERROR : ", err)
	}
	log.Println("response from createUser rpc : ", res)
}
