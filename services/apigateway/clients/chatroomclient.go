package client

import (
	"context"
	"log"
	"os"

	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/gateway/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatRoomClient struct {
	C messagev1.MessageGatewayServiceClient
}

func NewChatRoomClient() *MessageClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("CHATROOM"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := messagev1.NewMessageGatewayServiceClient(conn)

	return &MessageClient{C: client}
}
