package websocket

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/dietzy1/chatapp/services/apigateway/clients"
	messageclientv1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"github.com/gorilla/websocket"
)

type manager struct {
	//Key is userID
	clients map[string]*client

	//Key is chadroomID
	active map[string][]string

	upgrader websocket.Upgrader
	mu       sync.RWMutex

	//Redis pubsub service
	broker Broker

	//Message service client
	messageClient messageclientv1.MessageServiceClient
}

func newManager(broker Broker, messageClient messageclientv1.MessageServiceClient) *manager {
	return &manager{
		clients:       make(map[string]*client),
		active:        make(map[string][]string),
		mu:            sync.RWMutex{},
		broker:        broker,
		messageClient: messageClient,

		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

type activity struct {
	active          map[string][]string
	activityChannel chan []byte
	mu              sync.RWMutex
}

type id struct {
	chatroom string
	channel  string
	user     string
}

func (m *manager) upgradeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket connection established")
	id := id{
		chatroom: r.URL.Query().Get("chatroom"),
		channel:  r.URL.Query().Get("channel"),
		user:     r.URL.Query().Get("user"),
	}
	if id.chatroom == "" || id.channel == "" || id.user == "" {
		log.Println("Missing chatroom, user or channel")
		return
	}

	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection", err)
		return
	}

	ac := &activity{
		active:          m.active,
		activityChannel: make(chan []byte),
		mu:              sync.RWMutex{},
	}

	ws := newConnection(&connOptions{conn: conn, activity: ac})

	client := newClient(&clientOptions{
		conn:          ws,
		id:            id,
		broker:        m.broker,
		messageClient: m.messageClient,
	},
	)

	m.addClient(client, &id)
	client.updateClientActivity(id.chatroom)
	defer m.removeClient(client, &id)
	defer client.updateClientActivity(id.chatroom)

	client.run()

	//Right now I should have the active count for every chatroom+channel
	//So that means I need to count the active users for each combination of chatroom+channel
}

func (m *manager) addClient(c *client, id *id) {
	m.mu.Lock()
	defer m.mu.Unlock()
	log.Println(id)

	m.clients[id.user] = c
	m.active[id.chatroom] = append(m.active[id.chatroom], id.user)

	log.Println("Active Users: ", m.active)
}

func (m *manager) removeClient(c *client, id *id) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, id.user)

	ok := m.active[id.chatroom]
	//Remove element from slice that contains user id
	//TODO: veryfy that this logic works
	for i, v := range ok {
		if v == id.user {
			m.active[id.chatroom] = append(ok[:i], ok[i+1:]...)
		}
	}

	//remove element from slice that contains user id

	log.Println("Removing users, active users left: ", m.active)

}

func Start() {

	broker := newBroker()

	//New messageService
	messageClient := clients.NewMessageClient()

	//manager should accept the broker interface
	m := newManager(broker, *messageClient)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Websocket connection request received")
		m.upgradeHandler(w, r)
	})

	err := http.ListenAndServe(os.Getenv("WS"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Println("Websocket server started on port " + os.Getenv("WS"))
}
