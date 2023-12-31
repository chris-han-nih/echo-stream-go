package models

import (
	"github.com/google/uuid"
)

type Attachment struct {
	MessageId     uuid.UUID `json:"message_id,omitempty"`
	Color         string    `json:"color,omitempty"`
	Fallback      string    `json:"fallback,omitempty"`
	CallbackID    string    `json:"callback_id,omitempty"`
	AuthorID      string    `json:"author_id,omitempty"`
	AuthorName    string    `json:"author_name,omitempty"`
	AuthorSubname string    `json:"author_subname,omitempty"`
	AuthorLink    string    `json:"author_link,omitempty"`
	AuthorIcon    string    `json:"author_icon,omitempty"`
	Title         string    `json:"title,omitempty"`
	TitleLink     string    `json:"title_link,omitempty"`
	Pretext       string    `json:"pretext,omitempty"`
	Text          string    `json:"text,omitempty"`
	ImageURL      string    `json:"image_url,omitempty"`
	ThumbURL      string    `json:"thumb_url,omitempty"`
}
