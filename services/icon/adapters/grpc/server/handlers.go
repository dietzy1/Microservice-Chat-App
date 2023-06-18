package server

import (
	"bytes"
	"context"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	iconv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
)

func (s *server) UploadIcon(stream iconv1.IconService_UploadIconServer) error {

	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Internal, "Error receiving stream: %v", err)
	}

	iconType := req.GetInfo().Kindof

	//Check if the avatar type is equal to png or jpg
	if iconType != ".png" && iconType != ".jpg" && iconType != ".jpeg" && iconType != ".gif" && iconType != ".webp" {
		return status.Errorf(codes.InvalidArgument, "Invalid image type")
	}

	imageData := bytes.Buffer{}
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "Error receiving stream: %v", err)
		}
		chunk := req.GetChunkData()
		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "Error writing image data: %v", err)
		}

	}
	log.Println("finished receiving data")
	log.Println(imageData.Len())

	//send ImageData to domain
	res, err := s.domain.UploadIcon(stream.Context(), imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "Error uploading avatar: %v", err)
	}

	err = stream.SendAndClose(&iconv1.UploadIconResponse{
		Link: res.Link,
		Uuid: res.Uuid,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "Error sending response: %v", err)
	}

	return nil
}

func (s *server) GetIcons(ctx context.Context, req *iconv1.GetIconsRequest) (*iconv1.GetIconsResponse, error) {

	icons, err := s.domain.GetIcons(ctx)
	if err != nil {
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
