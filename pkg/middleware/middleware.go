package middleware

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingMiddleware(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info("gRPC method", zap.String("method", info.FullMethod))
		resp, err := handler(ctx, req)
		return resp, err
	}
}
