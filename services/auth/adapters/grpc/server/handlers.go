package server

import (
	"context"
	"log"
	"net/http"

	"github.com/dietzy1/chatapp/services/auth/domain"
	authv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	"google.golang.org/grpc/status"
)

//It needs a domain interface it can use to call the domain functions

type Auth interface {
	Login(ctx context.Context, cred domain.Credentials) (domain.Credentials, error)
	Register(ctx context.Context, cred domain.Credentials) (domain.Credentials, error)
	Logout(ctx context.Context, session string, useruuid string) error
	Authenticate(ctx context.Context, session string, useruuid string) (domain.Credentials, error)
	Invalidate(ctx context.Context, userUUid string) error
}

// Pass down the paramters to the interface functions
func (s *server) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//create a credentials struct
	log.Println("LOGIN CALLED")
	cred := domain.Credentials{
		Username: req.Username,
		Password: req.Password,
	}

	//Validate that the fields aren't empty
	if req.Username == "" || req.Password == "" {
		return &authv1.LoginResponse{
			Session:  "",
			UserUuid: "",
		}, nil
	}

	//perform client side call to the authentication service
	creds, err := s.auth.Login(ctx, cred)
	if err != nil {
		log.Println(err)
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

func (s *server) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	cred := domain.Credentials{
		Username: req.Username,
		Password: req.Password,
	}

	//Validate that the fields aren't empty
	if req.Username == "" || req.Password == "" {
		return &authv1.RegisterResponse{
			Session:  "",
			UserUuid: "",
		}, status.Errorf(http.StatusBadRequest, "username or password is empty")
	}

	//perform client side call to the authentication service
	creds, err := s.auth.Register(ctx, cred)
	if err != nil {
		log.Println(err)
		return &authv1.RegisterResponse{
			Session:  "",
			UserUuid: "",
		}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.RegisterResponse{
		Session:  creds.Session,
		UserUuid: creds.Uuid,
	}, nil
}

func (s *server) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	if req.Session == "" || req.UserUuid == "" {
		return &authv1.LogoutResponse{}, status.Errorf(http.StatusBadRequest, "session or user uuid is empty")
	}

	log.Println("Logout GRPC endpoint called")
	err := s.auth.Logout(ctx, req.Session, req.UserUuid)
	if err != nil {
		log.Println(err)
		return &authv1.LogoutResponse{}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.LogoutResponse{}, nil
}

func (s *server) Authenticate(ctx context.Context, req *authv1.AuthenticateRequest) (*authv1.AuthenticateResponse, error) {
	if req.Session == "" || req.UserUuid == "" {
		log.Println("session is or user uuid is empty session: ", req.Session, "userID", req.UserUuid)
		return &authv1.AuthenticateResponse{
			Session:  "",
			UserUuid: "",
		}, nil
	}

	log.Println("session: ", req.Session, "userID", req.UserUuid)

	creds, err := s.auth.Authenticate(ctx, req.Session, req.UserUuid)
	if err != nil {
		log.Println(err)
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
	if req.UserUuid == "" {
		return &authv1.InvalidateResponse{}, status.Errorf(http.StatusBadRequest, "user uuid is empty")
	}

	err := s.auth.Invalidate(ctx, req.UserUuid)
	if err != nil {
		log.Println(err)
		return &authv1.InvalidateResponse{}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.InvalidateResponse{}, nil
}
