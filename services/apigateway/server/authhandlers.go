package server

import (
	"context"
	"log"
	"time"

	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
	"github.com/dietzy1/chatapp/services/apigateway/metrics"
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

	start := time.Now()
	metrics.AuthRequestCounter.Inc()

	creds := authclientv1.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	//perform client side call to the authentication service
	login, err := s.authClient.Login(ctx, &creds)
	if err != nil {
		log.Println(err)
		//return error code
		return &authv1.LoginResponse{}, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	log.Println(login)

	//add the session token to the metadata
	md := metadata.Pairs("session_token", login.Session, "uuid_token", login.UserUuid)

	//log the metadata
	log.Println(md)
	grpc.SendHeader(ctx, md)

	roundtrip := time.Since(start).Milliseconds()
	metrics.AuthRequestLatency.Observe(float64(roundtrip))

	return &authv1.LoginResponse{}, nil

}

func (s *server) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {

	start := time.Now()
	metrics.AuthRequestCounter.Inc()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
		return &authv1.LogoutResponse{}, status.Errorf(codes.Unauthenticated, "no metadata")

	}

	//extract the token from the metadata
	if len(md["session_token"]) == 0 || len(md["uuid_token"]) == 0 {
		log.Println("no session token")
		return &authv1.LogoutResponse{}, status.Errorf(codes.Unauthenticated, "no session or uuid token")
	}
	session := md["session_token"][0]
	userUuid := md["uuid_token"][0]

	logout, err := s.authClient.Logout(ctx, &authclientv1.LogoutRequest{
		Session:  session,
		UserUuid: userUuid,
	})
	if err != nil {
		log.Println(err)
		return &authv1.LogoutResponse{}, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	log.Println(logout)

	md = metadata.Pairs("session_token", "", "uuid_token", "")
	grpc.SendHeader(ctx, md)

	roundtrip := time.Since(start).Milliseconds()
	metrics.AuthRequestLatency.Observe(float64(roundtrip))

	return &authv1.LogoutResponse{}, nil
}

func (s *server) Authenticate(ctx context.Context, req *authv1.AuthenticateRequest) (*authv1.AuthenticateResponse, error) {

	start := time.Now()
	metrics.AuthRequestCounter.Inc()

	//implement whatever logic needs to be implemented
	log.Println("Authenticate called")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
		return &authv1.AuthenticateResponse{}, status.Errorf(codes.Unauthenticated, "no metadata")
	}

	//TODO: is potential bug here idk why the fuck it keeps on calling no session token on the 2nd request
	//extract the token from the metadata
	if len(md["session_token"]) == 0 || len(md["uuid_token"]) == 0 {
		log.Println("no session token")
		return &authv1.AuthenticateResponse{}, status.Errorf(codes.Unauthenticated, "no session or uuid token")
	}
	session := md["session_token"][0]
	uuid := md["uuid_token"][0]

	log.Println("session token", session)
	log.Println("uuid_token", uuid)

	authenticate, err := s.authClient.Authenticate(ctx, &authclientv1.AuthenticateRequest{
		Session:  session,
		UserUuid: uuid,
	})
	if err != nil {
		return &authv1.AuthenticateResponse{}, status.Errorf(codes.Unauthenticated, "invalid credentials")

	}
	log.Println(authenticate)

	roundtrip := time.Since(start).Milliseconds()
	metrics.AuthRequestLatency.Observe(float64(roundtrip))

	return &authv1.AuthenticateResponse{}, nil
}
