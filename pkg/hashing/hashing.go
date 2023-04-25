package hashing

import (
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

// First parameter is the hash, second is the password
func CompareHash(hashedPassword, unhashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(unhashedPassword))
}

func GenerateToken() string {
	return uuid.New().String()
}

func GenerateName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return nameGenerator.Generate()
}
