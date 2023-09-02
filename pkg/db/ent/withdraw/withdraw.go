// Code generated by ent, DO NOT EDIT.

package withdraw

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the withdraw type in the database.
	Label = "withdraw"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPlatformTransactionID holds the string denoting the platform_transaction_id field in the database.
	FieldPlatformTransactionID = "platform_transaction_id"
	// FieldChainTransactionID holds the string denoting the chain_transaction_id field in the database.
	FieldChainTransactionID = "chain_transaction_id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldReviewID holds the string denoting the review_id field in the database.
	FieldReviewID = "review_id"
	// Table holds the table name of the withdraw in the database.
	Table = "withdraws"
)

// Columns holds all SQL columns for withdraw fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldCoinTypeID,
	FieldAccountID,
	FieldAddress,
	FieldPlatformTransactionID,
	FieldChainTransactionID,
	FieldState,
	FieldAmount,
	FieldReviewID,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultAccountID holds the default value on creation for the "account_id" field.
	DefaultAccountID func() uuid.UUID
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
	// DefaultPlatformTransactionID holds the default value on creation for the "platform_transaction_id" field.
	DefaultPlatformTransactionID func() uuid.UUID
	// DefaultChainTransactionID holds the default value on creation for the "chain_transaction_id" field.
	DefaultChainTransactionID string
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultReviewID holds the default value on creation for the "review_id" field.
	DefaultReviewID func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
