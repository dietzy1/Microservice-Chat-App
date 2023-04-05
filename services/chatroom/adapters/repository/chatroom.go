package repository

import (
	"context"
	"fmt"
	"log"
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

const database = "Chatroom-Database"
const collection = "Chatrooms"

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

func (a *Db) CreateChatroom(ctx context.Context, chatroom domain.Chatroom) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.InsertOne(ctx, chatroom)
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) DeleteChatroom(ctx context.Context, chatroomUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": chatroomUuid})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) GetChatroom(ctx context.Context, chatroomUuid string) (domain.Chatroom, error) {
	log.Println(chatroomUuid)

	collection := a.client.Database(database).Collection(collection)
	var chatroom domain.Chatroom
	err := collection.FindOne(ctx, bson.M{"uuid": chatroomUuid}).Decode(&chatroom)
	if err != nil {
		return chatroom, err
	}
	return chatroom, nil
}

func (a *Db) GetChatrooms(ctx context.Context, chatroomUuids []string) ([]domain.Chatroom, error) {
	collection := a.client.Database(database).Collection(collection)
	var chatrooms []domain.Chatroom
	//Retrieve all documents where uuid is in chatroomUuids
	cursor, err := collection.Find(ctx, bson.M{"uuid": bson.M{"$in": chatroomUuids}})
	if err != nil {
		return chatrooms, err
	}
	//Decode all documents into chatrooms
	err = cursor.All(ctx, &chatrooms)
	if err != nil {
		return chatrooms, err
	}
	return chatrooms, nil

}

// ---------------------------------------------------------------------------

func (a *Db) CreateChannel(ctx context.Context, channel domain.Channel) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": channel.ChatroomUuid}, bson.M{"$push": bson.M{"channels": channel}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) DeleteChannel(ctx context.Context, channelUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"channels.uuid": channelUuid}, bson.M{"$pull": bson.M{"channels": bson.M{"uuid": channelUuid}}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) GetChannel(ctx context.Context, channelUuid string) (domain.Channel, error) {
	collection := a.client.Database(database).Collection(collection)
	var channel domain.Channel
	err := collection.FindOne(ctx, bson.M{"channels.uuid": channelUuid}).Decode(&channel)
	if err != nil {
		return channel, err
	}
	return channel, nil
}

// ---------------------------------------------------------------------------

func (a *Db) InviteUser(ctx context.Context, chatroomUuid string, userUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$push": bson.M{"invited": userUuid}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) RemoveUser(ctx context.Context, chatroomUuid string, userUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$pull": bson.M{"users": userUuid}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) AddUser(ctx context.Context, chatroomUuid string, userUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	//Perform check if user is in the invited array
	count, err := collection.CountDocuments(ctx, bson.M{"uuid": chatroomUuid, "invited": userUuid})
	if err != nil || count == 0 {
		return err
	}

	//Remove user from invited array
	_, err = collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$pull": bson.M{"invited": userUuid}})
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$push": bson.M{"users": userUuid}})
	if err != nil {
		return err
	}
	return nil
}
