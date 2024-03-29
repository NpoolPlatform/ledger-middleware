// Code generated by ent, DO NOT EDIT.

package statement

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the statement type in the database.
	Label = "statement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldIoType holds the string denoting the io_type field in the database.
	FieldIoType = "io_type"
	// FieldIoSubType holds the string denoting the io_sub_type field in the database.
	FieldIoSubType = "io_sub_type"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldIoExtra holds the string denoting the io_extra field in the database.
	FieldIoExtra = "io_extra"
	// FieldIoExtraV1 holds the string denoting the io_extra_v1 field in the database.
	FieldIoExtraV1 = "io_extra_v1"
	// Table holds the table name of the statement in the database.
	Table = "details"
)

// Columns holds all SQL columns for statement fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldCoinTypeID,
	FieldIoType,
	FieldIoSubType,
	FieldAmount,
	FieldIoExtra,
	FieldIoExtraV1,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultIoType holds the default value on creation for the "io_type" field.
	DefaultIoType string
	// DefaultIoSubType holds the default value on creation for the "io_sub_type" field.
	DefaultIoSubType string
	// DefaultIoExtra holds the default value on creation for the "io_extra" field.
	DefaultIoExtra string
	// IoExtraValidator is a validator for the "io_extra" field. It is called by the builders before save.
	IoExtraValidator func(string) error
	// DefaultIoExtraV1 holds the default value on creation for the "io_extra_v1" field.
	DefaultIoExtraV1 string
	// IoExtraV1Validator is a validator for the "io_extra_v1" field. It is called by the builders before save.
	IoExtraV1Validator func(string) error
)
