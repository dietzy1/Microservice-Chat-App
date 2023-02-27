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

func UploadIcon() {
	//Upload icon to CDN
	//Store link to icon in database
}

func DeleteIcon() {
	//Delete icon from CDN
	//Delete icon from database
}

func GetIcon() {
	//Get icon from database
	//Get icon from CDN
}
