package client

import (
	"context"
	"log"
	"os"

	userv1 "github.com/dietzy1/chatapp/services/apigateway/usergateway/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	C userv1.UserGatewayServiceClient
}

func NewUserClient() *UserClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("USER"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := userv1.NewUserGatewayServiceClient(conn)

	return &UserClient{C: client}
}
