package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-practice/proto/hellopb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	fmt.Println("gRPC server running on :50051")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}