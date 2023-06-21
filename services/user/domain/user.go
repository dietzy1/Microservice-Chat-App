package domain

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"

	iconv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
)

func (d Domain) CreateUser(ctx context.Context, username string, uuid string) error {

	//Move validation to here
	if username == "" || uuid == "" {
		d.logger.Info("Username or uuid is empty")
		return errors.New("username or uuid is empty")
	}

	//here we need to call into the icon service and get a random avatar
	icons, err := d.iconClient.GetEmojiIcons(ctx, &iconv1.GetEmojiIconsRequest{})
	if err != nil {
		d.logger.Info("Error getting icons from icon service", zap.Error(err))
		//return error with message
		return fmt.Errorf("error getting icons from icon service: %w", err)
	}

	//Convert icon slice to a slice of icons
	avatars := make([]Icon, len(icons.Icons))
	for i, icon := range icons.Icons {
		avatars[i] = Icon{
			Link: icon.Link,
			Uuid: icon.Uuid,
		}
	}

	//Choose a random icon from the list of avatars
	icon := avatars[rand.Intn(len(avatars))]

	//Create a new user struct
	user := User{
		Name: username,
		Uuid: uuid,
		Icon: Icon{
			Link: icon.Link,
			Uuid: icon.Uuid,
		},
		Description: "No description",
		JoinDate:    time.Now().Format("02 January 2006"),
		ChatServers: []string{"5cd69ca7-7fbf-4693-99a7-62ceb4e6a395"},
	}

	//Do call to database and request that the user is added to the database
	if err := d.user.AddUser(ctx, user); err != nil {
		d.logger.Info("Error adding user to database", zap.Error(err))
		return fmt.Errorf("error adding user to database: %w", err)
	}

	return nil
}

func (d Domain) AddChatServer(ctx context.Context, uuid string, serverUuid string) error {

	if uuid == "" || serverUuid == "" {
		d.logger.Info("user uuid or serveruuid is empty")
		return errors.New("user uuid or serveruuid is empty")
	}

	//Do call to database and request that the user uuid has the serverUuid added to the chatServers array
	if err := d.user.AddChatServer(ctx, uuid, serverUuid); err != nil {
		d.logger.Info("Error adding chat server to user", zap.Error(err))
		return fmt.Errorf("error adding chat server to user: %w", err)
	}
	return nil
}

func (d Domain) RemoveChatServer(ctx context.Context, uuid string, serverUuid string) error {

	if uuid == "" || serverUuid == "" {
		d.logger.Info("uuid or server uuid is empty")
		return errors.New("uuid or server uuid is empty")
	}

	//Do call to database and request that the user uuid has the serverUuid removed from the chatServers array
	if err := d.user.RemoveChatServer(ctx, uuid, serverUuid); err != nil {
		d.logger.Info("Error removing chat server from user", zap.Error(err))
		return fmt.Errorf("error removing chat server from user: %w", err)
	}

	return nil
}

func (d Domain) EditDescription(ctx context.Context, uuid string, description string) error {

	if uuid == "" || description == "" {
		d.logger.Info("uuid or description is empty")
		return errors.New("uuid or description is empty")
	}

	//Do call to database and request that the user uuid has the description changed
	if err := d.user.EditDescription(ctx, uuid, description); err != nil {
		d.logger.Info("Error editing description", zap.Error(err))
		return fmt.Errorf("error editing description: %w", err)
	}

	//Do call to database and request that the user uuid has the description changed
	return nil
}

func (d Domain) GetUser(ctx context.Context, uuid string) (User, error) {

	if uuid == "" {
		d.logger.Info("uuid is empty")
		return User{}, errors.New("uuid is empty")
	}

	//Do call to database and request a full user struct based on the uuid
	user, err := d.user.GetUser(ctx, uuid)
	if err != nil {
		d.logger.Info("Error getting user from database", zap.Error(err))
		return User{}, fmt.Errorf("error getting user from database: %w", err)
	}
	return user, nil
}

func (d Domain) GetUsers(ctx context.Context, uuids []string) ([]User, error) {

	if len(uuids) == 0 {
		d.logger.Info("uuids is empty")
		return []User{}, errors.New("uuids is empty")
	}

	//Do call to database and request a full user struct based on the uuid
	users, err := d.user.GetUsers(ctx, uuids)
	if err != nil {
		d.logger.Info("Error getting users from database", zap.Error(err))
		return []User{}, fmt.Errorf("error getting users from database: %w", err)
	}
	return users, nil
}

func (d Domain) ChangeAvatar(ctx context.Context, userUuid string, iconUuid string) error {

	if userUuid == "" || iconUuid == "" {
		d.logger.Info("user uuid or icon uuid is empty")
		return errors.New("user uuid or icon uuid is empty")
	}

	//Ask the icon service for the link to the icon
	icon, err := d.iconClient.GetIcon(ctx, &iconv1.GetIconRequest{Uuid: iconUuid})
	if err != nil {
		d.logger.Info("Error getting icon from icon service", zap.Error(err))
		return fmt.Errorf("error getting icon from icon service: %w", err)
	}

	//Do call to database and request that the user uuid has the avatar changed
	if err := d.user.ChangeAvatar(ctx, userUuid, Icon{Link: icon.Icon.Link, Uuid: icon.Icon.Uuid}); err != nil {
		d.logger.Info("Error changing avatar", zap.Error(err))
		return fmt.Errorf("error changing avatar: %w", err)
	}

	return nil
}
