package client

import (
	"context"
	"log"
	"os"

	chatroomclientv1 "github.com/dietzy1/chatapp/services/apigateway/clients/chatroom/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// func NewChatRoomClient() *chatroomclientv1.ChatroomGatewayServiceClient {
func NewChatRoomClient() *chatroomclient.v1 {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("CHATROOM"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := chatroomclientv1.NewChatroomGatewayServiceClient(conn)

	return &client
}
