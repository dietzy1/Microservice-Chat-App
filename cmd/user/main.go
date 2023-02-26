package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"

	"github.com/dietzy1/chatapp/services/user/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/user/adapters/repository"
	"github.com/dietzy1/chatapp/services/user/adapters/rest/cdn"

	"github.com/dietzy1/chatapp/services/user/domain"
)

func main() {

	config.ReadEnvfile()

	repo, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	redis, err := repository.NewRedis()
	if err != nil {
		log.Fatal(err)
	}

	_ = redis

	cdn := cdn.New()

	domain := domain.New(repo, repo, cdn)

	server.Start(domain)

}
