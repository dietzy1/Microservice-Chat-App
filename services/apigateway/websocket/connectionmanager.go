package websocket

import (
	"log"
	"net/http"
	"os"

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
	clients    map[*ws]bool //this part should be changed to a redis database
	broadcast  chan Message //Before this was []byte
	register   chan *ws
	unregister chan *ws
}

func Start() {
	connectionManager := NewConnectionManger()
	go connectionManager.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	err := http.ListenAndServe(os.Getenv("WS"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//It might be the solution to use a map of channels where the channel id is the key for broadcasting channel

// Might need to inject some dependencies here
func NewConnectionManger() *ConnectionManager {
	return &ConnectionManager{
		broadcast:  make(chan Message), //before this was []byte //Now it should actually be a slice of broadcast channels
		register:   make(chan *ws),     //should be a slice of register channels
		unregister: make(chan *ws),     //should be a slice of unregister channels
		clients:    make(map[*ws]bool), //should be a slice of clients
	}
}

// Might need to somehow take in which chatroom the client is in
func serveWs( /* cm *ConnectionManager, */ w http.ResponseWriter, r *http.Request) {
	opts := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := opts.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//vars := chi.URLParam(r, "guildID")

	cm := make(map[string]*ConnectionManager)
	if _, ok := cm[vars]; !ok {
		cm[vars] = New()
		go cm[vars].Run()
	}

	ws := &ws{hub: cm[vars], conn: conn, send: make(chan Message, 256)}
	ws.hub.register <- ws

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	//this is the entry point for the map that  returns the connection manager

	go ws.writePump()
	go ws.readPump()
}

//Usually communication would happen through interfaces
//But in this case all communication is done through channels

//Hub is the central point of the websocket server
//Is in charge of registering, unregistering and broadcasting messages to all clients
//Should send messages to the application layer to request that data is stored in the database

//Should perhabs call on the application layer to perform verification of the user/data

func (cm *ConnectionManager) Run() {
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
					//Perform call to application to store message in database

				default:
					close(client.send)
					delete(cm.clients, client)
				}
			}
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
