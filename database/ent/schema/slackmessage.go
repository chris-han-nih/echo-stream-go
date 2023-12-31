package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// SlackMessage holds the schema definition for the SlackMessage entity.
type SlackMessage struct {
	ent.Schema
}

func (SlackMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("CreatedAt"),
	}
}

// Fields of the SlackMessage.
func (SlackMessage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("ApplicationId", uuid.UUID{}),
		field.String("UserName").Optional(),
		field.String("IconEmoji").Optional(),
		field.String("Channel").NotEmpty(),
		field.String("Text").Optional(),
		field.String("ThreadTs").Optional(),
		field.String("State").Optional().Default("created"),
		field.Int16("RetryCount").Optional().Default(0),
		field.Time("CreatedAt").Immutable().Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SlackMessage.
func (SlackMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("slack_message_attachments", SlackMessageAttachment.Type).StorageKey(edge.Column("message_id")),
	}
}
