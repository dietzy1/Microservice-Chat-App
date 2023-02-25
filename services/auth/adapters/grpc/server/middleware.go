package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func loggingMiddleware(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Printf("gRPC method: %s", info.FullMethod)
	resp, err = handler(ctx, req)
	return
}
