package websocket

//The websocket should be where the server is listening for messages from the client

//then the server should pass along the information to the application layer

import (
	"context"
	"fmt"
	"log"
	"time"

	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

type ws struct {
	hub           *ConnectionManager
	conn          *websocket.Conn
	send          chan *messagev1.CreateMessageRequest
	messageClient messagev1.MessageServiceClient
}

//I could implement a write only here which is going to continuesly send updates to the client
//About online // offline users

// I need to implement some error handling so it doesn't crash the server on incorrect proto format
func (ws *ws) readPump() {
	defer func() {
		ws.hub.unregister <- ws
		ws.conn.Close()
	}()
	ws.conn.SetReadLimit(maxMessageSize)
	ws.conn.SetReadDeadline(time.Now().Add(pongWait))
	ws.conn.SetPongHandler(func(string) error { ws.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Println("CLIENT READ", ws.hub.clients)
		msg, err := unmarshal(message)
		if err != nil {
			log.Println(err)
		}
		log.Printf("SENT msg: %s", msg)

		ws.hub.broadcast <- msg

	}
}

func (ws *ws) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		ws.conn.Close()
	}()
	for {
		select {
		case message, ok := <-ws.send:
			ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				ws.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//This was websocket.TextMessage beforehand
			w, err := ws.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			fmt.Println("message: ", message)

			//Here we contact the message service and send the message to the database
			_, err = ws.messageClient.CreateMessage(context.TODO(), message)
			if err != nil {
				//If we are unable to store messages then we shouldn't be able to send messages right
				log.Println(err)
				return
			}

			log.Println("CLIENT WRITE: ", ws.hub.clients)

			log.Printf("RECIEVED msg: %+v", message)
			msg, err := marshal(message)
			if err != nil {
				log.Println(err)
			}
			w.Write(msg)

			// Add queued chat messages to the current websocket message.
			n := len(ws.send)
			for i := 0; i < n; i++ {
				ok := <-ws.send
				msg, err := marshal(ok)
				log.Println(msg)
				if err != nil {
					log.Println(err)
				}

				w.Write(msg)
				//w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		}

	}
}

//When the message object comes in then it should be of type
//type msg struct{
// author string
// message string
// chatroomid string
// authorid string
// }

//then message object should be passed to the application layer and have added a timestamp and uuid to it
//type msg struct{
// author string
// message string
// chatroomid string
// authorid string
// timestamp string
// uuid string
// }
