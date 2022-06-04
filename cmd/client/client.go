package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/lucasleonardobs/go-grpc-server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	AddUserVerbose(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "Alexia",
		Email: "alexia@mo.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), request)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not receive response %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}
