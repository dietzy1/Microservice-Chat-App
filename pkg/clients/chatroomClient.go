package clients

import (
	"context"
	"log"
	"os"

	chatroomclientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// returns a pointer to a client that can be used to call the chatroom service
func NewChatRoomClient() *chatroomclientv1.ChatroomServiceClient {
	log.Println("Connecting to chatroom service", os.Getenv("CHATROOM"))
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("CHATROOMSERVICE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	//client := chatroomclientv1.ChatroomServiceClient(conn)
	client := chatroomclientv1.NewChatroomServiceClient(conn)

	return &client
}
