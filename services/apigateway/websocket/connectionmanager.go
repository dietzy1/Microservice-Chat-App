package websocket

import (
	"log"
	"net/http"
	"os"

	client "github.com/dietzy1/chatapp/services/apigateway/clients"
	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"github.com/gorilla/websocket"
)

//TODO:
//I think the proper solution to the problem
//is to have a map that returns a struct of connectionManger struct
//if the key doesn't exist then the new connection manager should be called
//if the key already exists then the connection manager struct should be returned
//So its possible to perform operations on the channel
//make (map[string]ConnectionManager)
//Should probaly create a seperate Run() goroutine for each seperate ConnectionManger struct that is created
//logic stays contained in websocket package and should not need to be exposed to the main package

type ConnectionManager struct {
	clients    map[*ws]bool
	broadcast  chan *messagev1.CreateMessageRequest //Before this was []byte
	register   chan *ws
	unregister chan *ws
	//TODO: experimental
	activeCheck chan *ws
}

// Initialize the connection manager
func Start() {
	connectionManager := NewConnectionManger()

	go connectionManager.Run()

	//Create a message client
	messageClient := client.NewMessageClient()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Websocket connection request received")
		serveWs(w, r, *messageClient)
	})

	err := http.ListenAndServe(os.Getenv("WS"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Println("Websocket server started on port " + os.Getenv("WS"))

}

//It might be the solution to use a map of channels where the channel id is the key for broadcasting channel

// Might need to inject some dependencies here
func NewConnectionManger() *ConnectionManager {
	return &ConnectionManager{
		broadcast:  make(chan *messagev1.CreateMessageRequest),
		register:   make(chan *ws),
		unregister: make(chan *ws),
		clients:    make(map[*ws]bool),
		//TODO: experimental
		activeCheck: make(chan *ws),
	}
}

type id struct {
	chatroom string
	channel  string
}

// TODO: this solution is really cursed
// Might need to somehow take in which chatroom the client is in
func serveWs(w http.ResponseWriter, r *http.Request, client messagev1.MessageServiceClient) {
	log.Println("Websocket connection established")
	opts := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	id := &id{
		chatroom: r.URL.Query().Get("chatroom"),
		channel:  r.URL.Query().Get("channel"),
	}

	conn, err := opts.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	cm := make(map[string]*ConnectionManager)
	if _, ok := cm[id.chatroom+id.channel]; !ok {
		cm[id.chatroom+id.channel] = NewConnectionManger()
		go cm[id.chatroom+id.channel].Run()
	}

	ws := &ws{hub: cm[id.chatroom+id.channel], conn: conn, send: make(chan *messagev1.CreateMessageRequest, 256), messageClient: client}
	ws.hub.register <- ws

	go ws.writePump()
	go ws.readPump()
}

func (cm *ConnectionManager) Run() {
	log.Println("Run function called")
	for {
		//if a select statement is ready to run it will
		select {

		case Client := <-cm.register:
			cm.clients[Client] = true

		case client := <-cm.unregister:
			if _, ok := cm.clients[client]; ok {
				delete(cm.clients, client)
				close(client.send)
			}

		case message := <-cm.broadcast:
			for client := range cm.clients {
				select {
				case client.send <- message:

				default:
					close(client.send)
					delete(cm.clients, client)
				}
			}
			//Here I could add a case for state changes to the online/offline clients
			//Ideally the state itself should be contained within the map itself since its a map of clients with a boolean value
			/* case stateChange := <-cm.activeCheck:
			//If stateChanged update the state of the client

			//If an event is recieved here we should send a message to the client
			for client := range cm.clients {
				select {
				case client.send <- message:

				default:
					close(client.send)
					delete(cm.clients, client)
				}

			} */
		}

	}

}

//https://pkg.go.dev/github.com/golang/protobuf/jsonpb

//https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson

//Registration involves the underlying structs

//registering is equal to subsribing to output from a specific server

//if a user is logged in client side then the client has stored information about uuid username and chatrooms
//This data should be used to subscribe to the correct websocket server

/* type User struct {
	Uuid      string   `json:"uuid" bson:"uuid"`
	Username  string   `json:"username" bson:"username"`
	ChatRooms []string `json:"chatrooms" bson:"chatrooms"`
}

type ChatRoom struct {
	Users []User
	Uuid  string
	Name  string
} */

//I need to figure out how to manually convert a json string to a protofile
//I think I need to use the protofile to convert the json string to a protofile
