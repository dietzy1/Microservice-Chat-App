package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/pkg/logger"
	"github.com/dietzy1/chatapp/services/icon/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/icon/adapters/repository"
	"github.com/dietzy1/chatapp/services/icon/adapters/rest/cdn"
	"github.com/dietzy1/chatapp/services/icon/domain"
	"go.uber.org/zap"
)

func main() {

	config.ReadEnvfile()

	logger := logger.New()

	repo, err := repository.New()
	if err != nil {
		logger.Fatal("failed to create repository", zap.Error(err))
	}

	cdn := cdn.New()

	domain := domain.New(logger, repo, cdn)

	server.Start(logger, domain)

}
