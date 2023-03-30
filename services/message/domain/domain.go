package domain

import (
	"context"

	"github.com/dietzy1/chatapp/services/message/core"
)

type Domain struct {
}

func New() *Domain {
	return &Domain{}
}

// add the interface to the domain struct
type MsgRepo interface {
	GetMsg(ctx context.Context, room core.ChatRoom) (core.Message, error)
	AddMsg(ctx context.Context, room core.ChatRoom, msg core.Message) error
	UpdateMsg(ctx context.Context, room core.ChatRoom, msg core.Message) error
	DeleteMsg(ctx context.Context, room core.ChatRoom, id string) error
}

//Implement the methods for the domain layer and create the interface and inject it into the grpc layer

type Message struct {
	Author       string
	Content      string
	AuthorUuid   string
	ChatRoomUuid string
	ChannelUuid  string
	MessageUuid  string
	Timestamp    string
}

type MessageRequest struct {
	Author         string
	Content        string
	Author_uuid    string
	Chat_room_uuid string
	Channel_uuid   string
}
