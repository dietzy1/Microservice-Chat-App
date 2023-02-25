package domain

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type cdn interface {
	UploadFile(ctx context.Context, icon Icon, buf bytes.Buffer) (string, error)
	DeleteFile(ctx context.Context, uuid string) error
	GetFile(ctx context.Context, uuid string) (string, error)
}

type Icon struct {
	Link string
	Uuid string
}

type icon interface {
	StoreIcon(ctx context.Context, icon Icon) error
	GetIcon(ctx context.Context, uuid string) (Icon, error)
	UpdateIcon(ctx context.Context, icon Icon) error
	DeleteIcon(ctx context.Context, uuid string) error
}

// I need to add calls incase CDN or database fails
func (a Domain) setIcon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//Parameters that should be recieved from grpc
	file, _, err := r.FormFile("file")
	if err != nil {

	}
	defer file.Close()

	buf := new(bytes.Buffer)
	err = ConvertToJPEG(buf, file)
	if err != nil {

	}
	//I want to generate a new icon uuid and add the new link to the icon struct

	icon := Icon{
		Uuid: uuid.New().String(),
		//Link: fmt.Sprintf("https://ik.imagekit.io/your_imagekit_id/icons/%s.jpg", icon.Uuid),
	}
	icon.Link, err = a.cdn.UploadFile(ctx, icon, *buf)
	if err != nil {

	}
	if err = a.cdn.DeleteFile(ctx, icon.Uuid); err != nil {

	}
	//Do call to database

	if err = a.icon.UpdateIcon(ctx, icon); err != nil {

	}
	w.WriteHeader(http.StatusOK)
}

func (a Domain) deleteIcon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	//Do call to CDN to delete icon
	//Do call to database to delete icon
	if err := a.cdn.DeleteFile(ctx, uuid); err != nil {

	}
}

func (a Domain) GetIcon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")

	//Do call to database to get icon
	icon, err := a.icon.GetIcon(ctx, uuid)
	if err != nil {

	}

	_ = icon

}

func ConvertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 95})
}
