package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	pb "my_project/api/user"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type userService struct {
	pb.UnimplementedUserServer
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{
		db: db,
	}
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

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannnot connect mysql: %+v", err)
	}

	return db
}

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	s := grpc.NewServer()
	db := NewDB()
	service := NewUserService(db)
	pb.RegisterUserServer(s, service)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("err: %v\n", err)
	}
}
