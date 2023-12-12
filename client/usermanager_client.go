package main

import (
	"context"
	"log"
	"time"

	pb "github.com/FadyGamilM/usermanagement/usermanager"

	"google.golang.org/grpc"
)

const (
	server_addrs = "0.0.0.0:50050"
)

func main() {
	conn, err := grpc.Dial(server_addrs, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("[info] ➜ failed to establish a connection with the gRPC server : %v \n", err)
	}
	defer conn.Close()

	// get a client instance
	client := pb.NewUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// define users to be sent to server as a request
	users := make(map[string]int32, 0)
	users["fadyGamil"] = 25
	users["ahmedRushdy"] = 26
	for username, age := range users {
		req := &pb.NewUser{
			UserInfo: &pb.UserDetails{Username: username, Age: age},
		}
		createdUser, err := client.CreateUser(ctx, req)
		if err != nil {
			log.Printf("[info] ➜ error creating user with username %v , error : %v \n", username, err)
		}
		log.Printf("user created successfully \t\t\n username : %v, \t\t\n age : %v, \t\t\n user_id : %v \t\t\n", createdUser.UserInfo.GetUsername(), createdUser.UserInfo.Age, createdUser.UserInfo)
	}
}
