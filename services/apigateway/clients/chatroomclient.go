package client

import (
	"context"
	"log"
	"os"

	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/chatroomgateway/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatRoomClient struct {
	C chatroomv1.ChatroomGatewayServiceClient
}

func NewChatRoomClient() *ChatRoomClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("CHATROOM"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := chatroomv1.NewChatroomGatewayServiceClient(conn)

	return &ChatRoomClient{C: client}
}
