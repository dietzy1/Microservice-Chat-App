package server

import (
	"context"

	"github.com/dietzy1/chatapp/services/apigateway/metrics"
	userv1 "github.com/dietzy1/chatapp/services/apigateway/usergateway/v1"
	iconclientv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
	userclientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {

	metrics.UserRequestCounter.Inc()

	user, err := s.userClient.GetUser(ctx, &userclientv1.GetUserRequest{
		UserUuid: req.UserUuid,
	})
	if err != nil {
		return &userv1.GetUserResponse{}, status.Errorf(codes.Internal, "Error getting user: %v", err)
	}

	//I need to convert the user.Icon to a userclientv1.Icon type
	return &userv1.GetUserResponse{
		Name: user.Name,
		Uuid: user.Uuid,
		Icon: &userv1.Icon{
			Uuid: user.Icon.Uuid,
			Link: user.Icon.Link,
		},
		Description: user.Description,
		JoinDate:    user.JoinDate,
		ChatServers: user.ChatServers,
	}, nil

}

func (s *server) GetUsers(ctx context.Context, req *userv1.GetUsersRequest) (*userv1.GetUsersResponse, error) {

	metrics.UserRequestCounter.Inc()

	users, err := s.userClient.GetUsers(ctx, &userclientv1.GetUsersRequest{
		UserUuids: req.UserUuids,
	})
	if err != nil {
		return &userv1.GetUsersResponse{}, status.Errorf(codes.Internal, "Error getting users: %v", err)
	}

	//I need to convert the users.Users to a []*userv1.GetUserResponse type
	resp := &userv1.GetUsersResponse{}

	for _, u := range users.Users {
		resp.Users = append(resp.Users, &userv1.GetUserResponse{
			Name: u.Name,
			Uuid: u.Uuid,
			Icon: &userv1.Icon{
				Uuid: u.Icon.Uuid,
				Link: u.Icon.Link,
			},
			Description: u.Description,
			JoinDate:    u.JoinDate,
			ChatServers: u.ChatServers,
		})
	}

	return resp, nil

}

func (s *server) EditDescription(ctx context.Context, req *userv1.EditDescriptionRequest) (*userv1.EditDescriptionResponse, error) {

	metrics.UserRequestCounter.Inc()

	_, err := s.userClient.EditDescription(ctx, &userclientv1.EditDescriptionRequest{
		UserUuid:    req.UserUuid,
		Description: req.Description,
	})
	if err != nil {
		return &userv1.EditDescriptionResponse{}, status.Errorf(codes.Internal, "Error editing description: %v", err)
	}

	return &userv1.EditDescriptionResponse{}, nil

}

func (s *server) ChangeAvatar(ctx context.Context, req *userv1.ChangeAvatarRequest) (*userv1.ChangeAvatarResponse, error) {

	metrics.UserRequestCounter.Inc()

	_, err := s.userClient.ChangeAvatar(ctx, &userclientv1.ChangeAvatarRequest{
		UserUuid: req.UserUuid,
		IconUuid: req.IconUuid,
	})
	if err != nil {
		return &userv1.ChangeAvatarResponse{}, status.Errorf(codes.Internal, "Error changing avatar: %v", err)
	}

	return &userv1.ChangeAvatarResponse{}, nil

}

func (s *server) GetAvatars(ctx context.Context, req *userv1.GetAvatarsRequest) (*userv1.GetAvatarsResponse, error) {

	metrics.UserRequestCounter.Inc()

	avatars, err := s.iconClient.GetIcons(ctx, &iconclientv1.GetIconsRequest{})
	if err != nil {
		return &userv1.GetAvatarsResponse{}, status.Errorf(codes.Internal, "Error getting avatars: %v", err)
	}

	//I need to convert the avatars.Icons to a []*userv1.Icon type
	var icons []*userv1.Icon
	for _, icon := range avatars.Icons {
		icons = append(icons, &userv1.Icon{
			Uuid: icon.Uuid,
			Link: icon.Link,
		})
	}

	return &userv1.GetAvatarsResponse{
		Icons: icons,
	}, nil

}
