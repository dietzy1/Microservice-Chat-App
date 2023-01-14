package server

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	authv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
)

type server struct {
	authv1.UnimplementedAuthServiceServer
	auth Auth
}

func newServer(auth Auth) *server {
	return &server{auth: auth}
}

func Start(auth Auth) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("AUTH")
	//addr := ":9000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	//Inject dependencies into the server
	s := grpc.NewServer()

	dependencies := newServer(auth)

	authv1.RegisterAuthServiceServer(s, dependencies)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

	//it fails to connect because I did not implement the server for the authentication service yet

}
