package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	pb "my_project/api/user"
	"net"

	"google.golang.org/grpc"
)

type userService struct {
	pb.UnimplementedUserServer
}

// Create 创建用户
func (s *userService) Create(ctx context.Context,
	in *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	fmt.Printf("in: %+v\n", in)
	if len(in.Name) == 0 {
		return nil, errors.New("Name不能为空")
	}
	return &pb.CreateUserReply{Id: 11111}, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &userService{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("err: %v\n", err)
	}
}
