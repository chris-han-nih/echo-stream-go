package models

import "github.com/google/uuid"

type Message struct {
	ApplicationId uuid.UUID    `json:"application_id" validate:"required,max=36"`
	UserName      string       `json:"username,omitempty" validate:"max=30"`
	IconEmoji     string       `json:"icon_emoji,omitempty"`
	Channel       string       `json:"channel" validate:"required,min=10,max=10"`
	Text          string       `json:"text,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}
