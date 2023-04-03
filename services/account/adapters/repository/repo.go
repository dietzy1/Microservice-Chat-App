package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dietzy1/chatapp/services/auth/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const database = "Credential-Database"
const collection = "Credentials"

type Db struct {
	client *mongo.Client
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

	a := &Db{client: client}
	return a, nil
}

func (a *Db) NewIndex(database string, collectionName string, field string, unique bool) {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := a.client.Database(database).Collection(collectionName)

	index, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Created new index:", index)
}

func (a *Db) CheckUsername(ctx context.Context, username string) error {
	collection := a.client.Database(database).Collection(collection)
	var result domain.Credentials
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) ChangeUsername(ctx context.Context, userUuid string, username string) error {
	collection := a.client.Database(database).Collection(collection)
	err := collection.FindOneAndUpdate(ctx, bson.M{"uuid": userUuid}, bson.M{"$set": bson.M{"username": username}}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) CheckPassword(ctx context.Context, userUuid string) (string, error) {
	collection := a.client.Database(database).Collection(collection)
	var result domain.Credentials
	err := collection.FindOne(ctx, bson.M{"uuid": userUuid}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Password, nil
}

func (a *Db) UpdatePassword(ctx context.Context, userUuid string, password string) error {
	collection := a.client.Database(database).Collection(collection)
	err := collection.FindOneAndUpdate(ctx, bson.M{"uuid": userUuid}, bson.M{"$set": bson.M{"password": password}}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) DeleteAccount(ctx context.Context, userUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	err := collection.FindOneAndDelete(ctx, bson.M{"uuid": userUuid}).Err()
	if err != nil {
		return err
	}
	return nil
}
