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

	userclientv1 "github.com/dietzy1/chatapp/services/user/proto/user/v1"
)

//

// Purpose of this function is to upload an avatar to the user service
func main() {
	config.ReadEnvfile()

	client := connect()

	uploadAvatar(*client, "test.png")

	//Look into the upload folder and look for new files that doesn't match the files in uploaded folder

}

func uploadAvatar(client userclientv1.UserServiceClient, imagePath string) {
	//Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Retrieve the stream object from the server
	stream, err := client.UploadAvatar(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	log.Println(filepath.Ext(imagePath))

	//Construct the singular request object
	req := &userclientv1.UploadAvatarRequest{
		Data: &userclientv1.UploadAvatarRequest_Info{
			Info: &userclientv1.ImageInfo{
				ImageType: filepath.Ext(imagePath),
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

		req := &userclientv1.UploadAvatarRequest{
			Data: &userclientv1.UploadAvatarRequest_ChunkData{
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

}

func connect() *userclientv1.UserServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("USERSERVICE"),
		//"localhost:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := userclientv1.NewUserServiceClient(conn)
	return &client
}

/* cdn := cdn.New()

buf := new(bytes.Buffer)

buf.ReadFrom(image)

ok, err := cdn.UploadFile(context.TODO(), test, *buf)
if err != nil {
	panic(err)
} */
