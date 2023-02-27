package server

import (
	"fmt"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	userv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

type server struct {
	userv1.UnimplementedUserServiceServer
	user User
}

func newServer(user User) *server {
	return &server{user: user}
}

func Start(user User) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("USERSERVICE")

	fmt.Println("USER addr: ", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingMiddleware),
	)
	//Inject dependencies into the server
	dependencies := newServer(user)

	//Register the server
	userv1.RegisterUserServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
