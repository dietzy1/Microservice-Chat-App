package client

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authv1 "github.com/dietzy1/chatapp/services/apigateway/authgateway/v1"
)

//I need to generate a client somewhere around here
//And use the client in the handlers file to call the authentication service

type AuthClient struct {
	C authv1.AuthGatewayServiceClient
}

func NewAuthClient() *AuthClient {
	conn, err := grpc.DialContext(
		context.Background(),
		/* "dns:///0.0.0.0:8080", */
		"dns:///0.0.0.0"+os.Getenv("AUTH"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := authv1.NewAuthGatewayServiceClient(conn)

	return &AuthClient{C: client}
}

//The issue might be that its trying to create a connection before the apigateway has been set up correctly
