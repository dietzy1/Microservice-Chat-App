package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/icon/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const iconDatabase = "Icon-Database"
const iconCollection = "Icons"

const emojiType = "emoji"

// insert icon into database
func (a *Db) StoreIcon(ctx context.Context, icon domain.Icon) error {
	collection := a.client.Database(iconDatabase).Collection(iconCollection)
	_, err := collection.InsertOne(ctx, icon)
	if err != nil {
		return err
	}
	return nil
}

// TODO: veryify validatiy of this
// Retrieve all icons from database associated with owner
func (a *Db) GetIcons(ctx context.Context, ownerUuid string) ([]domain.Icon, error) {
	collection := a.client.Database(iconDatabase).Collection(iconCollection)
	Icons := []domain.Icon{}

	cursor, err := collection.Find(ctx, bson.M{"ownerUuid": ownerUuid})
	if err != nil {
		return Icons, err
	}
	err = cursor.All(ctx, &Icons)
	if err != nil {
		return Icons, err
	}
	return Icons, nil
}

// TODO: verify validity of this
// Retrieve all icons that are emojis
func (a *Db) GetEmojiIcons(ctx context.Context) ([]domain.Icon, error) {
	collection := a.client.Database(iconDatabase).Collection(iconCollection)
	Icons := []domain.Icon{}

	cursor, err := collection.Find(ctx, bson.M{"kindOf": emojiType})
	if err != nil {
		return Icons, err
	}
	err = cursor.All(ctx, &Icons)
	if err != nil {
		return Icons, err
	}
	return Icons, nil
}

// get icon based on uuid
func (a *Db) GetIcon(ctx context.Context, uuid string) (domain.Icon, error) {
	//Get icon from database
	collection := a.client.Database(iconDatabase).Collection(iconCollection)
	Icon := domain.Icon{}

	err := collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&Icon)
	if err != nil {
		return Icon, err
	}
	return Icon, nil
}

func (a *Db) DeleteIcon(ctx context.Context, uuid string) error {
	//Delete icon from database
	collection := a.client.Database(iconDatabase).Collection(iconCollection)

	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}
