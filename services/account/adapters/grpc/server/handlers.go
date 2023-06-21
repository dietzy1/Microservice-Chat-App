package server

import (
	"context"
	"log"
	"net/http"

	"github.com/dietzy1/chatapp/services/account/domain"
	accountv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	"google.golang.org/grpc/status"
)

type Account interface {
	ChangeUsername(ctx context.Context, userUuid string, username string) error
	ChangePassword(ctx context.Context, userUuid string, password string, newPassword string) error

	DeleteAccount(ctx context.Context, userUuid string, password string) error
	RegisterAccount(ctx context.Context, creds domain.Credentials) (domain.Credentials, error)
	DemoUserRegister(ctx context.Context) (domain.Credentials, error)
	UpgradeDemoUser(ctx context.Context, creds domain.Credentials) error
}

func (s *server) ChangeUsername(ctx context.Context, req *accountv1.ChangeUsernameRequest) (*accountv1.ChangeUsernameResponse, error) {

	//Call domain
	err := s.domain.ChangeUsername(ctx, req.UserUuid, req.Username)
	if err != nil {
		return &accountv1.ChangeUsernameResponse{}, status.Errorf(400, "Error changing username %s", err.Error())
	}

	return &accountv1.ChangeUsernameResponse{}, nil

}

func (s *server) ChangePassword(ctx context.Context, req *accountv1.ChangePasswordRequest) (*accountv1.ChangePasswordResponse, error) {

	//Call domain
	err := s.domain.ChangePassword(ctx, req.UserUuid, req.Password, req.NewPassword)
	if err != nil {
		return &accountv1.ChangePasswordResponse{}, status.Errorf(400, "Error changing password %s", err.Error())
	}

	return &accountv1.ChangePasswordResponse{}, nil

}

func (s *server) DeleteAccount(ctx context.Context, req *accountv1.DeleteAccountRequest) (*accountv1.DeleteAccountResponse, error) {

	//Call domain
	err := s.domain.DeleteAccount(ctx, req.UserUuid, req.Password)
	if err != nil {
		return &accountv1.DeleteAccountResponse{}, status.Errorf(400, "Error deleting account %s", err.Error())
	}

	return &accountv1.DeleteAccountResponse{}, nil
}

func (s *server) RegisterAccount(ctx context.Context, req *accountv1.RegisterAccountRequest) (*accountv1.RegisterAccountResponse, error) {

	cred := domain.Credentials{
		Username: req.Username,
		Password: req.Password,
	}

	//perform client side call to the authentication service
	creds, err := s.domain.RegisterAccount(ctx, cred)
	if err != nil {
		return &accountv1.RegisterAccountResponse{
			Session:  "",
			UserUuid: "",
		}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &accountv1.RegisterAccountResponse{
		Session:  creds.Session,
		UserUuid: creds.Uuid,
	}, nil
}

func (s *server) DemoUserRegister(ctx context.Context, req *accountv1.DemoUserRegisterRequest) (*accountv1.DemoUserRegisterResponse, error) {

	creds, err := s.domain.DemoUserRegister(ctx)
	if err != nil {
		log.Println(err)
		return &accountv1.DemoUserRegisterResponse{
			Session:  "",
			UserUuid: "",
		}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &accountv1.DemoUserRegisterResponse{
		Session:  creds.Session,
		UserUuid: creds.Uuid,
	}, nil
}

func (s *server) UpgradeDemoUser(ctx context.Context, req *accountv1.UpgradeDemoUserRequest) (*accountv1.UpgradeDemoUserResponse, error) {

	//Create credentials struct
	creds := domain.Credentials{
		Username: req.Username,
		Password: req.Password,
		Uuid:     req.UserUuid,
	}

	err := s.domain.UpgradeDemoUser(ctx, creds)
	if err != nil {
		return &accountv1.UpgradeDemoUserResponse{}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &accountv1.UpgradeDemoUserResponse{}, nil

}
