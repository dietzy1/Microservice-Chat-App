package server

import (
	"context"
	"log"
	"net/http"

	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

func (s *server) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {

	creds := authclientv1.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	//perform client side call to the authentication service
	login, err := s.authClient.Login(ctx, &creds)
	if err != nil {
		log.Println(err)
		//return error code
		return &authv1.LoginResponse{
			Status: http.StatusForbidden,
			Error:  "invalid credentials",
		}, err
	}
	log.Println(login)

	//add the session token to the metadata
	md := metadata.Pairs("session_token", login.Session)
	grpc.SendHeader(ctx, md)

	return &authv1.LoginResponse{
		Status: 200,
		Error:  "no error",
	}, nil

}

func (s *server) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {

	creds := authclientv1.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}

	//perform client side call to the authentication service
	register, err := s.authClient.Register(ctx, &creds)
	if err != nil {
		log.Println(err)
		//return error code
		return &authv1.RegisterResponse{
			Status: http.StatusForbidden,
			Error:  "invalid credentials",
		}, err
	}
	log.Println("Register object", register)

	//Return a session token to the client so the client is authenticated and logged in
	md := metadata.Pairs("session_token", register.Session)
	grpc.SendHeader(ctx, md)

	return &authv1.RegisterResponse{
		Status: 200,
		Error:  "no error",
	}, nil
}

func (s *server) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
		return &authv1.LogoutResponse{
			Status: http.StatusForbidden,
			Error:  "no metadata",
		}, status.Errorf(codes.Unauthenticated, "no metadata")

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
