package cdn

import (
	//"image/jpeg"

	//"image/jpeg"

	"bytes"
	"context"
	"fmt"

	"io"
	"os"

	"github.com/dietzy1/chatapp/services/chatroom/domain"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

//Imagekit client implementation
//Imlements methods on the port type CdnPort interface
//This file is responcible for crud operations for storing the file bytes itself at the imagekit CDN.

type cdn struct {
	client *imagekit.ImageKit
}

//I need to define an interface for this

// Purpose of this for the user to be able to upload custom icons to their profile
func New() *cdn {
	client := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  os.Getenv("PRIVATE_KEY"),
		PublicKey:   os.Getenv("PUBLIC_KEY"),
		UrlEndpoint: os.Getenv("URL_ENDPOINT"),
	})
	return &cdn{client: client}
}

// sends a POST http request that stores the image bytes with a path of uuid.jpg at the CDN.
func (f *cdn) UploadFile(ctx context.Context, icon domain.Icon, buf bytes.Buffer) (string, error) {
	params := uploader.UploadParam{
		FileName:          icon.Uuid + ".jpg",
		UseUniqueFileName: newFalse(),
		Folder:            "/chatroom/",
		IsPrivateFile:     newFalse(),
		ResponseFields:    "filepath",
	}
	res, err := f.client.Uploader.Upload(ctx, io.ByteReader(&buf), params)
	if err != nil {
		return "", err
	}
	return res.Data.Url, nil
}

// Sends a DELETE http request that deletes the image bytes at the CDN.
func (f *cdn) DeleteFile(ctx context.Context, uuid string) error {
	fileid, err := f.GetFile(ctx, uuid)
	if err != nil {
		return err
	}
	_, err = f.client.Media.DeleteFile(ctx, fileid)
	if err != nil {
		return err
	}
	return nil
}

// helper function to enable deletefile and update file. Sends a GET request that locates the image bytes at the CDN.
func (f *cdn) GetFile(ctx context.Context, uuid string) (string, error) {
	query := fmt.Sprintf(`name = "%s"`, uuid+".jpg")
	res, err := f.client.Media.Files(ctx, media.FilesParam{
		SearchQuery: query,
	})
	if err != nil {
		return "", err
	}
	return res.Data[0].FileId, nil
}

// Hack to circumvent poor client library implementation.
func newFalse() *bool {
	b := false
	return &b
}
