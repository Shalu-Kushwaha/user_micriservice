package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "user_microservice/proto"
	"user_microservice/service/user_service"
)

func main() {
	// Listen for incoming connections on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the UserService implementation with the gRPC server
	pb.RegisterUserServiceServer(grpcServer, user_service.NewUserService())

	// Register the reflection service on the gRPC server for debugging and exploration
	reflection.Register(grpcServer)

	// Print a message indicating that the server is running
	fmt.Println("Server is running on port 50051...")

	// Start serving gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
