// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeApplications holds the string denoting the applications edge name in mutations.
	EdgeApplications = "applications"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ApplicationsTable is the table that holds the applications relation/edge.
	ApplicationsTable = "applications"
	// ApplicationsInverseTable is the table name for the Application entity.
	// It exists in this package in order to avoid circular dependency with the "application" package.
	ApplicationsInverseTable = "applications"
	// ApplicationsColumn is the table column denoting the applications relation/edge.
	ApplicationsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldEmail,
	FieldPassword,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUserId holds the default value on creation for the "UserId" field.
	DefaultUserId func() uuid.UUID
	// EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserId orders the results by the UserId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByEmail orders the results by the Email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the Password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByApplicationsCount orders the results by applications count.
func ByApplicationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newApplicationsStep(), opts...)
	}
}

// ByApplications orders the results by applications terms.
func ByApplications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApplicationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newApplicationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApplicationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ApplicationsTable, ApplicationsColumn),
	)
}
