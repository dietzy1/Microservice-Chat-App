package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/user/core"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *Db) AddUser(ctx context.Context, user core.User) error {
	collection := a.mClient.Database("credentials").Collection("public")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// Should take in field of what needs to be changed and the new value
func (a *Db) ChangeUser(ctx context.Context, uuid string, key string, value string) error {
	collection := a.mClient.Database("credentials").Collection("public")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		return err
	}
	return nil
}

// Authentication is done prior so it should not be nessescary to do agai
func (a *Db) DeleteUser(ctx context.Context, uuid string) error {
	collection := a.mClient.Database("credentials").Collection("public")
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}

// Idk if this actually works needs to check
func (a *Db) UpdateChatServers(ctx context.Context, uuid string, server string) error {
	collection := a.mClient.Database("credentials").Collection("public")
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$push": bson.M{"chatservers": server}})
	if err != nil {
		return err
	}
	return nil
}
