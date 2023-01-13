package core

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password " bson:"password"`
	Uuid     string `json:"uuid" bson:"uuid"`
	Session  string `json:"session" bson:"session"`
}

func GenerateHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CompareHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateToken() string {
	return uuid.New().String()
}
