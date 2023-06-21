package clients

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	messageclientv1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
)

// returns a pointer to a client that can be used to call the message service
func NewMessageClient() *messageclientv1.MessageServiceClient {
	log.Println("Connecting to message service", os.Getenv("MESSAGE"))
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("MESSAGESERVICE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := messageclientv1.NewMessageServiceClient(conn)

	return &client
}
