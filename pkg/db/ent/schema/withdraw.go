package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/withdraw"
)

// Withdraw holds the schema definition for the Withdraw entity.
type Withdraw struct {
	ent.Schema
}

func (Withdraw) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Withdraw.
func (Withdraw) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("address").
			Optional().
			Default(""),
		field.
			UUID("platform_transaction_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("chain_transaction_id").
			Optional().
			Default(""),
		field.
			String("state").
			Optional().
			Default(withdraw.WithdrawState_DefaultWithdrawState.String()),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
	}
}

// Edges of the Withdraw.
func (Withdraw) Edges() []ent.Edge {
	return nil
}
