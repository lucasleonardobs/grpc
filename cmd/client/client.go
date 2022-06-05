package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	AddUserStreamBoth(client)
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

func AddUsers(client pb.UserServiceClient) {
	users := []*pb.User{
		{
			Id:    "1",
			Name:  "Alexia",
			Email: "alexia@mo.com",
		},
		{
			Id:    "2",
			Name:  "Chuchi",
			Email: "chuchi@mo.com",
		},
		{
			Id:    "3",
			Name:  "Lele",
			Email: "lele@mo.com",
		},
		{
			Id:    "4",
			Name:  "Lexi",
			Email: "lexi@mo.com",
		},
		{
			Id:    "5",
			Name:  "Carol",
			Email: "carol@mo.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Could not request to gRPC server: %v", err)
	}

	for _, user := range users {
		stream.Send(user)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not receive response data: %v", err)
	}

	fmt.Println(res.GetUser())
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Could not create the request: %v", err)
	}

	requests := []*pb.User{
		{
			Id:    "1",
			Name:  "Alexia",
			Email: "alexia@mo.com",
		},
		{
			Id:    "2",
			Name:  "Chuchi",
			Email: "chuchi@mo.com",
		},
		{
			Id:    "3",
			Name:  "Lele",
			Email: "lele@mo.com",
		},
		{
			Id:    "4",
			Name:  "Lexi",
			Email: "lexi@mo.com",
		},
		{
			Id:    "5",
			Name:  "Carol",
			Email: "carol@mo.com",
		},
	}

	wait := make(chan bool)

	go func() {
		for _, request := range requests {
			fmt.Println("Sending user: ", request.Name)
			stream.Send(request)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not receive data: %v", err)
			}

			fmt.Println("Receiving user ... ", response.GetUser(), " - ", response.GetStatus())
		}

		close(wait)
	}()

	<-wait
}
