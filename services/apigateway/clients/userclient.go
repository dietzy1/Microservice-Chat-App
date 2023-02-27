package client

import (
	"context"
	"log"
	"os"

	userclientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient() *userclientv1.UserServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("USERSERVICE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := userclientv1.NewUserServiceClient(conn)

	return &client
}
