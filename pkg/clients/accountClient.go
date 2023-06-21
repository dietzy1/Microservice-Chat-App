package clients

import (
	"context"
	"log"
	"os"

	accountclientv1 "github.com/dietzy1/chatapp/services/account/proto/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// returns a pointer to a client that can be used to call the account service
func NewAccountClient() *accountclientv1.AccountServiceClient {
	log.Println("Connecting to account service", os.Getenv("ACCOUNT"))
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("ACCOUNTSERVICE"),
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
