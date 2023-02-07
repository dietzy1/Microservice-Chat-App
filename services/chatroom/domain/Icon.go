package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"net/http"

	"github.com/dietzy1/chatapp/services/user/core"
	"github.com/dietzy1/chatapp/services/user/domain"
	"github.com/google/uuid"
)

type Icon struct {
	Link string `json:"link" bson:"link"`
	Uuid string `json:"uuid" bson:"uuid"`
}

// Injected into domain struct
type cdn interface {
	UploadFile(ctx context.Context, icon domain.Icon, buf bytes.Buffer) (string, error)
	DeleteFile(ctx context.Context, uuid string) error
	GetFile(ctx context.Context, uuid string) (string, error)
}

func ConvertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
}

// I need to define the methods that takes in the file and converts it
func (a Icon) StoreIcon(ctx context.Context, icon core.Icon) error {
	return nil
}

func (a Icon) GetIcon(ctx context.Context, uuid string) (Icon, error) {
	return Icon{}, nil
}

func (a Icon) UpdateIcon(ctx context.Context, icon core.Icon) error {
	return nil
}

func (a Icon) DeleteIcon(ctx context.Context, uuid string) error {
	return nil
}

// I need to add calls incase CDN or database fails
func (a Domain) setIcon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//Parameters that should be recieved from grpc
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to parse file data. Here is the error value:", core.Errconv(err)})
		return
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	err = core.ConvertToJPEG(buf, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to convert file to jpg. Here is the error value:", core.Errconv(err)})
		return
	}
	//I want to generate a new icon uuid and add the new link to the icon struct

	icon := core.Icon{
		Uuid: uuid.New().String(),
		//Link: fmt.Sprintf("https://ik.imagekit.io/your_imagekit_id/icons/%s.jpg", icon.Uuid),
	}
	icon.Link, err = a.cdn.UploadFile(ctx, icon, *buf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to upload file to CDN. Here is the error value:", core.Errconv(err)})
		return
	}
	if err = a.cdn.DeleteFile(ctx, icon.Uuid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to delete file from CDN. Here is the error value:", core.Errconv(err)})
		return
	}
	//Do call to database

	if err = a.icon.UpdateIcon(ctx, icon); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to update icon in database. Here is the error value:", core.Errconv(err)})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a Domain) deleteIcon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	//Do call to CDN to delete icon
	//Do call to database to delete icon
	if err := a.cdn.DeleteFile(ctx, uuid); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to delete file from CDN. Here is the error value:", core.Errconv(err)})
		return
	}
}

func (a Domain) GetIcon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")

	//Do call to database to get icon
	icon, err := a.icon.GetIcon(ctx, uuid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode([]any{"Unable to get icon from database. Here is the error value:", core.Errconv(err)})
		return
	}
	// return 200
	json.NewEncoder(w).Encode(icon)
	w.WriteHeader(http.StatusOK)
}
