package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/chatroom/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const database = "Chatroom-Database"
const collection = "Chatrooms"

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

// Skips the check if user is in the invited array
func (a *Db) ForceAddUser(ctx context.Context, chatroomUuid string, userUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$push": bson.M{"users": userUuid}})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) ChangeIcon(ctx context.Context, chatroomUuid string, chatroomIcon domain.Icon) error {
	collection := a.client.Database(database).Collection(collection)
	//Use the chatroomUuid to locate the objct we want to update and then insert the domain.Icon object
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": chatroomUuid}, bson.M{"$set": bson.M{"icon": chatroomIcon}})
	if err != nil {
		return err
	}
	return nil
}
