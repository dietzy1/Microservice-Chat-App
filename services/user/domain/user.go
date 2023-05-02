package domain

import (
	"context"
	"math/rand"
	"time"
)

type User struct {
	Name        string   `bson:"name"`
	Uuid        string   `bson:"uuid"`
	Icon        Icon     `bson:"icon"`
	Description string   `bson:"description"`
	JoinDate    string   `bson:"joindate"`
	ChatServers []string `bson:"chatservers"`
}

type user interface {
	AddUser(ctx context.Context, user User) error
	AddChatServer(ctx context.Context, uuid string, serveruuid string) error
	RemoveChatServer(ctx context.Context, uuid string, serveruuid string) error
	EditDescription(ctx context.Context, uuid string, description string) error
	GetUser(ctx context.Context, uuid string) (User, error)
	GetUsers(ctx context.Context, uuids []string) ([]User, error)
	ChangeAvatar(ctx context.Context, userUuid string, avatarUuid string) error
}

func (d Domain) CreateUser(ctx context.Context, username string, uuid string) error {
	// Retrieve all the icons from the database
	avatars, err := d.icon.GetAllIcons(ctx)
	if err != nil {
		return err
	}
	//Choose a random avatar from the list of avatars
	avatar := avatars[rand.Intn(len(avatars))]

	//Create a new user struct
	user := User{
		Name: username,
		Uuid: uuid,
		Icon: Icon{
			Link: avatar.Link,
			Uuid: avatar.Uuid,
		},
		Description: "No description",
		JoinDate:    time.Now().Format("02 January 2006"),
		ChatServers: []string{"5cd69ca7-7fbf-4693-99a7-62ceb4e6a395"},
	}

	//Do call to database and request that the user is added to the database
	if err := d.user.AddUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (d Domain) AddChatServer(ctx context.Context, uuid string, serverUuid string) error {

	//Do call to database and request that the user uuid has the serverUuid added to the chatServers array
	if err := d.user.AddChatServer(ctx, uuid, serverUuid); err != nil {
		return err
	}
	return nil
}

func (d Domain) RemoveChatServer(ctx context.Context, uuid string, server string) error {

	//Do call to database and request that the user uuid has the serverUuid removed from the chatServers array
	if err := d.user.RemoveChatServer(ctx, uuid, server); err != nil {
		return err
	}

	return nil
}

func (d Domain) EditDescription(ctx context.Context, uuid string, description string) error {

	//Do call to database and request that the user uuid has the description changed
	if err := d.user.EditDescription(ctx, uuid, description); err != nil {
		return err
	}

	//Do call to database and request that the user uuid has the description changed
	return nil
}

func (d Domain) GetUser(ctx context.Context, uuid string) (User, error) {

	//Do call to database and request a full user struct based on the uuid
	user, err := d.user.GetUser(ctx, uuid)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (d Domain) GetUsers(ctx context.Context, uuids []string) ([]User, error) {

	//Do call to database and request a full user struct based on the uuid
	users, err := d.user.GetUsers(ctx, uuids)
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (d Domain) ChangeAvatar(ctx context.Context, userUuid string, iconUuid string) error {

	//Do call to database and request that the user uuid has the avatar changed
	if err := d.user.ChangeAvatar(ctx, userUuid, iconUuid); err != nil {
		return err
	}

	return nil
}
