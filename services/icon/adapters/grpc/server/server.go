package server

import (
	"bytes"
	"context"

	"github.com/dietzy1/chatapp/services/icon/domain"
	iconv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
	"go.uber.org/zap"
)

type Domain interface {
	UploadIcon(ctx context.Context, imageData bytes.Buffer) (domain.Icon, error)
	GetIcons(ctx context.Context) ([]domain.Icon, error)
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
