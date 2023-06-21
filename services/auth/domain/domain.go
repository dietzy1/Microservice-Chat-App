package domain

import (
	"context"

	"go.uber.org/zap"
)

type Auth interface {
	Login(ctx context.Context, username string) (string, error)
	Logout(ctx context.Context, userUuid string) error
	Authenticate(ctx context.Context, userUuid string) (string, error)
	UpdateToken(ctx context.Context, username string, token string) (string, error)
}

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type domain struct {
	logger *zap.Logger
	auth   Auth
	cache  Cache
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password " bson:"password"`
	Uuid     string `json:"uuid" bson:"uuid"`
	Session  string `json:"session" bson:"session"`
	Demo     bool   `json:"demo" bson:"demo"`
}

func New(logger *zap.Logger, auth Auth, cache Cache) *domain {
	return &domain{logger: logger, auth: auth, cache: cache}
}
