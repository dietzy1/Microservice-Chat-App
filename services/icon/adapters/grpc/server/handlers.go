package server

import (
	"bytes"
	"context"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	iconv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
)

/* service IconService {
	//get by uuid
	rpc GetIcon(GetIconRequest) returns (GetIconResponse) {}

	//get by owner uuid
	rpc GetIcons(GetIconsRequest) returns (GetIconsResponse) {}

	//return all emoji icons
	rpc GetEmojiIcons(GetEmojiIconsRequest) returns (GetEmojiIconsResponse) {}

	rpc DeleteIcon(DeleteIconRequest) returns (DeleteIconResponse) {}

	rpc UploadIcon(stream UploadIconRequest) returns (UploadIconResponse) {}
  } */

func (s *server) GetIcon(ctx context.Context, req *iconv1.GetIconRequest) (*iconv1.GetIconResponse, error) {

	//Veryfy that the uuid is not empty
	if req.Uuid == "" {
		s.logger.Error("Invalid uuid")
		return &iconv1.GetIconResponse{}, status.Errorf(codes.InvalidArgument, "Invalid uuid")
	}

	icon, err := s.domain.GetIcon(ctx, req.Uuid)
	if err != nil {
		s.logger.Error("Error getting avatar %v", zap.Error(err))
		return &iconv1.GetIconResponse{}, status.Errorf(codes.Internal, "Error getting avatar: %v", err)
	}

	return &iconv1.GetIconResponse{
		Icon: &iconv1.Icon{
			Uuid:      icon.Uuid,
			Link:      icon.Link,
			Kindof:    icon.Kindof,
			OwnerUuid: icon.OwnerUuid,
		},
	}, nil

}

func (s *server) GetIcons(ctx context.Context, req *iconv1.GetIconsRequest) (*iconv1.GetIconsResponse, error) {

	//Verify that the owner uuid is not empty
	if req.OwnerUuid == "" {
		s.logger.Error("Invalid owner uuid")
		return &iconv1.GetIconsResponse{}, status.Errorf(codes.InvalidArgument, "Invalid owner uuid")
	}

	icons, err := s.domain.GetIcons(ctx, req.OwnerUuid)
	if err != nil {
		s.logger.Error("Error getting avatars %v", zap.Error(err))
		return &iconv1.GetIconsResponse{}, status.Errorf(codes.Internal, "Error getting avatars: %v", err)
	}

	//convert domain avatars to proto avatars
	protoIcons := []*iconv1.Icon{}
	for _, icon := range icons {
		protoIcons = append(protoIcons, &iconv1.Icon{
			Uuid:      icon.Uuid,
			Link:      icon.Link,
			Kindof:    icon.Kindof,
			OwnerUuid: icon.OwnerUuid,
		})
	}

	return &iconv1.GetIconsResponse{
		Icons: protoIcons,
	}, nil

}

func (s *server) GetEmojiIcons(ctx context.Context, req *iconv1.GetEmojiIconsRequest) (*iconv1.GetEmojiIconsResponse, error) {

	icons, err := s.domain.GetEmojiIcons(ctx)
	if err != nil {
		s.logger.Error("Error getting avatars %v", zap.Error(err))
		return &iconv1.GetEmojiIconsResponse{}, status.Errorf(codes.Internal, "Error getting avatars: %v", err)
	}

	//convert domain avatars to proto avatars
	protoIcons := []*iconv1.Icon{}
	for _, icon := range icons {
		protoIcons = append(protoIcons, &iconv1.Icon{
			Uuid:      icon.Uuid,
			Link:      icon.Link,
			Kindof:    icon.Kindof,
			OwnerUuid: icon.OwnerUuid,
		})
	}

	return &iconv1.GetEmojiIconsResponse{
		Icons: protoIcons,
	}, nil

}

func (s *server) DeleteIcon(ctx context.Context, req *iconv1.DeleteIconRequest) (*iconv1.DeleteIconResponse, error) {

	//Verify that the uuid is not empty
	if req.Uuid == "" {
		s.logger.Error("Invalid uuid")
		return &iconv1.DeleteIconResponse{}, status.Errorf(codes.InvalidArgument, "Invalid uuid")
	}

	err := s.domain.DeleteIcon(ctx, req.Uuid)
	if err != nil {
		s.logger.Error("Error deleting avatar %v", zap.Error(err))
		return &iconv1.DeleteIconResponse{}, status.Errorf(codes.Internal, "Error deleting avatar: %v", err)
	}

	return &iconv1.DeleteIconResponse{}, nil

}

func (s *server) UploadIcon(stream iconv1.IconService_UploadIconServer) error {

	req, err := stream.Recv()
	if err != nil {
		s.logger.Error("Error receiving stream")
		return status.Errorf(codes.Internal, "Error receiving stream: %v", err)
	}

	iconType := req.GetInfo().Kindof

	//Check if the avatar type is equal to png or jpg
	if iconType != ".png" && iconType != ".jpg" && iconType != ".jpeg" && iconType != ".gif" && iconType != ".webp" {
		s.logger.Error("Invalid image type")
		return status.Errorf(codes.InvalidArgument, "Invalid image type")
	}

	imageData := bytes.Buffer{}
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			s.logger.Info("finished receiving data")
			break
		}
		if err != nil {
			s.logger.Error("Error receiving stream")
			return status.Errorf(codes.Internal, "Error receiving stream: %v", err)
		}
		chunk := req.GetChunkData()
		_, err = imageData.Write(chunk)
		if err != nil {
			s.logger.Error("Error writing image data")
			return status.Errorf(codes.Internal, "Error writing image data: %v", err)
		}

	}
	s.logger.Info("Image data received")

	//send ImageData to domain
	res, err := s.domain.UploadIcon(stream.Context(), imageData)
	if err != nil {
		s.logger.Error("Error uploading avatar")
		return status.Errorf(codes.Internal, "Error uploading avatar: %v", err)
	}

	err = stream.SendAndClose(&iconv1.UploadIconResponse{
		Link: res.Link,
		Uuid: res.Uuid,
	})
	if err != nil {
		s.logger.Error("Error sending response: %v", zap.Error(err))
		return status.Errorf(codes.Internal, "Error sending response: %v", err)
	}

	return nil
}
