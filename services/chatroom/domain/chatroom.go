package domain

import (
	"context"
	"errors"
	"fmt"

	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
	"go.uber.org/zap"

	"github.com/google/uuid"
)

// Need to pass in owner uuid
// Name of the chatroom
// Description of the chatroom
// Tags of the chatroom
// FIXME:
func (d *Domain) CreateRoom(ctx context.Context, chatroom Chatroom) (string, error) {

	if chatroom.Name == "" || chatroom.Owner == "" {
		d.logger.Info("name and OwnerUuid cannot be empty")
		return "", errors.New("name and OwnerUuid cannot be empty")
	}

	//Name is set
	chatroom.Uuid = uuid.New().String()

	//THIS PART WE NEED TO FIX
	chatroom.Icon.Link = "https://ik.imagekit.io/imageAPI/user/a884bf0f-4b26-4a4f-a454-d0e0286882e4.png"
	chatroom.Icon.Uuid = "a884bf0f-4b26-4a4f-a454-d0e0286882e4"
	//Owner is set
	chatroom.Users = append(chatroom.Users, chatroom.Owner)

	//I need to create a standart channel sortof
	channel := Channel{
		ChannelUuid:  uuid.New().String(),
		Name:         "general",
		ChatroomUuid: chatroom.Uuid,
		Owner:        chatroom.Owner,
	}
	chatroom.Channels = append(chatroom.Channels, channel)

	//Use the chatroom object to create a new chatroom in the database
	err := d.repo.CreateChatroom(ctx, chatroom)
	if err != nil {
		d.logger.Info("Error creating chatroom", zap.Error(err))
		return "", fmt.Errorf("error creating chatroom: %w", err)
	}

	//if everything passed then we need to contact the user service and add the chatroom to the owner
	_, err = d.userClient.AddChatServer(ctx, &userClientv1.AddChatServerRequest{
		UserUuid:       chatroom.Owner,
		ChatserverUuid: chatroom.Uuid,
	})
	if err != nil {
		d.logger.Info("Error adding chatroom to user", zap.Error(err))
		return "", fmt.Errorf("error adding chatroom to user: %w", err)

	}

	return chatroom.Uuid, nil
	//Should prolly come up with some default values incase nothing is provided
}

// This call should also delete the messages in the chatroom
func (d *Domain) DeleteRoom(ctx context.Context, chatroom Chatroom) error {
	//The owner uuid should be passed through the cookie so it should already be authenticated prior to this call

	if chatroom.Uuid == "" || chatroom.Owner == "" {
		d.logger.Info("Uuid and OwnerUuid cannot be empty")
		return errors.New("uuid and OwnerUuid cannot be empty")
	}

	//Check if owner uuid is the owner of the chatroom
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {

		d.logger.Info("Chatroom does not exist", zap.Error(err))
		return fmt.Errorf("chatroom does not exist: %w", err)
	}
	if room.Owner != chatroom.Owner {
		//User is not the owner of the chatroom
		d.logger.Info("User is not the owner of the chatroom", zap.Error(err))
		return errors.New("user is not the owner of the chatroom")
	}

	//Delete chatroom
	err = d.repo.DeleteChatroom(ctx, chatroom.Uuid)
	if err != nil {
		d.logger.Info("Error deleting chatroom", zap.Error(err))
		return fmt.Errorf("error deleting chatroom: %w", err)
	}
	return nil
}

func (d *Domain) GetRoom(ctx context.Context, chatroom Chatroom) (Chatroom, error) {

	if chatroom.Uuid == "" {
		d.logger.Info("Uuid cannot be empty")
		return Chatroom{}, errors.New("uuid cannot be empty")
	}

	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		d.logger.Info("Error getting chatroom", zap.Error(err))
		return Chatroom{}, fmt.Errorf("error getting chatroom: %w", err)
	}

	return room, nil
}

func (d *Domain) GetRooms(ctx context.Context, chatroomUuids []string) ([]Chatroom, error) {

	if len(chatroomUuids) == 0 {
		d.logger.Info("Uuids cannot be empty")
		return []Chatroom{}, errors.New("uuids cannot be empty")
	}

	rooms, err := d.repo.GetChatrooms(ctx, chatroomUuids)
	if err != nil {
		d.logger.Info("Error getting chatrooms", zap.Error(err))
		return []Chatroom{}, fmt.Errorf("error getting chatrooms: %w", err)
	}

	return rooms, nil
}

//---------------------------------------------------------------------------------

func (d *Domain) CreateChannel(ctx context.Context, channel Channel) (string, error) {

	if channel.Name == "" || channel.Owner == "" || channel.ChatroomUuid == "" {
		d.logger.Info("Name, OwnerUuid and ChatroomUuid cannot be empty")
		return "", errors.New("name, owneruuid and chatroomuuid cannot be empty")
	}

	//The other fields are already set from the grpc call
	channel.ChannelUuid = uuid.New().String()
	//Use the channel object to create a new channel in the database
	err := d.repo.CreateChannel(ctx, channel)
	if err != nil {
		d.logger.Info("Error creating channel", zap.Error(err))
		return "", fmt.Errorf("error creating channel: %w", err)
	}

	return channel.ChannelUuid, nil
	//Should prolly come up with some default values incase nothing is provided
}

func (d *Domain) DeleteChannel(ctx context.Context, channel Channel) error {
	//The owner uuid should be passed through the cookie so it should already be authenticated prior to this call

	if channel.ChannelUuid == "" || channel.Owner == "" || channel.ChatroomUuid == "" {
		d.logger.Info("channel, chatroom uuids and OwnerUuid cannot be empty")
		return errors.New("channel, chatroom uuids and OwnerUuid cannot be empty")
	}

	//Check if owner uuid is the owner of the chatroom
	room, err := d.repo.GetChatroom(ctx, channel.ChatroomUuid)
	if err != nil {
		//Chatroom does not exist
		d.logger.Info("Chatroom does not exist", zap.Error(err))
		return fmt.Errorf("chatroom does not exist: %w", err)

	}
	if room.Owner != channel.Owner {
		//User is not the owner of the chatroom
		d.logger.Info("User is not the owner of the chatroom", zap.Error(err))
		return errors.New("user is not the owner of the chatroom")
	}

	//Delete channel
	err = d.repo.DeleteChannel(ctx, channel.ChannelUuid)
	if err != nil {
		d.logger.Info("Error deleting channel", zap.Error(err))
		return fmt.Errorf("error deleting channel: %w", err)
	}
	return nil
}

func (d *Domain) GetChannel(ctx context.Context, channel Channel) (Channel, error) {

	if channel.ChatroomUuid == "" || channel.ChannelUuid == "" {
		d.logger.Info("ChatroomUuid and ChannelUuid cannot be empty")
		return Channel{}, errors.New("chatroomuuid and channeluuid cannot be empty")
	}

	channel, err := d.repo.GetChannel(ctx, channel.ChannelUuid)
	if err != nil {
		d.logger.Info("Error getting channel", zap.Error(err))
		return Channel{}, fmt.Errorf("error getting channel: %w", err)
	}

	return Channel{}, nil
}

//---------------------------------------------------------------------------------

func (d *Domain) InviteUser(ctx context.Context, chatroom Chatroom, userUuid string) error {

	if userUuid == "" || chatroom.Uuid == "" || chatroom.Owner == "" {
		d.logger.Info("UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
		return errors.New("UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
	}

	//Check if user is already in the chatroom
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		d.logger.Info("Chatroom does not exist", zap.Error(err))
		return fmt.Errorf("chatroom does not exist: %w", err)
	}
	//Check if user is already in the chatroom
	for _, user := range room.Users {
		if user == userUuid {
			d.logger.Info("User is already in the chatroom")
			return nil
		}
	}

	//Add user to invited list
	room.Invited = append(chatroom.Invited, userUuid)

	//Update chatroom
	err = d.repo.InviteUser(ctx, chatroom.Uuid, userUuid)
	if err != nil {
		d.logger.Info("Error inviting user", zap.Error(err))
		return fmt.Errorf("error inviting user: %w", err)
	}

	return nil
}

func (d *Domain) RemoveUser(ctx context.Context, chatroom Chatroom, userUuid string) error {

	if userUuid == "" || chatroom.Uuid == "" || chatroom.Owner == "" {
		d.logger.Info("UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
		return errors.New("UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
	}

	//Check if owner uuid is the owner of the chatroom
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		//Chatroom does not exist
		d.logger.Info("Chatroom does not exist", zap.Error(err))
		return fmt.Errorf("chatroom does not exist: %w", err)
	}
	if room.Owner != chatroom.Owner {
		//User is not the owner of the chatroom
		d.logger.Info("User is not the owner of the chatroom", zap.Error(err))
		return errors.New("user is not the owner of the chatroom")
	}

	//Update chatroom
	err = d.repo.RemoveUser(ctx, chatroom.Uuid, userUuid)
	if err != nil {
		d.logger.Info("Error removing user", zap.Error(err))
		return fmt.Errorf("error removing user: %w", err)
	}

	return nil
}

func (d *Domain) AddUser(ctx context.Context, chatroom Chatroom, userUuid string) error {

	if userUuid == "" || chatroom.Uuid == "" || chatroom.Owner == "" {
		d.logger.Info("UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
		return errors.New("UserUuid, ChatroomUuid and OwnerUuid cannot be empty")
	}

	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		d.logger.Info("Chatroom does not exist", zap.Error(err))
		return fmt.Errorf("chatroom does not exist: %w", err)
	}
	//Check if user is in the invited list
	for _, user := range room.Invited {
		if user == userUuid {
			//Add user to chatroom
			err = d.repo.AddUser(ctx, chatroom.Uuid, userUuid)
			if err != nil {
				d.logger.Info("Error adding user", zap.Error(err))
				return fmt.Errorf("error adding user: %w", err)
			}
		}
	}

	return nil
}

// internal use
func (d *Domain) ForceAddUser(ctx context.Context, userUuid string) error {

	if userUuid == "" {
		d.logger.Info("UserUuid cannot be empty")
		return errors.New("UserUuid cannot be empty")
	}

	const chatroomUuid = "5cd69ca7-7fbf-4693-99a7-62ceb4e6a395"

	err := d.repo.ForceAddUser(ctx, chatroomUuid, userUuid)
	if err != nil {
		d.logger.Info("Error adding user", zap.Error(err))
		return fmt.Errorf("error adding user: %w", err)
	}
	return nil
}

func (d *Domain) ChangeIcon(ctx context.Context, chatroom Chatroom) error {

	if chatroom.Uuid == "" || chatroom.Icon.Uuid == "" || chatroom.Icon.Link == "" {
		d.logger.Info("ChatroomUuid, IconUuid and IconLink cannot be empty")
		return errors.New("ChatroomUuid, IconUuid and IconLink cannot be empty")
	}

	err := d.repo.ChangeIcon(ctx, chatroom.Uuid, chatroom.Icon)
	if err != nil {
		d.logger.Info("Error changing icon", zap.Error(err))
		return fmt.Errorf("error changing icon: %w", err)
	}

	return nil
}
