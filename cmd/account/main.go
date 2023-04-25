package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/pkg/clients"
	"github.com/dietzy1/chatapp/services/account/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/account/adapters/repository"
	"github.com/dietzy1/chatapp/services/account/domain"
)

func main() {

	config.ReadEnvfile()

	repo, err := repository.New()
	if err != nil {
		panic(err)
	}

	userClient := clients.NewUserClient()

	chatroomClient := clients.NewChatRoomClient()

	domain := domain.New(repo, *userClient, *chatroomClient)

	server.Start(domain)
}
