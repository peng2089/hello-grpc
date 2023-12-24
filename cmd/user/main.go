package main

import (
	"context"
	"log"
	v1 "my_project/api/user/v1"
	"net"

	"google.golang.org/grpc"
)

type userService struct {
	v1.UnimplementedUserServer
}

func (s *userService) Create(ctx context.Context, in *v1.CreateUserReply) (*v1.CreateUserReply, error) {
	return &v1.CreateUserReply{Id: 11111}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &userService{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("err: %v\n", err)
	}
}
