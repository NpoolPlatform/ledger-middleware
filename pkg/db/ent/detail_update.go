// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// DetailUpdate is the builder for updating Detail entities.
type DetailUpdate struct {
	config
	hooks    []Hook
	mutation *DetailMutation
}

// Where appends a list predicates to the DetailUpdate builder.
func (du *DetailUpdate) Where(ps ...predicate.Detail) *DetailUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DetailUpdate) SetCreatedAt(u uint32) *DetailUpdate {
	du.mutation.ResetCreatedAt()
	du.mutation.SetCreatedAt(u)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCreatedAt(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetCreatedAt(*u)
	}
	return du
}

// AddCreatedAt adds u to the "created_at" field.
func (du *DetailUpdate) AddCreatedAt(u int32) *DetailUpdate {
	du.mutation.AddCreatedAt(u)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DetailUpdate) SetUpdatedAt(u uint32) *DetailUpdate {
	du.mutation.ResetUpdatedAt()
	du.mutation.SetUpdatedAt(u)
	return du
}

// AddUpdatedAt adds u to the "updated_at" field.
func (du *DetailUpdate) AddUpdatedAt(u int32) *DetailUpdate {
	du.mutation.AddUpdatedAt(u)
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DetailUpdate) SetDeletedAt(u uint32) *DetailUpdate {
	du.mutation.ResetDeletedAt()
	du.mutation.SetDeletedAt(u)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DetailUpdate) SetNillableDeletedAt(u *uint32) *DetailUpdate {
	if u != nil {
		du.SetDeletedAt(*u)
	}
	return du
}

// AddDeletedAt adds u to the "deleted_at" field.
func (du *DetailUpdate) AddDeletedAt(u int32) *DetailUpdate {
	du.mutation.AddDeletedAt(u)
	return du
}

// SetAppID sets the "app_id" field.
func (du *DetailUpdate) SetAppID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetAppID(u)
	return du
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableAppID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetAppID(*u)
	}
	return du
}

// ClearAppID clears the value of the "app_id" field.
func (du *DetailUpdate) ClearAppID() *DetailUpdate {
	du.mutation.ClearAppID()
	return du
}

// SetUserID sets the "user_id" field.
func (du *DetailUpdate) SetUserID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetUserID(u)
	return du
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableUserID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetUserID(*u)
	}
	return du
}

// ClearUserID clears the value of the "user_id" field.
func (du *DetailUpdate) ClearUserID() *DetailUpdate {
	du.mutation.ClearUserID()
	return du
}

// SetCoinTypeID sets the "coin_type_id" field.
func (du *DetailUpdate) SetCoinTypeID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetCoinTypeID(u)
	return du
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCoinTypeID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetCoinTypeID(*u)
	}
	return du
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (du *DetailUpdate) ClearCoinTypeID() *DetailUpdate {
	du.mutation.ClearCoinTypeID()
	return du
}

// SetIoType sets the "io_type" field.
func (du *DetailUpdate) SetIoType(s string) *DetailUpdate {
	du.mutation.SetIoType(s)
	return du
}

// SetNillableIoType sets the "io_type" field if the given value is not nil.
func (du *DetailUpdate) SetNillableIoType(s *string) *DetailUpdate {
	if s != nil {
		du.SetIoType(*s)
	}
	return du
}

// ClearIoType clears the value of the "io_type" field.
func (du *DetailUpdate) ClearIoType() *DetailUpdate {
	du.mutation.ClearIoType()
	return du
}

// SetIoSubType sets the "io_sub_type" field.
func (du *DetailUpdate) SetIoSubType(s string) *DetailUpdate {
	du.mutation.SetIoSubType(s)
	return du
}

// SetNillableIoSubType sets the "io_sub_type" field if the given value is not nil.
func (du *DetailUpdate) SetNillableIoSubType(s *string) *DetailUpdate {
	if s != nil {
		du.SetIoSubType(*s)
	}
	return du
}

// ClearIoSubType clears the value of the "io_sub_type" field.
func (du *DetailUpdate) ClearIoSubType() *DetailUpdate {
	du.mutation.ClearIoSubType()
	return du
}

// SetAmount sets the "amount" field.
func (du *DetailUpdate) SetAmount(d decimal.Decimal) *DetailUpdate {
	du.mutation.ResetAmount()
	du.mutation.SetAmount(d)
	return du
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (du *DetailUpdate) SetNillableAmount(d *decimal.Decimal) *DetailUpdate {
	if d != nil {
		du.SetAmount(*d)
	}
	return du
}

// AddAmount adds d to the "amount" field.
func (du *DetailUpdate) AddAmount(d decimal.Decimal) *DetailUpdate {
	du.mutation.AddAmount(d)
	return du
}

// ClearAmount clears the value of the "amount" field.
func (du *DetailUpdate) ClearAmount() *DetailUpdate {
	du.mutation.ClearAmount()
	return du
}

// SetFromCoinTypeID sets the "from_coin_type_id" field.
func (du *DetailUpdate) SetFromCoinTypeID(u uuid.UUID) *DetailUpdate {
	du.mutation.SetFromCoinTypeID(u)
	return du
}

// SetNillableFromCoinTypeID sets the "from_coin_type_id" field if the given value is not nil.
func (du *DetailUpdate) SetNillableFromCoinTypeID(u *uuid.UUID) *DetailUpdate {
	if u != nil {
		du.SetFromCoinTypeID(*u)
	}
	return du
}

// ClearFromCoinTypeID clears the value of the "from_coin_type_id" field.
func (du *DetailUpdate) ClearFromCoinTypeID() *DetailUpdate {
	du.mutation.ClearFromCoinTypeID()
	return du
}

// SetCoinUsdCurrency sets the "coin_usd_currency" field.
func (du *DetailUpdate) SetCoinUsdCurrency(d decimal.Decimal) *DetailUpdate {
	du.mutation.ResetCoinUsdCurrency()
	du.mutation.SetCoinUsdCurrency(d)
	return du
}

// SetNillableCoinUsdCurrency sets the "coin_usd_currency" field if the given value is not nil.
func (du *DetailUpdate) SetNillableCoinUsdCurrency(d *decimal.Decimal) *DetailUpdate {
	if d != nil {
		du.SetCoinUsdCurrency(*d)
	}
	return du
}

// AddCoinUsdCurrency adds d to the "coin_usd_currency" field.
func (du *DetailUpdate) AddCoinUsdCurrency(d decimal.Decimal) *DetailUpdate {
	du.mutation.AddCoinUsdCurrency(d)
	return du
}

// ClearCoinUsdCurrency clears the value of the "coin_usd_currency" field.
func (du *DetailUpdate) ClearCoinUsdCurrency() *DetailUpdate {
	du.mutation.ClearCoinUsdCurrency()
	return du
}

// SetIoExtra sets the "io_extra" field.
func (du *DetailUpdate) SetIoExtra(s string) *DetailUpdate {
	du.mutation.SetIoExtra(s)
	return du
}

// SetNillableIoExtra sets the "io_extra" field if the given value is not nil.
func (du *DetailUpdate) SetNillableIoExtra(s *string) *DetailUpdate {
	if s != nil {
		du.SetIoExtra(*s)
	}
	return du
}

// ClearIoExtra clears the value of the "io_extra" field.
func (du *DetailUpdate) ClearIoExtra() *DetailUpdate {
	du.mutation.ClearIoExtra()
	return du
}

// Mutation returns the DetailMutation object of the builder.
func (du *DetailUpdate) Mutation() *DetailMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := du.defaults(); err != nil {
		return 0, err
	}
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DetailUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DetailUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DetailUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DetailUpdate) defaults() error {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		if detail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (du *DetailUpdate) check() error {
	if v, ok := du.mutation.IoExtra(); ok {
		if err := detail.IoExtraValidator(v); err != nil {
			return &ValidationError{Name: "io_extra", err: fmt.Errorf(`ent: validator failed for field "Detail.io_extra": %w`, err)}
		}
	}
	return nil
}

func (du *DetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := du.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldAppID,
		})
	}
	if du.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldAppID,
		})
	}
	if value, ok := du.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldUserID,
		})
	}
	if du.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldUserID,
		})
	}
	if value, ok := du.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldCoinTypeID,
		})
	}
	if du.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldCoinTypeID,
		})
	}
	if value, ok := du.mutation.IoType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldIoType,
		})
	}
	if du.mutation.IoTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldIoType,
		})
	}
	if value, ok := du.mutation.IoSubType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldIoSubType,
		})
	}
	if du.mutation.IoSubTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldIoSubType,
		})
	}
	if value, ok := du.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldAmount,
		})
	}
	if value, ok := du.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldAmount,
		})
	}
	if du.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: detail.FieldAmount,
		})
	}
	if value, ok := du.mutation.FromCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldFromCoinTypeID,
		})
	}
	if du.mutation.FromCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldFromCoinTypeID,
		})
	}
	if value, ok := du.mutation.CoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldCoinUsdCurrency,
		})
	}
	if value, ok := du.mutation.AddedCoinUsdCurrency(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldCoinUsdCurrency,
		})
	}
	if du.mutation.CoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: detail.FieldCoinUsdCurrency,
		})
	}
	if value, ok := du.mutation.IoExtra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldIoExtra,
		})
	}
	if du.mutation.IoExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldIoExtra,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{detail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DetailUpdateOne is the builder for updating a single Detail entity.
type DetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DetailMutation
}

// SetCreatedAt sets the "created_at" field.
func (duo *DetailUpdateOne) SetCreatedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetCreatedAt()
	duo.mutation.SetCreatedAt(u)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCreatedAt(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetCreatedAt(*u)
	}
	return duo
}

// AddCreatedAt adds u to the "created_at" field.
func (duo *DetailUpdateOne) AddCreatedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddCreatedAt(u)
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DetailUpdateOne) SetUpdatedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetUpdatedAt()
	duo.mutation.SetUpdatedAt(u)
	return duo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (duo *DetailUpdateOne) AddUpdatedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddUpdatedAt(u)
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DetailUpdateOne) SetDeletedAt(u uint32) *DetailUpdateOne {
	duo.mutation.ResetDeletedAt()
	duo.mutation.SetDeletedAt(u)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableDeletedAt(u *uint32) *DetailUpdateOne {
	if u != nil {
		duo.SetDeletedAt(*u)
	}
	return duo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (duo *DetailUpdateOne) AddDeletedAt(u int32) *DetailUpdateOne {
	duo.mutation.AddDeletedAt(u)
	return duo
}

// SetAppID sets the "app_id" field.
func (duo *DetailUpdateOne) SetAppID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetAppID(u)
	return duo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableAppID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetAppID(*u)
	}
	return duo
}

// ClearAppID clears the value of the "app_id" field.
func (duo *DetailUpdateOne) ClearAppID() *DetailUpdateOne {
	duo.mutation.ClearAppID()
	return duo
}

// SetUserID sets the "user_id" field.
func (duo *DetailUpdateOne) SetUserID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetUserID(u)
	return duo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableUserID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetUserID(*u)
	}
	return duo
}

// ClearUserID clears the value of the "user_id" field.
func (duo *DetailUpdateOne) ClearUserID() *DetailUpdateOne {
	duo.mutation.ClearUserID()
	return duo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (duo *DetailUpdateOne) SetCoinTypeID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetCoinTypeID(u)
	return duo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetCoinTypeID(*u)
	}
	return duo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (duo *DetailUpdateOne) ClearCoinTypeID() *DetailUpdateOne {
	duo.mutation.ClearCoinTypeID()
	return duo
}

// SetIoType sets the "io_type" field.
func (duo *DetailUpdateOne) SetIoType(s string) *DetailUpdateOne {
	duo.mutation.SetIoType(s)
	return duo
}

// SetNillableIoType sets the "io_type" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableIoType(s *string) *DetailUpdateOne {
	if s != nil {
		duo.SetIoType(*s)
	}
	return duo
}

// ClearIoType clears the value of the "io_type" field.
func (duo *DetailUpdateOne) ClearIoType() *DetailUpdateOne {
	duo.mutation.ClearIoType()
	return duo
}

// SetIoSubType sets the "io_sub_type" field.
func (duo *DetailUpdateOne) SetIoSubType(s string) *DetailUpdateOne {
	duo.mutation.SetIoSubType(s)
	return duo
}

// SetNillableIoSubType sets the "io_sub_type" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableIoSubType(s *string) *DetailUpdateOne {
	if s != nil {
		duo.SetIoSubType(*s)
	}
	return duo
}

// ClearIoSubType clears the value of the "io_sub_type" field.
func (duo *DetailUpdateOne) ClearIoSubType() *DetailUpdateOne {
	duo.mutation.ClearIoSubType()
	return duo
}

// SetAmount sets the "amount" field.
func (duo *DetailUpdateOne) SetAmount(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.ResetAmount()
	duo.mutation.SetAmount(d)
	return duo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableAmount(d *decimal.Decimal) *DetailUpdateOne {
	if d != nil {
		duo.SetAmount(*d)
	}
	return duo
}

// AddAmount adds d to the "amount" field.
func (duo *DetailUpdateOne) AddAmount(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.AddAmount(d)
	return duo
}

// ClearAmount clears the value of the "amount" field.
func (duo *DetailUpdateOne) ClearAmount() *DetailUpdateOne {
	duo.mutation.ClearAmount()
	return duo
}

// SetFromCoinTypeID sets the "from_coin_type_id" field.
func (duo *DetailUpdateOne) SetFromCoinTypeID(u uuid.UUID) *DetailUpdateOne {
	duo.mutation.SetFromCoinTypeID(u)
	return duo
}

// SetNillableFromCoinTypeID sets the "from_coin_type_id" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableFromCoinTypeID(u *uuid.UUID) *DetailUpdateOne {
	if u != nil {
		duo.SetFromCoinTypeID(*u)
	}
	return duo
}

// ClearFromCoinTypeID clears the value of the "from_coin_type_id" field.
func (duo *DetailUpdateOne) ClearFromCoinTypeID() *DetailUpdateOne {
	duo.mutation.ClearFromCoinTypeID()
	return duo
}

// SetCoinUsdCurrency sets the "coin_usd_currency" field.
func (duo *DetailUpdateOne) SetCoinUsdCurrency(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.ResetCoinUsdCurrency()
	duo.mutation.SetCoinUsdCurrency(d)
	return duo
}

// SetNillableCoinUsdCurrency sets the "coin_usd_currency" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableCoinUsdCurrency(d *decimal.Decimal) *DetailUpdateOne {
	if d != nil {
		duo.SetCoinUsdCurrency(*d)
	}
	return duo
}

// AddCoinUsdCurrency adds d to the "coin_usd_currency" field.
func (duo *DetailUpdateOne) AddCoinUsdCurrency(d decimal.Decimal) *DetailUpdateOne {
	duo.mutation.AddCoinUsdCurrency(d)
	return duo
}

// ClearCoinUsdCurrency clears the value of the "coin_usd_currency" field.
func (duo *DetailUpdateOne) ClearCoinUsdCurrency() *DetailUpdateOne {
	duo.mutation.ClearCoinUsdCurrency()
	return duo
}

// SetIoExtra sets the "io_extra" field.
func (duo *DetailUpdateOne) SetIoExtra(s string) *DetailUpdateOne {
	duo.mutation.SetIoExtra(s)
	return duo
}

// SetNillableIoExtra sets the "io_extra" field if the given value is not nil.
func (duo *DetailUpdateOne) SetNillableIoExtra(s *string) *DetailUpdateOne {
	if s != nil {
		duo.SetIoExtra(*s)
	}
	return duo
}

// ClearIoExtra clears the value of the "io_extra" field.
func (duo *DetailUpdateOne) ClearIoExtra() *DetailUpdateOne {
	duo.mutation.ClearIoExtra()
	return duo
}

// Mutation returns the DetailMutation object of the builder.
func (duo *DetailUpdateOne) Mutation() *DetailMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DetailUpdateOne) Select(field string, fields ...string) *DetailUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Detail entity.
func (duo *DetailUpdateOne) Save(ctx context.Context) (*Detail, error) {
	var (
		err  error
		node *Detail
	)
	if err := duo.defaults(); err != nil {
		return nil, err
	}
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Detail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DetailUpdateOne) SaveX(ctx context.Context) *Detail {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DetailUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DetailUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DetailUpdateOne) defaults() error {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		if detail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized detail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := detail.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (duo *DetailUpdateOne) check() error {
	if v, ok := duo.mutation.IoExtra(); ok {
		if err := detail.IoExtraValidator(v); err != nil {
			return &ValidationError{Name: "io_extra", err: fmt.Errorf(`ent: validator failed for field "Detail.io_extra": %w`, err)}
		}
	}
	return nil
}

func (duo *DetailUpdateOne) sqlSave(ctx context.Context) (_node *Detail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Detail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, detail.FieldID)
		for _, f := range fields {
			if !detail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != detail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: detail.FieldDeletedAt,
		})
	}
	if value, ok := duo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldAppID,
		})
	}
	if duo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldAppID,
		})
	}
	if value, ok := duo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldUserID,
		})
	}
	if duo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldUserID,
		})
	}
	if value, ok := duo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldCoinTypeID,
		})
	}
	if duo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldCoinTypeID,
		})
	}
	if value, ok := duo.mutation.IoType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldIoType,
		})
	}
	if duo.mutation.IoTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldIoType,
		})
	}
	if value, ok := duo.mutation.IoSubType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldIoSubType,
		})
	}
	if duo.mutation.IoSubTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldIoSubType,
		})
	}
	if value, ok := duo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldAmount,
		})
	}
	if value, ok := duo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldAmount,
		})
	}
	if duo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: detail.FieldAmount,
		})
	}
	if value, ok := duo.mutation.FromCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: detail.FieldFromCoinTypeID,
		})
	}
	if duo.mutation.FromCoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: detail.FieldFromCoinTypeID,
		})
	}
	if value, ok := duo.mutation.CoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldCoinUsdCurrency,
		})
	}
	if value, ok := duo.mutation.AddedCoinUsdCurrency(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: detail.FieldCoinUsdCurrency,
		})
	}
	if duo.mutation.CoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: detail.FieldCoinUsdCurrency,
		})
	}
	if value, ok := duo.mutation.IoExtra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: detail.FieldIoExtra,
		})
	}
	if duo.mutation.IoExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: detail.FieldIoExtra,
		})
	}
	_node = &Detail{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{detail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
