package clients

import (
	"context"
	"log"
	"os"

	iconclientv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewIconClient() *iconclientv1.IconServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("ICONSERVICE"),
		//"localhost:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := iconclientv1.NewIconServiceClient(conn)
	return &client
}
