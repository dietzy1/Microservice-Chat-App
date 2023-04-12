package websocket

import (
	//"github.com/golang/protobuf/proto"
	"log"

	"google.golang.org/protobuf/proto"

	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
)

func marshal(message *messagev1.CreateMessageResponse) ([]byte, error) {

	msg, err := proto.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return nil, err
	}
	return msg, nil
}

func unmarshal(msg []byte) (*messagev1.CreateMessageRequest, error) {

	message := &messagev1.CreateMessageRequest{}

	err := proto.Unmarshal(msg, message)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return nil, err
	}
	log.Println(message)
	return message, nil
}
