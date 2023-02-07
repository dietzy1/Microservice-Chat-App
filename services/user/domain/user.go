package domain

import (
	"context"
	"time"

	"github.com/dietzy1/chatapp/services/user/core"
)

//Icon of the user
//Preferred name of the user
//Join date
//Chat servers the user is in
//Friends
//Blocked users
//User settings
//User permissions

type User struct {
	Icon        Icon //A link to the users icon
	Name        string
	Uuid        string
	Discription string
	JoinDate    string
	ChatServers []string
}

type user interface {
	AddUser(ctx context.Context, user core.User) error
	ChangeUser(ctx context.Context, uuid string, key string, value string) error
	DeleteUser(ctx context.Context, uuid string) error
	UpdateChatServers(ctx context.Context, uuid string, server string) error
}

func (a Domain) AddUser(ctx context.Context, username string, uuid string) error {
	user := core.User{
		Icon: core.Icon{
			Link: "123",
			Uuid: "123",
		},
		Name:        username,
		Uuid:        uuid,
		Discription: "Click to add a discription",
		JoinDate:    time.Now().Format("2006-01-02 15:04:05"),
		Roles:       []string{"user"},
		ChatServers: []string{"123"},
		Reports:     0,
	}
	if err := a.user.AddUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (a Domain) setName(context context.Context) {
	name := "newName"
	uuid := "uuid"

	err := a.user.ChangeUser(context, "name", uuid, name)
	if err != nil {
		return
	}
}

func (a Domain) setChatServers(context context.Context) {

	chatServer := "chatserver"
	uuid := "uuid"

	err := a.user.ChangeUser(context, "chatServers", uuid, chatServer)
	if err != nil {
		return
	}
}

//Should be able to change icon - default is some random icon
//Should be able to change display name -- default is username
//UUID should be recived from the registration process
//Join date is recieved from the registration process
//Chat servers should be recieved from a grpc call from the message server
