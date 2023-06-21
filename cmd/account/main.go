package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/pkg/clients"
	"github.com/dietzy1/chatapp/pkg/logger"
	"github.com/dietzy1/chatapp/services/account/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/account/adapters/repository"
	"github.com/dietzy1/chatapp/services/account/domain"
)

func main() {

	config.ReadEnvfile()

	logger := logger.New()

	repo, err := repository.New()
	if err != nil {
		panic(err)
	}

	userClient := clients.NewUserClient()

	chatroomClient := clients.NewChatRoomClient()

	domain := domain.New(logger, repo, *userClient, *chatroomClient)

	server.Start(logger, domain)
}
