package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/user/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const iconDatabase = "Icon-Database"
const iconCollection = "Icons"

func (a *Db) StoreIcon(ctx context.Context, icon domain.Icon) error {
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)
	_, err := collection.InsertOne(ctx, icon)
	if err != nil {
		return err
	}
	return nil
}

func (a *Db) GetIcon(ctx context.Context, uuid string) (domain.Icon, error) {
	//Get icon from database
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)
	Icon := domain.Icon{}

	err := collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&Icon)
	if err != nil {
		return Icon, err
	}
	return Icon, nil
}

func (a *Db) UpdateIcon(ctx context.Context, icon domain.Icon) error {
	//Update icon in database
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)

	_, err := collection.UpdateOne(ctx, bson.M{"uuid": icon.Uuid}, bson.M{"$set": bson.M{"icon": icon}})
	if err != nil {
		return err
	}

	return nil
}

func (a *Db) DeleteIcon(ctx context.Context, uuid string) error {
	//Delete icon from database
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)

	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}
