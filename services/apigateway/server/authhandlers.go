package server

import (
	"context"
	"log"

	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	"google.golang.org/grpc/metadata"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

func (s *server) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//implement whatever logic needs to be implemented
	log.Println("Login called")
	//set a cookie and add it to the grpc response
	/* cookie := &http.Cookie{
		Name:  "session_token",
		Value: "value",
	} */

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
	}
	log.Println(md.Get("session_token"))

	log.Println(req.Username, req.Password)
	//perform client side call to the authentication service

	login, err := s.authClient.C.Login(ctx, req)
	if err != nil {
		log.Println(err)
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

	register, err := s.authClient.C.Register(ctx, req)
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

	logout, err := s.authClient.C.Logout(ctx, req)
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

	authenticate, err := s.authClient.C.Authenticate(ctx, req)
	if err != nil {
		log.Println(err)

	}
	log.Println(authenticate)

	return &authv1.AuthenticateResponse{
		Status: 200,
		Error:  "no error",
	}, nil
}
