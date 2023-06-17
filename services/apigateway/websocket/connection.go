package websocket

import (
	"log"
	"sync"
	"time"

	chatroomv1 "github.com/dietzy1/chatapp/services/apigateway/chatroomgateway/v1"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
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
	conn             *websocket.Conn
	sendChannel      chan []byte
	receiveChannel   chan []byte
	activeChannel    chan []string
	heartbeatChannel chan []byte
	shutdown         chan any
	cleanupOnce      *sync.Once

	logger *zap.Logger

	lastPongAt time.Time
}

type connOptions struct {
	conn *websocket.Conn
}

func newConnection(o *connOptions) *conn {
	return &conn{
		conn:             o.conn,
		sendChannel:      make(chan []byte, sendBufferSize),
		receiveChannel:   make(chan []byte, recieveBufferSize),
		activeChannel:    make(chan []string, recieveBufferSize),
		heartbeatChannel: make(chan []byte, recieveBufferSize),
		shutdown:         make(chan interface{}),
		cleanupOnce:      &sync.Once{},
	}
}

func (c *conn) writePump() {
	c.logger.Info("Write pump started")

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.logger.Info("Write pump stopped")
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

				c.logger.Info("Send channel closed")
				return
			}

			if err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
				c.logger.Error("Failed to set write deadline", zap.Error(err))
				return
			}

			if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {

				c.logger.Error("Failed to write message", zap.Error(err))
				return
			}
			c.logger.Info("Message sent")

		case active, ok := <-c.activeChannel:
			if !ok {
				return
			}
			c.logger.Info("Active channel received", zap.Any("active", active))

			//Construct a protobuf message of activity
			activity := &chatroomv1.Activity{
				OnlineUsers: active,
			}

			//Marshal the protobuf message
			marshaled, err := marshalActivity(activity)
			if err != nil {
				c.logger.Error("Failed to marshal activity", zap.Error(err))
				return
			}

			c.logger.Info("Activity marshaled", zap.Any("marshaled", marshaled))
			//Send the array of active users to the client

			err = c.conn.WriteMessage(websocket.BinaryMessage, marshaled)
			if err != nil {
				c.logger.Error("Failed to write message to client", zap.Error(err))
				return
			}

		case _, ok := <-c.heartbeatChannel:
			if !ok {
				return
			}
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
				c.logger.Error("Failed to set write deadline", zap.Error(err))
				return
			}

			log.Println("conn: ", c.conn.RemoteAddr().String())
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				c.logger.Error("Failed to write ping message", zap.Error(err))
				return
			}

		}
	}
}

func (c *conn) readPump() {
	c.logger.Info("Read pump started")
	defer func() {
		c.logger.Info("Read pump stopped")
		c.cleanup()
	}()

	c.conn.SetReadLimit(maxInboundMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(idleTimeout)); err != nil {
		c.logger.Error("Failed to set read deadline", zap.Error(err))
	}

	//Here I need to implement some sort of ping pong mechanism to keep the connection alive
	c.conn.SetPongHandler(func(input string) error {

		c.logger.Info("Received pong", zap.String("input", input))

		c.lastPongAt = time.Now()
		if err := c.conn.SetReadDeadline(time.Now().Add(idleTimeout)); err != nil {
			c.logger.Error("Failed to set read deadline", zap.Error(err))
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

			c.logger.Info("Received message", zap.Any("message", bs))
			select {
			case c.receiveChannel <- bs:

			default:
				// If the receive channel is full, we should assume the client is disconnected
				// and close the connection.
				c.logger.Info("Receive channel full")
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

				c.logger.Info("Websocket connection closed unexpectedly")
			} else {
				// Otherwise we failed to read from the connection, increment metric for failure
				c.logger.Error("Failed to read from connection", zap.Error(err))
			}
			return
		}
	}
}

func (c *conn) heartBeat() {

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.logger.Info("Heartbeat stopped cleaning up the connection")
		ticker.Stop()
		c.cleanup()
	}()

	for {
		select {
		case <-ticker.C:
			c.logger.Info("Sending heartbeat")
			c.heartbeatChannel <- []byte{}

		case <-c.shutdown:

			c.logger.Info("Forcing shutdown")
			return
		}
	}
}

func (c *conn) cleanup() {

	c.cleanupOnce.Do(func() {

		c.activeChannel <- []string{}
		close(c.shutdown)
		close(c.receiveChannel)
		close(c.sendChannel)
		//FIXME: maybe this
		//close(c.activeChannel)
		close(c.heartbeatChannel)

		// Close the underlying websocket connection.
		err := c.conn.Close()
		if err != nil {

			c.logger.Error("Failed to close connection", zap.Error(err))
		}
		c.logger.Info("Connection closed")
	})

}

func (c *conn) run() {
	c.logger.Info("New connection starting read, write and heartbeat pumps")

	go c.readPump()
	go c.writePump()
	go c.heartBeat()

}
