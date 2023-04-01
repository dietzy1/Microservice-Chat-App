package domain

import (
	"context"

	"github.com/google/uuid"
)

type Domain struct {
	repo chatroom
	//cdn  cdn
}

func New(repo chatroom /* , cdn cdn */) *Domain {
	return &Domain{repo: repo /* , cdn: cdn */}
}

// this is the repository interface
type chatroom interface {
	CreateChatroom(ctx context.Context, chatroom Chatroom) error
	GetChatroom(ctx context.Context, chatroomUuid string) (Chatroom, error)
	DeleteChatroom(ctx context.Context, chatroomUuid string) error
	AddUserToChatroom(ctx context.Context, chatroomUuid string, userUuid string) error
	RemoveUserFromChatroom(ctx context.Context, chatroomUuid string, userUuid string) error
	/* ChangeDescription()
	ChangeName()
	ChangeTags() */
	//these should be implemented later but its not nessesary for the first version

}

type Chatroom struct {
	Name     string   `json:"name" bson:"name"`
	Uuid     string   `json:"uuid" bson:"uuid"`
	Icon     Icon     `json:"icon" bson:"icon"`
	Owner    string   `json:"owner" bson:"owner"`
	Users    []string `json:"users" bson:"users"`
	Channels []string `json:"channels" bson:"channels"`
	Invited  []string `json:"invited" bson:"invited"`
}

type Channel struct {
	ChannelUuid  string `json:"uuid" bson:"uuid"`
	Name         string `json:"name" bson:"name"`
	ChatroomUuid string `json:"chatroom" bson:"chatroom"`
	Owner        string `json:"owner" bson:"owner"`
}

// Need to pass in owner uuid
// Name of the chatroom
// Description of the chatroom
// Tags of the chatroom
func (d *Domain) CreateRoom(ctx context.Context, chatroom Chatroom) (string, error) {

	//Name is set
	chatroom.Uuid = uuid.New().String()
	chatroom.Icon.Link = "Some standart shit"
	chatroom.Icon.Uuid = "Some other standart shit"
	//Owner is set
	chatroom.Users = append(chatroom.Users, chatroom.Owner)

	//I need to create a standart channel sortof
	channel := Channel{
		ChannelUuid:  uuid.New().String(),
		Name:         "general",
		ChatroomUuid: chatroom.Uuid,
		Owner:        chatroom.Owner,
	}
	chatroom.Channels = append(chatroom.Channels, channel.ChannelUuid)

	//Use the chatroom object to create a new chatroom in the database
	err := d.repo.CreateChatroom(ctx, chatroom)
	if err != nil {
		return "", err
	}

	err = d.repo.CreateChannel(ctx, channel)
	if err != nil {
		return "", err
	}

	return chatroom.Uuid, nil
	//Should prolly come up with some default values incase nothing is provided
}

// This call should also delete the messages in the chatroom
func (d *Domain) DeleteRoom(ctx context.Context, chatroom Chatroom) error {
	//The owner uuid should be passed through the cookie so it should already be authenticated prior to this call

	//Check if owner uuid is the owner of the chatroom
	room, err := d.GetRoom(ctx, chatroom.Uuid)
	if err != nil {
		//Chatroom does not exist
		return err
	}
	if room.Owner != chatroom.Owner {
		//User is not the owner of the chatroom
		return err
	}

	//Delete chatroom
	err = d.repo.DeleteChatroom(ctx, chatroom.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (d *Domain) GetRoom(ctx context.Context, chatroom Chatroom) (Chatroom, error) {

	return Chatroom{}, nil
}

//---------------------------------------------------------------------------------

func (d *Domain) CreateChannel(ctx context.Context, channel Channel) (string, error) {
	//Check if channel name is taken
	_, err := d.GetChannel(ctx, channel.Name)
	if err != nil {
		//return incase the name is already taken
		return "", err
	}

	//The other fields are already set from the grpc call
	channel.ChannelUuid = uuid.New().String()

	//Use the channel object to create a new channel in the database
	err = d.repo.CreateChannel(ctx, channel)
	if err != nil {
		return "", err
	}

	return channel.ChannelUuid, nil
	//Should prolly come up with some default values incase nothing is provided
}

func (d *Domain) DeleteChannel(ctx context.Context, channel Channel) error {
	//The owner uuid should be passed through the cookie so it should already be authenticated prior to this call

	//Check if owner uuid is the owner of the chatroom
	room, err := d.GetRoom(ctx, channel.ChatroomUuid)
	if err != nil {
		//Chatroom does not exist
		return err
	}
	if room.Owner != channel.Owner {
		//User is not the owner of the chatroom
		return err
	}

	//Delete channel
	err = d.repo.DeleteChannel(ctx, channel.ChannelUuid)
	if err != nil {
		return err
	}
	return nil
}

func (d *Domain) GetChannel(ctx context.Context, channel Channel) (Channel, error) {

	return Channel{}, nil
}

//---------------------------------------------------------------------------------

func (d *Domain) InviteUser(ctx context.Context, chatroom Chatroom, userUuid string) (string, error) {
	//Check if user is already in the chatroom
	for _, user := range chatroom.Users {
		if user == userUuid {
			return nil
		}
	}

	//Check if user is already invited
	for _, user := range chatroom.Invited {
		if user == userUuid {
			return nil
		}
	}

	//Add user to invited list
	chatroom.Invited = append(chatroom.Invited, userUuid)

	//Update chatroom
	err := d.repo.UpdateChatroom(ctx, chatroom)
	if err != nil {
		return err
	}

	return nil
}

func (d *Domain) RemoveUser(ctx context.Context, chatroom Chatroom, userUuid string) error {
	//Check if user is in the chatroom
	for i, user := range chatroom.Users {
		if user == userUuid {
			chatroom.Users = append(chatroom.Users[:i], chatroom.Users[i+1:]...)
			break
		}
	}

	//Update chatroom
	err := d.repo.UpdateChatroom(ctx, chatroom)
	if err != nil {
		return err
	}

	return nil
}

func (d *Domain) AddUser(ctx context.Context, chatroom Chatroom, userUuid string) (string, error) {

	return nil
}
