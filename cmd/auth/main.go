package main

import (
	"github.com/dietzy1/chatapp/config"
	client "github.com/dietzy1/chatapp/services/apigateway/clients"
	"github.com/dietzy1/chatapp/services/auth/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/auth/adapters/repository"
	"github.com/dietzy1/chatapp/services/auth/domain"
)

func main() {

	config.ReadEnvfile()

	repo := repository.NewRepository()

	userClient := client.NewUserClient()

	domain := domain.New(repo.Auth, repo.Caching, *userClient)

	//I prolly need to inject the domain into the GRPC server here

	//start the GRPC server
	server.Start(domain)
}
