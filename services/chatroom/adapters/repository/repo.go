package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	client *mongo.Client
}

func New() *Db {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(("DB_URI"))))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	a := &Db{client: client}
	return a

}

func (a *Db) newIndex(database string, collectionName string, field string, unique bool) {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := a.client.Database(database).Collection(collectionName)

	index, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Created new index:", index)
}
