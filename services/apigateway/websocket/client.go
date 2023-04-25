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

			//case active, ok := <-c.conn.activity.activityChannel:
			/* case active, ok := <-c.conn.activeChannel:
				if !ok {
					return
				}
				log.Println("ACTIVE CHANNEL", active)

				//Construct a protobuf message of activity
				activity := &chatroomv1.Activity{
					OnlineUsers: active,
				}

				//Marshal the protobuf message
				marshaled, err := marshalActivity(activity)
				if err != nil {
					log.Println("Failed to marshal activity")
					return
				}

				log.Println("SENDING ACTIVITY", activity)
				//Send the array of active users to the client

				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				//FIXME: THIS IS CAUSING A DATARACE
				err = c.conn.conn.WriteMessage(websocket.BinaryMessage, marshaled)
				if err != nil {
					log.Println("Failed to write message to client")
					return
				}

			} */
		}
	}

}

func (c *client) updateClientActivity(chatroom string) {
	// return array of who is active in the chatroom
	active := c.conn.activity.active[chatroom]
	//Take the slice and convert it to a byte array
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

	//Close the pubsub channel
	c.broker.unsubscribe(context.TODO(), pubsub)
}
