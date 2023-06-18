package domain

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"io"

	"github.com/dietzy1/chatapp/services/user/domain"
)

type Icon struct {
	Link      string `json:"link" bson:"link"`
	Uuid      string `json:"uuid" bson:"uuid"`
	Kindof    string `json:"kindof" bson:"kindof"`
	OwnerUuid string `json:"owneruuid" bson:"owneruuid"`
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
func (a Icon) StoreIcon(ctx context.Context, icon Icon) error {
	return nil
}

func (a Icon) GetIcon(ctx context.Context, uuid string) (Icon, error) {
	return Icon{}, nil
}

func (a Icon) UpdateIcon(ctx context.Context, icon Icon) error {
	return nil
}

func (a Icon) DeleteIcon(ctx context.Context, uuid string) error {
	return nil
}

// I need to add calls incase CDN or database fails
