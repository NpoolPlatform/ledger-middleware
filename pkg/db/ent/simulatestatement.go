// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/simulatestatement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SimulateStatement is the model entity for the SimulateStatement schema.
type SimulateStatement struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// IoType holds the value of the "io_type" field.
	IoType string `json:"io_type,omitempty"`
	// IoSubType holds the value of the "io_sub_type" field.
	IoSubType string `json:"io_sub_type,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// IoExtra holds the value of the "io_extra" field.
	IoExtra string `json:"io_extra,omitempty"`
	// SendCoupon holds the value of the "send_coupon" field.
	SendCoupon bool `json:"send_coupon,omitempty"`
	// Cashable holds the value of the "cashable" field.
	Cashable bool `json:"cashable,omitempty"`
	// CashUsed holds the value of the "cash_used" field.
	CashUsed bool `json:"cash_used,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SimulateStatement) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case simulatestatement.FieldAmount:
			values[i] = new(decimal.Decimal)
		case simulatestatement.FieldSendCoupon, simulatestatement.FieldCashable, simulatestatement.FieldCashUsed:
			values[i] = new(sql.NullBool)
		case simulatestatement.FieldID, simulatestatement.FieldCreatedAt, simulatestatement.FieldUpdatedAt, simulatestatement.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case simulatestatement.FieldIoType, simulatestatement.FieldIoSubType, simulatestatement.FieldIoExtra:
			values[i] = new(sql.NullString)
		case simulatestatement.FieldEntID, simulatestatement.FieldAppID, simulatestatement.FieldUserID, simulatestatement.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SimulateStatement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SimulateStatement fields.
func (ss *SimulateStatement) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case simulatestatement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ss.ID = uint32(value.Int64)
		case simulatestatement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ss.CreatedAt = uint32(value.Int64)
			}
		case simulatestatement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ss.UpdatedAt = uint32(value.Int64)
			}
		case simulatestatement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ss.DeletedAt = uint32(value.Int64)
			}
		case simulatestatement.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ss.EntID = *value
			}
		case simulatestatement.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ss.AppID = *value
			}
		case simulatestatement.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ss.UserID = *value
			}
		case simulatestatement.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ss.CoinTypeID = *value
			}
		case simulatestatement.FieldIoType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field io_type", values[i])
			} else if value.Valid {
				ss.IoType = value.String
			}
		case simulatestatement.FieldIoSubType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field io_sub_type", values[i])
			} else if value.Valid {
				ss.IoSubType = value.String
			}
		case simulatestatement.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				ss.Amount = *value
			}
		case simulatestatement.FieldIoExtra:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field io_extra", values[i])
			} else if value.Valid {
				ss.IoExtra = value.String
			}
		case simulatestatement.FieldSendCoupon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field send_coupon", values[i])
			} else if value.Valid {
				ss.SendCoupon = value.Bool
			}
		case simulatestatement.FieldCashable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cashable", values[i])
			} else if value.Valid {
				ss.Cashable = value.Bool
			}
		case simulatestatement.FieldCashUsed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cash_used", values[i])
			} else if value.Valid {
				ss.CashUsed = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SimulateStatement.
// Note that you need to call SimulateStatement.Unwrap() before calling this method if this SimulateStatement
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *SimulateStatement) Update() *SimulateStatementUpdateOne {
	return (&SimulateStatementClient{config: ss.config}).UpdateOne(ss)
}

// Unwrap unwraps the SimulateStatement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ss *SimulateStatement) Unwrap() *SimulateStatement {
	_tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("ent: SimulateStatement is not a transactional entity")
	}
	ss.config.driver = _tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *SimulateStatement) String() string {
	var builder strings.Builder
	builder.WriteString("SimulateStatement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ss.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ss.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ss.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ss.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.UserID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ss.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("io_type=")
	builder.WriteString(ss.IoType)
	builder.WriteString(", ")
	builder.WriteString("io_sub_type=")
	builder.WriteString(ss.IoSubType)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", ss.Amount))
	builder.WriteString(", ")
	builder.WriteString("io_extra=")
	builder.WriteString(ss.IoExtra)
	builder.WriteString(", ")
	builder.WriteString("send_coupon=")
	builder.WriteString(fmt.Sprintf("%v", ss.SendCoupon))
	builder.WriteString(", ")
	builder.WriteString("cashable=")
	builder.WriteString(fmt.Sprintf("%v", ss.Cashable))
	builder.WriteString(", ")
	builder.WriteString("cash_used=")
	builder.WriteString(fmt.Sprintf("%v", ss.CashUsed))
	builder.WriteByte(')')
	return builder.String()
}

// SimulateStatements is a parsable slice of SimulateStatement.
type SimulateStatements []*SimulateStatement

func (ss SimulateStatements) config(cfg config) {
	for _i := range ss {
		ss[_i].config = cfg
	}
}
