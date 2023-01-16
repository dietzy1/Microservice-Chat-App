package domain

import "context"

type Domain struct {
	repo chatroom
}

func New(repo chatroom) *Domain {
	return &Domain{repo: repo}
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

	//Check if tags are valid
	//Check if description is valid
	//Check if name is valid
	//Create chatroom

	return "", nil
	//Should prolly come up with some default values incase nothing is provided
}

func (d *Domain) DeleteRoom(ctx context.Context, chatroom Chatroom) (string, error) {
	//Check if chatroom exists
	//Check if owner uuid is valid
	//Check if owner uuid is the owner of the chatroom
	//Delete chatroom

	return "", nil
}

func (d *Domain) JoinRoom(ctx context.Context, chatroom Chatroom) (string, error) {
	//Check if chatroom exists
	//Check if user uuid is valid
	//Check if user is already in chatroom
	//Add user to chatroom

	return "", nil

}

func (d *Domain) LeaveRoom(ctx context.Context, chatroom Chatroom) (string, error) {
	//Check if chatroom exists
	//Check if user uuid is valid
	//Check if user is in chatroom
	//Remove user from chatroom

	return "", nil

}

func (d *Domain) GetRoom(ctx context.Context, chatroomUuid string) (Chatroom, error) {

	return Chatroom{}, nil
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
