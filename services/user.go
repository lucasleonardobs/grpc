package services

import (
	"context"

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
