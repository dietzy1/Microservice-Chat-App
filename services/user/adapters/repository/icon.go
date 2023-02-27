package repository

import (
	"context"

	"github.com/dietzy1/chatapp/services/user/domain"
	"go.mongodb.org/mongo-driver/bson"
)

const iconDatabase = "Icon-Database"
const iconCollection = "Icons"

// insert icon into database
func (a *Db) StoreIcon(ctx context.Context, icon domain.Icon) error {
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)
	_, err := collection.InsertOne(ctx, icon)
	if err != nil {
		return err
	}
	return nil
}

// get icon based on uuid
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

func (a *Db) DeleteIcon(ctx context.Context, uuid string) error {
	//Delete icon from database
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)

	_, err := collection.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		return err
	}
	return nil
}

// look through the icon database and return a random icon
func (a *Db) GetRandomIcon(ctx context.Context) (domain.Icon, error) {
	collection := a.mClient.Database(iconDatabase).Collection(iconCollection)
	Icon := domain.Icon{}

	err := collection.FindOne(ctx, bson.M{}).Decode(&Icon)
	if err != nil {
		return Icon, err
	}
	return Icon, nil
}

//TODO: test validity of the GetRandomIcon function
//the following code is from the imageapi project and might be more correct
/* cursor, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": 1}}})
if err != nil {
	return nil, err
}
if err = cursor.All(ctx, &images); err != nil {
	return nil, err
}
return &images[0], nil */
