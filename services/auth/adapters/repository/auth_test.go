package repository

import (
	"context"
	"log"
	"testing"

	"github.com/dietzy1/chatapp/services/auth/core"
	"go.mongodb.org/mongo-driver/bson"
)

type mongoMock struct {
	client *auth
}

func newMongoMock() *mongoMock {
	mongo, err := newAuth()
	if err != nil {
		log.Fatal(err)
	}
	return &mongoMock{client: mongo}
}

func (m *mongoMock) createMockData() core.Credentials {

	mockData := core.Credentials{
		Username: "test",
		Password: "test",
		Uuid:     "test",
		Token:    "test",
	}
	//Add the mock data to the database
	collection := m.client.client.Database("Credential-Database").Collection("Credentials")
	_, err := collection.InsertOne(context.Background(), mockData)
	if err != nil {
		log.Fatal(err)
	}
	return mockData
}

func (m *mongoMock) deleteMockData(mockData core.Credentials) {
	// delete the mock data from the database
	collection := m.client.client.Database("Credential-Database").Collection("Credentials")
	_, err := collection.DeleteOne(context.Background(), mockData)
	if err != nil {
		log.Fatal(err)
	}
}

func TestLogin(t *testing.T) {
	mongoMock := newMongoMock()
	mockData := mongoMock.createMockData()

	// I need to verify that the password returned is the same as the password in the mock data based on the input username
	password, err := mongoMock.client.Login(context.Background(), mockData.Username)
	if err != nil {
		t.Errorf("Error unable to retrieve password: %v", err)
	}
	if password != mockData.Password {
		t.Errorf("Error password does not match: %v", err)
	}
	// delete the mock data from the database
	mongoMock.deleteMockData(mockData)
}

func TestAuthenticate(t *testing.T) {
	mongoMock := newMongoMock()
	mockData := mongoMock.createMockData()

	//I need to verify that the username returned is the same as the username in the mock data based on the input token
	username, err := mongoMock.client.Authenticate(context.Background(), mockData.Token)
	if err != nil {
		t.Errorf("Error unable to retrieve username: %v", err)
	}
	if username != mockData.Username {
		t.Errorf("Error username does not match: %v", err)
	}
	// delete the mock data from the database
	mongoMock.deleteMockData(mockData)
}

func TestLogout(t *testing.T) {
	mongoMock := newMongoMock()
	mockData := mongoMock.createMockData()

	//I need to verify that the token returned is empty
	err := mongoMock.client.Logout(context.Background(), mockData.Username)
	if err != nil {
		t.Errorf("Error unable to logout: %v", err)
	}

	if err := mongoMock.getToken(mockData.Token); err == nil {
		t.Errorf("Error token is not empty: %v, expected token to be %s", err, mockData.Token)
	}

	// delete the mock data from the database
	mongoMock.deleteMockData(mockData)
}

// Helper mock function to test logout and test update token
func (m *mongoMock) getToken(token string) error {
	// get the token from the database
	collection := m.client.client.Database("Credential-Database").Collection("Credentials")
	// take in username and use that to update the token to empty
	var result core.Credentials
	if err := collection.FindOne(context.Background(), bson.M{"token": token}).Decode(&result); err != nil {
		return err
	}
	return nil
}

func TestUpdateToken(t *testing.T) {
	mongoMock := newMongoMock()
	mockData := mongoMock.createMockData()

	//I need to verify that the token is updated to the new token based on the username input
	if err := mongoMock.client.UpdateToken(context.Background(), mockData.Username, "newToken"); err != nil {
		t.Errorf("Error unable to update token: %v", err)
	}

	if err := mongoMock.getToken("newToken"); err != nil {
		t.Errorf("Error token is not updated: %v, expected token to be %s", err, "newToken")
	}
	// delete the mock data from the database
	mongoMock.deleteMockData(mockData)
}
