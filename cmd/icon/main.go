package main

import (
	"github.com/dietzy1/chatapp/pkg/logger"
	"github.com/dietzy1/chatapp/services/auth/domain"
	"github.com/dietzy1/chatapp/services/icon/adapters/grpc/server"
	"github.com/dietzy1/chatapp/services/icon/adapters/repository"
	"github.com/dietzy1/chatapp/services/icon/adapters/rest/cdn"
)

func main() {

	log := logger.New()

	repo := repository.New()

	cdn := cdn.New()

	domain.New(log, repo, cdn)

	server.Start(log, domain)

}
