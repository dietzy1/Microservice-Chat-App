package core

// Public
/* type Message struct {
	Message      string `json:"message" bson:"message"`
	User         User   `json:"user" bson:"user"`
	ChatRoomUuid string `json:"chatroomuuid" bson:"chatroomuuid"`
	Time         string `json:"time" bson:"time"`
} */

//Need to reimplement these types

// Public
type User struct {
	Username  string   `json:"username" bson:"username"`
	Uuid      string   `json:"uuid" bson:"uuid"` //Provided by auth0
	ChatRooms []string `json:"chatrooms" bson:"chatrooms"`
}

type ChatRoom struct {
	Users []User `json:"users" bson:"users"`
	Uuid  string `json:"uuid" bson:"uuid"`
	Name  string `json:"name" bson:"name"`
	Owner string `json:"owner" bson:"owner"`
}

type Message struct {
	Author       string `json:"author" bson:"author"`
	Message      string `json:"message" bson:"message"`
	Authoruuid   string `json:"authoruuid" bson:"authoruuid"`
	ChatRoomUuid string `json:"chatroomuuid" bson:"chatroomuuid"`
	Timestamp    string `json:"timestamp,omitempty" bson:"timestamp"`
}

/* func Encode(msg Message) []byte {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	return jsonData
}

func Decode(jsonData []byte) Message {
	msg := Message{}
	err := json.Unmarshal([]byte(jsonData), &msg)
	if err != nil {
		log.Println(err)
	}
	return msg
} */
