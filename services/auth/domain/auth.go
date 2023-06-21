package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/dietzy1/chatapp/pkg/hashing"
	"go.uber.org/zap"
)

// If someone is trying to login to the application the session token should be returned
func (d domain) Login(ctx context.Context, cred Credentials) (Credentials, error) {

	// Verify username and password isn't empty
	if cred.Username == "" || cred.Password == "" {
		d.logger.Info("username or password is empty")
		return Credentials{}, errors.New("username or password is empty")
	}

	//Check password in database
	password, err := d.auth.Login(ctx, cred.Username)
	if err != nil {
		d.logger.Info(fmt.Sprintf("username %s not found", cred.Username), zap.Error(err))
		return Credentials{}, fmt.Errorf("username %s not found", cred.Username)
	}

	//Perform bcrypt check to see if password is correct
	if err = hashing.CompareHash(password, cred.Password); err != nil {
		d.logger.Info(fmt.Sprintf("password for username %s is incorrect", cred.Username), zap.Error(err))
		return Credentials{}, fmt.Errorf("password for username %s is incorrect", cred.Username)
	}
	//if someone logins then the session token should be regenerated and returned
	token := hashing.GenerateToken()

	uuid, err := d.auth.UpdateToken(ctx, cred.Username, token)
	if err != nil {
		d.logger.Info(fmt.Sprintf("error updating token for username %s", cred.Username), zap.Error(err))
		return Credentials{}, fmt.Errorf("error updating token for username %s", cred.Username)
	}

	//add the user to the cache
	if err = d.cache.Set(uuid, token); err != nil {
		d.logger.Info(fmt.Sprintf("error adding user to cache for username %s", cred.Username), zap.Error(err))
		return Credentials{}, fmt.Errorf("error adding user to cache for username %s", cred.Username)
	}

	return Credentials{Session: token, Uuid: uuid}, nil
}

// For now I am unsure if I actually need to do anything with the session token
func (d domain) Logout(ctx context.Context, session string, userUuid string) error {

	if session == "" || userUuid == "" {
		d.logger.Info("session or user uuid is empty")
		return errors.New("session or user uuid is empty")
	}

	token, err := d.cache.Get(userUuid)
	if err == nil {
		//check the session token against the database if its there it will be deleted
		token, err = d.auth.Authenticate(ctx, userUuid)
		if err != nil {
			d.logger.Info(fmt.Sprintf("error authenticating user %s", userUuid), zap.Error(err))
			return fmt.Errorf("error authenticating user %s", userUuid)
		}

	}

	if token != session {
		d.logger.Info("invalid session token when logging out")
		return errors.New("invalid session token when logging out")
	}

	//delete session token from cache
	if err := d.cache.Set(userUuid, ""); err == nil {
		//check the session token against the database if its there it will be deleted
		if err := d.auth.Logout(ctx, userUuid); err != nil {
			d.logger.Info(fmt.Sprintf("error logging out user %s", userUuid), zap.Error(err))
			return fmt.Errorf("error logging out user %s", userUuid)
		}
	}
	return nil
}

// this method needs to be cached for certain it will be under alot of pressure
func (d domain) Authenticate(ctx context.Context, session string, userUuid string) (Credentials, error) {

	if session == "" || userUuid == "" {
		d.logger.Info("session or user uuid is empty")
		return Credentials{}, errors.New("session or user uuid is empty")
	}

	//Check if token is in cache
	token, err := d.cache.Get(userUuid)
	if err != nil {
		d.logger.Info(fmt.Sprintf("cache miss for user %s", userUuid), zap.Error(err))

		//if token is not in cache, check if token is in database
		token, err = d.auth.Authenticate(ctx, userUuid)
		if err != nil {
			d.logger.Info(fmt.Sprintf("error authenticating user %s", userUuid), zap.Error(err))
			return Credentials{}, fmt.Errorf("error authenticating user %s", userUuid)
		}
	}

	//perform check to see if the token is the same as the one in the database
	if token != session {
		d.logger.Info("invalid session token when authenticating")
		return Credentials{}, errors.New("invalid session token")
	}

	//if token is in database, add token to cache
	if err := d.cache.Set(userUuid, session); err != nil {
		d.logger.Info(fmt.Sprintf("error adding user to cache for user %s", userUuid), zap.Error(err))
		return Credentials{}, fmt.Errorf("error adding user to cache for user %s", userUuid)
	}
	d.logger.Info(fmt.Sprintf("user %s authenticated", userUuid))

	return Credentials{Session: session, Uuid: userUuid}, nil
}

// This function doesn't care about validating the session token it should always delete
func (d domain) Invalidate(ctx context.Context, userUuid string) error {

	if userUuid == "" {
		d.logger.Info("user uuid is empty")
		return errors.New("user uuid is empty")
	}

	if err := d.cache.Set(userUuid, ""); err != nil {
		d.logger.Info(fmt.Sprintf("error invalidating user %s", userUuid), zap.Error(err))
		return fmt.Errorf("error invalidating user %s", userUuid)
	}

	//set the token in the database to an empty string
	_, err := d.auth.UpdateToken(ctx, userUuid, "")
	if err != nil {
		d.logger.Info(fmt.Sprintf("error invalidating user %s", userUuid), zap.Error(err))
		return fmt.Errorf("error invalidating user %s", userUuid)

	}
	return nil
}
