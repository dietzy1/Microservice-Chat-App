package websocket

import (
	"context"
	"log"

	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"github.com/go-redis/redis/v8"
)

type client struct {
	conn *conn
	id   id
	//Passed down from manager
	broker Broker

	//passed down from manager
	messageClient messagev1.MessageServiceClient
}

type clientOptions struct {
	conn          *conn
	id            id
	broker        Broker
	messageClient messagev1.MessageServiceClient
}

func newClient(o *clientOptions) *client {
	return &client{

		conn:          o.conn,
		id:            o.id,
		broker:        o.broker,
		messageClient: o.messageClient,
	}
}

func (c *client) handleMessages(ch <-chan *redis.Message) {
	for {
		select {
		//Messages recieved from the recieve channel is from the user itself
		case msg, ok := <-c.conn.receiveChannel:
			if !ok {
				return
			}
			//Unmarshal byte into protobuf message
			message, err := unmarshal(msg)
			if err != nil {
				log.Println("Failed to unmarshal message")
				return
			}

			//Call the message service to create a message
			resp, err := c.messageClient.CreateMessage(context.TODO(), message)
			if err != nil {
				log.Println("Failed to create message")
				return
			}

			message1, err := marshal(resp)
			if err != nil {
				log.Println("Failed to unmarshal message")
				return
			}

			c.broker.Publish(context.TODO(), c.id.channel, message1)

			//check if userID is same as Author ID -- to avoid duplicate messages

			if c.id.user != message.AuthorUuid {
				c.conn.sendChannel <- message1
			}

		case msg, ok := <-ch:
			if !ok {
				return
			}

			//convert msg.Payload to slice of bytes
			c.conn.sendChannel <- []byte(msg.Payload)

		}
	}

}

func (c *client) updateClientActivity(chatroom string, active []string) {

	log.Println("UPDATING CLIENT ACTIVITY")
	c.conn.activeChannel <- active

}

func (c *client) run() {
	//Start read, write and heartbeat goroutines
	c.conn.run()

	pubsub, err := c.broker.Subscribe(context.TODO(), c.id.channel)
	if err != nil {
		log.Println("Failed to subcribe to channelID:", c.id.channel)
		log.Println(err)
		return
	}

	//Blocking call to handle messages
	c.handleMessages(pubsub.Channel())

	log.Println("Client run stopped")

	//Close the pubsub channel
	c.broker.unsubscribe(context.TODO(), pubsub)
}
