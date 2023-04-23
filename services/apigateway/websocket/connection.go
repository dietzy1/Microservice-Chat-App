package websocket

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// wait time for message sends to succeed.
	writeTimeout = 10 * time.Second
	// close connections where we haven't received a ping in `idleTimeout`.
	idleTimeout = 70 * time.Second
	// How often we ping clients.
	pingPeriod = 30 * time.Second
	// Max size of inbound message, in bytes.
	maxInboundMessageSize = 40 * 1024
	// Max number of messages queued in send buffer.
	sendBufferSize = 200
	// Max number of messages queued in receive buffer.
	recieveBufferSize = 200
)

type conn struct {
	conn           *websocket.Conn
	sendChannel    chan []byte
	receiveChannel chan []byte
	activeChannel  chan []string
	shutdown       chan any
	cleanupOnce    *sync.Once
	activity       *activity
	lastPongAt     time.Time
}

type connOptions struct {
	conn     *websocket.Conn
	activity *activity
}

func newConnection(o *connOptions) *conn {
	return &conn{
		conn:           o.conn,
		sendChannel:    make(chan []byte, sendBufferSize),
		receiveChannel: make(chan []byte, recieveBufferSize),
		activeChannel:  make(chan []string, recieveBufferSize),
		shutdown:       make(chan interface{}),
		cleanupOnce:    &sync.Once{},
		activity:       o.activity,
	}
}

func (c *conn) writePump() {
	log.Println("Write pump started")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		log.Println("Write pump stopped")
		ticker.Stop()
		c.cleanup()
	}()

	for {
		select {
		case <-c.shutdown:
			return
		case msg, ok := <-c.sendChannel:
			// If the channel is closed, we should stop the goroutine.
			if !ok {
				log.Println("Send channel closed")
				return
			}

			if err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
				log.Println("Failed to set write deadline")
				return
			}

			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				log.Println("Failed to write message")
				return
			}
			log.Println("Sent message")

		}
	}
}

func (c *conn) readPump() {

	log.Println("Read pump started")
	defer func() {
		log.Println("Read pump stopped")
		c.cleanup()
	}()

	c.conn.SetReadLimit(maxInboundMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(idleTimeout)); err != nil {
		log.Println("Failed to set read deadline")
	}

	//Here I need to implement some sort of ping pong mechanism to keep the connection alive
	c.conn.SetPongHandler(func(input string) error {
		log.Println("Received pong", input)

		c.lastPongAt = time.Now()
		if err := c.conn.SetReadDeadline(time.Now().Add(idleTimeout)); err != nil {

			log.Println("Failed to set read deadline", err)
			c.cleanup()
		}
		return nil
	})

	errC := make(chan error, 1)

	// Start reading the incoming messages from the websocket connection.
	go func() {
		for {
			_, bs, err := c.conn.ReadMessage()
			if err != nil {
				errC <- err
				return
			}

			log.Println("Received message")
			select {
			case c.receiveChannel <- bs:

			default:
				// If the receive channel is full, we should assume the client is disconnected
				// and close the connection.
				log.Println("Receive channel full")
				go c.cleanup()
				return
			}
		}
	}()

	for {
		select {
		case <-c.shutdown:
			// Increment metric for closing
			return
		case err := <-errC:
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// If the connection is closed, increment metric for closing

				log.Println("Websocket connection closed unexpectedly")
			} else {
				// Otherwise we failed to read from the connection, increment metric for failure

				log.Println("THIS ERROR IS IMPORTANT: ", err)
				log.Println("Failed to read from connection")
			}
			return
		}
	}
}

func (c *conn) heartBeat() {

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		log.Println("Heartbeat stopped cleaning up the connection")
		ticker.Stop()
		c.cleanup()
	}()

	for {
		select {
		case <-ticker.C:

			if err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
				log.Println("Failed to set write deadline")
				return
			}

			log.Println("conn: ", c.conn.RemoteAddr().String())
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("Failed to write ping message")
				return
			}
		case <-c.shutdown:
			log.Println("Forcing shutdown")
			return
		}
	}
}

func (c *conn) cleanup() {

	c.cleanupOnce.Do(func() {
		close(c.shutdown)
		close(c.receiveChannel)
		close(c.sendChannel)

		// Close the underlying websocket connection.
		err := c.conn.Close()
		if err != nil {
			log.Println("Failed to close connection")
		}
		log.Println("Connection closed")
	})

}

func (c *conn) run() {
	log.Println("New connection")

	go c.readPump()
	go c.writePump()
	go c.heartBeat()

}
