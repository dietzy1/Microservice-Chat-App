package server

import (
	"context"
	"log"

	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	"google.golang.org/grpc/metadata"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

func (s *server) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//implement whatever logic needs to be implemented
	log.Println("Login called")
	//perform client side call to the authentication service
	creds := authclientv1.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	login, err := s.authClient.Login(ctx, &creds)
	if err != nil {
		log.Println(err)
		log.Println("ERROR")
	}
	log.Println(login)

	return &authv1.LoginResponse{
		Status: 200,
		Error:  "no error",
	}, nil

}

func (s *server) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	//implement whatever logic needs to be implemented
	log.Println("Register called")

	register, err := s.authClient.Register(ctx, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println(register)

	return &authv1.RegisterResponse{
		Status: 200,
		Error:  "no error",
	}, nil
}

func (s *server) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	//implement whatever logic needs to be implemented
	log.Println("Logout called")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
	}
	log.Println(md)
	//extract the token from the metadata
	session := md["session_token"][0]

	logout, err := s.authClient.Logout(ctx, &authclientv1.LogoutRequest{
		Session: session,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(logout)

	return &authv1.LogoutResponse{
		Status: 200,
		Error:  "no error",
	}, nil
}

func (s *server) Authenticate(ctx context.Context, req *authv1.AuthenticateRequest) (*authv1.AuthenticateResponse, error) {
	//implement whatever logic needs to be implemented
	log.Println("Authenticate called")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
	}
	log.Println(md)
	//extract the token from the metadata
	session := md["session_token"][0]

	authenticate, err := s.authClient.Authenticate(ctx, &authclientv1.AuthenticateRequest{
		Session: session})
	if err != nil {
		log.Println(err)

	}
	log.Println(authenticate)

	return &authv1.AuthenticateResponse{
		Status: 200,
		Error:  "no error",
	}, nil
}
