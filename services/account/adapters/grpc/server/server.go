package server

import (
	"io"
	"net"
	"os"

	"github.com/dietzy1/chatapp/pkg/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	accountv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
)

type server struct {
	accountv1.UnimplementedAccountServiceServer
	domain Account
}

func newServer(a Account) *server {
	return &server{domain: a}
}

func Start(a Account) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("ACCOUNT")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingMiddleware),
	)
	//Inject dependencies into the server
	dependencies := newServer(a)

	//Register the server
	accountv1.RegisterAccountServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
