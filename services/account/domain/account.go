package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/dietzy1/chatapp/pkg/hashing"
	"go.uber.org/zap"

	chatroomClientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

func (d domain) ChangeUsername(ctx context.Context, userUuid string, username string) error {

	if userUuid == "" || username == "" {
		d.logger.Info("userUuid or username is empty")
		return errors.New("userUuid or username is empty")
	}

	// check if username is in database
	err := d.repo.CheckUsername(ctx, username)
	if err != nil {
		d.logger.Info("username already exists %v", zap.Error(err))
		return fmt.Errorf("username %s already exists", username)
	}

	// change username in database
	err = d.repo.ChangeUsername(ctx, userUuid, username)
	if err != nil {
		d.logger.Info("error changing username %v", zap.Error(err))
		return fmt.Errorf("error changing username %s", err)
	}

	return nil
}

func (d domain) ChangePassword(ctx context.Context, userUuid string, password string, newPassword string) error {

	//retrieve password from database
	storedPassword, err := d.repo.CheckPassword(ctx, userUuid)
	if err != nil {
		d.logger.Info("error retrieving password %v", zap.Error(err))
		return fmt.Errorf("error retrieving password for %s", userUuid)
	}

	//check if password is correct
	if err = hashing.CompareHash(storedPassword, password); err != nil {
		d.logger.Info("password is incorrect %v", zap.Error(err))
		return fmt.Errorf("password is incorrect for %s", userUuid)
	}

	//hash new password
	hashedPassword, err := hashing.GenerateHash(newPassword)
	if err != nil {
		d.logger.Info("error hashing password %v", zap.Error(err))
		return fmt.Errorf("error hashing password %s for %s", err, userUuid)
	}

	//update password in database
	err = d.repo.UpdatePassword(ctx, userUuid, hashedPassword)
	if err != nil {
		d.logger.Info("error changing password %v", zap.Error(err))
		return fmt.Errorf("error changing password %s for %s", err, userUuid)
	}

	return nil
}

func (d domain) DeleteAccount(ctx context.Context, userUuid string, password string) error {

	if userUuid == "" || password == "" {
		d.logger.Info("userUuid or password is empty")
		return errors.New("userUuid or password is empty")
	}

	//retrieve password from database
	storedPassword, err := d.repo.CheckPassword(ctx, userUuid)
	if err != nil {
		d.logger.Info("error retrieving password %v", zap.Error(err))
		return fmt.Errorf("error retrieving password for %s", userUuid)
	}

	//check if password is correct
	if err = hashing.CompareHash(storedPassword, password); err != nil {
		d.logger.Info("password is incorrect %v", zap.Error(err))
		return fmt.Errorf("password is incorrect for %s", userUuid)
	}

	//delete account from database
	err = d.repo.DeleteAccount(ctx, userUuid)
	if err != nil {
		d.logger.Info("error deleting account %v", zap.Error(err))
		return fmt.Errorf("error deleting account for %s", userUuid)
	}

	return nil
}

func (d domain) RegisterAccount(ctx context.Context, cred Credentials) (Credentials, error) {

	if cred.Username == "" || cred.Password == "" {
		d.logger.Info("username or password is empty")
		return Credentials{}, errors.New("username or password is empty")
	}

	//check if username is in database
	err := d.repo.CheckUsername(ctx, cred.Username)
	if err != nil {
		d.logger.Info("username already exists %v", zap.Error(err))
		return Credentials{}, fmt.Errorf("username %s already exists", cred.Username)
	}

	//perform check if password is of atleast 8 characters
	if len(cred.Password) < 8 {
		d.logger.Info("password is too short")
		return Credentials{}, errors.New("password is too short")
	}

	//hash password
	cred.Password, err = hashing.GenerateHash(cred.Password)
	if err != nil {
		d.logger.Info("error hashing password %v", zap.Error(err))
		return Credentials{}, fmt.Errorf("error hashing password %s", err)
	}

	cred.Uuid = hashing.GenerateToken()
	cred.Session = hashing.GenerateToken()
	cred.DemoUser = false
	//add user to database
	session, err := d.repo.Register(ctx, cred)
	if err != nil {
		d.logger.Info("error registering user %v", zap.Error(err))
		return Credentials{}, fmt.Errorf("error registering user %s", err)
	}

	//Call the user service to create a user

	_, err = d.userClient.CreateUser(ctx, &userClientv1.CreateUserRequest{
		Username: cred.Username,
		UserUuid: cred.Uuid,
	})
	if err != nil {
		d.logger.Info("error creating user %v", zap.Error(err))

		//if the user service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			d.logger.Info("error deleting user %v", zap.Error(err))
			return Credentials{}, fmt.Errorf("error deleting user %s", err)
		}
		return Credentials{}, fmt.Errorf("error creating user %s", err)
	}

	//Call the chatroom service to add the user to the default chatroom
	//TODO:this here is causing the issue

	_, err = d.chatroomClient.ForceAddUser(ctx, &chatroomClientv1.ForceAddUserRequest{
		UserUuid: cred.Uuid,
	})
	if err != nil {
		d.logger.Info("error adding user to default chatroom %v", zap.Error(err))

		//if the chatroom service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			d.logger.Info("error deleting user %v", zap.Error(err))
			return Credentials{}, fmt.Errorf("error deleting user %s", err)
		}
		return Credentials{}, fmt.Errorf("error adding user to default chatroom %s", err)

		//Potentially need to add delete user from user service
	}

	return Credentials{Session: session, Uuid: cred.Uuid}, nil
}

func (d domain) DemoUserRegister(ctx context.Context) (Credentials, error) {
	cred := Credentials{
		Username: hashing.GenerateName(),
		Password: "",
	}

	// check if username is in database
	err := d.repo.CheckUsername(ctx, cred.Username)
	if err != nil {
		d.logger.Info("username already exists %v", zap.Error(err))
		cred.Username = hashing.GenerateName()
	}

	cred.Uuid = hashing.GenerateToken()
	cred.Session = hashing.GenerateToken()
	cred.DemoUser = true
	// add user to database
	session, err := d.repo.Register(ctx, cred)
	if err != nil {
		d.logger.Info("error registering user %v", zap.Error(err))
		return Credentials{}, fmt.Errorf("error registering user %s", err)
	}

	//Call the user service to create a user
	_, err = d.userClient.CreateUser(ctx, &userClientv1.CreateUserRequest{
		Username: cred.Username,
		UserUuid: cred.Uuid,
	})
	if err != nil {
		d.logger.Info("error creating user %v", zap.Error(err))

		//if the user service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			d.logger.Info("error deleting user %v", zap.Error(err))
			return Credentials{}, err
		}
		return Credentials{}, err
	}

	//Add user to main chatroom
	_, err = d.chatroomClient.ForceAddUser(ctx, &chatroomClientv1.ForceAddUserRequest{
		UserUuid: cred.Uuid,
	})
	if err != nil {
		d.logger.Info("error adding user to default chatroom %v", zap.Error(err))

		//if the chatroom service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			d.logger.Info("error deleting user %v", zap.Error(err))
			return Credentials{}, fmt.Errorf("error deleting user %s", err)
		}
		return Credentials{}, fmt.Errorf("error adding user to default chatroom %s", err)

		//Potentially need to add delete user from user service
	}

	return Credentials{Session: session, Uuid: cred.Uuid}, nil
}

func (d domain) UpgradeDemoUser(ctx context.Context, creds Credentials) error {

	if creds.Uuid == "" || creds.Password == "" || creds.Username == "" {
		d.logger.Info("uuid, password or username is empty")
		return errors.New("uuid, password or username is empty")
	}

	//Validate password is long enough
	if len(creds.Password) < 8 {
		d.logger.Info("password is too short")
		return errors.New("password is too short")
	}

	//hash password
	hashedPassword, err := hashing.GenerateHash(creds.Password)
	if err != nil {
		d.logger.Info("error hashing password %v", zap.Error(err))
		return fmt.Errorf("error hashing password %s", err)
	}

	creds.Password = hashedPassword
	creds.DemoUser = false

	//Update user in database
	err = d.repo.UpdateAccount(ctx, creds)
	if err != nil {
		d.logger.Info("error updating user %v", zap.Error(err))
		return fmt.Errorf("error updating user %s", err)
	}

	return nil
}
