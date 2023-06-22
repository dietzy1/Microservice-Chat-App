package websocket

import (
	"bytes"
	"context"
	"encoding/gob"

	messagev1 "github.com/dietzy1/chatapp/services/message/proto/message/v1"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type client struct {
	conn *conn
	id   id

	logger *zap.Logger
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
	logger        *zap.Logger
}

func newClient(o *clientOptions) *client {
	return &client{

		conn:          o.conn,
		id:            o.id,
		broker:        o.broker,
		messageClient: o.messageClient,
		logger:        o.logger,
	}
}

func (c *client) handleMessages(ch <-chan *redis.Message, activityCh <-chan *redis.Message) {
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
				c.logger.Error("Failed to unmarshal message", zap.Error(err))
				return
			}

			//Call the message service to create a message
			resp, err := c.messageClient.CreateMessage(context.TODO(), message)
			if err != nil {
				c.logger.Error("Failed to create message", zap.Error(err))
				return
			}

			message1, err := marshal(resp)
			if err != nil {
				c.logger.Error("Failed to unmarshal message", zap.Error(err))
				return
			}

			c.broker.Publish(context.TODO(), c.id.channel, message1)
			if err != nil {
				c.logger.Error("Failed to publish message", zap.Error(err))
				return
			}

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

			//FIXME: this part here is not getting hit on connection
		case msg, ok := <-activityCh:
			if !ok {
				return
			}
			c.logger.Info("Recieved activity message", zap.String("message", msg.Payload))

			buf := []byte(msg.Payload)

			strs2 := []string{}
			gob.NewDecoder(bytes.NewReader(buf)).Decode(&strs2)

			c.logger.Info("Decoded activity message", zap.Strings("message", strs2))

			c.conn.activeChannel <- strs2

			//c.conn.activeChannel <- msg.PayloadSlice

		}

	}

}

// FIXME:
func (c *client) updateClientActivity(chatroom string, active []string) {

	c.logger.Info("Updating client activity", zap.String("chatroom", chatroom), zap.Strings("active", active))

	//I need to find out how to convert from []string to []byte
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(active)

	c.broker.Publish(context.TODO(), chatroom, buf.Bytes())
}

func (c *client) run(chatroom string, active []string) {
	//Start read, write and heartbeat goroutines
	c.conn.run()

	pubsub, err := c.broker.Subscribe(context.TODO(), c.id.channel)
	if err != nil {
		c.logger.Error("Failed to subscribe to channel", zap.Error(err))
		return
	}

	activityPubsub, err := c.broker.Subscribe(context.TODO(), c.id.chatroom)
	if err != nil {
		c.logger.Error("Failed to subscribe to channel", zap.Error(err))

		return
	}

	//Send updated activity
	c.updateClientActivity(c.id.chatroom, active)

	//Blocking call to handle messages
	c.handleMessages(pubsub.Channel(), activityPubsub.Channel())

	c.logger.Info("Client run stopped")

	//Close the pubsub channel
	err = c.broker.unsubscribe(context.TODO(), pubsub)
	if err != nil {
		c.logger.Error("Failed to unsubscribe from channel", zap.Error(err))
		return
	}
	pubsub.Close()

	err = c.broker.unsubscribe(context.TODO(), activityPubsub)
	if err != nil {
		c.logger.Error("Failed to unsubscribe from channel", zap.Error(err))
		return
	}
	activityPubsub.Close()

	c.logger.Info("Client unsubscribed from pubsub channels")
	//list

}

//The way I think about activity is flawed
//I need to implement it using redis pubsub so the user subscribes to a server and then acvitity is published to that server
