package domain

import (
	"context"

	"github.com/dietzy1/chatapp/services/auth/core"
	"github.com/google/uuid"
)

type Account interface {
	Register(ctx context.Context, creds core.Credentials) error
	Delete(ctx context.Context, creds core.Credentials) error
	ChangeUsername(ctx context.Context, creds core.Credentials) error
	ChangePassword(ctx context.Context, creds core.Credentials) error
	CheckUsername(ctx context.Context, username string) bool
	CheckPassword(ctx context.Context, uuid string) (string, error)
}

/* type Credentials struct {
	Username string
	Password string
	Uuid     string
	Token    string
} */

func (a Domain) Register(ctx context.Context) {
	//recieve password and and username from grpc
	username := "username"
	password := "password"

	// Perform check against database if username is taken
	if a.acc.CheckUsername(ctx, username) {
		return //Should return error code that entry already exists
	}
	//if not taken create a new user
	hashPassword, err := core.GenerateHash(password)
	if err != nil {
		return //Should return error code that hash failed
	}
	creds := core.Credentials{
		Username: username,
		Password: hashPassword, //needs to add bcrypt hash
		Uuid:     uuid.New().String(),
	}
	a.acc.Register(ctx, creds)

	//Should do a call to the database to create a new user
	// needs to be passed down username and uuid

}

func (a Domain) Delete(ctx context.Context) {
	creds := core.Credentials{
		Username: "username",
		Password: "password",
	}

	a.acc.Delete(ctx, creds)
	a.user.DeleteUser(ctx, creds.Uuid)
}

// recieve password and username from grpc
func (a Domain) ChangeUsername(ctx context.Context) {
	creds := core.Credentials{
		Username: "username",
		Password: "password",
	}
	//perform check if username is taken
	if a.acc.CheckUsername(ctx, creds.Username) {
		return //Should return error code that entry already exists
	}

	a.acc.ChangeUsername(ctx, creds)
}

func (a Domain) ChangePassword(ctx context.Context) {
	creds := core.Credentials{
		Username: "username",
		Password: "password",
	}
	hashPassword, err := a.acc.CheckPassword(ctx, creds.Uuid)
	if err != nil {
		return
	}
	if core.CompareHash(creds.Password, hashPassword) {
		return //Should return error code that password does not match
	}

	//Must perform check if original password match

	a.acc.ChangePassword(ctx, creds)
}
