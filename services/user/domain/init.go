package domain

// Application is the main application struct
type Domain struct {
	acc  Account
	user User
	cdn  Cdn
	icon Icon
}

// New creates a new application
func New(acc Account, user User, cdn Cdn, icon Icon) *Domain {
	return &Domain{acc: acc, user: user, cdn: cdn}
}

//All of the Application methods should accept grpc parameters

//User and Credentail UUID should be shared
