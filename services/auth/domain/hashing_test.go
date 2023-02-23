package domain

import (
	"log"
	"testing"
)

func TestCompareHash(t *testing.T) {
	password := "test"
	//Generate a hash from the password
	hash, err := GenerateHash(password)
	if err != nil {
		log.Fatal(err)
	}
	//Compare the hash
	err = CompareHash(hash, password)
	if err != nil {
		log.Fatal(err)
	}

	err = CompareHash(hash, "test2")
	if err == nil {
		log.Fatal("Error password does not match")
	}

}

/* func GenerateHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CompareHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
} */
