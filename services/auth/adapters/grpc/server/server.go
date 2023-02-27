package server

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

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

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingMiddleware),
	)

	//Inject dependencies into the server
	dependencies := newServer(auth)

	//Register the server
	authv1.RegisterAuthServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
