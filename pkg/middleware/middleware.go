package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LoggingMiddleware(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Printf("gRPC method: %s", info.FullMethod)
	resp, err = handler(ctx, req)
	return
}
