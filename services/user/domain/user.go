package domain

import (
	"context"
	"time"
)

type User struct {
	Name        string   `bson:"name"`
	Uuid        string   `bson:"uuid"`
	Icon        Icon     `bson:"icon"`
	Description string   `bson:"description"`
	JoinDate    string   `bson:"joinDate"`
	ChatServers []string `bson:"chatServers"`
}

type user interface {
	AddUser(ctx context.Context, user User) error
	AddChatServer(ctx context.Context, uuid string, serveruuid string) error
	RemoveChatServer(ctx context.Context, uuid string, serveruuid string) error
	EditDescription(ctx context.Context, uuid string, description string) error
	GetUser(ctx context.Context, uuid string) (User, error)
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
