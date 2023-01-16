package grpc

import (
	"context"

	"github.com/dietzy1/chatapp/services/chatroom/domain"
)

type chatroom interface {
	CreateRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	DeleteRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	JoinRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	LeaveRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	GetRoom(ctx context.Context, chatroom domain.Chatroom) (domain.Chatroom, error)
}

//define methods on the struct that uses grpc to implement the interface

func (s *server) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error) {
	chatroom := domain.Chatroom{
		Name:      req.Name,
		OwnerUuid: req.OwnerUuid,
	}

	chatroomUuid, err := s.chatroom.CreateRoom(ctx, chatroom)
	if err != nil {
		return &CreateRoomResponse{
			Error: err.Error(),
		}, nil
	}

	return &CreateRoomResponse{
		ChatroomUuid: chatroomUuid,
	}, nil
}
