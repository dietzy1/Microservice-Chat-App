package server

import (
	"context"

	userv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User interface {
	CreateUser(ctx context.Context, username string, uuid string) error
	AddChatServer(ctx context.Context, uuid string, serverUuid string) error
	RemoveChatServer(ctx context.Context, uuid string, server string) error
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

	err := s.user.AddChatServer(ctx, req.UserUuid, req.ChatserverUuid)
	if err != nil {
		return &userv1.AddChatServerResponse{}, status.Errorf(codes.Internal, "Error adding chat server: %v", err)
	}

	return &userv1.AddChatServerResponse{}, nil
}

func (s *server) RemoveChatServer(ctx context.Context, req *userv1.RemoveChatServerRequest) (*userv1.RemoveChatServerResponse, error) {

	err := s.user.RemoveChatServer(ctx, req.UserUuid, req.ChatserverUuid)
	if err != nil {
		return &userv1.RemoveChatServerResponse{}, status.Errorf(codes.Internal, "Error removing chat server: %v", err)
	}

	return &userv1.RemoveChatServerResponse{}, nil
}
