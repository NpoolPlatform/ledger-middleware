package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/mixin"
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
	}
}

// Fields of the LedgerLock.
func (LedgerLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
	}
}

// Edges of the LedgerLock.
func (LedgerLock) Edges() []ent.Edge {
	return nil
}
