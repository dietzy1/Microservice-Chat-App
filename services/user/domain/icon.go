package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/dietzy1/chatapp/services/user/core"
	"github.com/google/uuid"
)

type Cdn interface {
	UploadFile(ctx context.Context, icon core.Icon, buf bytes.Buffer) (string, error)
	DeleteFile(ctx context.Context, uuid string) error
	GetFile(ctx context.Context, uuid string) (string, error)
}

type Icon interface {
	StoreIcon(ctx context.Context, icon core.Icon) error
	GetIcon(ctx context.Context, uuid string) (Icon, error)
	UpdateIcon(ctx context.Context, icon core.Icon) error
	DeleteIcon(ctx context.Context, uuid string) error
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
