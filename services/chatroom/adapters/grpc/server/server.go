package server

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	chatroomv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
)

type server struct {
	chatroomv1.UnimplementedChatroomServiceServer

	chatroom chatroom
}

func newServer(chatroom chatroom) *server {
	return &server{chatroom: chatroom}
}

func Start(chatroom chatroom) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	//addr := ":9000"
	addr := os.Getenv("CHATROOM")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingMiddleware),
	)
	//Inject dependencies into the server
	dependencies := newServer(chatroom)

	//Register the server
	chatroomv1.RegisterChatroomServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
