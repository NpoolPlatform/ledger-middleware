package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// LedgerLock holds the schema definition for the LedgerLock entity.
type LedgerLock struct {
	ent.Schema
}

func (LedgerLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the LedgerLock.
func (LedgerLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("ledger_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("statement_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			String("lock_state").
			Optional().
			Default(types.LedgerLockState_LedgerLockLocked.String()),
	}
}

// Edges of the LedgerLock.
func (LedgerLock) Edges() []ent.Edge {
	return nil
}
