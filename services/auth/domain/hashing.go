package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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
