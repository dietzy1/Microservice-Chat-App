package domain

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"io"
	"log"

	_ "image/gif"
	_ "image/jpeg"

	"github.com/google/uuid"
	_ "golang.org/x/image/webp"
)

type cdn interface {
	UploadIcon(ctx context.Context, icon Icon, buf bytes.Buffer) (string, error)
	DeleteIcon(ctx context.Context, uuid string) error
	GetIcon(ctx context.Context, uuid string) (string, error)
}

type Icon struct {
	Link string
	Uuid string
}

type icon interface {
	StoreIcon(ctx context.Context, icon Icon) error
	GetIcon(ctx context.Context, uuid string) (Icon, error)
	GetAllIcons(ctx context.Context) ([]Icon, error)
	DeleteIcon(ctx context.Context, uuid string) error
}

func (d Domain) UploadAvatar(ctx context.Context, image bytes.Buffer) (Icon, error) {
	//convert to jpeg - Accepts webp, png, jpeg and gif
	buf := bytes.Buffer{}
	err := ConvertToPng(&buf, &image)
	if err != nil {
		return Icon{}, err
	}

	icon := Icon{
		Link: "",
		Uuid: uuid.New().String(),
	}

	//upload icon to CDN
	icon.Link, err = d.cdn.UploadIcon(ctx, icon, buf)
	if err != nil {
		return Icon{}, err
	}

	//store icon in database
	err = d.icon.StoreIcon(ctx, icon)
	if err != nil {
		return Icon{}, err
	}

	return Icon{}, nil
}

func (d Domain) DeleteIcon(ctx context.Context, uuid string) error {
	//Delete icon from CDN
	err := d.cdn.DeleteIcon(ctx, uuid)
	if err != nil {
		return err
	}

	//Delete icon from database
	err = d.icon.DeleteIcon(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (d Domain) GetAllIcons(ctx context.Context) ([]Icon, error) {
	//Get icons from database
	icons, err := d.icon.GetAllIcons(ctx)
	if err != nil {
		return []Icon{}, err
	}
	return icons, nil

}

func (d Domain) GetIcon(ctx context.Context, uuid string) (Icon, error) {
	//Get icon from database
	icon, err := d.icon.GetIcon(ctx, uuid)
	if err != nil {
		return Icon{}, err
	}

	//Get icon from CDN
	icon.Uuid, err = d.cdn.GetIcon(ctx, uuid)
	if err != nil {
		return Icon{}, err
	}
	return icon, nil
}

// Accepts formats of webp, png, jpeg and gif
func ConvertToPng(w io.Writer, r io.Reader) error {
	img, imageType, err := image.Decode(r)
	if err != nil {
		return err
	}
	log.Println("Encoding the image of type: ", imageType, " to png")
	return png.Encode(w, img)
}

//The way I think I want this logic to work is that I upload icons to the CDN and then store the link to the icon in the database.
//I want to be able to delete the icon from the CDN and the database. I also want to be able to update the icon in the database and the CDN.

/* cdn := cdn.New()

buf := new(bytes.Buffer)

buf.ReadFrom(image)

ok, err := cdn.UploadFile(context.TODO(), test, *buf)
if err != nil {
	panic(err)
} */
