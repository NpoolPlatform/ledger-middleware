// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodLedger is the model entity for the GoodLedger schema.
type GoodLedger struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// ToPlatform holds the value of the "to_platform" field.
	ToPlatform decimal.Decimal `json:"to_platform,omitempty"`
	// ToUser holds the value of the "to_user" field.
	ToUser decimal.Decimal `json:"to_user,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodLedger) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodledger.FieldAmount, goodledger.FieldToPlatform, goodledger.FieldToUser:
			values[i] = new(decimal.Decimal)
		case goodledger.FieldCreatedAt, goodledger.FieldUpdatedAt, goodledger.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case goodledger.FieldID, goodledger.FieldGoodID, goodledger.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodLedger", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodLedger fields.
func (gl *GoodLedger) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodledger.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gl.ID = *value
			}
		case goodledger.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gl.CreatedAt = uint32(value.Int64)
			}
		case goodledger.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gl.UpdatedAt = uint32(value.Int64)
			}
		case goodledger.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gl.DeletedAt = uint32(value.Int64)
			}
		case goodledger.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gl.GoodID = *value
			}
		case goodledger.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				gl.CoinTypeID = *value
			}
		case goodledger.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				gl.Amount = *value
			}
		case goodledger.FieldToPlatform:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field to_platform", values[i])
			} else if value != nil {
				gl.ToPlatform = *value
			}
		case goodledger.FieldToUser:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field to_user", values[i])
			} else if value != nil {
				gl.ToUser = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodLedger.
// Note that you need to call GoodLedger.Unwrap() before calling this method if this GoodLedger
// was returned from a transaction, and the transaction was committed or rolled back.
func (gl *GoodLedger) Update() *GoodLedgerUpdateOne {
	return (&GoodLedgerClient{config: gl.config}).UpdateOne(gl)
}

// Unwrap unwraps the GoodLedger entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gl *GoodLedger) Unwrap() *GoodLedger {
	_tx, ok := gl.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodLedger is not a transactional entity")
	}
	gl.config.driver = _tx.drv
	return gl
}

// String implements the fmt.Stringer.
func (gl *GoodLedger) String() string {
	var builder strings.Builder
	builder.WriteString("GoodLedger(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gl.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", gl.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", gl.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", gl.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", gl.GoodID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", gl.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", gl.Amount))
	builder.WriteString(", ")
	builder.WriteString("to_platform=")
	builder.WriteString(fmt.Sprintf("%v", gl.ToPlatform))
	builder.WriteString(", ")
	builder.WriteString("to_user=")
	builder.WriteString(fmt.Sprintf("%v", gl.ToUser))
	builder.WriteByte(')')
	return builder.String()
}

// GoodLedgers is a parsable slice of GoodLedger.
type GoodLedgers []*GoodLedger

func (gl GoodLedgers) config(cfg config) {
	for _i := range gl {
		gl[_i].config = cfg
	}
}
