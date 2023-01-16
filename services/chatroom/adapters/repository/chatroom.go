package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dietzy1/chatapp/services/chatroom/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
	return a, err

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

func (a *Db) CreateChatroom(ctx context.Context, chatroom domain.Chatroom) error {
	collection := a.client.Database("Chatroom-Database").Collection("Chatrooms")
	_, err := collection.InsertOne(ctx, chatroom)
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) GetChatroom(ctx context.Context, chatroomUuid string) (domain.Chatroom, error) {
	collection := a.client.Database("Chatroom-Database").Collection("Chatrooms")
	var chatroom domain.Chatroom
	err := collection.FindOne(ctx, bson.M{"uuid": chatroomUuid}).Decode(&chatroom)
	if err != nil {
		return chatroom, err
	}
	return chatroom, nil
}

func (a *Db) DeleteChatroom(ctx context.Context, chatroomUuid string) error {
	collection := a.client.Database("Chatroom-Database").Collection("Chatrooms")
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": chatroomUuid})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) AddUserToChatroom(ctx context.Context, chatroomUuid string, userUuid string) error {
	collection := a.client.Database("Chatroom-Database").Collection("Chatrooms")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$push": bson.M{"users": userUuid}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) RemoveUserFromChatroom(ctx context.Context, chatroomUuid string, userUuid string) error {
	collection := a.client.Database("Chatroom-Database").Collection("Chatrooms")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$pull": bson.M{"users": userUuid}})
	if err != nil {
		return err
	}
	return nil
}
