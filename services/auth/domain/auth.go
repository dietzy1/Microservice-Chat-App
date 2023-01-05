package domain

import (
	"context"

	"github.com/dietzy1/chatapp/services/auth/core"
)

type Auth interface {
	Login(ctx context.Context, username string) (string, error)
	Logout(ctx context.Context, username string) error
	Authenticate(ctx context.Context, token string) (string, error)
	UpdateToken(ctx context.Context, username string, token string) error
}

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type domain struct {
	auth  Auth
	cache Cache
}

func New(auth Auth, cache Cache) *domain {
	return &domain{auth, cache}
}

// If someone is trying to login to the application the session token should be returned
func (d *domain) Login(ctx context.Context) (string, error) {
	//recieve data from GRPC
	username := "grpc username"
	inputPassword := "grpc password"

	//perform check against database to see if user exists
	password, err := d.auth.Login(ctx, username)
	if err != nil {
		//handle error
		//Unable to find a user with that username return error back to GRPC gateway
	}

	//Perform bcrypt check to see if password is correct
	if err = core.CompareHash(password, inputPassword); err != nil {
		//handle error
		//Password was incorrect return error back to GRPC gateway
	}
	//if someone logins then the session token should be regenerated and returned
	token := core.GenerateToken()
	if err = d.auth.UpdateToken(ctx, username, token); err != nil {
		//handle error
		//Unable to update token in database return error back to GRPC gateway
	}
	return token, nil
}

func (d domain) Logout(ctx context.Context) error {
	username := "grpc username"

	if err := d.auth.Logout(ctx, username); err != nil {
		//handle error
	}
	return nil
}

func (d domain) Authenticate(ctx context.Context) (string, error) {
	// extract token from GRPC
	token := "grpc token"

	// check if token is in cache
	uuid, err := d.cache.Get(token)
	if err != nil {

		//if token is not in cache, check if token is in database
		uuid, err = d.auth.Authenticate(ctx, token)
		if err != nil {
			//handle error
			//token was not found in database return error back to GRPC gateway
		}

		//if token is in database, add token to cache
		if err := d.cache.Set(token, uuid); err != nil {
			//handle error
			//unable to add token to cache return error back to GRPC gateway
		}
	}
	return uuid, nil
}
