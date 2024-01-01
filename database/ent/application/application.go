// Code generated by ent, DO NOT EDIT.

package application

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the application type in the database.
	Label = "application"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldApplicationId holds the string denoting the applicationid field in the database.
	FieldApplicationId = "application_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUserID holds the string denoting the user_id edge name in mutations.
	EdgeUserID = "user_id"
	// Table holds the table name of the application in the database.
	Table = "applications"
	// UserIDTable is the table that holds the user_id relation/edge.
	UserIDTable = "applications"
	// UserIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserIDInverseTable = "users"
	// UserIDColumn is the table column denoting the user_id relation/edge.
	UserIDColumn = "user_id"
)

// Columns holds all SQL columns for application fields.
var Columns = []string{
	FieldID,
	FieldApplicationId,
	FieldName,
	FieldDescription,
	FieldSecret,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "applications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultApplicationId holds the default value on creation for the "ApplicationId" field.
	DefaultApplicationId func() uuid.UUID
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "Description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// SecretValidator is a validator for the "Secret" field. It is called by the builders before save.
	SecretValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Application queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByApplicationId orders the results by the ApplicationId field.
func ByApplicationId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationId, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the Description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySecret orders the results by the Secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserIDField orders the results by user_id field.
func ByUserIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserIDStep(), sql.OrderByField(field, opts...))
	}
}
func newUserIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
	)
}
