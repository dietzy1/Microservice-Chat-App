package domain

import (
	"bytes"
	"context"

	"log"

	"github.com/google/uuid"
)

/* service IconService {
	//get by uuid
	rpc GetIcon(GetIconRequest) returns (GetIconResponse) {}

	//get by owner uuid
	rpc GetIcons(GetIconsRequest) returns (GetIconsResponse) {}

	//return all emoji icons
	rpc GetEmojiIcons(GetEmojiIconsRequest) returns (GetEmojiIconsResponse) {}

	rpc DeleteIcon(DeleteIconRequest) returns (DeleteIconResponse) {}

	rpc UploadIcon(stream UploadIconRequest) returns (UploadIconResponse) {}
  } */

// Method which retrieves an icon based on uuid
func (d Domain) GetIcon(ctx context.Context, uuid string) (Icon, error) {
	icon, err := d.repo.GetIcon(ctx, uuid)
	if err != nil {
		return Icon{}, err
	}
	return icon, nil
}

// Method which retrieves all icons accociated with owner uuid
func (d Domain) GetIcons(ctx context.Context, ownerUuid string) ([]Icon, error) {
	icons, err := d.repo.GetIcons(ctx, ownerUuid)
	if err != nil {
		return []Icon{}, err
	}
	return icons, nil
}

func (d Domain) GetEmojiIcons(ctx context.Context) ([]Icon, error) {
	icons, err := d.repo.GetEmojiIcons(ctx)
	if err != nil {
		return []Icon{}, err
	}
	return icons, nil
}

func (d Domain) UploadIcon(ctx context.Context, image bytes.Buffer, info Icon) (Icon, error) {
	//convert to jpeg - Accepts webp, png, jpeg and gif
	buf := bytes.Buffer{}
	err := ConvertToPng(&buf, &image)
	if err != nil {
		return Icon{}, err
	}

	icon := Icon{
		Link:      "",
		Uuid:      uuid.New().String(),
		Kindof:    info.Kindof,
		OwnerUuid: info.OwnerUuid,
	}

	folder := ""
	switch icon.Kindof {
	case "user":
		folder = userIconFolder
	case "chatroom":
		folder = chatroomIconFolder
	case "emoji":
		folder = emojiIconFolder
	default:
		log.Println("Invalid icon type")
		return Icon{}, nil
	}

	//upload icon to CDN
	icon.Link, err = d.cdn.UploadIcon(ctx, icon, buf, folder)
	if err != nil {
		return Icon{}, err
	}

	//store icon in database
	err = d.repo.StoreIcon(ctx, icon)
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
	err = d.repo.DeleteIcon(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
