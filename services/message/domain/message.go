package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (d *Domain) CreateMessage(ctx context.Context, msg Message) (Message, error) {

	// perform similar message is empty check on msg
	if msg.Author == "" || msg.Content == "" || msg.AuthorUuid == "" || msg.ChatRoomUuid == "" || msg.ChannelUuid == "" {
		d.logger.Info("Message is empty")
		return Message{}, errors.New("Message is empty")
	}

	//Add the required UUIDS and timestamps
	msg.MessageUuid = uuid.New().String()
	msg.Timestamp = formatTimestamp()

	//Call the repository layer
	err := d.repo.AddMessage(ctx, msg)
	if err != nil {
		d.logger.Info("Error adding message to database", zap.Error(err))
		return Message{}, fmt.Errorf("error adding message to database: %w", err)
	}

	return msg, nil
}

func (d *Domain) GetMessages(ctx context.Context, chatroomUuid string, channelUuid string) ([]Message, error) {

	if chatroomUuid == "" || channelUuid == "" {
		d.logger.Info("Message is empty")
		return []Message{}, errors.New("Message is empty")
	}

	//Call the repository layer'
	messages, err := d.repo.GetMessages(ctx, chatroomUuid, channelUuid)
	if err != nil {
		d.logger.Info("Error getting messages from database", zap.Error(err))
		return []Message{}, fmt.Errorf("error getting messages from database: %w", err)
	}

	return messages, nil
}

func (d *Domain) EditMessage(ctx context.Context, msg Message) (Message, error) {

	if msg.Author == "" || msg.Content == "" || msg.AuthorUuid == "" || msg.ChatRoomUuid == "" || msg.ChannelUuid == "" || msg.MessageUuid == "" {
		d.logger.Info("Message is empty")
		return Message{}, errors.New("Message is empty")
	}

	err := d.repo.UpdateMessage(ctx, msg)
	if err != nil {
		d.logger.Info("Error updating message in database", zap.Error(err))
		return Message{}, fmt.Errorf("error updating message in database: %w", err)
	}

	return Message{}, nil
}

func (d *Domain) DeleteMessage(ctx context.Context, msg Message) error {

	//check if msg is empty
	if msg.Author == "" || msg.Content == "" || msg.AuthorUuid == "" || msg.ChatRoomUuid == "" || msg.ChannelUuid == "" || msg.MessageUuid == "" {
		d.logger.Info("Message is empty")
		return errors.New("Message is empty")
	}

	err := d.repo.DeleteMessage(ctx, msg)
	if err != nil {
		d.logger.Info("Error deleting message from database", zap.Error(err))
		return fmt.Errorf("error deleting message from database: %w", err)
	}

	return nil
}
