package domain

import (
	"context"
	"time"
)

/* type User struct {
	Icon        Icon //A link to the users icon
	Name        string
	Uuid        string
	Discription string
	JoinDate    string
	Roles       []string
	ChatServers []string
	Reports     int
} */
//Icon of the user
//Preferred name of the user
//Join date
//Chat servers the user is in
//Friends
//Blocked users
//User settings
//User permissions

type User struct {
	Name        string
	Uuid        string
	Icon        Icon //A link to the users icon
	Description string
	JoinDate    string
	ChatServers []string
}

type user interface {
	AddUser(ctx context.Context, user User) error
	AddChatServer(ctx context.Context, uuid string, serveruuid string) error
	RemoveChatServer(ctx context.Context, uuid string, serveruuid string) error
}

func (d Domain) CreateUser(ctx context.Context, username string, uuid string) error {

	user := User{
		Name: username,
		Uuid: uuid,
		Icon: Icon{
			Link: "https://ik.imagekit.io/imageAPI/user/defaultAvatar.jpg?ik-sdk-version=javascript-1.4.3&updatedAt=1677417476895",
			Uuid: "defaultAvatar",
		},
		Description: "No description",
		JoinDate:    time.Now().Format("02 January 2006"),
		ChatServers: []string{"MANUALLY INPUT THE STANDART CHAT SERVER"},
	}

	_ = user
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
	return nil
}
