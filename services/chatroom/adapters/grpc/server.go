package grpc

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type server struct {
	chatroomv1.UnimplementedChatroomServiceServer
	chatroom Chatroom
}

func newServer(chatroom Chatroom) *server {
	return &server{chatroom: chatroom}
}

func Start(chatroom Chatroom) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	//addr := ":9000"
	addr := os.Getenv("CHATROOM")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	//Inject dependencies into the server
	dependencies := newServer(chatroom)

	//Register the server
	chatroomv1.RegisterAuthServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}
