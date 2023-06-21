package domain

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

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

	//Validate data
	if ownerUuid == "" {
		d.logger.Info("ownerUuid is empty")
		return []Icon{}, errors.New("ownerUuid is empty")
	}

	icons, err := d.repo.GetIcons(ctx, ownerUuid)
	if err != nil {
		d.logger.Info("error retrieving icons", zap.Error(err))
		return []Icon{}, fmt.Errorf("error retrieving icons: %w", err)
	}
	return icons, nil
}

func (d Domain) GetEmojiIcons(ctx context.Context) ([]Icon, error) {
	icons, err := d.repo.GetEmojiIcons(ctx)
	if err != nil {
		d.logger.Info("error retrieving icons", zap.Error(err))
		return []Icon{}, fmt.Errorf("error retrieving emoji icons: %w", err)
	}
	return icons, nil
}

func (d Domain) UploadIcon(ctx context.Context, image bytes.Buffer, info Icon) (Icon, error) {
	//convert to jpeg - Accepts webp, png, jpeg and gif
	buf := bytes.Buffer{}
	err := ConvertToPng(&buf, &image)
	if err != nil {
		d.logger.Info("error converting image to png", zap.Error(err))
		return Icon{}, fmt.Errorf("error converting image to png: %w", err)
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
		d.logger.Info("invalid kindof", zap.String("kindof", icon.Kindof))
		return Icon{}, errors.New("invalid kindof")
	}

	//upload icon to CDN
	icon.Link, err = d.cdn.UploadIcon(ctx, icon, buf, folder)
	if err != nil {
		d.logger.Info("error uploading icon to cdn", zap.Error(err))
		return Icon{}, fmt.Errorf("error uploading icon to cdn: %w", err)
	}

	//store icon in database
	err = d.repo.StoreIcon(ctx, icon)
	if err != nil {
		d.logger.Info("error storing icon in database", zap.Error(err))
		return Icon{}, fmt.Errorf("error storing icon in database: %w", err)
	}

	return Icon{}, nil
}

func (d Domain) DeleteIcon(ctx context.Context, uuid string) error {

	if uuid == "" {
		d.logger.Info("uuid is empty")
		return errors.New("uuid is empty")
	}

	//Delete icon from CDN
	err := d.cdn.DeleteIcon(ctx, uuid)
	if err != nil {
		d.logger.Info("error deleting icon from cdn", zap.Error(err))
		return fmt.Errorf("error deleting icon from cdn: %w", err)
	}

	//Delete icon from database
	err = d.repo.DeleteIcon(ctx, uuid)
	if err != nil {
		d.logger.Info("error deleting icon from database", zap.Error(err))
		return fmt.Errorf("error deleting icon from database: %w", err)
	}
	return nil
}
