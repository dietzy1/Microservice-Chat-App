package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/pkg/clients"
	"github.com/dietzy1/chatapp/pkg/logger"

	"github.com/dietzy1/chatapp/services/user/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/user/adapters/repository"

	"github.com/dietzy1/chatapp/services/user/domain"
)

func main() {

	config.ReadEnvfile()

	logger := logger.New()

	repo, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	iconClient := clients.NewIconClient()

	domain := domain.New(logger, repo, *iconClient)

	server.Start(logger, domain)

}
