// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApplicationsColumns holds the columns for the "applications" table.
	ApplicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 30},
		{Name: "description", Type: field.TypeString},
		{Name: "secret", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ApplicationsTable holds the schema information for the "applications" table.
	ApplicationsTable = &schema.Table{
		Name:       "applications",
		Columns:    ApplicationsColumns,
		PrimaryKey: []*schema.Column{ApplicationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "applications_users_applications",
				Columns:    []*schema.Column{ApplicationsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SlackMessagesColumns holds the columns for the "slack_messages" table.
	SlackMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "application_id", Type: field.TypeUUID},
		{Name: "user_name", Type: field.TypeString, Nullable: true},
		{Name: "icon_emoji", Type: field.TypeString, Nullable: true},
		{Name: "channel", Type: field.TypeString},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "thread_ts", Type: field.TypeString, Nullable: true},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "created"},
		{Name: "retry_count", Type: field.TypeInt16, Nullable: true, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SlackMessagesTable holds the schema information for the "slack_messages" table.
	SlackMessagesTable = &schema.Table{
		Name:       "slack_messages",
		Columns:    SlackMessagesColumns,
		PrimaryKey: []*schema.Column{SlackMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "slackmessage_created_at",
				Unique:  false,
				Columns: []*schema.Column{SlackMessagesColumns[9]},
			},
		},
	}
	// SlackMessageAttachmentsColumns holds the columns for the "slack_message_attachments" table.
	SlackMessageAttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "color", Type: field.TypeString, Nullable: true},
		{Name: "fallback", Type: field.TypeString, Nullable: true},
		{Name: "callback_id", Type: field.TypeString, Nullable: true},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
		{Name: "author_name", Type: field.TypeString, Nullable: true},
		{Name: "author_subname", Type: field.TypeString, Nullable: true},
		{Name: "author_link", Type: field.TypeString, Nullable: true},
		{Name: "author_icon", Type: field.TypeString, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "title_link", Type: field.TypeString, Nullable: true},
		{Name: "pretext", Type: field.TypeString, Nullable: true},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "thumb_url", Type: field.TypeString, Nullable: true},
		{Name: "message_id", Type: field.TypeInt},
	}
	// SlackMessageAttachmentsTable holds the schema information for the "slack_message_attachments" table.
	SlackMessageAttachmentsTable = &schema.Table{
		Name:       "slack_message_attachments",
		Columns:    SlackMessageAttachmentsColumns,
		PrimaryKey: []*schema.Column{SlackMessageAttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "slack_message_attachments_slack_messages_slack_message_attachments",
				Columns:    []*schema.Column{SlackMessageAttachmentsColumns[15]},
				RefColumns: []*schema.Column{SlackMessagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApplicationsTable,
		SlackMessagesTable,
		SlackMessageAttachmentsTable,
		UsersTable,
	}
)

func init() {
	ApplicationsTable.ForeignKeys[0].RefTable = UsersTable
	SlackMessageAttachmentsTable.ForeignKeys[0].RefTable = SlackMessagesTable
}
