package server

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	iconclientv1 "github.com/dietzy1/"
)

// REST endpoint which accepts an image and contacts the icon service to upload it
func (s *server) uploadIconHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to parse multipartform")
		return
	}

	file, _, err := r.FormFile("icon")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to parse file")
		return
	}
	defer file.Close()

	//Construct new timeout to ensure stability
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Retrieve the stream object from the server
	stream, err := s.iconClient.UploadIcon(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	log.Println(filepath.Ext(imagePath))

	//Construct the singular request object
	req := &iconclientv1.UploadAvatarRequest{
		Data: &iconclientv1.UploadAvatarRequest_Info{
			Info: &iconclientv1.ImageInfo{
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

	return nil

}
