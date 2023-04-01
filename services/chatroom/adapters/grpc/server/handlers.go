package server

import (
	"context"
	"log"

	"github.com/dietzy1/chatapp/services/chatroom/domain"
	"google.golang.org/grpc/status"

	chatroomv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
)

type chatroom interface {
	CreateRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	DeleteRoom(ctx context.Context, chatroom domain.Chatroom) error
	GetRoom(ctx context.Context, chatroom domain.Chatroom) (domain.Chatroom, error)

	CreateChannel(ctx context.Context, channel domain.Channel) (string, error)
	DeleteChannel(ctx context.Context, channel domain.Channel) error
	GetChannel(ctx context.Context, channel domain.Channel) (domain.Channel, error)

	InviteUser(ctx context.Context, chatroom domain.Chatroom, userUuid string) (string, error)
	RemoveUser(ctx context.Context, chatroom domain.Chatroom, userUuid string) error
	AddUser(ctx context.Context, chatroom domain.Chatroom, userUuid string) (string, error)
}

func (s *server) CreateRoom(ctx context.Context, req *chatroomv1.CreateRoomRequest) (*chatroomv1.CreateRoomResponse, error) {

	//Perform check if name and owner uuid is empty
	if req.Name == "" || req.OwnerUuid == "" {
		return &chatroomv1.CreateRoomResponse{}, status.Error(400, "Name and OwnerUuid cannot be empty")
	}

	chatroom := domain.Chatroom{
		Name:  req.Name,
		Owner: req.OwnerUuid,
	}

	_, err := s.chatroom.CreateRoom(ctx, chatroom)
	if err != nil {
		log.Println(err)
		return &chatroomv1.CreateRoomResponse{}, status.Error(500, "Internal Server Error")
	}
	return &chatroomv1.CreateRoomResponse{}, nil
}

func (s *server) DeleteRoom(ctx context.Context, req *chatroomv1.DeleteRoomRequest) (*chatroomv1.DeleteRoomResponse, error) {

	// Perform check if uuid or Owneruuid is empty
	if req.Uuid == "" || req.OwnerUuid == "" {
		return &chatroomv1.DeleteRoomResponse{}, status.Error(400, "Uuid and OwnerUuid cannot be empty")
	}

	chatroom := domain.Chatroom{
		Uuid:  req.Uuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.DeleteRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.DeleteRoomResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.DeleteRoomResponse{}, nil
}

func (s *server) GetRoom(ctx context.Context, req *chatroomv1.GetRoomRequest) (*chatroomv1.GetRoomResponse, error) {

	// Perform check if uuid is empty
	if req.Uuid == "" {
		return &chatroomv1.GetRoomResponse{}, status.Error(400, "Uuid cannot be empty")
	}

	chatroom := domain.Chatroom{
		Uuid: req.Uuid,
	}

	chatroom, err := s.chatroom.GetRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.GetRoomResponse{}, status.Error(500, "Internal Server Error")
	}

	//Convert chatroom.Users
	users := make([]string, len(chatroom.Users))
	users = append(users, chatroom.Users...)

	//Convert chatroom.channel_uuids to getroomresponse.channel_uuids
	channels := make([]string, len(chatroom.Channels))
	channels = append(channels, chatroom.Channels...)

	return &chatroomv1.GetRoomResponse{
		Name:         chatroom.Name,
		Uuid:         chatroom.Uuid,
		OwnerUuid:    chatroom.Owner,
		UserUuids:    users,
		ChannelUuids: channels,
	}, nil
}

//----------

func (s *server) CreateChannel(ctx context.Context, req *chatroomv1.CreateChannelRequest) (*chatroomv1.CreateChannelResponse, error) {

	//Perform check if name owneruuid and chatroom_uuid is empty
	if req.Name == "" || req.OwnerUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.CreateChannelResponse{}, status.Error(400, "Name, OwnerUuid and ChatroomUuid cannot be empty")
	}

	channel := domain.Channel{
		ChannelUuid:  req.OwnerUuid,
		Name:         req.Name,
		ChatroomUuid: req.ChatroomUuid,
	}

	_, err := s.chatroom.CreateChannel(ctx, channel)
	if err != nil {
		return &chatroomv1.CreateChannelResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.CreateChannelResponse{}, nil
}

func (s *server) DeleteChannel(ctx context.Context, req *chatroomv1.DeleteChannelRequest) (*chatroomv1.DeleteChannelResponse, error) {

	// Perform check if uuid or Owneruuid is empty
	if req.ChannelUuid == "" || req.OwnerUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.DeleteChannelResponse{}, status.Error(400, "Uuid and OwnerUuid cannot be empty")
	}

	channel := domain.Channel{
		ChannelUuid:  req.ChannelUuid,
		ChatroomUuid: req.ChatroomUuid,
		Owner:        req.OwnerUuid,
	}

	err := s.chatroom.DeleteChannel(ctx, channel)
	if err != nil {
		return &chatroomv1.DeleteChannelResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.DeleteChannelResponse{}, nil
}

func (s *server) GetChannel(ctx context.Context, req *chatroomv1.GetChannelRequest) (*chatroomv1.GetChannelResponse, error) {

	//Check if chatroom and channel uuid is empty
	if req.ChatroomUuid == "" || req.ChannelUuid == "" {
		return &chatroomv1.GetChannelResponse{}, status.Error(400, "ChatroomUuid and ChannelUuid cannot be empty")
	}

	channel := domain.Channel{
		ChannelUuid:  req.ChannelUuid,
		ChatroomUuid: req.ChatroomUuid,
	}

	channel, err := s.chatroom.GetChannel(ctx, channel)
	if err != nil {
		return &chatroomv1.GetChannelResponse{}, status.Error(500, "Internal Server Error")
	}

	//Convert channel to getchannelresponse
	return &chatroomv1.GetChannelResponse{
		Name:         channel.Name,
		ChatroomUuid: channel.ChatroomUuid,
		ChannelUuid:  channel.ChannelUuid,
	}, nil

}

//---------------------------------------------------------------------------------

func (s *server) InviteUser(ctx context.Context, req *chatroomv1.InviteUserRequest) (*chatroomv1.InviteUserResponse, error) {

	// Check if user_uuid, chatroomuuid and owner uuid is empty
	if req.UserUuid == "" || req.ChatroomUuid == "" || req.OwnerUuid == "" {
		return &chatroomv1.InviteUserResponse{}, status.Error(400, "UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
	}

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	_, err := s.chatroom.InviteUser(ctx, chatroom, req.UserUuid)
	if err != nil {
		return &chatroomv1.InviteUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.InviteUserResponse{}, nil
}

func (s *server) RemoveUser(ctx context.Context, req *chatroomv1.RemoveUserRequest) (*chatroomv1.RemoveUserResponse, error) {

	// Check if user_uuid, chatroomuuid and owner uuid is empty
	if req.UserUuid == "" || req.ChatroomUuid == "" || req.OwnerUuid == "" {
		return &chatroomv1.RemoveUserResponse{}, status.Error(400, "UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
	}

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.RemoveUser(ctx, chatroom, req.UserUuid)
	if err != nil {
		return &chatroomv1.RemoveUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.RemoveUserResponse{}, nil
}

func (s *server) AddUser(ctx context.Context, req *chatroomv1.AddUserRequest) (*chatroomv1.AddUserResponse, error) {

	// Check if user_uuid, chatroomuuid and owner uuid is empty
	if req.UserUuid == "" || req.ChatroomUuid == "" || req.OwnerUuid == "" {
		return &chatroomv1.AddUserResponse{}, status.Error(400, "UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
	}

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	_, err := s.chatroom.AddUser(ctx, chatroom, req.UserUuid)
	if err != nil {
		return &chatroomv1.AddUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.AddUserResponse{}, nil
}
