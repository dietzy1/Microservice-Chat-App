package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dietzy1/chatapp/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	iconclientv1 "github.com/dietzy1/chatapp/services/icon/proto/icon/v1"
)

//https://emojipedia.org/anxious-face-with-sweat/
//https://getemoji.com/

// Purpose of this function is to upload an avatar to the user service
func main() {
	config.ReadEnvfile()

	log.Println(discoverFolder())

	client := connect()

	for _, file := range discoverFolder() {
		err := uploadAvatar(*client, file)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("uploaded file: ", file)

		//take the file out of the folder and put it into uploaded
		err = os.Rename("upload/"+file, "uploaded/"+file)
		if err != nil {
			log.Println(err)
			return
		}

	}

}

// function that looks into a folder and returns a list of file names
func discoverFolder() []string {
	files, err := os.ReadDir("upload")
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, file := range files {

		if file.Name() != ".DS_Store" {

			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames
}

func uploadAvatar(client iconclientv1.IconServiceClient, imagePath string) error {
	//Open the image file in the upload folder

	file, err := os.Open("upload/" + imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Retrieve the stream object from the server
	stream, err := client.UploadIcon(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	log.Println(filepath.Ext(imagePath))

	//Construct the singular request object
	req := &iconclientv1.UploadIconRequest{
		Data: &iconclientv1.UploadIconRequest_Info{
			Info: &iconclientv1.ImageInfo{
				OwnerUuid: "1234",
				Kindof:    "avatar",
			},
		},
	}

	//Send the singular request object to the server
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	//Read the image file in chunks
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &iconclientv1.UploadIconRequest{
			Data: &iconclientv1.UploadIconRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err)
		}
	}

	//Close the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot close stream: ", err)
	}
	log.Println("Upload avatar response: ", res)

	return nil
}

func connect() *iconclientv1.IconServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("ICONSERVICE"),
		//"localhost:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := iconclientv1.NewIconServiceClient(conn)
	return &client
}
