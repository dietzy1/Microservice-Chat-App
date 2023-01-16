package core

type ChatRoom struct {
	Users       []string `json:"users" bson:"users"`
	Uuid        string   `json:"uuid" bson:"uuid"`
	Name        string   `json:"name" bson:"name"`
	Owner       string   `json:"owner" bson:"owner"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
}

//copilot what other fields would a chatroom contain?
