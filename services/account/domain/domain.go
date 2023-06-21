package domain

import (
	"context"

	"go.uber.org/zap"

	chatroomClientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

type domain struct {
	logger *zap.Logger
	repo   Repo

	userClient     userClientv1.UserServiceClient
	chatroomClient chatroomClientv1.ChatroomServiceClient
}

type Repo interface {
	CheckUsername(ctx context.Context, username string) error
	ChangeUsername(ctx context.Context, userUuid string, username string) error
	CheckPassword(ctx context.Context, userUuid string) (string, error)
	UpdatePassword(ctx context.Context, userUuid string, password string) error
	DeleteAccount(ctx context.Context, userUuid string) error
	Register(ctx context.Context, cred Credentials) (string, error)
	Unregister(ctx context.Context, userUuid string) error
	UpdateAccount(ctx context.Context, cred Credentials) error
}

func New(logger *zap.Logger, repo Repo, userClient userClientv1.UserServiceClient, chatroomClient chatroomClientv1.ChatroomServiceClient) *domain {
	return &domain{logger: logger, repo: repo, userClient: userClient, chatroomClient: chatroomClient}
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password " bson:"password"`
	Uuid     string `json:"uuid" bson:"uuid"`
	Session  string `json:"session" bson:"session"`
	DemoUser bool   `json:"demouser" bson:"demouser"`
}
