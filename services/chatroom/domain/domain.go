package domain

import (
	"context"

	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"go.uber.org/zap"
)

//TODO: I need to implement the icon change service for the server stuff here aswell

type Domain struct {
	logger     *zap.Logger
	repo       chatroom
	userClient userClientv1.UserServiceClient
}

func New(logger *zap.Logger, repo chatroom, userClient userClientv1.UserServiceClient) *Domain {
	return &Domain{repo: repo, userClient: userClient}
}

// this is the repository interface
type chatroom interface {
	CreateChatroom(ctx context.Context, chatroom Chatroom) error
	DeleteChatroom(ctx context.Context, chatroomUuid string) error
	GetChatroom(ctx context.Context, chatroomUuid string) (Chatroom, error)
	GetChatrooms(ctx context.Context, chatroomUuids []string) ([]Chatroom, error)

	CreateChannel(ctx context.Context, channel Channel) error
	DeleteChannel(ctx context.Context, channelUuid string) error
	GetChannel(ctx context.Context, channelUuid string) (Channel, error)

	InviteUser(ctx context.Context, chatroomUuid string, userUuid string) error
	RemoveUser(ctx context.Context, chatroomUuid string, userUuid string) error
	AddUser(ctx context.Context, chatroomUuid string, userUuid string) error

	//internal use
	ForceAddUser(ctx context.Context, chatroomUuid string, userUuid string) error
	ChangeIcon(ctx context.Context, chatroomUuid string, chatroomIcon Icon) error
}

type Chatroom struct {
	Name     string    `json:"name" bson:"name"`
	Uuid     string    `json:"uuid" bson:"uuid"`
	Icon     Icon      `json:"icon" bson:"icon"`
	Owner    string    `json:"owner" bson:"owner"`
	Users    []string  `json:"users" bson:"users"`
	Channels []Channel `json:"channels" bson:"channels"`
	Invited  []string  `json:"invited" bson:"invited"`
}

type Channel struct {
	ChannelUuid  string `json:"channeluuid" bson:"channeluuid"`
	Name         string `json:"name" bson:"name"`
	ChatroomUuid string `json:"chatroomuuid" bson:"chatroomuuid"`
	Owner        string `json:"owner" bson:"owner"`
}

type Icon struct {
	Link string `json:"link" bson:"link"`
	Uuid string `json:"uuid" bson:"uuid"`
}
