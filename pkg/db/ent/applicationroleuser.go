// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationroleuser"
	"github.com/google/uuid"
)

// ApplicationRoleUser is the model entity for the ApplicationRoleUser schema.
type ApplicationRoleUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID uuid.UUID `json:"role_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApplicationRoleUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case applicationroleuser.FieldCreateAt, applicationroleuser.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case applicationroleuser.FieldAppID:
			values[i] = new(sql.NullString)
		case applicationroleuser.FieldID, applicationroleuser.FieldRoleID, applicationroleuser.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApplicationRoleUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApplicationRoleUser fields.
func (aru *ApplicationRoleUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case applicationroleuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				aru.ID = *value
			}
		case applicationroleuser.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				aru.AppID = value.String
			}
		case applicationroleuser.FieldRoleID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value != nil {
				aru.RoleID = *value
			}
		case applicationroleuser.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				aru.UserID = *value
			}
		case applicationroleuser.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				aru.CreateAt = value.Int64
			}
		case applicationroleuser.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				aru.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ApplicationRoleUser.
// Note that you need to call ApplicationRoleUser.Unwrap() before calling this method if this ApplicationRoleUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (aru *ApplicationRoleUser) Update() *ApplicationRoleUserUpdateOne {
	return (&ApplicationRoleUserClient{config: aru.config}).UpdateOne(aru)
}

// Unwrap unwraps the ApplicationRoleUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aru *ApplicationRoleUser) Unwrap() *ApplicationRoleUser {
	tx, ok := aru.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApplicationRoleUser is not a transactional entity")
	}
	aru.config.driver = tx.drv
	return aru
}

// String implements the fmt.Stringer.
func (aru *ApplicationRoleUser) String() string {
	var builder strings.Builder
	builder.WriteString("ApplicationRoleUser(")
	builder.WriteString(fmt.Sprintf("id=%v", aru.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(aru.AppID)
	builder.WriteString(", role_id=")
	builder.WriteString(fmt.Sprintf("%v", aru.RoleID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", aru.UserID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", aru.CreateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", aru.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// ApplicationRoleUsers is a parsable slice of ApplicationRoleUser.
type ApplicationRoleUsers []*ApplicationRoleUser

func (aru ApplicationRoleUsers) config(cfg config) {
	for _i := range aru {
		aru[_i].config = cfg
	}
}