package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/pkg/clients"
	"github.com/dietzy1/chatapp/pkg/logger"
	"github.com/dietzy1/chatapp/services/chatroom/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/chatroom/adapters/repository"
	"github.com/dietzy1/chatapp/services/chatroom/domain"
)

func main() {

	config.ReadEnvfile()

	logger := logger.New()

	repo := repository.New()

	userClient := clients.NewUserClient()

	domain := domain.New(logger, repo, *userClient)

	server.Start(logger, domain)
}
