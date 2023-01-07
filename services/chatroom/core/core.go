package core

type ChatRoom struct {
	Users []User `json:"users" bson:"users"`
	Uuid  string `json:"uuid" bson:"uuid"`
	Name  string `json:"name" bson:"name"`
	Owner string `json:"owner" bson:"owner"`
}
