// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/miningdetail"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// MiningDetailUpdate is the builder for updating MiningDetail entities.
type MiningDetailUpdate struct {
	config
	hooks    []Hook
	mutation *MiningDetailMutation
}

// Where appends a list predicates to the MiningDetailUpdate builder.
func (mdu *MiningDetailUpdate) Where(ps ...predicate.MiningDetail) *MiningDetailUpdate {
	mdu.mutation.Where(ps...)
	return mdu
}

// SetCreatedAt sets the "created_at" field.
func (mdu *MiningDetailUpdate) SetCreatedAt(u uint32) *MiningDetailUpdate {
	mdu.mutation.ResetCreatedAt()
	mdu.mutation.SetCreatedAt(u)
	return mdu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mdu *MiningDetailUpdate) SetNillableCreatedAt(u *uint32) *MiningDetailUpdate {
	if u != nil {
		mdu.SetCreatedAt(*u)
	}
	return mdu
}

// AddCreatedAt adds u to the "created_at" field.
func (mdu *MiningDetailUpdate) AddCreatedAt(u int32) *MiningDetailUpdate {
	mdu.mutation.AddCreatedAt(u)
	return mdu
}

// SetUpdatedAt sets the "updated_at" field.
func (mdu *MiningDetailUpdate) SetUpdatedAt(u uint32) *MiningDetailUpdate {
	mdu.mutation.ResetUpdatedAt()
	mdu.mutation.SetUpdatedAt(u)
	return mdu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (mdu *MiningDetailUpdate) AddUpdatedAt(u int32) *MiningDetailUpdate {
	mdu.mutation.AddUpdatedAt(u)
	return mdu
}

// SetDeletedAt sets the "deleted_at" field.
func (mdu *MiningDetailUpdate) SetDeletedAt(u uint32) *MiningDetailUpdate {
	mdu.mutation.ResetDeletedAt()
	mdu.mutation.SetDeletedAt(u)
	return mdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mdu *MiningDetailUpdate) SetNillableDeletedAt(u *uint32) *MiningDetailUpdate {
	if u != nil {
		mdu.SetDeletedAt(*u)
	}
	return mdu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (mdu *MiningDetailUpdate) AddDeletedAt(u int32) *MiningDetailUpdate {
	mdu.mutation.AddDeletedAt(u)
	return mdu
}

// SetGoodID sets the "good_id" field.
func (mdu *MiningDetailUpdate) SetGoodID(u uuid.UUID) *MiningDetailUpdate {
	mdu.mutation.SetGoodID(u)
	return mdu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (mdu *MiningDetailUpdate) SetNillableGoodID(u *uuid.UUID) *MiningDetailUpdate {
	if u != nil {
		mdu.SetGoodID(*u)
	}
	return mdu
}

// ClearGoodID clears the value of the "good_id" field.
func (mdu *MiningDetailUpdate) ClearGoodID() *MiningDetailUpdate {
	mdu.mutation.ClearGoodID()
	return mdu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (mdu *MiningDetailUpdate) SetCoinTypeID(u uuid.UUID) *MiningDetailUpdate {
	mdu.mutation.SetCoinTypeID(u)
	return mdu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (mdu *MiningDetailUpdate) SetNillableCoinTypeID(u *uuid.UUID) *MiningDetailUpdate {
	if u != nil {
		mdu.SetCoinTypeID(*u)
	}
	return mdu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (mdu *MiningDetailUpdate) ClearCoinTypeID() *MiningDetailUpdate {
	mdu.mutation.ClearCoinTypeID()
	return mdu
}

// SetAmount sets the "amount" field.
func (mdu *MiningDetailUpdate) SetAmount(d decimal.Decimal) *MiningDetailUpdate {
	mdu.mutation.ResetAmount()
	mdu.mutation.SetAmount(d)
	return mdu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mdu *MiningDetailUpdate) SetNillableAmount(d *decimal.Decimal) *MiningDetailUpdate {
	if d != nil {
		mdu.SetAmount(*d)
	}
	return mdu
}

// AddAmount adds d to the "amount" field.
func (mdu *MiningDetailUpdate) AddAmount(d decimal.Decimal) *MiningDetailUpdate {
	mdu.mutation.AddAmount(d)
	return mdu
}

// ClearAmount clears the value of the "amount" field.
func (mdu *MiningDetailUpdate) ClearAmount() *MiningDetailUpdate {
	mdu.mutation.ClearAmount()
	return mdu
}

// SetBenefitDate sets the "benefit_date" field.
func (mdu *MiningDetailUpdate) SetBenefitDate(u uint32) *MiningDetailUpdate {
	mdu.mutation.ResetBenefitDate()
	mdu.mutation.SetBenefitDate(u)
	return mdu
}

// SetNillableBenefitDate sets the "benefit_date" field if the given value is not nil.
func (mdu *MiningDetailUpdate) SetNillableBenefitDate(u *uint32) *MiningDetailUpdate {
	if u != nil {
		mdu.SetBenefitDate(*u)
	}
	return mdu
}

// AddBenefitDate adds u to the "benefit_date" field.
func (mdu *MiningDetailUpdate) AddBenefitDate(u int32) *MiningDetailUpdate {
	mdu.mutation.AddBenefitDate(u)
	return mdu
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (mdu *MiningDetailUpdate) ClearBenefitDate() *MiningDetailUpdate {
	mdu.mutation.ClearBenefitDate()
	return mdu
}

// Mutation returns the MiningDetailMutation object of the builder.
func (mdu *MiningDetailUpdate) Mutation() *MiningDetailMutation {
	return mdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mdu *MiningDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := mdu.defaults(); err != nil {
		return 0, err
	}
	if len(mdu.hooks) == 0 {
		affected, err = mdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiningDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mdu.mutation = mutation
			affected, err = mdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mdu.hooks) - 1; i >= 0; i-- {
			if mdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mdu *MiningDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := mdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mdu *MiningDetailUpdate) Exec(ctx context.Context) error {
	_, err := mdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdu *MiningDetailUpdate) ExecX(ctx context.Context) {
	if err := mdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mdu *MiningDetailUpdate) defaults() error {
	if _, ok := mdu.mutation.UpdatedAt(); !ok {
		if miningdetail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized miningdetail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := miningdetail.UpdateDefaultUpdatedAt()
		mdu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (mdu *MiningDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   miningdetail.Table,
			Columns: miningdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: miningdetail.FieldID,
			},
		},
	}
	if ps := mdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mdu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldCreatedAt,
		})
	}
	if value, ok := mdu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldCreatedAt,
		})
	}
	if value, ok := mdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldUpdatedAt,
		})
	}
	if value, ok := mdu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldUpdatedAt,
		})
	}
	if value, ok := mdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldDeletedAt,
		})
	}
	if value, ok := mdu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldDeletedAt,
		})
	}
	if value, ok := mdu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: miningdetail.FieldGoodID,
		})
	}
	if mdu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: miningdetail.FieldGoodID,
		})
	}
	if value, ok := mdu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: miningdetail.FieldCoinTypeID,
		})
	}
	if mdu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: miningdetail.FieldCoinTypeID,
		})
	}
	if value, ok := mdu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: miningdetail.FieldAmount,
		})
	}
	if value, ok := mdu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: miningdetail.FieldAmount,
		})
	}
	if mdu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: miningdetail.FieldAmount,
		})
	}
	if value, ok := mdu.mutation.BenefitDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldBenefitDate,
		})
	}
	if value, ok := mdu.mutation.AddedBenefitDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldBenefitDate,
		})
	}
	if mdu.mutation.BenefitDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: miningdetail.FieldBenefitDate,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{miningdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MiningDetailUpdateOne is the builder for updating a single MiningDetail entity.
type MiningDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MiningDetailMutation
}

// SetCreatedAt sets the "created_at" field.
func (mduo *MiningDetailUpdateOne) SetCreatedAt(u uint32) *MiningDetailUpdateOne {
	mduo.mutation.ResetCreatedAt()
	mduo.mutation.SetCreatedAt(u)
	return mduo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mduo *MiningDetailUpdateOne) SetNillableCreatedAt(u *uint32) *MiningDetailUpdateOne {
	if u != nil {
		mduo.SetCreatedAt(*u)
	}
	return mduo
}

// AddCreatedAt adds u to the "created_at" field.
func (mduo *MiningDetailUpdateOne) AddCreatedAt(u int32) *MiningDetailUpdateOne {
	mduo.mutation.AddCreatedAt(u)
	return mduo
}

// SetUpdatedAt sets the "updated_at" field.
func (mduo *MiningDetailUpdateOne) SetUpdatedAt(u uint32) *MiningDetailUpdateOne {
	mduo.mutation.ResetUpdatedAt()
	mduo.mutation.SetUpdatedAt(u)
	return mduo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (mduo *MiningDetailUpdateOne) AddUpdatedAt(u int32) *MiningDetailUpdateOne {
	mduo.mutation.AddUpdatedAt(u)
	return mduo
}

// SetDeletedAt sets the "deleted_at" field.
func (mduo *MiningDetailUpdateOne) SetDeletedAt(u uint32) *MiningDetailUpdateOne {
	mduo.mutation.ResetDeletedAt()
	mduo.mutation.SetDeletedAt(u)
	return mduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mduo *MiningDetailUpdateOne) SetNillableDeletedAt(u *uint32) *MiningDetailUpdateOne {
	if u != nil {
		mduo.SetDeletedAt(*u)
	}
	return mduo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (mduo *MiningDetailUpdateOne) AddDeletedAt(u int32) *MiningDetailUpdateOne {
	mduo.mutation.AddDeletedAt(u)
	return mduo
}

// SetGoodID sets the "good_id" field.
func (mduo *MiningDetailUpdateOne) SetGoodID(u uuid.UUID) *MiningDetailUpdateOne {
	mduo.mutation.SetGoodID(u)
	return mduo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (mduo *MiningDetailUpdateOne) SetNillableGoodID(u *uuid.UUID) *MiningDetailUpdateOne {
	if u != nil {
		mduo.SetGoodID(*u)
	}
	return mduo
}

// ClearGoodID clears the value of the "good_id" field.
func (mduo *MiningDetailUpdateOne) ClearGoodID() *MiningDetailUpdateOne {
	mduo.mutation.ClearGoodID()
	return mduo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (mduo *MiningDetailUpdateOne) SetCoinTypeID(u uuid.UUID) *MiningDetailUpdateOne {
	mduo.mutation.SetCoinTypeID(u)
	return mduo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (mduo *MiningDetailUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *MiningDetailUpdateOne {
	if u != nil {
		mduo.SetCoinTypeID(*u)
	}
	return mduo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (mduo *MiningDetailUpdateOne) ClearCoinTypeID() *MiningDetailUpdateOne {
	mduo.mutation.ClearCoinTypeID()
	return mduo
}

// SetAmount sets the "amount" field.
func (mduo *MiningDetailUpdateOne) SetAmount(d decimal.Decimal) *MiningDetailUpdateOne {
	mduo.mutation.ResetAmount()
	mduo.mutation.SetAmount(d)
	return mduo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mduo *MiningDetailUpdateOne) SetNillableAmount(d *decimal.Decimal) *MiningDetailUpdateOne {
	if d != nil {
		mduo.SetAmount(*d)
	}
	return mduo
}

// AddAmount adds d to the "amount" field.
func (mduo *MiningDetailUpdateOne) AddAmount(d decimal.Decimal) *MiningDetailUpdateOne {
	mduo.mutation.AddAmount(d)
	return mduo
}

// ClearAmount clears the value of the "amount" field.
func (mduo *MiningDetailUpdateOne) ClearAmount() *MiningDetailUpdateOne {
	mduo.mutation.ClearAmount()
	return mduo
}

// SetBenefitDate sets the "benefit_date" field.
func (mduo *MiningDetailUpdateOne) SetBenefitDate(u uint32) *MiningDetailUpdateOne {
	mduo.mutation.ResetBenefitDate()
	mduo.mutation.SetBenefitDate(u)
	return mduo
}

// SetNillableBenefitDate sets the "benefit_date" field if the given value is not nil.
func (mduo *MiningDetailUpdateOne) SetNillableBenefitDate(u *uint32) *MiningDetailUpdateOne {
	if u != nil {
		mduo.SetBenefitDate(*u)
	}
	return mduo
}

// AddBenefitDate adds u to the "benefit_date" field.
func (mduo *MiningDetailUpdateOne) AddBenefitDate(u int32) *MiningDetailUpdateOne {
	mduo.mutation.AddBenefitDate(u)
	return mduo
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (mduo *MiningDetailUpdateOne) ClearBenefitDate() *MiningDetailUpdateOne {
	mduo.mutation.ClearBenefitDate()
	return mduo
}

// Mutation returns the MiningDetailMutation object of the builder.
func (mduo *MiningDetailUpdateOne) Mutation() *MiningDetailMutation {
	return mduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mduo *MiningDetailUpdateOne) Select(field string, fields ...string) *MiningDetailUpdateOne {
	mduo.fields = append([]string{field}, fields...)
	return mduo
}

// Save executes the query and returns the updated MiningDetail entity.
func (mduo *MiningDetailUpdateOne) Save(ctx context.Context) (*MiningDetail, error) {
	var (
		err  error
		node *MiningDetail
	)
	if err := mduo.defaults(); err != nil {
		return nil, err
	}
	if len(mduo.hooks) == 0 {
		node, err = mduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiningDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mduo.mutation = mutation
			node, err = mduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mduo.hooks) - 1; i >= 0; i-- {
			if mduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MiningDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MiningDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mduo *MiningDetailUpdateOne) SaveX(ctx context.Context) *MiningDetail {
	node, err := mduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mduo *MiningDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := mduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mduo *MiningDetailUpdateOne) ExecX(ctx context.Context) {
	if err := mduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mduo *MiningDetailUpdateOne) defaults() error {
	if _, ok := mduo.mutation.UpdatedAt(); !ok {
		if miningdetail.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized miningdetail.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := miningdetail.UpdateDefaultUpdatedAt()
		mduo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (mduo *MiningDetailUpdateOne) sqlSave(ctx context.Context) (_node *MiningDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   miningdetail.Table,
			Columns: miningdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: miningdetail.FieldID,
			},
		},
	}
	id, ok := mduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MiningDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, miningdetail.FieldID)
		for _, f := range fields {
			if !miningdetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != miningdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mduo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldCreatedAt,
		})
	}
	if value, ok := mduo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldCreatedAt,
		})
	}
	if value, ok := mduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldUpdatedAt,
		})
	}
	if value, ok := mduo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldUpdatedAt,
		})
	}
	if value, ok := mduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldDeletedAt,
		})
	}
	if value, ok := mduo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldDeletedAt,
		})
	}
	if value, ok := mduo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: miningdetail.FieldGoodID,
		})
	}
	if mduo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: miningdetail.FieldGoodID,
		})
	}
	if value, ok := mduo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: miningdetail.FieldCoinTypeID,
		})
	}
	if mduo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: miningdetail.FieldCoinTypeID,
		})
	}
	if value, ok := mduo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: miningdetail.FieldAmount,
		})
	}
	if value, ok := mduo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: miningdetail.FieldAmount,
		})
	}
	if mduo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: miningdetail.FieldAmount,
		})
	}
	if value, ok := mduo.mutation.BenefitDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldBenefitDate,
		})
	}
	if value, ok := mduo.mutation.AddedBenefitDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: miningdetail.FieldBenefitDate,
		})
	}
	if mduo.mutation.BenefitDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: miningdetail.FieldBenefitDate,
		})
	}
	_node = &MiningDetail{config: mduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{miningdetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}