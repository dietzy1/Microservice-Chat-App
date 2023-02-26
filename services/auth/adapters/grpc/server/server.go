package server

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/dietzy1/chatapp/services/auth/adapters/grpc/client"

	authv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

type server struct {
	authv1.UnimplementedAuthServiceServer
	auth Auth

	userClient userClientv1.UserServiceClient
}

func newServer(auth Auth, userClient userClientv1.UserServiceClient) *server {
	return &server{auth: auth, userClient: userClient}
}

func Start(auth Auth) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	//addr := ":9000"
	addr := os.Getenv("AUTH")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingMiddleware),
	)

	userClient := client.NewUserClient()

	//Inject dependencies into the server
	dependencies := newServer(auth, *userClient)

	//Register the server
	authv1.RegisterAuthServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
