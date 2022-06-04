package services

import (
	"context"
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
