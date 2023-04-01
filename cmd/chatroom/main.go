package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/chatroom/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/chatroom/adapters/repository"
	"github.com/dietzy1/chatapp/services/chatroom/domain"
)

func main() {

	config.ReadEnvfile()

	//cdn := cdn.New()

	repo := repository.New()

	domain := domain.New(repo)

	server.Start(domain)
}
