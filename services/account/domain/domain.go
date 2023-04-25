package domain

import (
	"context"
	"errors"
	"log"

	"github.com/dietzy1/chatapp/pkg/hashing"

	chatroomClientv1 "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1"
	userClientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

type domain struct {
	repo Repo

	userClient     userClientv1.UserServiceClient
	chatroomClient chatroomClientv1.ChatroomServiceClient
}

type Repo interface {
	CheckUsername(ctx context.Context, username string) error
	ChangeUsername(ctx context.Context, userUuid string, username string) error
	CheckPassword(ctx context.Context, userUuid string) (string, error)
	UpdatePassword(ctx context.Context, userUuid string, password string) error
	DeleteAccount(ctx context.Context, userUuid string) error
	Register(ctx context.Context, cred Credentials) (string, error)
	Unregister(ctx context.Context, userUuid string) error
	UpdateAccount(ctx context.Context, cred Credentials) error
}

func New(repo Repo, userClient userClientv1.UserServiceClient, chatroomClient chatroomClientv1.ChatroomServiceClient) *domain {
	return &domain{repo: repo, userClient: userClient, chatroomClient: chatroomClient}
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password " bson:"password"`
	Uuid     string `json:"uuid" bson:"uuid"`
	Session  string `json:"session" bson:"session"`
	DemoUser bool   `json:"demouser" bson:"demouser"`
}

func (d domain) ChangeUsername(ctx context.Context, userUuid string, username string) error {
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

func (d domain) ChangePassword(ctx context.Context, userUuid string, password string, newPassword string) error {

	//retrieve password from database
	storedPassword, err := d.repo.CheckPassword(ctx, userUuid)
	if err != nil {
		log.Printf("error retrieving password %s", err)
		return errors.New("error retrieving password")
	}

	//check if password is correct
	if err = hashing.CompareHash(storedPassword, password); err != nil {
		log.Println(err)
		return errors.New("password is incorrect")
	}

	//hash new password
	hashedPassword, err := hashing.GenerateHash(newPassword)
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

func (d domain) DeleteAccount(ctx context.Context, userUuid string, password string) error {

	//retrieve password from database
	storedPassword, err := d.repo.CheckPassword(ctx, userUuid)
	if err != nil {
		log.Printf("error retrieving password %s", err)
		return errors.New("error retrieving password")
	}

	//check if password is correct
	if err = hashing.CompareHash(storedPassword, password); err != nil {
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

func (d domain) RegisterAccount(ctx context.Context, cred Credentials) (Credentials, error) {
	//check if username is in database
	err := d.repo.CheckUsername(ctx, cred.Username)
	if err != nil {
		log.Printf("username %s already exists", cred.Username)
		return Credentials{}, errors.New("username already exists")
	}

	//perform check if password is of atleast 8 characters
	if len(cred.Password) < 8 {
		log.Println("password is too short")
		return Credentials{}, errors.New("password is too short")
	}

	//hash password
	cred.Password, err = hashing.GenerateHash(cred.Password)
	if err != nil {
		return Credentials{}, err
	}

	cred.Uuid = hashing.GenerateToken()
	cred.Session = hashing.GenerateToken()
	cred.DemoUser = false
	//add user to database
	session, err := d.repo.Register(ctx, cred)
	if err != nil {
		return Credentials{}, err
	}

	//Call the user service to create a user

	_, err = d.userClient.CreateUser(ctx, &userClientv1.CreateUserRequest{
		Username: cred.Username,
		UserUuid: cred.Uuid,
	})
	if err != nil {
		log.Println(err)

		//if the user service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			log.Println(err)
			return Credentials{}, err
		}
		return Credentials{}, err
	}
	log.Println("Everything went through fine until here")

	//Call the chatroom service to add the user to the default chatroom
	//TODO:this here is causing the issue

	_, err = d.chatroomClient.ForceAddUser(ctx, &chatroomClientv1.ForceAddUserRequest{
		UserUuid: cred.Uuid,
	})
	if err != nil {
		log.Println(err)

		//if the chatroom service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			log.Println(err)
			return Credentials{}, err
		}
		return Credentials{}, err

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
		log.Printf("username %s already exists", cred.Username)
		cred.Username = hashing.GenerateName()
	}

	cred.Uuid = hashing.GenerateToken()
	cred.Session = hashing.GenerateToken()
	cred.DemoUser = true
	// add user to database
	session, err := d.repo.Register(ctx, cred)
	if err != nil {
		log.Println("Failed to register user: ", err)
		return Credentials{}, err
	}

	//Call the user service to create a user
	_, err = d.userClient.CreateUser(ctx, &userClientv1.CreateUserRequest{
		Username: cred.Username,
		UserUuid: cred.Uuid,
	})
	if err != nil {
		log.Println(err)

		//if the user service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			log.Println(err)
			return Credentials{}, err
		}
		return Credentials{}, err
	}

	//Add user to main chatroom
	_, err = d.chatroomClient.ForceAddUser(ctx, &chatroomClientv1.ForceAddUserRequest{
		UserUuid: cred.Uuid,
	})
	if err != nil {
		log.Println(err)

		//if the chatroom service fails to create a user then the user should be deleted from the auth service
		if err := d.repo.Unregister(ctx, cred.Uuid); err != nil {
			log.Println(err)
			return Credentials{}, err
		}
		return Credentials{}, err

		//Potentially need to add delete user from user service
	}

	return Credentials{Session: session, Uuid: cred.Uuid}, nil
}

func (d domain) UpgradeDemoUser(ctx context.Context, creds Credentials) error {

	//Validate password is long enough
	if len(creds.Password) < 8 {
		log.Println("password is too short")
		return errors.New("password is too short")
	}

	//hash password
	hashedPassword, err := hashing.GenerateHash(creds.Password)
	if err != nil {
		log.Println("Failed to hash password: ", err)
		return err
	}

	creds.Password = hashedPassword
	creds.DemoUser = false

	//Update user in database
	err = d.repo.UpdateAccount(ctx, creds)
	if err != nil {
		log.Println("Failed to update user: ", err)
		return err
	}

	return nil
}
