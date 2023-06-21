package server

import (
	"context"
	"net/http"

	"github.com/dietzy1/chatapp/services/auth/domain"
	authv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	"google.golang.org/grpc/status"
)

//It needs a domain interface it can use to call the domain functions

type Auth interface {
	Login(ctx context.Context, cred domain.Credentials) (domain.Credentials, error)
	Logout(ctx context.Context, session string, useruuid string) error
	Authenticate(ctx context.Context, session string, useruuid string) (domain.Credentials, error)
	Invalidate(ctx context.Context, userUUid string) error
}

// Pass down the paramters to the interface functions
func (s *server) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//create a credentials struct

	cred := domain.Credentials{
		Username: req.Username,
		Password: req.Password,
	}

	//perform client side call to the authentication service
	creds, err := s.auth.Login(ctx, cred)
	if err != nil {

		return &authv1.LoginResponse{
			Session:  "",
			UserUuid: "",
		}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.LoginResponse{
		Session:  creds.Session,
		UserUuid: creds.Uuid,
	}, nil
}

func (s *server) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {

	err := s.auth.Logout(ctx, req.Session, req.UserUuid)
	if err != nil {
		return &authv1.LogoutResponse{}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.LogoutResponse{}, nil
}

func (s *server) Authenticate(ctx context.Context, req *authv1.AuthenticateRequest) (*authv1.AuthenticateResponse, error) {

	creds, err := s.auth.Authenticate(ctx, req.Session, req.UserUuid)
	if err != nil {
		return &authv1.AuthenticateResponse{
			Session:  "",
			UserUuid: "",
		}, err
	}

	return &authv1.AuthenticateResponse{
		Session:  creds.Session,
		UserUuid: creds.Uuid,
	}, nil
}

func (s *server) Invalidate(ctx context.Context, req *authv1.InvalidateRequest) (*authv1.InvalidateResponse, error) {

	err := s.auth.Invalidate(ctx, req.UserUuid)
	if err != nil {
		return &authv1.InvalidateResponse{}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.InvalidateResponse{}, nil
}
