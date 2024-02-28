// Code generated by ent, DO NOT EDIT.

package ledgerlock

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// LedgerID applies equality check predicate on the "ledger_id" field. It's identical to LedgerIDEQ.
func LedgerID(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLedgerID), v))
	})
}

// StatementID applies equality check predicate on the "statement_id" field. It's identical to StatementIDEQ.
func StatementID(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatementID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// LockState applies equality check predicate on the "lock_state" field. It's identical to LockStateEQ.
func LockState(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockState), v))
	})
}

// ExLockID applies equality check predicate on the "ex_lock_id" field. It's identical to ExLockIDEQ.
func ExLockID(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExLockID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// LedgerIDEQ applies the EQ predicate on the "ledger_id" field.
func LedgerIDEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLedgerID), v))
	})
}

// LedgerIDNEQ applies the NEQ predicate on the "ledger_id" field.
func LedgerIDNEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLedgerID), v))
	})
}

// LedgerIDIn applies the In predicate on the "ledger_id" field.
func LedgerIDIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLedgerID), v...))
	})
}

// LedgerIDNotIn applies the NotIn predicate on the "ledger_id" field.
func LedgerIDNotIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLedgerID), v...))
	})
}

// LedgerIDGT applies the GT predicate on the "ledger_id" field.
func LedgerIDGT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLedgerID), v))
	})
}

// LedgerIDGTE applies the GTE predicate on the "ledger_id" field.
func LedgerIDGTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLedgerID), v))
	})
}

// LedgerIDLT applies the LT predicate on the "ledger_id" field.
func LedgerIDLT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLedgerID), v))
	})
}

// LedgerIDLTE applies the LTE predicate on the "ledger_id" field.
func LedgerIDLTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLedgerID), v))
	})
}

// LedgerIDIsNil applies the IsNil predicate on the "ledger_id" field.
func LedgerIDIsNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLedgerID)))
	})
}

// LedgerIDNotNil applies the NotNil predicate on the "ledger_id" field.
func LedgerIDNotNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLedgerID)))
	})
}

// StatementIDEQ applies the EQ predicate on the "statement_id" field.
func StatementIDEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatementID), v))
	})
}

// StatementIDNEQ applies the NEQ predicate on the "statement_id" field.
func StatementIDNEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatementID), v))
	})
}

// StatementIDIn applies the In predicate on the "statement_id" field.
func StatementIDIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatementID), v...))
	})
}

// StatementIDNotIn applies the NotIn predicate on the "statement_id" field.
func StatementIDNotIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatementID), v...))
	})
}

// StatementIDGT applies the GT predicate on the "statement_id" field.
func StatementIDGT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatementID), v))
	})
}

// StatementIDGTE applies the GTE predicate on the "statement_id" field.
func StatementIDGTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatementID), v))
	})
}

// StatementIDLT applies the LT predicate on the "statement_id" field.
func StatementIDLT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatementID), v))
	})
}

// StatementIDLTE applies the LTE predicate on the "statement_id" field.
func StatementIDLTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatementID), v))
	})
}

// StatementIDIsNil applies the IsNil predicate on the "statement_id" field.
func StatementIDIsNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatementID)))
	})
}

// StatementIDNotNil applies the NotNil predicate on the "statement_id" field.
func StatementIDNotNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatementID)))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...decimal.Decimal) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...decimal.Decimal) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v decimal.Decimal) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// LockStateEQ applies the EQ predicate on the "lock_state" field.
func LockStateEQ(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockState), v))
	})
}

// LockStateNEQ applies the NEQ predicate on the "lock_state" field.
func LockStateNEQ(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLockState), v))
	})
}

// LockStateIn applies the In predicate on the "lock_state" field.
func LockStateIn(vs ...string) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLockState), v...))
	})
}

// LockStateNotIn applies the NotIn predicate on the "lock_state" field.
func LockStateNotIn(vs ...string) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLockState), v...))
	})
}

// LockStateGT applies the GT predicate on the "lock_state" field.
func LockStateGT(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLockState), v))
	})
}

// LockStateGTE applies the GTE predicate on the "lock_state" field.
func LockStateGTE(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLockState), v))
	})
}

// LockStateLT applies the LT predicate on the "lock_state" field.
func LockStateLT(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLockState), v))
	})
}

// LockStateLTE applies the LTE predicate on the "lock_state" field.
func LockStateLTE(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLockState), v))
	})
}

// LockStateContains applies the Contains predicate on the "lock_state" field.
func LockStateContains(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLockState), v))
	})
}

// LockStateHasPrefix applies the HasPrefix predicate on the "lock_state" field.
func LockStateHasPrefix(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLockState), v))
	})
}

// LockStateHasSuffix applies the HasSuffix predicate on the "lock_state" field.
func LockStateHasSuffix(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLockState), v))
	})
}

// LockStateIsNil applies the IsNil predicate on the "lock_state" field.
func LockStateIsNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLockState)))
	})
}

// LockStateNotNil applies the NotNil predicate on the "lock_state" field.
func LockStateNotNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLockState)))
	})
}

// LockStateEqualFold applies the EqualFold predicate on the "lock_state" field.
func LockStateEqualFold(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLockState), v))
	})
}

// LockStateContainsFold applies the ContainsFold predicate on the "lock_state" field.
func LockStateContainsFold(v string) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLockState), v))
	})
}

// ExLockIDEQ applies the EQ predicate on the "ex_lock_id" field.
func ExLockIDEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExLockID), v))
	})
}

// ExLockIDNEQ applies the NEQ predicate on the "ex_lock_id" field.
func ExLockIDNEQ(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExLockID), v))
	})
}

// ExLockIDIn applies the In predicate on the "ex_lock_id" field.
func ExLockIDIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExLockID), v...))
	})
}

// ExLockIDNotIn applies the NotIn predicate on the "ex_lock_id" field.
func ExLockIDNotIn(vs ...uuid.UUID) predicate.LedgerLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExLockID), v...))
	})
}

// ExLockIDGT applies the GT predicate on the "ex_lock_id" field.
func ExLockIDGT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExLockID), v))
	})
}

// ExLockIDGTE applies the GTE predicate on the "ex_lock_id" field.
func ExLockIDGTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExLockID), v))
	})
}

// ExLockIDLT applies the LT predicate on the "ex_lock_id" field.
func ExLockIDLT(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExLockID), v))
	})
}

// ExLockIDLTE applies the LTE predicate on the "ex_lock_id" field.
func ExLockIDLTE(v uuid.UUID) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExLockID), v))
	})
}

// ExLockIDIsNil applies the IsNil predicate on the "ex_lock_id" field.
func ExLockIDIsNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExLockID)))
	})
}

// ExLockIDNotNil applies the NotNil predicate on the "ex_lock_id" field.
func ExLockIDNotNil() predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExLockID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LedgerLock) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LedgerLock) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LedgerLock) predicate.LedgerLock {
	return predicate.LedgerLock(func(s *sql.Selector) {
		p(s.Not())
	})
}
