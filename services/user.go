package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/lucasleonardobs/go-grpc-server/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {
	user := &pb.User{
		Id:    "123",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}

	return user, nil
}

func (us *UserService) AddUserVerbose(request *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  request.GetName(),
			Email: request.GetEmail(),
		},
	})

	return nil
}

func (us *UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Could not receive stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    request.GetId(),
			Name:  request.GetName(),
			Email: request.GetEmail(),
		})

		fmt.Println("Adding ", request.GetName())
	}
}

func (us *UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Could not receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   request,
		})
		if err != nil {
			log.Fatalf("Could not send request to the client: %v", err)
		}
	}
}
