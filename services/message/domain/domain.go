package domain

import (
	"context"

	"go.uber.org/zap"
)

type Domain struct {
	logger *zap.Logger
	repo   MsgRepo
}

func New(logger *zap.Logger, repo MsgRepo) *Domain {
	return &Domain{logger: logger, repo: repo}
}

// add the interface to the domain struct
type MsgRepo interface {
	GetMessages(ctx context.Context, chatroomUuid string, channelUuid string) ([]Message, error)
	AddMessage(ctx context.Context, msg Message) error
	UpdateMessage(ctx context.Context, msg Message) error
	DeleteMessage(ctx context.Context, msg Message) error
}

type Message struct {
	Author       string `json:"author" bson:"author"`
	Content      string `json:"content" bson:"content"`
	AuthorUuid   string `json:"authoruuid" bson:"authoruuid"`
	ChatRoomUuid string `json:"chatroomuuid" bson:"chatroomuuid"`
	ChannelUuid  string `json:"channeluuid" bson:"channeluuid"`
	MessageUuid  string `json:"messageuuid" bson:"messageuuid"`
	Timestamp    string `json:"timestamp" bson:"timestamp"`
}
