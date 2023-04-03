package server

import (
	"context"

	"google.golang.org/grpc/status"

	accountclientv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	accountv1 "github.com/dietzy1/chatapp/services/apigateway/accountgateway/v1"
)

func (s *server) ChangeUsername(ctx context.Context, req *accountv1.ChangeUsernameRequest) (*accountv1.ChangeUsernameResponse, error) {
	// Check if user_uuid and username is empty
	if req.UserUuid == "" || req.Username == "" {
		return &accountv1.ChangeUsernameResponse{}, status.Error(400, "UserUuid and Username cannot be empty")
	}

	//Reroute object
	username := &accountclientv1.ChangeUsernameRequest{
		UserUuid: req.UserUuid,
		Username: req.Username,
	}

	_, err := s.accountClient.ChangeUsername(ctx, username)
	if err != nil {
		return &accountv1.ChangeUsernameResponse{}, status.Error(500, "Internal Server Error")
	}

	return &accountv1.ChangeUsernameResponse{}, nil
}

func (s *server) ChangePassword(ctx context.Context, req *accountv1.ChangePasswordRequest) (*accountv1.ChangePasswordResponse, error) {
	// Check if user_uuid and password is empty
	if req.UserUuid == "" || req.Password == "" {
		return &accountv1.ChangePasswordResponse{}, status.Error(400, "UserUuid and Password cannot be empty")
	}

	//Reroute object
	password := &accountclientv1.ChangePasswordRequest{
		UserUuid: req.UserUuid,
		Password: req.Password,
	}

	_, err := s.accountClient.ChangePassword(ctx, password)
	if err != nil {
		return &accountv1.ChangePasswordResponse{}, status.Error(500, "Internal Server Error")
	}

	return &accountv1.ChangePasswordResponse{}, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *accountv1.DeleteAccountRequest) (*accountv1.DeleteAccountResponse, error) {
	// Check if user_uuid is empty
	if req.UserUuid == "" {
		return &accountv1.DeleteAccountResponse{}, status.Error(400, "UserUuid cannot be empty")
	}

	//Reroute object
	userUuid := &accountclientv1.DeleteAccountRequest{
		UserUuid: req.UserUuid,
	}

	_, err := s.accountClient.DeleteAccount(ctx, userUuid)
	if err != nil {
		return &accountv1.DeleteAccountResponse{}, status.Error(500, "Internal Server Error")
	}

	return &accountv1.DeleteAccountResponse{}, nil
}
