package server

import (
	"context"
	"log"

	"github.com/dietzy1/chatapp/services/message/domain"
	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type message interface {
	CreateMessage(ctx context.Context, msg domain.Message) (domain.Message, error)
	GetMessages(ctx context.Context, chatroomUuid string, channelUuid string) ([]domain.Message, error)
	EditMessage(ctx context.Context, msg domain.Message) (domain.Message, error)
	DeleteMessage(ctx context.Context, msg domain.Message) error
}

func (s *server) CreateMessage(ctx context.Context, req *messagev1.CreateMessageRequest) (*messagev1.CreateMessageResponse, error) {
	//Check if message is empty
	if req.Author == "" || req.Content == "" || req.AuthorUuid == "" || req.ChatRoomUuid == "" || req.ChannelUuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Message is empty")
	}

	//Create the message request
	msgReq := domain.Message{
		Author:       req.Author,
		Content:      req.Content,
		AuthorUuid:   req.AuthorUuid,
		ChatRoomUuid: req.ChatRoomUuid,
		ChannelUuid:  req.ChannelUuid,
	}

	//Create the message
	msg, err := s.domain.CreateMessage(ctx, msgReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	//Create the response
	msgRes := &messagev1.CreateMessageResponse{
		Author:       msg.Author,
		Content:      msg.Content,
		AuthorUuid:   msg.AuthorUuid,
		ChatRoomUuid: msg.ChatRoomUuid,
		ChannelUuid:  msg.ChannelUuid,
		MessageUuid:  msg.MessageUuid,
		Timestamp:    msg.Timestamp,
	}

	return msgRes, nil

}

func (s *server) GetMessage(ctx context.Context, req *messagev1.GetMessagesRequest) (*messagev1.GetMessagesResponse, error) {
	//Check if message is empty
	if req.ChatRoomUuid == "" || req.ChannelUuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Message is empty")
	}

	//Get the message
	msg, err := s.domain.GetMessages(ctx, req.ChatRoomUuid, req.ChannelUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	//Perform check if index 0 of the slice is empty
	if msg[0].Author == "" || msg[0].Content == "" || msg[0].AuthorUuid == "" || msg[0].ChatRoomUuid == "" || msg[0].ChannelUuid == "" || msg[0].MessageUuid == "" || msg[0].Timestamp == "" {
		return nil, status.Error(codes.NotFound, "Message not found")
	}

	//Create the response
	msgRes := &messagev1.GetMessagesResponse{
		Messages: []*messagev1.Msg{
			{
				Author:       msg[0].Author,
				Content:      msg[0].Content,
				AuthorUuid:   msg[0].AuthorUuid,
				ChatRoomUuid: msg[0].ChatRoomUuid,
				ChannelUuid:  msg[0].ChannelUuid,
				MessageUuid:  msg[0].MessageUuid,
				Timestamp:    msg[0].Timestamp,
			},
		},
	}
	return msgRes, nil
}

func (s *server) GetMessages(ctx context.Context, req *messagev1.GetMessagesRequest) (*messagev1.GetMessagesResponse, error) {
	log.Println("GetMessages called")
	//Check if message is empty
	if req.ChatRoomUuid == "" || req.ChannelUuid == "" {
		log.Println("Message is empty")
		return nil, status.Error(codes.InvalidArgument, "Message is empty")
	}

	//Get the message
	msg, err := s.domain.GetMessages(ctx, req.ChatRoomUuid, req.ChannelUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	//Create the response
	resp := &messagev1.GetMessagesResponse{}

	//Convert msg to []*Msg
	for _, m := range msg {
		resp.Messages = append(resp.Messages, &messagev1.Msg{
			Author:       m.Author,
			Content:      m.Content,
			AuthorUuid:   m.AuthorUuid,
			ChatRoomUuid: m.ChatRoomUuid,
			ChannelUuid:  m.ChannelUuid,
			MessageUuid:  m.MessageUuid,
			Timestamp:    m.Timestamp,
		})
	}

	return resp, nil
}

func (s *server) EditMessage(ctx context.Context, req *messagev1.EditMessageRequest) (*messagev1.EditMessageResponse, error) {
	//Check if message is empty

	if req.Author == "" || req.Content == "" || req.AuthorUuid == "" || req.ChatRoomUuid == "" || req.ChannelUuid == "" || req.MessageUuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Message is empty")
	}

	//Create the message request
	msgReq := domain.Message{
		Author:       req.Author,
		Content:      req.Content,
		AuthorUuid:   req.AuthorUuid,
		ChatRoomUuid: req.ChatRoomUuid,
		ChannelUuid:  req.ChannelUuid,
		MessageUuid:  req.MessageUuid,
		Timestamp:    req.Timestamp,
	}

	//Edit the msg
	msg, err := s.domain.EditMessage(ctx, msgReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	//Create the response
	msgRes := &messagev1.EditMessageResponse{
		Author:       msg.Author,
		Content:      msg.Content,
		AuthorUuid:   msg.AuthorUuid,
		ChatRoomUuid: msg.ChatRoomUuid,
		ChannelUuid:  msg.ChannelUuid,
		MessageUuid:  msg.MessageUuid,
		Timestamp:    msg.Timestamp,
	}

	return msgRes, nil

}

func (s *server) DeleteMessage(ctx context.Context, req *messagev1.DeleteMessageRequest) (*messagev1.DeleteMessageResponse, error) {
	//Check if message is empty
	if req.Author == "" || req.Content == "" || req.AuthorUuid == "" || req.ChatRoomUuid == "" || req.ChannelUuid == "" || req.MessageUuid == "" {
		return nil, status.Error(codes.InvalidArgument, "Message is empty")
	}

	//Delete the message
	msgReq := domain.Message{
		Author:       req.Author,
		Content:      req.Content,
		AuthorUuid:   req.AuthorUuid,
		ChatRoomUuid: req.ChatRoomUuid,
		ChannelUuid:  req.ChannelUuid,
		MessageUuid:  req.MessageUuid,
		Timestamp:    req.Timestamp,
	}

	err := s.domain.DeleteMessage(ctx, msgReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	//Create the response
	msgRes := &messagev1.DeleteMessageResponse{}

	return msgRes, nil
}
