package server

import (
	"context"

	userv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

func (s *server) EditDescription(ctx context.Context, req *userv1.EditDescriptionRequest) (*userv1.EditDescriptionResponse, error) {

	/* err := s.user.EditDescription(ctx, req.UserUuid, req.Description)
	if err != nil {
		return &userv1.EditDescriptionResponse{}, status.Errorf(codes.Internal, "Error editing description: %v", err)
	} */

	return &userv1.EditDescriptionResponse{}, nil
}

func (s *server) ChangeAvatar(ctx context.Context, req *userv1.ChangeAvatarRequest) (*userv1.ChangeAvatarResponse, error) {

	/* err := s.user.ChangeAvatar(ctx, req.UserUuid, req.AvatarUuid)
	if err != nil {
		return &userv1.ChangeAvatarResponse{}, status.Errorf(codes.Internal, "Error changing avatar: %v", err)
	} */

	return &userv1.ChangeAvatarResponse{}, nil
}
