package domain

import (
	"context"
	"errors"
	"log"

	"github.com/dietzy1/chatapp/services/auth/core"
)

type Auth interface {
	Login(ctx context.Context, username string) (string, error)
	Register(ctx context.Context, cred core.Credentials) (string, error)
	Logout(ctx context.Context, userUuid string) error
	Authenticate(ctx context.Context, userUuid string) (string, error)
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
	return &domain{auth: auth, cache: cache}
}

// If someone is trying to login to the application the session token should be returned
func (d domain) Login(ctx context.Context, cred core.Credentials) (string, error) {
	password, err := d.auth.Login(ctx, cred.Username)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//Perform bcrypt check to see if password is correct
	if err = core.CompareHash(password, cred.Password); err != nil {
		log.Println(err)
		return "", err
	}
	//if someone logins then the session token should be regenerated and returned
	token := core.GenerateToken()
	if err = d.auth.UpdateToken(ctx, cred.Username, token); err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (d domain) Register(ctx context.Context, cred core.Credentials) (string, error) {
	//check if username is in database
	_, err := d.auth.Login(ctx, cred.Username)
	if err == nil {
		log.Println(err)
		return "", err
	}

	//hash password
	cred.Password, err = core.GenerateHash(cred.Password)
	if err != nil {
		return "", err
	}

	//add user to database
	cred.Uuid = core.GenerateToken()
	cred.Session = core.GenerateToken()

	session, err := d.auth.Register(ctx, cred)
	if err != nil {
		return "", err
	}
	return session, nil
}

// For now I am unsure if I actually need to do anything with the session token
func (d domain) Logout(ctx context.Context, session string, userUuid string) error {
	uuid, err := d.cache.Get(userUuid)
	if err == nil {
		//check the session token against the database if its there it will be deleted
		uuid, err = d.auth.Authenticate(ctx, userUuid)
		if err != nil {
			log.Println(err)
			return err
		}

	}
	if uuid != session {
		log.Println("invalid session token")
		return errors.New("invalid session token")
	}

	//delete session token from cache
	if err := d.cache.Set(userUuid, ""); err == nil {
		//check the session token against the database if its there it will be deleted
		if err := d.auth.Logout(ctx, userUuid); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// this method needs to be cached for certain it will be under alot of pressure
func (d domain) Authenticate(ctx context.Context, session string, userUuid string) (string, error) {
	// check if token is in cache
	uuid, err := d.cache.Get(userUuid)
	if err != nil {
		//if token is not in cache, check if token is in database
		uuid, err = d.auth.Authenticate(ctx, userUuid)
		if err != nil {
			log.Println(err)
			return "", err

		}
	}
	//perform check to see if the token is the same as the one in the database
	if uuid != session {
		log.Println("invalid session token")
		return "", errors.New("invalid session token")
	}
	//if token is in database, add token to cache
	if err := d.cache.Set(userUuid, session); err != nil {
		log.Println(err)
		return "", err
	}
	return session, nil
}

// This function doesn't care about validating the session token it should always delete
func (d domain) Invalidate(ctx context.Context, userUuid string) error {
	if err := d.cache.Set(userUuid, ""); err != nil {
		log.Println(err)
		return err
	}

	//set the token in the database to an empty string
	if err := d.auth.UpdateToken(ctx, userUuid, ""); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
