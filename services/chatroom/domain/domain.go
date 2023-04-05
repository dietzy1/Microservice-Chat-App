package domain

import (
	"context"
	"errors"
	"log"

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
	DeleteChatroom(ctx context.Context, chatroomUuid string) error
	GetChatroom(ctx context.Context, chatroomUuid string) (Chatroom, error)
	GetChatrooms(ctx context.Context, chatroomUuids []string) ([]Chatroom, error)

	CreateChannel(ctx context.Context, channel Channel) error
	DeleteChannel(ctx context.Context, channelUuid string) error
	GetChannel(ctx context.Context, channelUuid string) (Channel, error)

	InviteUser(ctx context.Context, chatroomUuid string, userUuid string) error
	RemoveUser(ctx context.Context, chatroomUuid string, userUuid string) error
	AddUser(ctx context.Context, chatroomUuid string, userUuid string) error
}

type Chatroom struct {
	Name     string    `json:"name" bson:"name"`
	Uuid     string    `json:"uuid" bson:"uuid"`
	Icon     Icon      `json:"icon" bson:"icon"`
	Owner    string    `json:"owner" bson:"owner"`
	Users    []string  `json:"users" bson:"users"`
	Channels []Channel `json:"channels" bson:"channels"`
	Invited  []string  `json:"invited" bson:"invited"`
}

type Channel struct {
	ChannelUuid  string `json:"channeluuid" bson:"channeluuid"`
	Name         string `json:"name" bson:"name"`
	ChatroomUuid string `json:"chatroomuuid" bson:"chatroomuuid"`
	Owner        string `json:"owner" bson:"owner"`
}

// Need to pass in owner uuid
// Name of the chatroom
// Description of the chatroom
// Tags of the chatroom
func (d *Domain) CreateRoom(ctx context.Context, chatroom Chatroom) (string, error) {

	//Name is set
	chatroom.Uuid = uuid.New().String()
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
		log.Println(err)
		return "", err
	}

	err = d.repo.CreateChannel(ctx, channel)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return chatroom.Uuid, nil
	//Should prolly come up with some default values incase nothing is provided
}

// This call should also delete the messages in the chatroom
func (d *Domain) DeleteRoom(ctx context.Context, chatroom Chatroom) error {
	//The owner uuid should be passed through the cookie so it should already be authenticated prior to this call

	//Check if owner uuid is the owner of the chatroom
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		//Chatroom does not exist
		log.Println(err)
		return err
	}
	if room.Owner != chatroom.Owner {
		//User is not the owner of the chatroom
		log.Println("User is not the owner of the chatroom")
		return errors.New("user is not the owner of the chatroom")
	}

	//Delete chatroom
	err = d.repo.DeleteChatroom(ctx, chatroom.Uuid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *Domain) GetRoom(ctx context.Context, chatroom Chatroom) (Chatroom, error) {

	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		log.Println(err)
		return Chatroom{}, err
	}

	return room, nil
}

func (d *Domain) GetRooms(ctx context.Context, chatroomUuids []string) ([]Chatroom, error) {
	log.Println(chatroomUuids)
	rooms, err := d.repo.GetChatrooms(ctx, chatroomUuids)
	if err != nil {
		log.Println(err)
		return []Chatroom{}, err
	}

	return rooms, nil
}

//---------------------------------------------------------------------------------

func (d *Domain) CreateChannel(ctx context.Context, channel Channel) (string, error) {

	//The other fields are already set from the grpc call
	channel.ChannelUuid = uuid.New().String()
	//Use the channel object to create a new channel in the database
	err := d.repo.CreateChannel(ctx, channel)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return channel.ChannelUuid, nil
	//Should prolly come up with some default values incase nothing is provided
}

func (d *Domain) DeleteChannel(ctx context.Context, channel Channel) error {
	//The owner uuid should be passed through the cookie so it should already be authenticated prior to this call

	//Check if owner uuid is the owner of the chatroom
	room, err := d.repo.GetChatroom(ctx, channel.ChatroomUuid)
	if err != nil {
		//Chatroom does not exist
		log.Println(err)
		return err
	}
	if room.Owner != channel.Owner {
		//User is not the owner of the chatroom
		log.Println(err)
		return errors.New("user is not the owner of the chatroom")
	}

	//Delete channel
	err = d.repo.DeleteChannel(ctx, channel.ChannelUuid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *Domain) GetChannel(ctx context.Context, channel Channel) (Channel, error) {

	channel, err := d.repo.GetChannel(ctx, channel.ChannelUuid)
	if err != nil {
		log.Println(err)
		return Channel{}, err
	}

	return Channel{}, nil
}

//---------------------------------------------------------------------------------

func (d *Domain) InviteUser(ctx context.Context, chatroom Chatroom, userUuid string) error {
	//Check if user is already in the chatroom
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		log.Println(err)
		return err
	}
	//Check if user is already in the chatroom
	for _, user := range room.Users {
		if user == userUuid {
			return nil
		}
	}

	//Add user to invited list
	room.Invited = append(chatroom.Invited, userUuid)

	//Update chatroom
	err = d.repo.InviteUser(ctx, chatroom.Uuid, userUuid)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *Domain) RemoveUser(ctx context.Context, chatroom Chatroom, userUuid string) error {

	//Check if owner uuid is the owner of the chatroom
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		//Chatroom does not exist
		log.Println(err)
		return err
	}
	if room.Owner != chatroom.Owner {
		//User is not the owner of the chatroom
		log.Println(err)
		return errors.New("user is not the owner of the chatroom")
	}

	//Update chatroom
	err = d.repo.RemoveUser(ctx, chatroom.Uuid, userUuid)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *Domain) AddUser(ctx context.Context, chatroom Chatroom, userUuid string) error {
	room, err := d.repo.GetChatroom(ctx, chatroom.Uuid)
	if err != nil {
		log.Println(err)
		return err
	}
	//Check if user is in the invited list
	for _, user := range room.Invited {
		if user == userUuid {
			//Add user to chatroom
			err = d.repo.AddUser(ctx, chatroom.Uuid, userUuid)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	return nil
}
