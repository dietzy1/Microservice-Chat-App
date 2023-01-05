package domain

import (
	"testing"
)

// Mocked interface of the domain auth service
/* type auth interface {
	Login(ctx context.Context, username string) (string, error)
	Logout(ctx context.Context, username string) error
	Authenticate(ctx context.Context, token string) (string, error)
	UpdateToken(ctx context.Context, username string, token string) error
} */

// I should fake the result of the databases and simply ensure that the function calls the correct logic
func TestLogin(t *testing.T) {

}

/*
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
} */

func TestLogout(t *testing.T) {

}

func TestAuthenticate(t *testing.T) {

}

func TestUpdateToken(t *testing.T) {

}
