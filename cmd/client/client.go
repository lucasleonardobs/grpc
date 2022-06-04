package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lucasleonardobs/go-grpc-server/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "Alexia",
		Email: "alexia@mo.com",
	}

	response, err := client.AddUser(context.Background(), request)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(response)
}
