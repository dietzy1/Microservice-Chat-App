package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	accountclientv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	accountv1 "github.com/dietzy1/chatapp/services/apigateway/accountgateway/v1"
	"github.com/dietzy1/chatapp/services/apigateway/metrics"
)

func (s *server) ChangeUsername(ctx context.Context, req *accountv1.ChangeUsernameRequest) (*accountv1.ChangeUsernameResponse, error) {

	metrics.AccountRequestCounter.Inc()

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

	metrics.AccountRequestCounter.Inc()
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

	metrics.AccountRequestCounter.Inc()
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

func (s *server) RegisterAccount(ctx context.Context, req *accountv1.RegisterAccountRequest) (*accountv1.RegisterAccountResponse, error) {

	metrics.AccountRequestCounter.Inc()

	creds := accountclientv1.RegisterAccountRequest{
		Username: req.Username,
		Password: req.Password,
	}

	//perform client side call to the authentication service
	register, err := s.accountClient.RegisterAccount(ctx, &creds)
	if err != nil {
		log.Println(err)
		//return error code
		return &accountv1.RegisterAccountResponse{}, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	log.Println("Register object", register)

	//Return a session token to the client so the client is authenticated and logged in
	md := metadata.Pairs("session_token", register.Session)
	grpc.SendHeader(ctx, md)

	return &accountv1.RegisterAccountResponse{}, nil
}

func (s *server) DemoUserRegister(ctx context.Context, req *accountv1.DemoUserRegisterRequest) (*accountv1.DemoUserRegisterResponse, error) {

	metrics.AccountRequestCounter.Inc()
	//Reroute object
	register, err := s.accountClient.DemoUserRegister(ctx, &accountclientv1.DemoUserRegisterRequest{})
	if err != nil {
		return &accountv1.DemoUserRegisterResponse{}, status.Error(500, "Internal Server Error")
	}

	//Return a session token to the client so the client is authenticated and logged in
	md := metadata.Pairs("session_token", register.Session)
	grpc.SendHeader(ctx, md)

	return &accountv1.DemoUserRegisterResponse{}, nil
}

func (s *server) UpgradeDemoUser(ctx context.Context, req *accountv1.UpgradeDemoUserRequest) (*accountv1.UpgradeDemoUserResponse, error) {

	metrics.AccountRequestCounter.Inc()
	//Reroute object
	upgrade := &accountclientv1.UpgradeDemoUserRequest{
		UserUuid: req.UserUuid,
		Username: req.Username,
		Password: req.Password,
	}

	_, err := s.accountClient.UpgradeDemoUser(ctx, upgrade)
	if err != nil {
		return &accountv1.UpgradeDemoUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &accountv1.UpgradeDemoUserResponse{}, nil
}
