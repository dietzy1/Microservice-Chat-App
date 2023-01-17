package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/apigateway/server"
	"github.com/dietzy1/chatapp/services/auth/domain"
	"github.com/dietzy1/chatapp/services/user/adapters/repository"
)

func main() {

	config.ReadEnvfile()

	repo := repository.New()

	domain := domain.New(repo)

	server.Start(domain)
}
