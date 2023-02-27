package domain

type Domain struct {
	user user
	cdn  cdn
	icon icon
}

// New creates a new application
func New(user user, icon icon, cdn cdn) *Domain {
	return &Domain{user: user, icon: icon, cdn: cdn}
}
