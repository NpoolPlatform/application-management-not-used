// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/application-management/pkg/db/ent/applicationgroupuser"
	"github.com/google/uuid"
)

// ApplicationGroupUser is the model entity for the ApplicationGroupUser schema.
type ApplicationGroupUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID uuid.UUID `json:"group_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Annotation holds the value of the "annotation" field.
	Annotation string `json:"annotation,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApplicationGroupUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case applicationgroupuser.FieldCreateAt, applicationgroupuser.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case applicationgroupuser.FieldAppID, applicationgroupuser.FieldAnnotation:
			values[i] = new(sql.NullString)
		case applicationgroupuser.FieldID, applicationgroupuser.FieldGroupID, applicationgroupuser.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApplicationGroupUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApplicationGroupUser fields.
func (agu *ApplicationGroupUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case applicationgroupuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				agu.ID = *value
			}
		case applicationgroupuser.FieldGroupID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value != nil {
				agu.GroupID = *value
			}
		case applicationgroupuser.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				agu.AppID = value.String
			}
		case applicationgroupuser.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				agu.UserID = *value
			}
		case applicationgroupuser.FieldAnnotation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field annotation", values[i])
			} else if value.Valid {
				agu.Annotation = value.String
			}
		case applicationgroupuser.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				agu.CreateAt = value.Int64
			}
		case applicationgroupuser.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				agu.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ApplicationGroupUser.
// Note that you need to call ApplicationGroupUser.Unwrap() before calling this method if this ApplicationGroupUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (agu *ApplicationGroupUser) Update() *ApplicationGroupUserUpdateOne {
	return (&ApplicationGroupUserClient{config: agu.config}).UpdateOne(agu)
}

// Unwrap unwraps the ApplicationGroupUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (agu *ApplicationGroupUser) Unwrap() *ApplicationGroupUser {
	tx, ok := agu.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApplicationGroupUser is not a transactional entity")
	}
	agu.config.driver = tx.drv
	return agu
}

// String implements the fmt.Stringer.
func (agu *ApplicationGroupUser) String() string {
	var builder strings.Builder
	builder.WriteString("ApplicationGroupUser(")
	builder.WriteString(fmt.Sprintf("id=%v", agu.ID))
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", agu.GroupID))
	builder.WriteString(", app_id=")
	builder.WriteString(agu.AppID)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", agu.UserID))
	builder.WriteString(", annotation=")
	builder.WriteString(agu.Annotation)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", agu.CreateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", agu.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// ApplicationGroupUsers is a parsable slice of ApplicationGroupUser.
type ApplicationGroupUsers []*ApplicationGroupUser

func (agu ApplicationGroupUsers) config(cfg config) {
	for _i := range agu {
		agu[_i].config = cfg
	}
}