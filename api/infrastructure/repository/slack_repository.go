package repository

import (
	"context"
	"fmt"
	
	"github.com/echo-stream/api/domain/models"
	"github.com/echo-stream/database/ent"
)

type SlackRepository struct {
	client *ent.Client
}

func NewSlackRepository(client *ent.Client) *SlackRepository {
	return &SlackRepository{
		client: client,
	}
}

func (s *SlackRepository) Create(ctx context.Context, m *models.Message) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	
	message, err := tx.SlackMessage.Create().
		SetApplicationId(m.ApplicationId).
		SetUserName(m.UserName).
		SetIconEmoji(m.IconEmoji).
		SetChannel(m.Channel).
		SetText(m.Text).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	
	for _, attachment := range m.Attachments {
		_, err = tx.SlackMessageAttachment.Create().
			SetMessageID(message).
			SetColor(attachment.Color).
			SetFallback(attachment.Fallback).
			SetCallbackID(attachment.CallbackID).
			SetAuthorID(attachment.AuthorID).
			SetAuthorName(attachment.AuthorName).
			SetAuthorSubname(attachment.AuthorSubname).
			SetAuthorLink(attachment.AuthorLink).
			SetAuthorIcon(attachment.AuthorIcon).
			SetTitle(attachment.Title).
			SetTitleLink(attachment.TitleLink).
			SetPretext(attachment.Pretext).
			SetText(attachment.Text).
			SetImageURL(attachment.ImageURL).
			SetThumbURL(attachment.ThumbURL).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}
	
	err = tx.Commit()
	if err != nil {
		return err
	}
	
	return nil
}

func rollback(tx *ent.Tx, err error) error {
	if err := tx.Rollback(); err != nil {
		err = fmt.Errorf("%w: %v", err, err)
	}
	return err
}
