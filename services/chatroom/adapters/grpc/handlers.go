package grpc

import (
	"context"

	"github.com/dietzy1/chatapp/services/chatroom/domain"

	chatroomv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
)

type chatroom interface {
	CreateRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	DeleteRoom(ctx context.Context, chatroom domain.Chatroom) error
	JoinRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	LeaveRoom(ctx context.Context, chatroomUuid string, userUuid string) error
	GetRoom(ctx context.Context, chatroom domain.Chatroom) (domain.Chatroom, error)
}

func (s *server) CreateRoom(ctx context.Context, req *chatroomv1.CreateRoomRequest) (*chatroomv1.CreateRoomResponse, error) {
	chatroom := domain.Chatroom{
		Name:        req.Name,
		Owner:       req.OwnerUuid,
		Description: req.Description,
		Tags:        req.Tags,
	}

	chatroomUuid, err := s.chatroom.CreateRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.CreateRoomResponse{
			Error: err.Error(),
		}, nil
	}

	return &chatroomv1.CreateRoomResponse{
		Error:        "",
		ChatroomUuid: chatroomUuid,
	}, nil
}

func (s *server) DeleteRoom(ctx context.Context, req *chatroomv1.DeleteRoomRequest) (*chatroomv1.DeleteRoomResponse, error) {
	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.DeleteRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.DeleteRoomResponse{
			Error: err.Error(),
		}, nil
	}

	return &chatroomv1.DeleteRoomResponse{
		Error: "",
	}, nil
}

func (s *server) JoinRoom(ctx context.Context, req *chatroomv1.JoinRoomRequest) (*chatroomv1.JoinRoomResponse, error) {
	chatroom := domain.Chatroom{
		Uuid: req.ChatroomUuid,
	}

	chatroomUuid, err := s.chatroom.JoinRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.JoinRoomResponse{
			Error: err.Error(),
		}, nil
	}

	return &chatroomv1.JoinRoomResponse{
		Error:        "",
		ChatroomUuid: chatroomUuid,
	}, nil
}

func (s *server) LeaveRoom(ctx context.Context, req *chatroomv1.LeaveRoomRequest) (*chatroomv1.LeaveRoomResponse, error) {
	err := s.chatroom.LeaveRoom(ctx, req.ChatroomUuid, req.UserUuid)
	if err != nil {
		return &chatroomv1.LeaveRoomResponse{
			Error: err.Error(),
		}, nil
	}

	return &chatroomv1.LeaveRoomResponse{
		Error: "",
	}, nil
}

func (s *server) GetRoom(ctx context.Context, req *chatroomv1.GetRoomRequest) (*chatroomv1.GetRoomResponse, error) {
	chatroom := domain.Chatroom{
		Uuid: req.ChatroomUuid,
	}

	chatroom, err := s.chatroom.GetRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.GetRoomResponse{
			Error: err.Error(),
		}, nil
	}

	return &chatroomv1.GetRoomResponse{
		Error:       "",
		Name:        chatroom.Name,
		OwnerUuid:   chatroom.Owner,
		Users:       chatroom.Users,
		Description: chatroom.Description,
		Tags:        chatroom.Tags,
	}, nil
}
