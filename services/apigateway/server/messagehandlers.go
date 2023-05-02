package server

import (
	"context"

	messagev1 "github.com/dietzy1/chatapp/services/apigateway/messagegateway/v1"
	messageclientv1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetMessages(ctx context.Context, req *messagev1.GetMessagesRequest) (*messagev1.GetMessagesResponse, error) {

	//call message service

	res, err := s.messageClient.GetMessages(ctx, &messageclientv1.GetMessagesRequest{
		ChatRoomUuid: req.ChatRoomUuid,
		ChannelUuid:  req.ChannelUuid,
	})
	if err != nil {
		return &messagev1.GetMessagesResponse{
			Messages: nil,
		}, status.Errorf(codes.Internal, "Error getting messages: %v", err)
	}

	resp := &messagev1.GetMessagesResponse{}

	for _, msg := range res.Messages {
		//Take in the messages from the message service and convert them to the messagegatewayv1.msg type
		resp.Messages = append(resp.Messages, &messagev1.Msg{
			Author:       msg.Author,
			Content:      msg.Content,
			AuthorUuid:   msg.AuthorUuid,
			ChannelUuid:  msg.ChannelUuid,
			ChatRoomUuid: msg.ChatRoomUuid,
			MessageUuid:  msg.MessageUuid,
			Timestamp:    msg.Timestamp,
		})
	}

	return resp, nil

}

func (s *server) EditMessage(ctx context.Context, req *messagev1.EditMessageRequest) (*messagev1.EditMessageResponse, error) {

	res, err := s.messageClient.EditMessage(ctx, &messageclientv1.EditMessageRequest{
		ChatRoomUuid: req.ChatRoomUuid,
		ChannelUuid:  req.ChannelUuid,
		MessageUuid:  req.MessageUuid,
		Content:      req.Content,
		Author:       req.Author,
		AuthorUuid:   req.AuthorUuid,
		Timestamp:    req.Timestamp,
	})
	if err != nil {
		return &messagev1.EditMessageResponse{}, status.Errorf(codes.Internal, "Error editing message: %v", err)
	}

	//Convert res to a messagev1.EditMessageResponse
	return &messagev1.EditMessageResponse{
		Author:       res.Author,
		Content:      res.Content,
		AuthorUuid:   res.AuthorUuid,
		ChatRoomUuid: res.ChatRoomUuid,
		ChannelUuid:  res.ChannelUuid,
		MessageUuid:  res.MessageUuid,
		Timestamp:    res.Timestamp,
	}, nil

}

func (s *server) DeleteMessage(ctx context.Context, req *messagev1.DeleteMessageRequest) (*messagev1.DeleteMessageResponse, error) {

	_, err := s.messageClient.DeleteMessage(ctx, &messageclientv1.DeleteMessageRequest{
		Author:       req.Author,
		Content:      req.Content,
		AuthorUuid:   req.AuthorUuid,
		ChatRoomUuid: req.ChatRoomUuid,
		ChannelUuid:  req.ChannelUuid,
		MessageUuid:  req.MessageUuid,
		Timestamp:    req.Timestamp,
	})
	if err != nil {
		return &messagev1.DeleteMessageResponse{}, status.Errorf(codes.Internal, "Error deleting message: %v", err)
	}

	return &messagev1.DeleteMessageResponse{}, nil

}
