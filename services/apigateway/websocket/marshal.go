package websocket

//I need some sort of type of struct that can be decoded into protobuf struct

// I might be able to get the type of the message from the generated protobuf file

type Message struct {
	Author       string
	Content      string
	AuthorUuid   string
	ChatRoomUuid string
	Timestamp    string
}

// I need to find a way to convert the byte array into something that can be decoded into a protobuf struct
