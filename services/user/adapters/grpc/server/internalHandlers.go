package server

import (
	"context"

	userv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User interface {
	//CreateUser creates a new user
	CreateUser(ctx context.Context, username string, uuid string) error
}

//implement the handlers here

// CreateUser creates a new user
func (s *server) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {

	err := s.user.CreateUser(ctx, req.Username, req.UserUuid)
	if err != nil {
		return &userv1.CreateUserResponse{}, status.Errorf(codes.Internal, "Error creating user: %v", err)
	}

	return &userv1.CreateUserResponse{}, nil
}

func (s *server) AddChatServer(ctx context.Context, req *userv1.AddChatServerRequest) (*userv1.AddChatServerResponse, error) {
	panic("implement me")
}

func (s *server) RemoveChatServer(ctx context.Context, req *userv1.RemoveChatServerRequest) (*userv1.RemoveChatServerResponse, error) {
	panic("implement me")
}
