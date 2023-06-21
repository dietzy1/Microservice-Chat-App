package server

import (
	"bytes"
	"context"
	"io"
	"net"
	"os"

	"github.com/dietzy1/chatapp/pkg/middleware"
	"github.com/dietzy1/chatapp/services/icon/domain"
	iconv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type Domain interface {
	UploadIcon(ctx context.Context, imageData bytes.Buffer, icon domain.Icon) (domain.Icon, error)
	GetIcon(ctx context.Context, uuid string) (domain.Icon, error)
	GetIcons(ctx context.Context, ownerUuid string) ([]domain.Icon, error)
	GetEmojiIcons(ctx context.Context) ([]domain.Icon, error)
	DeleteIcon(ctx context.Context, uuid string) error
}

type server struct {
	iconv1.UnimplementedIconServiceServer
	logger *zap.Logger

	domain Domain
}

func New(logger *zap.Logger, domain Domain) iconv1.IconServiceServer {
	return &server{
		logger: logger,
		domain: domain,
	}
}

func Start(logger *zap.Logger, domain Domain) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	//addr := ":9000"
	addr := os.Getenv("ICONSERVICE")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingMiddleware(logger)),
	)
	//Inject dependencies into the server
	dependencies := New(logger, domain)

	//Register the server
	iconv1.RegisterIconServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))
}
