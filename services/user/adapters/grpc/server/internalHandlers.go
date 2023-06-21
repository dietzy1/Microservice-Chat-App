package server

import (
	"context"

	"github.com/dietzy1/chatapp/services/user/domain"
	userv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// I might want to split these up
type User interface {
	//internal
	CreateUser(ctx context.Context, username string, uuid string) error
	AddChatServer(ctx context.Context, uuid string, serverUuid string) error
	RemoveChatServer(ctx context.Context, uuid string, server string) error

	//external
	GetUser(ctx context.Context, uuid string) (domain.User, error)
	GetUsers(ctx context.Context, uuids []string) ([]domain.User, error)
	EditDescription(ctx context.Context, uuid string, description string) error
	ChangeAvatar(ctx context.Context, uuid string, iconUuid string) error
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

	//Call to domain layer
	err := s.user.AddChatServer(ctx, req.UserUuid, req.ChatserverUuid)
	if err != nil {
		return &userv1.AddChatServerResponse{}, status.Errorf(codes.Internal, "Error adding chat server: %v", err)
	}

	return &userv1.AddChatServerResponse{}, nil
}

func (s *server) RemoveChatServer(ctx context.Context, req *userv1.RemoveChatServerRequest) (*userv1.RemoveChatServerResponse, error) {

	//Call to domain layer
	err := s.user.RemoveChatServer(ctx, req.UserUuid, req.ChatserverUuid)
	if err != nil {
		return &userv1.RemoveChatServerResponse{}, status.Errorf(codes.Internal, "Error removing chat server: %v", err)
	}

	return &userv1.RemoveChatServerResponse{}, nil
}
