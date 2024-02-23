// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CouponWithdrawsColumns holds the columns for the "coupon_withdraws" table.
	CouponWithdrawsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "allocated_id", Type: field.TypeUUID, Nullable: true},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "Reviewing"},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "review_id", Type: field.TypeUUID, Nullable: true},
	}
	// CouponWithdrawsTable holds the schema information for the "coupon_withdraws" table.
	CouponWithdrawsTable = &schema.Table{
		Name:       "coupon_withdraws",
		Columns:    CouponWithdrawsColumns,
		PrimaryKey: []*schema.Column{CouponWithdrawsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "couponwithdraw_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CouponWithdrawsColumns[4]},
			},
		},
	}
	// MiningGeneralsColumns holds the columns for the "mining_generals" table.
	MiningGeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
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
		Indexes: []*schema.Index{
			{
				Name:    "goodledger_ent_id",
				Unique:  true,
				Columns: []*schema.Column{MiningGeneralsColumns[4]},
			},
		},
	}
	// MiningDetailsColumns holds the columns for the "mining_details" table.
	MiningDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "to_platform", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "to_user", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "technique_service_fee_amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "benefit_date", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// MiningDetailsTable holds the schema information for the "mining_details" table.
	MiningDetailsTable = &schema.Table{
		Name:       "mining_details",
		Columns:    MiningDetailsColumns,
		PrimaryKey: []*schema.Column{MiningDetailsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "goodstatement_ent_id",
				Unique:  true,
				Columns: []*schema.Column{MiningDetailsColumns[4]},
			},
		},
	}
	// GeneralsColumns holds the columns for the "generals" table.
	GeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
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
		Indexes: []*schema.Index{
			{
				Name:    "ledger_ent_id",
				Unique:  true,
				Columns: []*schema.Column{GeneralsColumns[4]},
			},
		},
	}
	// LedgerLocksColumns holds the columns for the "ledger_locks" table.
	LedgerLocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "ledger_id", Type: field.TypeUUID, Nullable: true},
		{Name: "statement_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "lock_state", Type: field.TypeString, Nullable: true, Default: "LedgerLockLocked"},
	}
	// LedgerLocksTable holds the schema information for the "ledger_locks" table.
	LedgerLocksTable = &schema.Table{
		Name:       "ledger_locks",
		Columns:    LedgerLocksColumns,
		PrimaryKey: []*schema.Column{LedgerLocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ledgerlock_ent_id",
				Unique:  true,
				Columns: []*schema.Column{LedgerLocksColumns[4]},
			},
		},
	}
	// ProfitsColumns holds the columns for the "profits" table.
	ProfitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
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
		Indexes: []*schema.Index{
			{
				Name:    "profit_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ProfitsColumns[4]},
			},
		},
	}
	// SimulateGeneralsColumns holds the columns for the "simulate_generals" table.
	SimulateGeneralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "incoming", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "outcoming", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// SimulateGeneralsTable holds the schema information for the "simulate_generals" table.
	SimulateGeneralsTable = &schema.Table{
		Name:       "simulate_generals",
		Columns:    SimulateGeneralsColumns,
		PrimaryKey: []*schema.Column{SimulateGeneralsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "simulateledger_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SimulateGeneralsColumns[4]},
			},
		},
	}
	// SimulateProfitsColumns holds the columns for the "simulate_profits" table.
	SimulateProfitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "incoming", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// SimulateProfitsTable holds the schema information for the "simulate_profits" table.
	SimulateProfitsTable = &schema.Table{
		Name:       "simulate_profits",
		Columns:    SimulateProfitsColumns,
		PrimaryKey: []*schema.Column{SimulateProfitsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "simulateprofit_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SimulateProfitsColumns[4]},
			},
		},
	}
	// SimulateDetailsColumns holds the columns for the "simulate_details" table.
	SimulateDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "io_type", Type: field.TypeString, Nullable: true, Default: "DefaultType"},
		{Name: "io_sub_type", Type: field.TypeString, Nullable: true, Default: "DefaultSubType"},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "io_extra", Type: field.TypeString, Nullable: true, Size: 512, Default: ""},
		{Name: "send_coupon", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "cashable", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "cash_used", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "cash_used_at", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// SimulateDetailsTable holds the schema information for the "simulate_details" table.
	SimulateDetailsTable = &schema.Table{
		Name:       "simulate_details",
		Columns:    SimulateDetailsColumns,
		PrimaryKey: []*schema.Column{SimulateDetailsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "simulatestatement_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SimulateDetailsColumns[4]},
			},
		},
	}
	// DetailsColumns holds the columns for the "details" table.
	DetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "io_type", Type: field.TypeString, Nullable: true, Default: "DefaultType"},
		{Name: "io_sub_type", Type: field.TypeString, Nullable: true, Default: "DefaultSubType"},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "io_extra", Type: field.TypeString, Nullable: true, Size: 512, Default: ""},
		{Name: "io_extra_v1", Type: field.TypeString, Nullable: true, Size: 512, Default: ""},
	}
	// DetailsTable holds the schema information for the "details" table.
	DetailsTable = &schema.Table{
		Name:       "details",
		Columns:    DetailsColumns,
		PrimaryKey: []*schema.Column{DetailsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "statement_ent_id",
				Unique:  true,
				Columns: []*schema.Column{DetailsColumns[4]},
			},
		},
	}
	// MiningUnsoldsColumns holds the columns for the "mining_unsolds" table.
	MiningUnsoldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "benefit_date", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "statement_id", Type: field.TypeUUID, Nullable: true},
	}
	// MiningUnsoldsTable holds the schema information for the "mining_unsolds" table.
	MiningUnsoldsTable = &schema.Table{
		Name:       "mining_unsolds",
		Columns:    MiningUnsoldsColumns,
		PrimaryKey: []*schema.Column{MiningUnsoldsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "unsoldstatement_ent_id",
				Unique:  true,
				Columns: []*schema.Column{MiningUnsoldsColumns[4]},
			},
		},
	}
	// WithdrawsColumns holds the columns for the "withdraws" table.
	WithdrawsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "platform_transaction_id", Type: field.TypeUUID, Nullable: true},
		{Name: "chain_transaction_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "Created"},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "review_id", Type: field.TypeUUID, Nullable: true},
	}
	// WithdrawsTable holds the schema information for the "withdraws" table.
	WithdrawsTable = &schema.Table{
		Name:       "withdraws",
		Columns:    WithdrawsColumns,
		PrimaryKey: []*schema.Column{WithdrawsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "withdraw_ent_id",
				Unique:  true,
				Columns: []*schema.Column{WithdrawsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CouponWithdrawsTable,
		MiningGeneralsTable,
		MiningDetailsTable,
		GeneralsTable,
		LedgerLocksTable,
		ProfitsTable,
		SimulateGeneralsTable,
		SimulateProfitsTable,
		SimulateDetailsTable,
		DetailsTable,
		MiningUnsoldsTable,
		WithdrawsTable,
	}
)

func init() {
	MiningGeneralsTable.Annotation = &entsql.Annotation{
		Table: "mining_generals",
	}
	MiningDetailsTable.Annotation = &entsql.Annotation{
		Table: "mining_details",
	}
	GeneralsTable.Annotation = &entsql.Annotation{
		Table: "generals",
	}
	SimulateGeneralsTable.Annotation = &entsql.Annotation{
		Table: "simulate_generals",
	}
	SimulateDetailsTable.Annotation = &entsql.Annotation{
		Table: "simulate_details",
	}
	DetailsTable.Annotation = &entsql.Annotation{
		Table: "details",
	}
	MiningUnsoldsTable.Annotation = &entsql.Annotation{
		Table: "mining_unsolds",
	}
}
