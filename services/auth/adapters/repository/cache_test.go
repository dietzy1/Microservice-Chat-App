package repository

import (
	"log"
	"testing"
)

type redismock struct {
	client *cache
}

// Mock test constructor
func newRedisMock() *redismock {
	redis, err := newCache()
	if err != nil {
		log.Fatal(err)
	}
	return &redismock{client: redis}
}

func TestGet(t *testing.T) {
	testKey := "testKey"
	testValue := "testValue"
	testIncorrectKey := "testIncorrectKey"

	redismock := newRedisMock()
	err := redismock.client.Set(testKey, testValue)
	if err != nil {
		t.Error(err, "Error setting value in redis")
	}

	//test key which exists
	value, err := redismock.client.Get("test")
	if err != nil {
		t.Errorf("Error getting value from redis: %s", err)
	}

	if value != testValue {
		t.Errorf("Expected %s, got %s", testValue, value)
	}

	// test key which doesn't exist
	value, err = redismock.client.Get(testIncorrectKey)
	if err != nil {
		t.Errorf("Error getting value from redis: %s", err)
	}

	if value != "" {
		t.Errorf("Expected %s, got %s", "", value)
	}

}

func TestSet(t *testing.T) {
	testKey := "testKey"
	testValue := "testValue"

	redismock := newRedisMock()
	err := redismock.client.Set(testKey, testValue)
	if err != nil {
		t.Error(err, "Error setting value in redis")
	}
}
