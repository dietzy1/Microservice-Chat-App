package websocket

import (
	//"github.com/golang/protobuf/proto"
	"log"

	"google.golang.org/protobuf/proto"

	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
)

//I need some sort of type of struct that can be decoded into protobuf struct

// I might be able to get the type of the message from the generated protobuf file

type Message struct {
	Author       string
	Content      string
	AuthorUuid   string
	ChatRoomUuid string
	Timestamp    string
}

// Take in the object and marshal it into a protobuf struct
func marshal(message *messagev1.CreateMessageRequest) ([]byte, error) {
	/* message := &messagev1.CreateMessageRequest{} */
	log.Println("marshaled msg")
	msg, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// I need to find a way to convert the byte array into something that can be decoded into a protobuf struct

func unmarshal(msg []byte) (*messagev1.CreateMessageRequest, error) {
	log.Println("unmarshaled msg")
	message := &messagev1.CreateMessageRequest{}

	err := proto.Unmarshal(msg, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
