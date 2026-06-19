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
	pb.UnimplementedUserServiceServer

	users map[int32]*pb.User

	nextID int32
}

func (s *server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	s.nextID++

	user := &pb.User{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}

	s.users[user.Id] = user

	return &pb.AddUserResponse{
		User: user,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, ok := s.users[req.Id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", req.Id)
	}

	return &pb.GetUserResponse{
		User: user,
	}, nil
}

func (s *server) ListUsers(
	ctx context.Context,
	req *pb.ListUsersRequest,
) (*pb.ListUsersResponse, error) {

	users := []*pb.User{}

	for _, user := range s.users {
		users = append(users, user)
	}

	return &pb.ListUsersResponse{
		Users: users,
	}, nil
}


func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, ok := s.users[req.Id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", req.Id)
	}

	delete(s.users, req.Id)

	return &pb.DeleteUserResponse{
		Message: fmt.Sprintf("user with id %d deleted successfully", req.Id),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // server will listen to this port
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()  // new server is created 
srv := &server{
	users: make(map[int32]*pb.User),
	nextID: 0,
}

pb.RegisterUserServiceServer(
	grpcServer,
	srv,
) // when RegisterGreeterServer is called then use the server created above

	fmt.Println("gRPC server running on :50051")

	err = grpcServer.Serve(lis)  
	if err != nil {
		log.Fatal(err)
	}
}