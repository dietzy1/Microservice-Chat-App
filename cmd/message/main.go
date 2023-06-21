package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/pkg/logger"
	"github.com/dietzy1/chatapp/services/message/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/message/adapters/repository"
	"github.com/dietzy1/chatapp/services/message/domain"
)

func main() {

	config.ReadEnvfile()

	logger := logger.New()

	repo, err := repository.New()
	if err != nil {
		log.Fatalln(err)
	}

	domain := domain.New(logger, repo)

	server.Start(logger, domain)
}
