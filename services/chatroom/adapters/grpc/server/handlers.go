package server

import (
	"context"

	"github.com/dietzy1/chatapp/services/chatroom/domain"
	"google.golang.org/grpc/status"

	chatroomv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
)

type chatroom interface {
	CreateRoom(ctx context.Context, chatroom domain.Chatroom) (string, error)
	DeleteRoom(ctx context.Context, chatroom domain.Chatroom) error
	GetRoom(ctx context.Context, chatroom domain.Chatroom) (domain.Chatroom, error)
	GetRooms(ctx context.Context, chatroomUuids []string) ([]domain.Chatroom, error)

	CreateChannel(ctx context.Context, channel domain.Channel) (string, error)
	DeleteChannel(ctx context.Context, channel domain.Channel) error
	GetChannel(ctx context.Context, channel domain.Channel) (domain.Channel, error)

	InviteUser(ctx context.Context, chatroom domain.Chatroom, userUuid string) error
	RemoveUser(ctx context.Context, chatroom domain.Chatroom, userUuid string) error
	AddUser(ctx context.Context, chatroom domain.Chatroom, userUuid string) error

	ChangeIcon(ctx context.Context, icon domain.Chatroom) error

	ForceAddUser(ctx context.Context, userUuid string) error
}

func (s *server) CreateRoom(ctx context.Context, req *chatroomv1.CreateRoomRequest) (*chatroomv1.CreateRoomResponse, error) {

	chatroom := domain.Chatroom{
		Name:  req.Name,
		Owner: req.OwnerUuid,
	}

	_, err := s.chatroom.CreateRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.CreateRoomResponse{}, status.Errorf(500, "Internal Server Error: %v", err)
	}
	return &chatroomv1.CreateRoomResponse{}, nil
}

func (s *server) DeleteRoom(ctx context.Context, req *chatroomv1.DeleteRoomRequest) (*chatroomv1.DeleteRoomResponse, error) {

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.DeleteRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.DeleteRoomResponse{}, status.Errorf(500, "Internal Server Error: %v", err)
	}

	return &chatroomv1.DeleteRoomResponse{}, nil
}

func (s *server) GetRoom(ctx context.Context, req *chatroomv1.GetRoomRequest) (*chatroomv1.GetRoomResponse, error) {

	chatroom := domain.Chatroom{
		Uuid: req.ChatroomUuid,
	}

	chatroom, err := s.chatroom.GetRoom(ctx, chatroom)
	if err != nil {
		return &chatroomv1.GetRoomResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	//Convert chatroom.Users
	users := make([]string, len(chatroom.Users))
	users = append(users, chatroom.Users...)

	//convert channel to []*chatroomv1.Channel
	channels := make([]*chatroomv1.Channel, len(chatroom.Channels))
	for i, channel := range chatroom.Channels {
		channels[i] = &chatroomv1.Channel{
			ChannelUuid:  channel.ChannelUuid,
			Name:         channel.Name,
			ChatroomUuid: channel.ChatroomUuid,
			OwnerUuid:    channel.Owner,
		}
	}

	return &chatroomv1.GetRoomResponse{
		Name:         chatroom.Name,
		ChatroomUuid: chatroom.Uuid,
		OwnerUuid:    chatroom.Owner,
		UserUuids:    users,
		Channels:     channels,
	}, nil
}

func (s *server) GetRooms(ctx context.Context, req *chatroomv1.GetRoomsRequest) (*chatroomv1.GetRoomsResponse, error) {

	chatrooms, err := s.chatroom.GetRooms(ctx, req.ChatroomUuids)
	if err != nil {
		return &chatroomv1.GetRoomsResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	//I have an array of chatrooms, I need to convert it to an array of getroomresponse
	rooms := make([]*chatroomv1.GetRoomResponse, len(chatrooms))
	for i, chatroom := range chatrooms {
		//Convert chatroom.Users
		users := make([]string, len(chatroom.Users))
		users = append(users, chatroom.Users...)

		//Convert chatroom.channel_uuids to getroomresponse.channel_uuids
		/* channels := make([]string, len(chatroom.Channels))
		channels = append(channels, chatroom.Channels...) */

		channels := make([]*chatroomv1.Channel, len(chatroom.Channels))
		for i, channel := range chatroom.Channels {
			channels[i] = &chatroomv1.Channel{
				ChannelUuid:  channel.ChannelUuid,
				Name:         channel.Name,
				ChatroomUuid: channel.ChatroomUuid,
				OwnerUuid:    channel.Owner,
			}
		}

		rooms[i] = &chatroomv1.GetRoomResponse{
			Name:         chatroom.Name,
			Icon:         chatroom.Icon.Link,
			ChatroomUuid: chatroom.Uuid,
			OwnerUuid:    chatroom.Owner,
			UserUuids:    users,
			Channels:     channels,
		}
	}

	return &chatroomv1.GetRoomsResponse{
		Rooms: rooms,
	}, nil
}

//----------

func (s *server) CreateChannel(ctx context.Context, req *chatroomv1.CreateChannelRequest) (*chatroomv1.CreateChannelResponse, error) {

	channel := domain.Channel{
		Name:         req.Name,
		ChatroomUuid: req.ChatroomUuid,
		Owner:        req.OwnerUuid,
	}

	_, err := s.chatroom.CreateChannel(ctx, channel)
	if err != nil {
		return &chatroomv1.CreateChannelResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.CreateChannelResponse{}, nil
}

func (s *server) DeleteChannel(ctx context.Context, req *chatroomv1.DeleteChannelRequest) (*chatroomv1.DeleteChannelResponse, error) {

	channel := domain.Channel{
		ChannelUuid:  req.ChannelUuid,
		ChatroomUuid: req.ChatroomUuid,
		Owner:        req.OwnerUuid,
	}

	err := s.chatroom.DeleteChannel(ctx, channel)
	if err != nil {
		return &chatroomv1.DeleteChannelResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.DeleteChannelResponse{}, nil
}

func (s *server) GetChannel(ctx context.Context, req *chatroomv1.GetChannelRequest) (*chatroomv1.GetChannelResponse, error) {

	channel := domain.Channel{
		ChannelUuid:  req.ChannelUuid,
		ChatroomUuid: req.ChatroomUuid,
	}

	channel, err := s.chatroom.GetChannel(ctx, channel)
	if err != nil {
		return &chatroomv1.GetChannelResponse{}, status.Errorf(500, "Internal Server Error %v", err)
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

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.InviteUser(ctx, chatroom, req.UserUuid)
	if err != nil {
		return &chatroomv1.InviteUserResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.InviteUserResponse{}, nil
}

func (s *server) RemoveUser(ctx context.Context, req *chatroomv1.RemoveUserRequest) (*chatroomv1.RemoveUserResponse, error) {

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.RemoveUser(ctx, chatroom, req.UserUuid)
	if err != nil {
		return &chatroomv1.RemoveUserResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.RemoveUserResponse{}, nil
}

func (s *server) AddUser(ctx context.Context, req *chatroomv1.AddUserRequest) (*chatroomv1.AddUserResponse, error) {

	chatroom := domain.Chatroom{
		Uuid:  req.ChatroomUuid,
		Owner: req.OwnerUuid,
	}

	err := s.chatroom.AddUser(ctx, chatroom, req.UserUuid)
	if err != nil {
		return &chatroomv1.AddUserResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.AddUserResponse{}, nil
}

//---------------------------------------------------------------------------------

// internal use only
func (s *server) ForceAddUser(ctx context.Context, req *chatroomv1.ForceAddUserRequest) (*chatroomv1.ForceAddUserResponse, error) {

	err := s.chatroom.ForceAddUser(ctx, req.UserUuid)
	if err != nil {
		return &chatroomv1.ForceAddUserResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.ForceAddUserResponse{}, nil
}

func (s *server) ChangeIcon(ctx context.Context, req *chatroomv1.ChangeIconRequest) (*chatroomv1.ChangeIconResponse, error) {

	chatroom := domain.Chatroom{
		Uuid: req.ChatroomUuid,
		Icon: domain.Icon{
			Link: req.Icon.Link,
			Uuid: req.Icon.Uuid,
		},
	}

	err := s.chatroom.ChangeIcon(ctx, chatroom)
	if err != nil {
		return &chatroomv1.ChangeIconResponse{}, status.Errorf(500, "Internal Server Error %v", err)
	}

	return &chatroomv1.ChangeIconResponse{}, nil
}
