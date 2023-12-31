// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/echo-stream/database/ent/application"
	"github.com/echo-stream/database/ent/user"
)

// Application is the model entity for the Application schema.
type Application struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Description holds the value of the "Description" field.
	Description string `json:"Description,omitempty"`
	// Secret holds the value of the "Secret" field.
	Secret string `json:"Secret,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationQuery when eager-loading is set.
	Edges        ApplicationEdges `json:"edges"`
	user_id      *int
	selectValues sql.SelectValues
}

// ApplicationEdges holds the relations/edges for other nodes in the graph.
type ApplicationEdges struct {
	// UserID holds the value of the user_id edge.
	UserID *User `json:"user_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserIDOrErr returns the UserID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationEdges) UserIDOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserID == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserID, nil
	}
	return nil, &NotLoadedError{edge: "user_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Application) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			values[i] = new(sql.NullInt64)
		case application.FieldName, application.FieldDescription, application.FieldSecret:
			values[i] = new(sql.NullString)
		case application.FieldCreatedAt, application.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case application.ForeignKeys[0]: // user_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Application fields.
func (a *Application) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case application.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case application.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case application.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Secret", values[i])
			} else if value.Valid {
				a.Secret = value.String
			}
		case application.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case application.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case application.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_id", value)
			} else if value.Valid {
				a.user_id = new(int)
				*a.user_id = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Application.
// This includes values selected through modifiers, order, etc.
func (a *Application) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUserID queries the "user_id" edge of the Application entity.
func (a *Application) QueryUserID() *UserQuery {
	return NewApplicationClient(a.config).QueryUserID(a)
}

// Update returns a builder for updating this Application.
// Note that you need to call Application.Unwrap() before calling this method if this Application
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Application) Update() *ApplicationUpdateOne {
	return NewApplicationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Application entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Application) Unwrap() *Application {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Application is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Application) String() string {
	var builder strings.Builder
	builder.WriteString("Application(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("Name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("Description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("Secret=")
	builder.WriteString(a.Secret)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Applications is a parsable slice of Application.
type Applications []*Application