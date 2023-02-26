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
			Link: "https://ik.imagekit.io/your_imagekit_id/icons/default.jpg",
			Uuid: "default",
		},
		Description: "No description yet",
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

/* user := core.User{
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
if err := d.user.AddUser(ctx, user); err != nil {
	return err
} */

/* func (a Domain) setName(context context.Context) {
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
} */

//Should be able to change icon - default is some random icon
//Should be able to change display name -- default is username
//UUID should be recived from the registration process
//Join date is recieved from the registration process
//Chat servers should be recieved from a grpc call from the message server
