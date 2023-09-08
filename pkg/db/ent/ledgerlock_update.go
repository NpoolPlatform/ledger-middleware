// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledgerlock"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/shopspring/decimal"
)

// LedgerLockUpdate is the builder for updating LedgerLock entities.
type LedgerLockUpdate struct {
	config
	hooks    []Hook
	mutation *LedgerLockMutation
}

// Where appends a list predicates to the LedgerLockUpdate builder.
func (llu *LedgerLockUpdate) Where(ps ...predicate.LedgerLock) *LedgerLockUpdate {
	llu.mutation.Where(ps...)
	return llu
}

// SetCreatedAt sets the "created_at" field.
func (llu *LedgerLockUpdate) SetCreatedAt(u uint32) *LedgerLockUpdate {
	llu.mutation.ResetCreatedAt()
	llu.mutation.SetCreatedAt(u)
	return llu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (llu *LedgerLockUpdate) SetNillableCreatedAt(u *uint32) *LedgerLockUpdate {
	if u != nil {
		llu.SetCreatedAt(*u)
	}
	return llu
}

// AddCreatedAt adds u to the "created_at" field.
func (llu *LedgerLockUpdate) AddCreatedAt(u int32) *LedgerLockUpdate {
	llu.mutation.AddCreatedAt(u)
	return llu
}

// SetUpdatedAt sets the "updated_at" field.
func (llu *LedgerLockUpdate) SetUpdatedAt(u uint32) *LedgerLockUpdate {
	llu.mutation.ResetUpdatedAt()
	llu.mutation.SetUpdatedAt(u)
	return llu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (llu *LedgerLockUpdate) AddUpdatedAt(u int32) *LedgerLockUpdate {
	llu.mutation.AddUpdatedAt(u)
	return llu
}

// SetDeletedAt sets the "deleted_at" field.
func (llu *LedgerLockUpdate) SetDeletedAt(u uint32) *LedgerLockUpdate {
	llu.mutation.ResetDeletedAt()
	llu.mutation.SetDeletedAt(u)
	return llu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (llu *LedgerLockUpdate) SetNillableDeletedAt(u *uint32) *LedgerLockUpdate {
	if u != nil {
		llu.SetDeletedAt(*u)
	}
	return llu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (llu *LedgerLockUpdate) AddDeletedAt(u int32) *LedgerLockUpdate {
	llu.mutation.AddDeletedAt(u)
	return llu
}

// SetAmount sets the "amount" field.
func (llu *LedgerLockUpdate) SetAmount(d decimal.Decimal) *LedgerLockUpdate {
	llu.mutation.ResetAmount()
	llu.mutation.SetAmount(d)
	return llu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (llu *LedgerLockUpdate) SetNillableAmount(d *decimal.Decimal) *LedgerLockUpdate {
	if d != nil {
		llu.SetAmount(*d)
	}
	return llu
}

// AddAmount adds d to the "amount" field.
func (llu *LedgerLockUpdate) AddAmount(d decimal.Decimal) *LedgerLockUpdate {
	llu.mutation.AddAmount(d)
	return llu
}

// ClearAmount clears the value of the "amount" field.
func (llu *LedgerLockUpdate) ClearAmount() *LedgerLockUpdate {
	llu.mutation.ClearAmount()
	return llu
}

// Mutation returns the LedgerLockMutation object of the builder.
func (llu *LedgerLockUpdate) Mutation() *LedgerLockMutation {
	return llu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (llu *LedgerLockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := llu.defaults(); err != nil {
		return 0, err
	}
	if len(llu.hooks) == 0 {
		affected, err = llu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LedgerLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			llu.mutation = mutation
			affected, err = llu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(llu.hooks) - 1; i >= 0; i-- {
			if llu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = llu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, llu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (llu *LedgerLockUpdate) SaveX(ctx context.Context) int {
	affected, err := llu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (llu *LedgerLockUpdate) Exec(ctx context.Context) error {
	_, err := llu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llu *LedgerLockUpdate) ExecX(ctx context.Context) {
	if err := llu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (llu *LedgerLockUpdate) defaults() error {
	if _, ok := llu.mutation.UpdatedAt(); !ok {
		if ledgerlock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ledgerlock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ledgerlock.UpdateDefaultUpdatedAt()
		llu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (llu *LedgerLockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ledgerlock.Table,
			Columns: ledgerlock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ledgerlock.FieldID,
			},
		},
	}
	if ps := llu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := llu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldCreatedAt,
		})
	}
	if value, ok := llu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldCreatedAt,
		})
	}
	if value, ok := llu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldUpdatedAt,
		})
	}
	if value, ok := llu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldUpdatedAt,
		})
	}
	if value, ok := llu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldDeletedAt,
		})
	}
	if value, ok := llu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldDeletedAt,
		})
	}
	if value, ok := llu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ledgerlock.FieldAmount,
		})
	}
	if value, ok := llu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ledgerlock.FieldAmount,
		})
	}
	if llu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: ledgerlock.FieldAmount,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, llu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ledgerlock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LedgerLockUpdateOne is the builder for updating a single LedgerLock entity.
type LedgerLockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LedgerLockMutation
}

// SetCreatedAt sets the "created_at" field.
func (lluo *LedgerLockUpdateOne) SetCreatedAt(u uint32) *LedgerLockUpdateOne {
	lluo.mutation.ResetCreatedAt()
	lluo.mutation.SetCreatedAt(u)
	return lluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lluo *LedgerLockUpdateOne) SetNillableCreatedAt(u *uint32) *LedgerLockUpdateOne {
	if u != nil {
		lluo.SetCreatedAt(*u)
	}
	return lluo
}

// AddCreatedAt adds u to the "created_at" field.
func (lluo *LedgerLockUpdateOne) AddCreatedAt(u int32) *LedgerLockUpdateOne {
	lluo.mutation.AddCreatedAt(u)
	return lluo
}

// SetUpdatedAt sets the "updated_at" field.
func (lluo *LedgerLockUpdateOne) SetUpdatedAt(u uint32) *LedgerLockUpdateOne {
	lluo.mutation.ResetUpdatedAt()
	lluo.mutation.SetUpdatedAt(u)
	return lluo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (lluo *LedgerLockUpdateOne) AddUpdatedAt(u int32) *LedgerLockUpdateOne {
	lluo.mutation.AddUpdatedAt(u)
	return lluo
}

// SetDeletedAt sets the "deleted_at" field.
func (lluo *LedgerLockUpdateOne) SetDeletedAt(u uint32) *LedgerLockUpdateOne {
	lluo.mutation.ResetDeletedAt()
	lluo.mutation.SetDeletedAt(u)
	return lluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lluo *LedgerLockUpdateOne) SetNillableDeletedAt(u *uint32) *LedgerLockUpdateOne {
	if u != nil {
		lluo.SetDeletedAt(*u)
	}
	return lluo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (lluo *LedgerLockUpdateOne) AddDeletedAt(u int32) *LedgerLockUpdateOne {
	lluo.mutation.AddDeletedAt(u)
	return lluo
}

// SetAmount sets the "amount" field.
func (lluo *LedgerLockUpdateOne) SetAmount(d decimal.Decimal) *LedgerLockUpdateOne {
	lluo.mutation.ResetAmount()
	lluo.mutation.SetAmount(d)
	return lluo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (lluo *LedgerLockUpdateOne) SetNillableAmount(d *decimal.Decimal) *LedgerLockUpdateOne {
	if d != nil {
		lluo.SetAmount(*d)
	}
	return lluo
}

// AddAmount adds d to the "amount" field.
func (lluo *LedgerLockUpdateOne) AddAmount(d decimal.Decimal) *LedgerLockUpdateOne {
	lluo.mutation.AddAmount(d)
	return lluo
}

// ClearAmount clears the value of the "amount" field.
func (lluo *LedgerLockUpdateOne) ClearAmount() *LedgerLockUpdateOne {
	lluo.mutation.ClearAmount()
	return lluo
}

// Mutation returns the LedgerLockMutation object of the builder.
func (lluo *LedgerLockUpdateOne) Mutation() *LedgerLockMutation {
	return lluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lluo *LedgerLockUpdateOne) Select(field string, fields ...string) *LedgerLockUpdateOne {
	lluo.fields = append([]string{field}, fields...)
	return lluo
}

// Save executes the query and returns the updated LedgerLock entity.
func (lluo *LedgerLockUpdateOne) Save(ctx context.Context) (*LedgerLock, error) {
	var (
		err  error
		node *LedgerLock
	)
	if err := lluo.defaults(); err != nil {
		return nil, err
	}
	if len(lluo.hooks) == 0 {
		node, err = lluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LedgerLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lluo.mutation = mutation
			node, err = lluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lluo.hooks) - 1; i >= 0; i-- {
			if lluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LedgerLock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LedgerLockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lluo *LedgerLockUpdateOne) SaveX(ctx context.Context) *LedgerLock {
	node, err := lluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lluo *LedgerLockUpdateOne) Exec(ctx context.Context) error {
	_, err := lluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lluo *LedgerLockUpdateOne) ExecX(ctx context.Context) {
	if err := lluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lluo *LedgerLockUpdateOne) defaults() error {
	if _, ok := lluo.mutation.UpdatedAt(); !ok {
		if ledgerlock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ledgerlock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ledgerlock.UpdateDefaultUpdatedAt()
		lluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (lluo *LedgerLockUpdateOne) sqlSave(ctx context.Context) (_node *LedgerLock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ledgerlock.Table,
			Columns: ledgerlock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ledgerlock.FieldID,
			},
		},
	}
	id, ok := lluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LedgerLock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ledgerlock.FieldID)
		for _, f := range fields {
			if !ledgerlock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ledgerlock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldCreatedAt,
		})
	}
	if value, ok := lluo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldCreatedAt,
		})
	}
	if value, ok := lluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldUpdatedAt,
		})
	}
	if value, ok := lluo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldUpdatedAt,
		})
	}
	if value, ok := lluo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldDeletedAt,
		})
	}
	if value, ok := lluo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ledgerlock.FieldDeletedAt,
		})
	}
	if value, ok := lluo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ledgerlock.FieldAmount,
		})
	}
	if value, ok := lluo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ledgerlock.FieldAmount,
		})
	}
	if lluo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: ledgerlock.FieldAmount,
		})
	}
	_node = &LedgerLock{config: lluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ledgerlock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
