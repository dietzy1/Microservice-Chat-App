package domain

import (
	"bytes"
	"context"
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
	DeleteIcon(ctx context.Context, uuid string) error
}

//The way I think I want this logic to work is that I upload icons to the CDN and then store the link to the icon in the database.
//I want to be able to delete the icon from the CDN and the database. I also want to be able to update the icon in the database and the CDN.
