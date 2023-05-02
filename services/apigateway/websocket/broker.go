package websocket

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// PubSub is an interface for a Redis Pub/Sub client.
type Broker interface {
	Subscribe(ctx context.Context, channel string) (*redis.PubSub, error)
	unsubscribe(ctx context.Context, pubsub *redis.PubSub) error
	Publish(ctx context.Context, channel string, message []byte) error
}

// RedisPubSub is an implementation of the PubSub interface using the go-redis library.
type broker struct {
	client *redis.Client
}

// NewRedisPubSub creates a new RedisPubSub instance.
func newBroker() *broker {
	otps, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("Failed to parse redis url: %v", err)
		return nil
	}
	client := redis.NewClient(otps)
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Printf("Failed to ping redis: %v", err)
		return nil
	}

	return &broker{
		client: client,
	}
}

// Subscribe subscribes to a channel and returns a PubSub instance for receiving messages.
func (b *broker) Subscribe(ctx context.Context, channel string) (*redis.PubSub, error) {
	pubsub := b.client.Subscribe(ctx, channel)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, err
	}

	return pubsub, nil
}

// Send message to channel
func (b *broker) Publish(ctx context.Context, channel string, message []byte) error {
	err := b.client.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (b *broker) unsubscribe(ctx context.Context, pubsub *redis.PubSub) error {
	err := pubsub.Unsubscribe(ctx)
	if err != nil {
		return err
	}
	log.Println("Unsubscribed from channel")
	return nil

}
