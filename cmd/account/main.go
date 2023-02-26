package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/dietzy1/chatapp/config"
	"github.com/dietzy1/chatapp/services/user/adapters/rest/cdn"
	"github.com/dietzy1/chatapp/services/user/domain"
)

func main() {

	config.ReadEnvfile()
	//Test some image uploading here

	//Read the image user.png from the current directory
	//Upload
	image, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}

	test := domain.Icon{
		Link: "defaultAvatar",
		Uuid: "defaultAvatar",
	}

	cdn := cdn.New()

	buf := new(bytes.Buffer)

	buf.ReadFrom(image)

	ok, err := cdn.UploadFile(context.TODO(), test, *buf)
	if err != nil {
		panic(err)
	}

	log.Println(ok)

}
