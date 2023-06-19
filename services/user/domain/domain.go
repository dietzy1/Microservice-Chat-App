package domain

type Domain struct {
	user user
}

// New creates a new application
func New(user user) *Domain {
	return &Domain{user: user}
}
