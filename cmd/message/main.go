package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/message/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/message/adapters/repository"
	"github.com/dietzy1/chatapp/services/message/domain"
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
