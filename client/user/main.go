package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "my_project/api/user"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Err: %+v\n", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, &pb.CreateUserRequest{Name: ""})
	if err != nil {
		log.Fatalf("count not create: %v", err)
	}
	log.Printf("Greeting: %d", r.GetId())
}
