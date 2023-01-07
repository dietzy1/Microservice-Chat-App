package main

import (
	"log"

	"github.com/dietzy1/chatapp/config"

	"github.com/dietzy1/chatapp/services/user/adapters/repository"
	"github.com/dietzy1/chatapp/services/user/adapters/server/rest/cdn"
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

	cdn := cdn.New()

	app := domain.New(repo, repo, cdn, redis)

	//Tempory fix to avoid errors
	_ = app

	//Should open the GRPC server here and pass the domain into it
}
