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
