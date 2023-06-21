package server

import (
	"io"
	"net"
	"os"

	"github.com/dietzy1/chatapp/pkg/middleware"
	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type server struct {
	messagev1.UnimplementedMessageServiceServer
	domain message
}

func newServer(domain message) *server {
	return &server{domain: domain}
}

func Start(logger *zap.Logger, domain message) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("MESSAGESERVICE")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingMiddleware(logger)),
	)
	//Inject dependencies into the server
	dependencies := newServer(domain)

	//Register the server
	messagev1.RegisterMessageServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
