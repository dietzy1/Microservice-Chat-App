package main

import (
	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/auth/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/auth/adapters/repository"
	"github.com/dietzy1/chatapp/services/auth/domain"
)

func main() {

	config.ReadEnvfile()

	repo := repository.NewRepository()

	domain := domain.New(repo.Auth, repo.Caching)

	_ = domain

	//I prolly need to inject the domain into the GRPC server here

	//start the GRPC server
	server.Start()
}
