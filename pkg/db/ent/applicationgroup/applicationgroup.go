// Code generated by entc, DO NOT EDIT.

package applicationgroup

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the applicationgroup type in the database.
	Label = "application_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldGroupName holds the string denoting the group_name field in the database.
	FieldGroupName = "group_name"
	// FieldGroupLogo holds the string denoting the group_logo field in the database.
	FieldGroupLogo = "group_logo"
	// FieldGroupOwner holds the string denoting the group_owner field in the database.
	FieldGroupOwner = "group_owner"
	// FieldAnnotation holds the string denoting the annotation field in the database.
	FieldAnnotation = "annotation"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the applicationgroup in the database.
	Table = "application_groups"
)

// Columns holds all SQL columns for applicationgroup fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldGroupName,
	FieldGroupLogo,
	FieldGroupOwner,
	FieldAnnotation,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
