package repository

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func newCache() (*cache, error) {
	otps, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	redisClient := redis.NewClient(otps)

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &cache{
		client: redisClient,
	}, nil
}

func (c *cache) Get(key string) (string, error) {
	val, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *cache) Set(key string, value string) error {
	return c.client.Set(context.Background(), key, value, 0).Err()
}
