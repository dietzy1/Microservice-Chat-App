package repository

import (
	"context"
	"os"
	"time"

	"github.com/dietzy1/chatapp/services/auth/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type auth struct {
	client *mongo.Client
}

func newAuth() (*auth, error) {
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

	a := &auth{client: client}
	return a, nil
}

// if login is called, the token is generated and the user is logged in
func (a *auth) Login(ctx context.Context, username string) (string, error) {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	var cred core.Credentials
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&cred)
	if err != nil {
		return cred.Password, err
	}
	return cred.Password, nil
}

// if logout is called, the token is invalidated and the user is logged out
func (a *auth) Logout(ctx context.Context, userUuid string) error {
	// perform update to the session token in the database
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	// take in username and use that to update the token to empty
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": userUuid}, bson.M{"$set": bson.M{"session": ""}})
	if err != nil {
		return err
	}
	return nil
}

func (a *auth) Authenticate(ctx context.Context, userUuid string) (string, error) {
	collection := a.client.Database("Credential-Database").Collection("Credentials")
	var cred core.Credentials
	err := collection.FindOne(ctx, bson.M{"uuid": userUuid}).Decode(&cred)
	if err != nil {
		//maybe I need to return uuid instead here?
		return cred.Session, err
	}
	return cred.Username, nil
}

func (a *auth) UpdateToken(ctx context.Context, userUuid string, session string) error {

	collection := a.client.Database("Credential-Database").Collection("Credentials")
	// take in username and use that to update the token to empty
	_, err := collection.UpdateOne(ctx, bson.M{"uuid": userUuid}, bson.M{"$set": bson.M{"session": session}})
	if err != nil {
		return err
	}
	return nil
}
