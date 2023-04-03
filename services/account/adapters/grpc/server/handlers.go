package server

import (
	"context"

	accountv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	"google.golang.org/grpc/status"
)

type Account interface {
	ChangeUsername(ctx context.Context, userUuid string, username string) error
	ChangePassword(ctx context.Context, userUuid string, username string) error
	DeleteAccount(ctx context.Context, userUuid string, password string) error
}

func (s *server) ChangeUsername(ctx context.Context, req *accountv1.ChangeUsernameRequest) (*accountv1.ChangeUsernameResponse, error) {
	//Check if useruuid and username is empty
	if req.UserUuid == "" || req.Username == "" {
		return &accountv1.ChangeUsernameResponse{}, status.Error(400, "UserUuid or Username is empty")
	}

	//Call domain
	err := s.account.ChangeUsername(ctx, req.UserUuid, req.Username)
	if err != nil {
		return &accountv1.ChangeUsernameResponse{}, status.Error(400, "Error changing username")
	}

	return &accountv1.ChangeUsernameResponse{}, nil

}

func (s *server) ChangePassword(ctx context.Context, req *accountv1.ChangePasswordRequest) (*accountv1.ChangePasswordResponse, error) {
	//Check if useruuid and password is empty
	if req.UserUuid == "" || req.Password == "" {
		return &accountv1.ChangePasswordResponse{}, status.Error(400, "UserUuid or Password is empty")
	}

	//Call domain
	err := s.account.ChangePassword(ctx, req.UserUuid, req.Password)
	if err != nil {
		return &accountv1.ChangePasswordResponse{}, status.Error(400, "Error changing password")
	}

	return &accountv1.ChangePasswordResponse{}, nil

}

func (s *server) DeleteAccount(ctx context.Context, req *accountv1.DeleteAccountRequest) (*accountv1.DeleteAccountResponse, error) {
	//Check if useruuid and password is empty
	if req.UserUuid == "" || req.Password == "" {
		return &accountv1.DeleteAccountResponse{}, status.Error(400, "UserUuid or Password is empty")
	}

	//Call domain
	err := s.account.DeleteAccount(ctx, req.UserUuid, req.Password)
	if err != nil {
		return &accountv1.DeleteAccountResponse{}, status.Error(400, "Error deleting account")
	}

	return &accountv1.DeleteAccountResponse{}, nil
}
