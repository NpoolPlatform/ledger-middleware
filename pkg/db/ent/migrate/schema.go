// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DetailsColumns holds the columns for the "details" table.
	DetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "io_type", Type: field.TypeString, Nullable: true, Default: "DefaultType"},
		{Name: "io_sub_type", Type: field.TypeString, Nullable: true, Default: "DefaultSubType"},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "from_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_usd_currency", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "io_extra", Type: field.TypeString, Nullable: true, Size: 512, Default: ""},
	}
	// DetailsTable holds the schema information for the "details" table.
	DetailsTable = &schema.Table{
		Name:       "details",
		Columns:    DetailsColumns,
		PrimaryKey: []*schema.Column{DetailsColumns[0]},
	}
	// GeneralsColumns holds the columns for the "generals" table.
	GeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "incoming", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "locked", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "outcoming", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "spendable", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// GeneralsTable holds the schema information for the "generals" table.
	GeneralsTable = &schema.Table{
		Name:       "generals",
		Columns:    GeneralsColumns,
		PrimaryKey: []*schema.Column{GeneralsColumns[0]},
	}
	// MiningDetailsColumns holds the columns for the "mining_details" table.
	MiningDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "benefit_date", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// MiningDetailsTable holds the schema information for the "mining_details" table.
	MiningDetailsTable = &schema.Table{
		Name:       "mining_details",
		Columns:    MiningDetailsColumns,
		PrimaryKey: []*schema.Column{MiningDetailsColumns[0]},
	}
	// MiningGeneralsColumns holds the columns for the "mining_generals" table.
	MiningGeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "to_platform", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "to_user", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// MiningGeneralsTable holds the schema information for the "mining_generals" table.
	MiningGeneralsTable = &schema.Table{
		Name:       "mining_generals",
		Columns:    MiningGeneralsColumns,
		PrimaryKey: []*schema.Column{MiningGeneralsColumns[0]},
	}
	// MiningUnsoldsColumns holds the columns for the "mining_unsolds" table.
	MiningUnsoldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "benefit_date", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// MiningUnsoldsTable holds the schema information for the "mining_unsolds" table.
	MiningUnsoldsTable = &schema.Table{
		Name:       "mining_unsolds",
		Columns:    MiningUnsoldsColumns,
		PrimaryKey: []*schema.Column{MiningUnsoldsColumns[0]},
	}
	// ProfitsColumns holds the columns for the "profits" table.
	ProfitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "incoming", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// ProfitsTable holds the schema information for the "profits" table.
	ProfitsTable = &schema.Table{
		Name:       "profits",
		Columns:    ProfitsColumns,
		PrimaryKey: []*schema.Column{ProfitsColumns[0]},
	}
	// WithdrawsColumns holds the columns for the "withdraws" table.
	WithdrawsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "platform_transaction_id", Type: field.TypeUUID, Nullable: true},
		{Name: "chain_transaction_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultWithdrawState"},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// WithdrawsTable holds the schema information for the "withdraws" table.
	WithdrawsTable = &schema.Table{
		Name:       "withdraws",
		Columns:    WithdrawsColumns,
		PrimaryKey: []*schema.Column{WithdrawsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DetailsTable,
		GeneralsTable,
		MiningDetailsTable,
		MiningGeneralsTable,
		MiningUnsoldsTable,
		ProfitsTable,
		WithdrawsTable,
	}
)

func init() {
}
