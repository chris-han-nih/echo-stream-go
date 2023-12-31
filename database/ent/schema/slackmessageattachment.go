package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SlackMessageAttachment holds the schema definition for the SlackMessageAttachment entity.
type SlackMessageAttachment struct {
	ent.Schema
}

// Fields of the SlackMessageAttachment.
func (SlackMessageAttachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("Color").Optional(),
		field.String("Fallback").Optional(),
		field.String("CallbackID").Optional(),
		field.String("AuthorID").Optional(),
		field.String("AuthorName").Optional(),
		field.String("AuthorSubname").Optional(),
		field.String("AuthorLink").Optional(),
		field.String("AuthorIcon").Optional(),
		field.String("Title").Optional(),
		field.String("TitleLink").Optional(),
		field.String("Pretext").Optional(),
		field.String("Text").Optional(),
		field.String("ImageURL").Optional(),
		field.String("ThumbURL").Optional(),
	}
}

// Edges of the SlackMessageAttachment.
func (SlackMessageAttachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message_id", SlackMessage.Type).
			Ref("slack_message_attachments").
			Unique().
			Required(),
	}
}
