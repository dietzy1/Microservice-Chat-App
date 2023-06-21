package domain

import (
	"context"

	"go.uber.org/zap"

	iconClientv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
)

type Domain struct {
	logger     *zap.Logger
	user       user
	iconClient iconClientv1.IconServiceClient
}

// New creates a new application
func New(logger *zap.Logger, user user, iconClient iconClientv1.IconServiceClient) *Domain {
	return &Domain{logger: logger, user: user, iconClient: iconClient}
}

type User struct {
	Name        string   `bson:"name"`
	Uuid        string   `bson:"uuid"`
	Icon        Icon     `bson:"icon"`
	Description string   `bson:"description"`
	JoinDate    string   `bson:"joindate"`
	ChatServers []string `bson:"chatservers"`
}

type Icon struct {
	Link string `json:"link" bson:"link"`
	Uuid string `json:"uuid" bson:"uuid"`
}

type user interface {
	AddUser(ctx context.Context, user User) error
	AddChatServer(ctx context.Context, uuid string, serveruuid string) error
	RemoveChatServer(ctx context.Context, uuid string, serveruuid string) error
	EditDescription(ctx context.Context, uuid string, description string) error
	GetUser(ctx context.Context, uuid string) (User, error)
	GetUsers(ctx context.Context, uuids []string) ([]User, error)
	ChangeAvatar(ctx context.Context, userUuid string, icon Icon) error
}
