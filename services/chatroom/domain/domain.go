package domain

import (
	"context"

	"github.com/google/uuid"
)

type Domain struct {
	repo chatroom
	cdn  cdn
}

func New(repo chatroom, cdn cdn) *Domain {
	return &Domain{repo: repo, cdn: cdn}
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
	Users       []string `json:"users" bson:"users"`
	Uuid        string   `json:"uuid" bson:"uuid"`
	Name        string   `json:"name" bson:"name"`
	Icon        Icon     `json:"icon" bson:"icon"`
	Owner       string   `json:"owner" bson:"owner"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
}

// Need to pass in owner uuid
// Name of the chatroom
// Description of the chatroom
// Tags of the chatroom
func (d *Domain) CreateRoom(ctx context.Context, chatroom Chatroom) (string, error) {
	//Check if chatroom name is taken
	_, err := d.GetRoom(ctx, chatroom.Name)
	if err != nil {
		//return incase the name is already taken
		return "", err
	}
	//perform check if tags and description are not empty
	if chatroom.Description == "" || chatroom.Tags == nil {
		return "", err
	}

	//perform call to cdn and upload the icon
	chatroom.Icon.Link = d.cdn.UploadIcon(ctx, chatroom.Icon)
	chatroom.Icon.Uuid = uuid.New().String()

	//The other fields are already set from the grpc call
	chatroom.Uuid = uuid.New().String()
	chatroom.Users = append(chatroom.Users, chatroom.Owner)

	//Use the chatroom object to create a new chatroom in the database
	err = d.repo.CreateChatroom(ctx, chatroom)
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

// For now make it so only the owner can add users to the chatroom
// Might want to change this to an invite link later though
func (d *Domain) JoinRoom(ctx context.Context, chatroom Chatroom) (string, error) {
	//Check if chatroom exists

	//Check if user uuid is valid
	//Check if user is already in chatroom
	//Add user to chatroom

	return "", nil

}

func (d *Domain) LeaveRoom(ctx context.Context, chatroomUuid string, userUuid string) error {
	//Check if chatroom exists

	//the passed in user uuid should be the one from the cookie

	err := d.repo.RemoveUserFromChatroom(ctx, chatroomUuid, userUuid)
	if err != nil {
		//unable to remove user from chatroom
		return err
	}

	//Check if user uuid is valid
	//Check if user is in chatroom
	//Remove user from chatroom

	return nil

}

// return the chatroom object from the database
func (d *Domain) GetRoom(ctx context.Context, chatroomUuid string) (Chatroom, error) {

	chatroom, err := d.repo.GetChatroom(ctx, chatroomUuid)
	if err != nil {
		//no chatroom was found
		return Chatroom{}, err
	}

	return chatroom, nil
}

/* message CreateRoomRequest {
   string name = 1;
   string owner_uuid = 2;
   }

   message CreateRoomResponse {
   string error = 1;
   string chatroom_uuid = 2;
   }

   message DeleteRoomRequest {
   string chatroom_uuid = 1;
   string owner_uuid = 2;
   }

   message DeleteRoomResponse {
   string error = 1;
   }

   message JoinRoomRequest {
   string chatroom_uuid = 1;
   string user_uuid = 2;
   }

   message JoinRoomResponse {
   string error = 2;
   }

   message LeaveRoomRequest {
   string chatroom_uuid = 1;
   string user_uuid = 2;
   }

   message LeaveRoomResponse {
   string error = 1;
   }

   message GetRoomRequest {
   string chatroom_uuid = 1;
   }

   message GetRoomResponse {
   string error = 1;
   string name = 2;
   string owner_uuid = 3;
   repeated string users = 4;
   } */
