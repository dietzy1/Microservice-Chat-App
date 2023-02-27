package domain

import (
	"context"
	"errors"
	"fmt"
	"log"

	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

type Auth interface {
	Login(ctx context.Context, username string) (string, error)
	Register(ctx context.Context, cred Credentials) (string, error)
	Unregister(ctx context.Context, userUuid string) error
	Logout(ctx context.Context, userUuid string) error
	Authenticate(ctx context.Context, userUuid string) (string, error)
	UpdateToken(ctx context.Context, username string, token string) (string, error)
}

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type domain struct {
	auth  Auth
	cache Cache

	userClient userClientv1.UserServiceClient
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password " bson:"password"`
	Uuid     string `json:"uuid" bson:"uuid"`
	Session  string `json:"session" bson:"session"`
}

func New(auth Auth, cache Cache, userClient userClientv1.UserServiceClient) *domain {
	return &domain{auth: auth, cache: cache, userClient: userClient}
}

// If someone is trying to login to the application the session token should be returned
func (d domain) Login(ctx context.Context, cred Credentials) (string, error) {
	//Check password in database
	password, err := d.auth.Login(ctx, cred.Username)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//Perform bcrypt check to see if password is correct
	if err = CompareHash(password, cred.Password); err != nil {
		log.Println(err)
		return "", err
	}
	//if someone logins then the session token should be regenerated and returned
	token := GenerateToken()

	uuid, err := d.auth.UpdateToken(ctx, cred.Username, token)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//add the user to the cache
	if err = d.cache.Set(uuid, token); err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (d domain) Register(ctx context.Context, cred Credentials) (string, error) {
	//check if username is in database
	_, err := d.auth.Login(ctx, cred.Username)
	if err == nil {
		log.Printf("username %s already exists", cred.Username)
		return "", errors.New("username already exists")
	}

	//perform check if password is of atleast 8 characters
	if len(cred.Password) < 8 {
		log.Println("password is too short")
		return "", errors.New("password is too short")
	}

	//hash password
	cred.Password, err = GenerateHash(cred.Password)
	if err != nil {
		return "", err
	}

	cred.Uuid = GenerateToken()
	cred.Session = GenerateToken()
	//add user to database
	session, err := d.auth.Register(ctx, cred)
	if err != nil {
		return "", err
	}

	//Call the user service to create a user

	_, err = d.userClient.CreateUser(ctx, &userClientv1.CreateUserRequest{
		Username: cred.Username,
		UserUuid: cred.Uuid,
	})
	if err != nil {
		log.Println(err)

		//if the user service fails to create a user then the user should be deleted from the auth service
		if err := d.auth.Unregister(ctx, cred.Uuid); err != nil {
			log.Println(err)
			return "", err
		}
		return "", err
	}

	return session, nil
}

// For now I am unsure if I actually need to do anything with the session token
func (d domain) Logout(ctx context.Context, session string, userUuid string) error {
	token, err := d.cache.Get(userUuid)
	if err == nil {
		//check the session token against the database if its there it will be deleted
		token, err = d.auth.Authenticate(ctx, userUuid)
		if err != nil {
			log.Println(err)
			return err
		}

	}
	if token != session {
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
	//Check if token is in cache
	token, err := d.cache.Get(userUuid)
	if err != nil {
		fmt.Println("cache miss")
		//if token is not in cache, check if token is in database
		token, err = d.auth.Authenticate(ctx, userUuid)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}

	//perform check to see if the token is the same as the one in the database
	if token != session {
		log.Println("invalid session token")
		return "", errors.New("invalid session token")
	}

	//if token is in database, add token to cache
	if err := d.cache.Set(userUuid, session); err != nil {
		log.Println(err)
		return "", err
	}
	log.Println("Valid session token")

	return session, nil
}

// This function doesn't care about validating the session token it should always delete
func (d domain) Invalidate(ctx context.Context, userUuid string) error {
	if err := d.cache.Set(userUuid, ""); err != nil {
		log.Println(err)
		return err
	}

	//set the token in the database to an empty string
	_, err := d.auth.UpdateToken(ctx, userUuid, "")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
