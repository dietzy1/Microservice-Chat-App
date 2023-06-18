package server

import (
	"context"
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

func (s *server) GetUsers(ctx context.Context, req *userv1.GetUsersRequest) (*userv1.GetUsersResponse, error) {

	//data validation
	if req.UserUuids == nil {
		return &userv1.GetUsersResponse{}, status.Errorf(codes.InvalidArgument, "Invalid arguments")
	}

	users, err := s.user.GetUsers(ctx, req.UserUuids)
	if err != nil {
		return &userv1.GetUsersResponse{}, status.Errorf(codes.Internal, "Error getting users: %v", err)
	}

	resp := &userv1.GetUsersResponse{}

	for _, u := range users {
		resp.Users = append(resp.Users, &userv1.GetUserResponse{
			Name:        u.Name,
			Uuid:        u.Uuid,
			Icon:        &userv1.Icon{Uuid: u.Icon.Uuid, Link: u.Icon.Link},
			Description: u.Description,
			JoinDate:    u.JoinDate,
			ChatServers: u.ChatServers,
		})
	}
	log.Println(resp.Users)

	return resp, nil

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

	//data validation
	if req.UserUuid == "" || req.IconUuid == "" {
		return &userv1.ChangeAvatarResponse{}, status.Errorf(codes.InvalidArgument, "Invalid arguments")
	}

	err := s.user.ChangeAvatar(ctx, req.UserUuid, req.IconUuid)
	if err != nil {
		return &userv1.ChangeAvatarResponse{}, status.Errorf(codes.Internal, "Error changing avatar: %v", err)
	}

	return &userv1.ChangeAvatarResponse{}, nil
}
