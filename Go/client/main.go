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

	//this the part where frontend is talking to backend
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // insecure  = dont use tls security on local
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // close the connection 

	c := pb.NewGreeterClient(conn) // this is object initializing the connection

	
ctx, cancel := context.WithTimeout( // if server didnt respond with in 1s second cut of the connection
	context.Background(),
	time.Second,
)

defer cancel()

res, err := c.SayHello(  // this part which is calling and sending the name like rest 
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
