package server

import (
	"bytes"
	"context"
	"io"
	"log"

	userv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {

	//data validation
	if req.UserUuid == "" {
		return &userv1.GetUserResponse{}, status.Errorf(codes.InvalidArgument, "Invalid arguments")
	}

	user, err := s.user.GetUser(ctx, req.UserUuid)
	if err != nil {
		return &userv1.GetUserResponse{}, status.Errorf(codes.Internal, "Error getting user: %v", err)
	}

	return &userv1.GetUserResponse{
		Name:        user.Name,
		Uuid:        user.Uuid,
		Icon:        &userv1.Icon{Uuid: user.Icon.Uuid, Link: user.Icon.Link},
		Description: user.Description,
		JoinDate:    user.JoinDate,
		ChatServers: user.ChatServers,
	}, nil

}

func (s *server) EditDescription(ctx context.Context, req *userv1.EditDescriptionRequest) (*userv1.EditDescriptionResponse, error) {

	//data validation
	if req.UserUuid == "" || req.Description == "" {
		return &userv1.EditDescriptionResponse{}, status.Errorf(codes.InvalidArgument, "Invalid arguments")
	}

	err := s.user.EditDescription(ctx, req.UserUuid, req.Description)
	if err != nil {
		return &userv1.EditDescriptionResponse{}, status.Errorf(codes.Internal, "Error editing description: %v", err)
	}

	return &userv1.EditDescriptionResponse{}, nil
}

func (s *server) ChangeAvatar(ctx context.Context, req *userv1.ChangeAvatarRequest) (*userv1.ChangeAvatarResponse, error) {

	/* err := s.user.ChangeAvatar(ctx, req.UserUuid, req.AvatarUuid)
	if err != nil {
		return &userv1.ChangeAvatarResponse{}, status.Errorf(codes.Internal, "Error changing avatar: %v", err)
	} */

	return &userv1.ChangeAvatarResponse{}, nil
}

func (s *server) UploadAvatar(stream userv1.UserService_UploadAvatarServer) error {

	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Internal, "Error receiving stream: %v", err)
	}

	avatarType := req.GetInfo().ImageType

	//Check if the avatar type is equal to png or jpg
	if avatarType != ".png" && avatarType != ".jpg" && avatarType != ".jpeg" && avatarType != ".gif" && avatarType != ".webp" {
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
	res, err := s.user.UploadAvatar(stream.Context(), imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "Error uploading avatar: %v", err)
	}

	err = stream.SendAndClose(&userv1.UploadAvatarResponse{
		Link: res.Link,
		Uuid: res.Uuid,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "Error sending response: %v", err)
	}

	return nil
}

func (s *server) GetAvatars(ctx context.Context, req *userv1.GetAvatarsRequest) (*userv1.GetAvatarsResponse, error) {

	avatars, err := s.user.GetAllIcons(ctx)
	if err != nil {
		return &userv1.GetAvatarsResponse{}, status.Errorf(codes.Internal, "Error getting avatars: %v", err)
	}

	//convert domain avatars to proto avatars
	protoAvatars := []*userv1.Icon{}
	for _, avatar := range avatars {
		protoAvatars = append(protoAvatars, &userv1.Icon{
			Uuid: avatar.Uuid,
			Link: avatar.Link,
		})
	}

	return &userv1.GetAvatarsResponse{
		Icons: protoAvatars,
	}, nil

}
