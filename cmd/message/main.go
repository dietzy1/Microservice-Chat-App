package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/apigateway/server"
	"github.com/dietzy1/chatapp/services/auth/domain"
	"github.com/dietzy1/chatapp/services/user/adapters/repository"
)

func main() {

	config.ReadEnvfile()

	repo, err := repository.New()
	if err != nil {
		log.Fatalln(err)
	}

	domain := domain.New(repo)

	server.Start(domain)
}
