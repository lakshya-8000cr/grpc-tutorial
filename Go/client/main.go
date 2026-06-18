package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-practice/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	
ctx, cancel := context.WithTimeout(
	context.Background(),
	time.Second,
)

defer cancel()

res, err := c.SayHello(
	ctx,
	&pb.HelloRequest{
		Name: "Lakshya",
	},
)

if err != nil {
	log.Fatal(err)
}

fmt.Println(res.Message)
}
