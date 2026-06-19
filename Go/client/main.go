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

c := pb.NewUserServiceClient(conn)

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

	

user1, err := c.AddUser(ctx, &pb.AddUserRequest{
	Name:  "Lakshya",
	Email: "lakshya@example.com",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println("Added:", user1.User)

user2, err := c.AddUser(ctx, &pb.AddUserRequest{
	Name:  "Ayush",
	Email: "ayush@example.com",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println("Added:", user2.User)

list, err := c.ListUsers(ctx, &pb.ListUsersRequest{})
if err != nil {
	log.Fatal(err)
}
fmt.Println("All users:", list.Users)

get, err := c.GetUser(ctx, &pb.GetUserRequest{
	Id: user1.User.Id,
})
if err != nil {
	log.Fatal(err)
}
fmt.Println("Fetched:", get.User)

del, err := c.DeleteUser(ctx, &pb.DeleteUserRequest{
	Id: user1.User.Id,
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(del.Message)

listAfterDelete, err := c.ListUsers(ctx, &pb.ListUsersRequest{})
if err != nil {
	log.Fatal(err)
}
fmt.Println("After delete:", listAfterDelete.Users)


if err != nil {
	log.Fatal(err)
}

}
