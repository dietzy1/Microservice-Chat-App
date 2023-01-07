package core

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
)

type User struct {
	Icon        Icon //A link to the users icon
	Name        string
	Uuid        string
	Discription string
	JoinDate    string
	Roles       []string
	ChatServers []string
	Reports     int
}

type Icon struct {
	Link string
	Uuid string
}

type Credentials struct {
	Username  string
	Password  string
	Uuid      string
	CSRFToken string
}

func ConvertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 95})
}

// Converts an error to a string
func Errconv(err error) string {
	return fmt.Sprintf("%s", err)
}
