package domain

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"io"

	_ "image/gif"
	_ "image/jpeg"

	"go.uber.org/zap"
	_ "golang.org/x/image/webp"
)

//ok so I need to figure out exactly how this is going to work
//There should be 3 different folders of icons on imagekit
//1. User Icons
//2. ServerIcons
//3. Standart emoji icons which I have preuploaded

//First step is probaly going in and editing the imagekit package so it supports these 3 different folders

// Enum to manage the different types of icon folders
const (
	userIconFolder     = "/usericons/"
	chatroomIconFolder = "/servericons/"
	emojiIconFolder    = "/emojiicons/"
)

type Icon struct {
	Link      string `json:"link" bson:"link"`
	Uuid      string `json:"uuid" bson:"uuid"`
	Kindof    string `json:"kindof" bson:"kindof"`
	OwnerUuid string `json:"owneruuid" bson:"owneruuid"`
}

type Domain struct {
	logger *zap.Logger
	repo   repo
	cdn    cdn
}

// Repository interface
type repo interface {
	StoreIcon(ctx context.Context, icon Icon) error
	GetIcon(ctx context.Context, uuid string) (Icon, error)
	GetIcons(ctx context.Context, ownerUuid string) ([]Icon, error)
	GetEmojiIcons(ctx context.Context) ([]Icon, error)
	DeleteIcon(ctx context.Context, uuid string) error
}

// Injected into domain struct
type cdn interface {
	UploadIcon(ctx context.Context, icon Icon, buf bytes.Buffer, folder string) (string, error)
	DeleteIcon(ctx context.Context, uuid string) error
	GetIcon(ctx context.Context, uuid string) (string, error)
}

// Constructor for the domain
func New(logger *zap.Logger, repo repo, cdn cdn) Domain {
	return Domain{
		logger: logger,
		repo:   repo,
		cdn:    cdn,
	}
}

// Accepts formats of webp, png, jpeg and gif
func ConvertToPng(w io.Writer, r io.Reader) error {
	img, imageType, err := image.Decode(r)
	if err != nil {
		return fmt.Errorf("error decoding image: %w It has the following type: %s", err, imageType)
	}
	return png.Encode(w, img)
}
