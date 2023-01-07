package websocket

//The websocket should be where the server is listening for messages from the client

//then the server should pass along the information to the application layer

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

type ws struct {
	hub  *ConnectionManager
	conn *websocket.Conn
	send chan Message //this here should prolly be changed to the message object
}

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
		//The issue is the convertion to the message object

		msg := Decode(message)

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

			w, err := ws.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message.Encode(message))
			log.Println(message)

			// Add queued chat messages to the current websocket message.
			n := len(ws.send)
			for i := 0; i < n; i++ {
				ok := <-ws.send
				w.Write(ok.Encode(<-ws.send))
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
