package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/user/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const userDatabase = "User-Database"
const UserCollection = "Users"

func (a *Db) AddUser(ctx context.Context, user domain.User) error {
	collection := a.mClient.Database(userDatabase).Collection(UserCollection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// TODO: verify that this works
// Function that locates a user by uuid and then adds the serveruuid to the array of chatservers
func (a *Db) AddChatServer(ctx context.Context, uuid string, serveruuid string) error {
	collection := a.mClient.Database(userDatabase).Collection(UserCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$push": bson.M{"chatservers": serveruuid}})
	if err != nil {
		return err
	}
	return nil
}

// TODO: veryify that this works
// Functuon that locates a user by uuid and then it removes the serveruuid from the array of chatservers
func (a *Db) RemoveChatServer(ctx context.Context, uuid string, serveruuid string) error {
	collection := a.mClient.Database(userDatabase).Collection(UserCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$pull": bson.M{"chatservers": serveruuid}})
	if err != nil {
		return err
	}
	return nil
}

// Should take in field of what needs to be changed and the new value
/* func (a *Db) ChangeUser(ctx context.Context, uuid string, key string, value string) error {
	collection := a.mClient.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		return err
	}
	return nil
} */

// Authentication is done prior so it should not be nessescary to do again
/* func (a *Db) DeleteUser(ctx context.Context, uuid string) error {
	collection := a.mClient.Database(database).Collection(collection)
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
} */
