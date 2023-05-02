package domain

import (
	"context"
	"log"

	"github.com/google/uuid"
)

type Domain struct {
	repo MsgRepo
}

func New(repo MsgRepo) *Domain {
	return &Domain{repo: repo}
}

// add the interface to the domain struct
type MsgRepo interface {
	GetMessages(ctx context.Context, chatroomUuid string, channelUuid string) ([]Message, error)
	AddMessage(ctx context.Context, msg Message) error
	UpdateMessage(ctx context.Context, msg Message) error
	DeleteMessage(ctx context.Context, msg Message) error
}

//Implement the methods for the domain layer and create the interface and inject it into the grpc layer

// add bson annotations to the message struct
type Message struct {
	Author       string `json:"author" bson:"author"`
	Content      string `json:"content" bson:"content"`
	AuthorUuid   string `json:"authoruuid" bson:"authoruuid"`
	ChatRoomUuid string `json:"chatroomuuid" bson:"chatroomuuid"`
	ChannelUuid  string `json:"channeluuid" bson:"channeluuid"`
	MessageUuid  string `json:"messageuuid" bson:"messageuuid"`
	Timestamp    string `json:"timestamp" bson:"timestamp"`
}

func (d *Domain) CreateMessage(ctx context.Context, msg Message) (Message, error) {

	//Add the required UUIDS and timestamps
	msg.MessageUuid = uuid.New().String()
	msg.Timestamp = formatTimestamp()

	//Call the repository layer
	err := d.repo.AddMessage(ctx, msg)
	if err != nil {
		log.Println(err)
		return Message{}, err
	}

	return msg, nil
}

func (d *Domain) GetMessages(ctx context.Context, chatroomUuid string, channelUuid string) ([]Message, error) {

	log.Println("Getting messages")
	//Call the repository layer'
	messages, err := d.repo.GetMessages(ctx, chatroomUuid, channelUuid)
	if err != nil {
		log.Println(err)
		return []Message{}, err
	}

	return messages, nil
}

func (d *Domain) EditMessage(ctx context.Context, msg Message) (Message, error) {

	log.Println("domain got hit")
	//The logic here is simple it we simply need to go in and replace the content of the message IDs stay the same
	err := d.repo.UpdateMessage(ctx, msg)
	if err != nil {
		log.Println(err)
		return Message{}, err
	}

	return Message{}, nil
}

func (d *Domain) DeleteMessage(ctx context.Context, msg Message) error {

	err := d.repo.DeleteMessage(ctx, msg)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//The logic I need to happen here is that I need to generate the required UUIDS and then I need to pass on all of the information
//To the repository layer

//Something that needs to be implemented layer is another middle layer inbetween the message service and the API gateway so multiple requests
//For the same channel and chatroom is grouped together into one request to the message service

//I need to figure out how to handle the messaging itself in the respository layer
// -- Chatroom
// -- Channel

//Each server could contain their own collection of chatrooms, and then the channel messages are stored randomly in the chatroom
//But this would require the chatroom to be stored in the channel as well
//Which could be non ideal as it would make it harder to query the chatroom
//TODO: A project for later is reducing the amount of data being sent over the wire in reality it should only be required to send a few select pieces of UUIDS instead of full messages
