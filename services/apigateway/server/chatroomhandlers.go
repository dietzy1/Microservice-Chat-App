package server

import (
	"context"
	"fmt"

	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/chatroomgateway/v1"
	chatroomclientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
	"google.golang.org/grpc/status"
)

// FIXME: This function currently has a bug where it doesn't contact the user service and adds the user to the newly created chatroom
func (s *server) CreateRoom(ctx context.Context, req *chatroomv1.CreateRoomRequest) (*chatroomv1.CreateRoomResponse, error) {
	// Check if name is empty
	if req.Name == "" || req.OwnerUuid == "" {
		return &chatroomv1.CreateRoomResponse{}, status.Error(400, "Name cannot be empty")
	}

	//Reroute object
	name := &chatroomclientv1.CreateRoomRequest{
		Name:      req.Name,
		OwnerUuid: req.OwnerUuid,
	}

	_, err := s.chatroomClient.CreateRoom(ctx, name)
	if err != nil {
		return &chatroomv1.CreateRoomResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.CreateRoomResponse{}, nil
}

func (s *server) DeleteRoom(ctx context.Context, req *chatroomv1.DeleteRoomRequest) (*chatroomv1.DeleteRoomResponse, error) {
	// Check if name is empty
	if req.ChatroomUuid == "" || req.OwnerUuid == "" {
		return &chatroomv1.DeleteRoomResponse{}, status.Error(400, "RoomUuid cannot be empty")
	}

	//Reroute object
	roomUuid := &chatroomclientv1.DeleteRoomRequest{
		ChatroomUuid: req.ChatroomUuid,
		OwnerUuid:    req.OwnerUuid,
	}

	_, err := s.chatroomClient.DeleteRoom(ctx, roomUuid)
	if err != nil {
		return &chatroomv1.DeleteRoomResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.DeleteRoomResponse{}, nil
}

func (s *server) GetRoom(ctx context.Context, req *chatroomv1.GetRoomRequest) (*chatroomv1.GetRoomResponse, error) {
	fmt.Println("GetRoom")
	// Check if name is empty
	if req.ChatroomUuid == "" {
		return &chatroomv1.GetRoomResponse{}, status.Error(400, "RoomUuid cannot be empty")
	}

	//Reroute object
	roomUuid := &chatroomclientv1.GetRoomRequest{
		ChatroomUuid: req.ChatroomUuid,
	}

	room, err := s.chatroomClient.GetRoom(ctx, roomUuid)
	if err != nil {
		return &chatroomv1.GetRoomResponse{}, status.Error(500, "Internal Server Error")
	}

	resp := &chatroomv1.GetRoomResponse{
		ChatroomUuid: room.ChatroomUuid,
		Name:         room.Name,
		Icon:         room.Icon,
		OwnerUuid:    room.OwnerUuid,
	}
	resp.UserUuids = append(resp.UserUuids, room.UserUuids...)

	//Convert from []*chatroomv1.Channel to []*chatroomgatewayv1.Channel
	for _, channel := range room.Channels {
		resp.Channel = append(resp.Channel, &chatroomv1.Channel{
			ChannelUuid:  channel.ChannelUuid,
			Name:         channel.Name,
			OwnerUuid:    channel.OwnerUuid,
			ChatroomUuid: channel.ChatroomUuid,
		})
	}

	return resp, nil
}

func (s *server) GetRooms(ctx context.Context, req *chatroomv1.GetRoomsRequest) (*chatroomv1.GetRoomsResponse, error) {

	//Check if chatroomUuid is empty
	if req.ChatroomUuids == nil {
		return &chatroomv1.GetRoomsResponse{}, status.Error(400, "ChatroomUuid cannot be empty")
	}

	//Reroute object
	rooms, err := s.chatroomClient.GetRooms(ctx, &chatroomclientv1.GetRoomsRequest{
		ChatroomUuids: req.ChatroomUuids,
	})
	if err != nil {
		return &chatroomv1.GetRoomsResponse{}, status.Error(500, "Internal Server Error")
	}

	resp := &chatroomv1.GetRoomsResponse{}
	//Convert rooms to resp
	for i, room := range rooms.Rooms {

		resp.Rooms = append(resp.Rooms, &chatroomv1.GetRoomResponse{
			ChatroomUuid: room.ChatroomUuid,
			Name:         room.Name,
			Icon:         room.Icon,
			OwnerUuid:    room.OwnerUuid,
			UserUuids:    room.UserUuids,
		})

		//Convert from []*chatroomv1.Channel to []*chatroomgatewayv1.Channel
		for _, channel := range rooms.Rooms[i].Channels {
			resp.Rooms[i].Channel = append(resp.Rooms[i].Channel, &chatroomv1.Channel{
				ChannelUuid:  channel.ChannelUuid,
				Name:         channel.Name,
				OwnerUuid:    channel.OwnerUuid,
				ChatroomUuid: channel.ChatroomUuid,
			})
		}

	}

	return resp, nil
}

// ------------------------------------------------------------------------------------------------------------------------------------

func (s *server) CreateChannel(ctx context.Context, req *chatroomv1.CreateChannelRequest) (*chatroomv1.CreateChannelResponse, error) {
	// Check if name is empty
	if req.Name == "" || req.OwnerUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.CreateChannelResponse{}, status.Error(400, "Name cannot be empty")
	}

	//Reroute object
	name := &chatroomclientv1.CreateChannelRequest{
		Name:         req.Name,
		OwnerUuid:    req.OwnerUuid,
		ChatroomUuid: req.ChatroomUuid,
	}

	_, err := s.chatroomClient.CreateChannel(ctx, name)
	if err != nil {
		return &chatroomv1.CreateChannelResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.CreateChannelResponse{}, nil
}

func (s *server) DeleteChannel(ctx context.Context, req *chatroomv1.DeleteChannelRequest) (*chatroomv1.DeleteChannelResponse, error) {
	// Check if name is empty
	if req.ChannelUuid == "" || req.OwnerUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.DeleteChannelResponse{}, status.Error(400, "ChannelUuid cannot be empty")
	}

	//Reroute object
	channelUuid := &chatroomclientv1.DeleteChannelRequest{
		ChannelUuid:  req.ChannelUuid,
		ChatroomUuid: req.ChatroomUuid,
		OwnerUuid:    req.OwnerUuid,
	}

	_, err := s.chatroomClient.DeleteChannel(ctx, channelUuid)
	if err != nil {
		return &chatroomv1.DeleteChannelResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.DeleteChannelResponse{}, nil
}

func (s *server) GetChannel(ctx context.Context, req *chatroomv1.GetChannelRequest) (*chatroomv1.GetChannelResponse, error) {
	// Check if name is empty
	if req.ChannelUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.GetChannelResponse{}, status.Error(400, "ChannelUuid cannot be empty")
	}

	//Reroute object
	channelUuid := &chatroomclientv1.GetChannelRequest{
		ChannelUuid:  req.ChannelUuid,
		ChatroomUuid: req.ChatroomUuid,
	}

	channel, err := s.chatroomClient.GetChannel(ctx, channelUuid)
	if err != nil {
		return &chatroomv1.GetChannelResponse{}, status.Error(500, "Internal Server Error")
	}

	resp := &chatroomv1.GetChannelResponse{
		ChannelUuid:  channel.ChannelUuid,
		Name:         channel.Name,
		ChatroomUuid: channel.ChatroomUuid,
	}

	return resp, nil
}

// ------------------------------------------------------------------------------------------------------------------------------------

func (s *server) InviteUser(ctx context.Context, req *chatroomv1.InviteUserRequest) (*chatroomv1.InviteUserResponse, error) {
	// Check if name is empty
	if req.UserUuid == "" || req.OwnerUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.InviteUserResponse{}, status.Error(400, "UserUuid cannot be empty")
	}

	//Reroute object
	userUuid := &chatroomclientv1.InviteUserRequest{
		UserUuid:     req.UserUuid,
		ChatroomUuid: req.ChatroomUuid,
		OwnerUuid:    req.OwnerUuid,
	}

	_, err := s.chatroomClient.InviteUser(ctx, userUuid)
	if err != nil {
		return &chatroomv1.InviteUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.InviteUserResponse{}, nil
}

func (s *server) RemoveUser(ctx context.Context, req *chatroomv1.RemoveUserRequest) (*chatroomv1.RemoveUserResponse, error) {
	// Check if name is empty
	if req.UserUuid == "" || req.OwnerUuid == "" || req.ChatroomUuid == "" {
		return &chatroomv1.RemoveUserResponse{}, status.Error(400, "UserUuid cannot be empty")
	}

	//Reroute object
	userUuid := &chatroomclientv1.RemoveUserRequest{
		UserUuid:     req.UserUuid,
		ChatroomUuid: req.ChatroomUuid,
		OwnerUuid:    req.OwnerUuid,
	}

	_, err := s.chatroomClient.RemoveUser(ctx, userUuid)
	if err != nil {
		return &chatroomv1.RemoveUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.RemoveUserResponse{}, nil
}

func (s *server) AddUser(ctx context.Context, req *chatroomv1.AddUserRequest) (*chatroomv1.AddUserResponse, error) {
	// Check if name is empty
	if req.UserUuid == "" || req.ChatroomUuid == "" || req.OwnerUuid == "" {
		return &chatroomv1.AddUserResponse{}, status.Error(400, "UserUuid cannot be empty")
	}

	//Reroute object
	userUuid := &chatroomclientv1.AddUserRequest{
		UserUuid:     req.UserUuid,
		ChatroomUuid: req.ChatroomUuid,
		OwnerUuid:    req.OwnerUuid,
	}

	_, err := s.chatroomClient.AddUser(ctx, userUuid)
	if err != nil {
		return &chatroomv1.AddUserResponse{}, status.Error(500, "Internal Server Error")
	}

	return &chatroomv1.AddUserResponse{}, nil
}

// ------------------------------------------------------------------------------------------------------------------------------------
