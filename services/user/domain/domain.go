package domain

type Domain struct {
	user user
	cdn  cdn
	icon icon
}

// New creates a new application
func New(user user, cdn cdn, icon icon) *Domain {
	return &Domain{user: user, cdn: cdn, icon: icon}
}
