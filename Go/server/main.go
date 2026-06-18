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
	return &pb.HelloResponse{         // this is the sayhello function which we have setup  .helloRequest is setup to get from client
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // server will listen to this port
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()  // new server is created 
	pb.RegisterGreeterServer(grpcServer, &server{}) // when RegisterGreeterServer is called then use the server created above

	fmt.Println("gRPC server running on :50051")

	err = grpcServer.Serve(lis)  
	if err != nil {
		log.Fatal(err)
	}
}