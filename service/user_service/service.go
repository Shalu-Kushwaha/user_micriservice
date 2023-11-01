package user_service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"log"
	pb "user_microservice/proto"
)

// User represents user details.
type User struct {
	ID      int32   `validate:"required"` // Unique identifier for the user.
	FName   string  `validate:"required"` // User's first name.
	City    string  `validate:"required"` // User's city.
	Phone   int64   // User's phone number.
	Height  float32 // User's height.
	Married bool    // Indicates whether the user is married.

}

// UserService is a gRPC service that provides user-related operations.
type UserService struct {
	pb.UnimplementedUserServiceServer // Embed the unimplemented server
	users                             []User
	validate                          *validator.Validate
}

// NewUserService creates a new instance of UserService.
func NewUserService() *UserService {
	return &UserService{
		users: []User{
			{ID: 1, FName: "Steve", City: "LA", Phone: 7632146873, Height: 5.2, Married: true},
			{ID: 2, FName: "John", City: "NY", Phone: 5645321345, Height: 4.8, Married: false},
			{ID: 3, FName: "Alice", City: "SF", Phone: 8753235432, Height: 6.3, Married: true},
		},
		validate: validator.New(),
	}
}

// GetUserById retrieves a user by ID.
func (us *UserService) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if err := us.validate.Struct(req); err != nil {
		log.Printf("Invalid request: %v \n", err)
		return nil, fmt.Errorf("invalid request: %v", err)
	}

	log.Printf("Received request for User ID: %d \n", req.Id)

	for _, user := range us.users {
		if user.ID == req.Id {
			log.Printf("User with ID %d found", req.Id)
			return &pb.UserResponse{Id: user.ID, Fname: user.FName, City: user.City, Phone: user.Phone, Height: user.Height, Married: user.Married}, nil
		}
	}
	log.Printf("User with ID %d not found \n", req.Id)
	return nil, fmt.Errorf("user with ID %d not found", req.Id)
}

// GetUsersByIds retrieves a list of users by a list of IDs.
func (us *UserService) GetUsersByIds(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	if err := us.validate.Struct(req); err != nil {
		log.Printf("Invalid request: %v", err)
		return nil, fmt.Errorf("invalid request: %v", err)
	}

	log.Printf("Received request for User IDs: %v \n", req.Ids)

	var response pb.UsersResponse
	for _, id := range req.Ids {
		for _, user := range us.users {
			if user.ID == id {
				log.Printf("User with ID %d found \n", id)
				response.Users = append(response.Users, &pb.UserResponse{Id: user.ID, Fname: user.FName, City: user.City, Phone: user.Phone, Height: user.Height, Married: user.Married})
			}
		}
	}
	return &response, nil
}
