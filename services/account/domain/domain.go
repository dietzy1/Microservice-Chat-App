package domain

import (
	"context"
	"errors"
	"log"
)

type domain struct {
	repo Repo
}

type Repo interface {
	CheckUsername(ctx context.Context, username string) error
	ChangeUsername(ctx context.Context, userUuid string, username string) error
	CheckPassword(ctx context.Context, userUuid string) (string, error)
	UpdatePassword(ctx context.Context, userUuid string, password string) error
	DeleteAccount(ctx context.Context, userUuid string) error
}

func New(repo Repo) *domain {
	return &domain{repo: repo}
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password " bson:"password"`
	Uuid     string `json:"uuid" bson:"uuid"`
	Session  string `json:"session" bson:"session"`
}

func (d *domain) ChangeUsername(ctx context.Context, userUuid string, username string) error {
	// check if username is in database
	err := d.repo.CheckUsername(ctx, username)
	if err != nil {
		log.Printf("username %s already exists", username)
		return errors.New("username already exists")
	}

	// change username in database
	err = d.repo.ChangeUsername(ctx, userUuid, username)
	if err != nil {
		log.Printf("error changing username %s", err)
		return errors.New("error changing username")
	}

	return nil
}

func (d *domain) ChangePassword(ctx context.Context, userUuid string, password string, newPassword string) error {

	//retrieve password from database
	storedPassword, err := d.repo.CheckPassword(ctx, userUuid)
	if err != nil {
		log.Printf("error retrieving password %s", err)
		return errors.New("error retrieving password")
	}

	//check if password is correct
	if err = CompareHash(storedPassword, password); err != nil {
		log.Println(err)
		return errors.New("password is incorrect")
	}

	//hash new password
	hashedPassword, err := GenerateHash(newPassword)
	if err != nil {
		log.Printf("error hashing password %s", err)
		return errors.New("error hashing password")
	}

	//update password in database
	err = d.repo.UpdatePassword(ctx, userUuid, hashedPassword)
	if err != nil {
		log.Printf("error changing password %s", err)
		return errors.New("error changing password")
	}

	return nil
}

func (d *domain) DeleteAccount(ctx context.Context, userUuid string, password string) error {

	//retrieve password from database
	storedPassword, err := d.repo.CheckPassword(ctx, userUuid)
	if err != nil {
		log.Printf("error retrieving password %s", err)
		return errors.New("error retrieving password")
	}

	//check if password is correct
	if err = CompareHash(storedPassword, password); err != nil {
		log.Println(err)
		return errors.New("password is incorrect")
	}

	//delete account from database
	err = d.repo.DeleteAccount(ctx, userUuid)
	if err != nil {
		log.Printf("error deleting account %s", err)
		return errors.New("error deleting account")
	}

	return nil
}
