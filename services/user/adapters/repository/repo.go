package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	mClient *mongo.Client
	rClient *redis.Client
}

func New() (*Db, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(("DB_URI"))))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	a := &Db{mClient: client}
	return a, nil
}

func NewRedis() (*Db, error) {
	otps, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	redisClient := redis.NewClient(otps)

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &Db{
		rClient: redisClient,
	}, nil
}

func (a *Db) NewIndex(database string, collectionName string, field string, unique bool) {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := a.mClient.Database(database).Collection(collectionName)

	index, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Created new index:", index)
}
