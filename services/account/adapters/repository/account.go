package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/account/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const database = "Credential-Database"
const collection = "Credentials"

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

func (a *Db) Register(ctx context.Context, cred domain.Credentials) (string, error) {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.InsertOne(ctx, cred)
	if err != nil {
		return "", err
	}
	return cred.Session, nil
}

func (a *Db) Unregister(ctx context.Context, userUuid string) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.DeleteOne(ctx, bson.M{"uuid": userUuid})
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) UpdateAccount(ctx context.Context, cred domain.Credentials) error {
	collection := a.client.Database(database).Collection(collection)
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": cred.Uuid}, bson.M{"$set": cred})
	if err != nil {
		return err
	}
	return nil
}
