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
	Login(ctx context.Context, cred domain.Credentials) (string, error)
	Register(ctx context.Context, cred domain.Credentials) (string, error)
	Logout(ctx context.Context, session string, useruuid string) error
	Authenticate(ctx context.Context, session string, useruuid string) (string, error)
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
			Session: "",
			Error:   "username or password is empty",
		}, nil
	}

	//perform client side call to the authentication service
	session, err := s.auth.Login(ctx, cred)
	if err != nil {
		log.Println(err)
		return &authv1.LoginResponse{
			Session: "",
			Error:   err.Error(),
		}, status.Errorf(http.StatusBadRequest, err.Error())
	}

	return &authv1.LoginResponse{
		Session: session,
		Error:   "",
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
			Session: "",
			Error:   "username or password is empty",
		}, status.Errorf(http.StatusBadRequest, "username or password is empty")
	}

	//perform client side call to the authentication service
	session, err := s.auth.Register(ctx, cred)
	if err != nil {
		log.Println(err)
		return &authv1.RegisterResponse{
			Session: "",
			Error:   err.Error(),
		}, status.Errorf(http.StatusBadRequest, err.Error())
	}
	log.Println(session)
	log.Println(err)

	return &authv1.RegisterResponse{
		Session: session,
		Error:   "no error",
	}, nil
}

func (s *server) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	if req.Session == "" || req.UserUuid == "" {
		return &authv1.LogoutResponse{
			Error: "session or user uuid is empty",
		}, status.Errorf(http.StatusBadRequest, "session or user uuid is empty")
	}

	log.Println("Logout GRPC endpoint called")
	err := s.auth.Logout(ctx, req.Session, req.UserUuid)
	if err != nil {
		log.Println(err)
		return &authv1.LogoutResponse{
			Error: err.Error(),
		}, err
	}

	return &authv1.LogoutResponse{
		Error: "",
	}, nil
}

func (s *server) Authenticate(ctx context.Context, req *authv1.AuthenticateRequest) (*authv1.AuthenticateResponse, error) {
	if req.Session == "" || req.UserUuid == "" {
		log.Println("session is or user uuid is empty session: ", req.Session, "userID", req.UserUuid)
		return &authv1.AuthenticateResponse{
			Error: "session is or user uuid is empty",
		}, nil
	}

	log.Println("session: ", req.Session, "userID", req.UserUuid)

	session, err := s.auth.Authenticate(ctx, req.Session, req.UserUuid)
	if err != nil {
		log.Println(err)
		return &authv1.AuthenticateResponse{
			Session: "",
			Error:   err.Error(),
		}, err
	}

	return &authv1.AuthenticateResponse{
		Session: session,
		Error:   "",
	}, nil
}

func (s *server) Invalidate(ctx context.Context, req *authv1.InvalidateRequest) (*authv1.InvalidateResponse, error) {
	if req.UserUuid == "" {
		return &authv1.InvalidateResponse{
			Error: "user uuid is empty",
		}, nil
	}

	err := s.auth.Invalidate(ctx, req.UserUuid)
	if err != nil {
		log.Println(err)
		return &authv1.InvalidateResponse{
			Error: err.Error(),
		}, err
	}

	return &authv1.InvalidateResponse{
		Error: "",
	}, nil
}
