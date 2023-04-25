package clients

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authclientv1 "github.com/dietzy1/chatapp/services/auth/proto/auth/v1"
)

// returns a pointer to a client that can be used to call the auth service
func NewAuthClient() *authclientv1.AuthServiceClient {
	log.Println("Connecting to auth service", os.Getenv("AUTH"))
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("AUTH"),
		//"localhost:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := authclientv1.NewAuthServiceClient(conn)
	return &client
}
