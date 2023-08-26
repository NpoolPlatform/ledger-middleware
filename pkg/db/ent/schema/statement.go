package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Statement holds the schema definition for the Statement entity.
type Statement struct {
	ent.Schema
}

func (Statement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Statement.
func (Statement) Fields() []ent.Field {
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
			String("io_type").
			Optional().
			Default(basetypes.IOType_DefaultType.String()),
		field.
			String("io_sub_type").
			Optional().
			Default(basetypes.IOSubType_DefaultSubType.String()),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			String("io_extra").
			Optional().
			Default("").
			MaxLen(512), //nolint
	}
}

// Edges of the Detail.
func (Statement) Edges() []ent.Edge {
	return nil
}

func (Statement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "details"},
	}
}
