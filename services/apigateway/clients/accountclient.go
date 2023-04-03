package client

import (
	"context"
	"log"
	"os"

	accountclientv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccountClient() *accountclientv1.AccountServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("ACCOUNT"),
		//"localhost:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := accountclientv1.NewAccountServiceClient(conn)
	return &client
}
