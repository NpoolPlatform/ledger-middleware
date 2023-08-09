// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UnsoldStatementCreate is the builder for creating a UnsoldStatement entity.
type UnsoldStatementCreate struct {
	config
	mutation *UnsoldStatementMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (usc *UnsoldStatementCreate) SetCreatedAt(u uint32) *UnsoldStatementCreate {
	usc.mutation.SetCreatedAt(u)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableCreatedAt(u *uint32) *UnsoldStatementCreate {
	if u != nil {
		usc.SetCreatedAt(*u)
	}
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UnsoldStatementCreate) SetUpdatedAt(u uint32) *UnsoldStatementCreate {
	usc.mutation.SetUpdatedAt(u)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableUpdatedAt(u *uint32) *UnsoldStatementCreate {
	if u != nil {
		usc.SetUpdatedAt(*u)
	}
	return usc
}

// SetDeletedAt sets the "deleted_at" field.
func (usc *UnsoldStatementCreate) SetDeletedAt(u uint32) *UnsoldStatementCreate {
	usc.mutation.SetDeletedAt(u)
	return usc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableDeletedAt(u *uint32) *UnsoldStatementCreate {
	if u != nil {
		usc.SetDeletedAt(*u)
	}
	return usc
}

// SetGoodID sets the "good_id" field.
func (usc *UnsoldStatementCreate) SetGoodID(u uuid.UUID) *UnsoldStatementCreate {
	usc.mutation.SetGoodID(u)
	return usc
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableGoodID(u *uuid.UUID) *UnsoldStatementCreate {
	if u != nil {
		usc.SetGoodID(*u)
	}
	return usc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (usc *UnsoldStatementCreate) SetCoinTypeID(u uuid.UUID) *UnsoldStatementCreate {
	usc.mutation.SetCoinTypeID(u)
	return usc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableCoinTypeID(u *uuid.UUID) *UnsoldStatementCreate {
	if u != nil {
		usc.SetCoinTypeID(*u)
	}
	return usc
}

// SetAmount sets the "amount" field.
func (usc *UnsoldStatementCreate) SetAmount(d decimal.Decimal) *UnsoldStatementCreate {
	usc.mutation.SetAmount(d)
	return usc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableAmount(d *decimal.Decimal) *UnsoldStatementCreate {
	if d != nil {
		usc.SetAmount(*d)
	}
	return usc
}

// SetBenefitDate sets the "benefit_date" field.
func (usc *UnsoldStatementCreate) SetBenefitDate(u uint32) *UnsoldStatementCreate {
	usc.mutation.SetBenefitDate(u)
	return usc
}

// SetNillableBenefitDate sets the "benefit_date" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableBenefitDate(u *uint32) *UnsoldStatementCreate {
	if u != nil {
		usc.SetBenefitDate(*u)
	}
	return usc
}

// SetID sets the "id" field.
func (usc *UnsoldStatementCreate) SetID(u uuid.UUID) *UnsoldStatementCreate {
	usc.mutation.SetID(u)
	return usc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (usc *UnsoldStatementCreate) SetNillableID(u *uuid.UUID) *UnsoldStatementCreate {
	if u != nil {
		usc.SetID(*u)
	}
	return usc
}

// Mutation returns the UnsoldStatementMutation object of the builder.
func (usc *UnsoldStatementCreate) Mutation() *UnsoldStatementMutation {
	return usc.mutation
}

// Save creates the UnsoldStatement in the database.
func (usc *UnsoldStatementCreate) Save(ctx context.Context) (*UnsoldStatement, error) {
	var (
		err  error
		node *UnsoldStatement
	)
	if err := usc.defaults(); err != nil {
		return nil, err
	}
	if len(usc.hooks) == 0 {
		if err = usc.check(); err != nil {
			return nil, err
		}
		node, err = usc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsoldStatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usc.check(); err != nil {
				return nil, err
			}
			usc.mutation = mutation
			if node, err = usc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(usc.hooks) - 1; i >= 0; i-- {
			if usc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, usc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UnsoldStatement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UnsoldStatementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UnsoldStatementCreate) SaveX(ctx context.Context) *UnsoldStatement {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UnsoldStatementCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UnsoldStatementCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UnsoldStatementCreate) defaults() error {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		if unsoldstatement.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized unsoldstatement.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := unsoldstatement.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		if unsoldstatement.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized unsoldstatement.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := unsoldstatement.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
	if _, ok := usc.mutation.DeletedAt(); !ok {
		if unsoldstatement.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized unsoldstatement.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := unsoldstatement.DefaultDeletedAt()
		usc.mutation.SetDeletedAt(v)
	}
	if _, ok := usc.mutation.GoodID(); !ok {
		if unsoldstatement.DefaultGoodID == nil {
			return fmt.Errorf("ent: uninitialized unsoldstatement.DefaultGoodID (forgotten import ent/runtime?)")
		}
		v := unsoldstatement.DefaultGoodID()
		usc.mutation.SetGoodID(v)
	}
	if _, ok := usc.mutation.CoinTypeID(); !ok {
		if unsoldstatement.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized unsoldstatement.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := unsoldstatement.DefaultCoinTypeID()
		usc.mutation.SetCoinTypeID(v)
	}
	if _, ok := usc.mutation.BenefitDate(); !ok {
		v := unsoldstatement.DefaultBenefitDate
		usc.mutation.SetBenefitDate(v)
	}
	if _, ok := usc.mutation.ID(); !ok {
		if unsoldstatement.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized unsoldstatement.DefaultID (forgotten import ent/runtime?)")
		}
		v := unsoldstatement.DefaultID()
		usc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usc *UnsoldStatementCreate) check() error {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UnsoldStatement.created_at"`)}
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UnsoldStatement.updated_at"`)}
	}
	if _, ok := usc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "UnsoldStatement.deleted_at"`)}
	}
	return nil
}

func (usc *UnsoldStatementCreate) sqlSave(ctx context.Context) (*UnsoldStatement, error) {
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (usc *UnsoldStatementCreate) createSpec() (*UnsoldStatement, *sqlgraph.CreateSpec) {
	var (
		_node = &UnsoldStatement{config: usc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: unsoldstatement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: unsoldstatement.FieldID,
			},
		}
	)
	_spec.OnConflict = usc.conflict
	if id, ok := usc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsoldstatement.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsoldstatement.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := usc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsoldstatement.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := usc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: unsoldstatement.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := usc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: unsoldstatement.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := usc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: unsoldstatement.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := usc.mutation.BenefitDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsoldstatement.FieldBenefitDate,
		})
		_node.BenefitDate = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UnsoldStatement.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UnsoldStatementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (usc *UnsoldStatementCreate) OnConflict(opts ...sql.ConflictOption) *UnsoldStatementUpsertOne {
	usc.conflict = opts
	return &UnsoldStatementUpsertOne{
		create: usc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UnsoldStatement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (usc *UnsoldStatementCreate) OnConflictColumns(columns ...string) *UnsoldStatementUpsertOne {
	usc.conflict = append(usc.conflict, sql.ConflictColumns(columns...))
	return &UnsoldStatementUpsertOne{
		create: usc,
	}
}

type (
	// UnsoldStatementUpsertOne is the builder for "upsert"-ing
	//  one UnsoldStatement node.
	UnsoldStatementUpsertOne struct {
		create *UnsoldStatementCreate
	}

	// UnsoldStatementUpsert is the "OnConflict" setter.
	UnsoldStatementUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UnsoldStatementUpsert) SetCreatedAt(v uint32) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateCreatedAt() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UnsoldStatementUpsert) AddCreatedAt(v uint32) *UnsoldStatementUpsert {
	u.Add(unsoldstatement.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UnsoldStatementUpsert) SetUpdatedAt(v uint32) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateUpdatedAt() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UnsoldStatementUpsert) AddUpdatedAt(v uint32) *UnsoldStatementUpsert {
	u.Add(unsoldstatement.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UnsoldStatementUpsert) SetDeletedAt(v uint32) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateDeletedAt() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UnsoldStatementUpsert) AddDeletedAt(v uint32) *UnsoldStatementUpsert {
	u.Add(unsoldstatement.FieldDeletedAt, v)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *UnsoldStatementUpsert) SetGoodID(v uuid.UUID) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateGoodID() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *UnsoldStatementUpsert) ClearGoodID() *UnsoldStatementUpsert {
	u.SetNull(unsoldstatement.FieldGoodID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UnsoldStatementUpsert) SetCoinTypeID(v uuid.UUID) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateCoinTypeID() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UnsoldStatementUpsert) ClearCoinTypeID() *UnsoldStatementUpsert {
	u.SetNull(unsoldstatement.FieldCoinTypeID)
	return u
}

// SetAmount sets the "amount" field.
func (u *UnsoldStatementUpsert) SetAmount(v decimal.Decimal) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateAmount() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *UnsoldStatementUpsert) AddAmount(v decimal.Decimal) *UnsoldStatementUpsert {
	u.Add(unsoldstatement.FieldAmount, v)
	return u
}

// ClearAmount clears the value of the "amount" field.
func (u *UnsoldStatementUpsert) ClearAmount() *UnsoldStatementUpsert {
	u.SetNull(unsoldstatement.FieldAmount)
	return u
}

// SetBenefitDate sets the "benefit_date" field.
func (u *UnsoldStatementUpsert) SetBenefitDate(v uint32) *UnsoldStatementUpsert {
	u.Set(unsoldstatement.FieldBenefitDate, v)
	return u
}

// UpdateBenefitDate sets the "benefit_date" field to the value that was provided on create.
func (u *UnsoldStatementUpsert) UpdateBenefitDate() *UnsoldStatementUpsert {
	u.SetExcluded(unsoldstatement.FieldBenefitDate)
	return u
}

// AddBenefitDate adds v to the "benefit_date" field.
func (u *UnsoldStatementUpsert) AddBenefitDate(v uint32) *UnsoldStatementUpsert {
	u.Add(unsoldstatement.FieldBenefitDate, v)
	return u
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (u *UnsoldStatementUpsert) ClearBenefitDate() *UnsoldStatementUpsert {
	u.SetNull(unsoldstatement.FieldBenefitDate)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UnsoldStatement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(unsoldstatement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UnsoldStatementUpsertOne) UpdateNewValues() *UnsoldStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(unsoldstatement.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UnsoldStatement.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UnsoldStatementUpsertOne) Ignore() *UnsoldStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UnsoldStatementUpsertOne) DoNothing() *UnsoldStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UnsoldStatementCreate.OnConflict
// documentation for more info.
func (u *UnsoldStatementUpsertOne) Update(set func(*UnsoldStatementUpsert)) *UnsoldStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UnsoldStatementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UnsoldStatementUpsertOne) SetCreatedAt(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UnsoldStatementUpsertOne) AddCreatedAt(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateCreatedAt() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UnsoldStatementUpsertOne) SetUpdatedAt(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UnsoldStatementUpsertOne) AddUpdatedAt(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateUpdatedAt() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UnsoldStatementUpsertOne) SetDeletedAt(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UnsoldStatementUpsertOne) AddDeletedAt(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateDeletedAt() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *UnsoldStatementUpsertOne) SetGoodID(v uuid.UUID) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateGoodID() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *UnsoldStatementUpsertOne) ClearGoodID() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UnsoldStatementUpsertOne) SetCoinTypeID(v uuid.UUID) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateCoinTypeID() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UnsoldStatementUpsertOne) ClearCoinTypeID() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *UnsoldStatementUpsertOne) SetAmount(v decimal.Decimal) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *UnsoldStatementUpsertOne) AddAmount(v decimal.Decimal) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateAmount() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *UnsoldStatementUpsertOne) ClearAmount() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearAmount()
	})
}

// SetBenefitDate sets the "benefit_date" field.
func (u *UnsoldStatementUpsertOne) SetBenefitDate(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetBenefitDate(v)
	})
}

// AddBenefitDate adds v to the "benefit_date" field.
func (u *UnsoldStatementUpsertOne) AddBenefitDate(v uint32) *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddBenefitDate(v)
	})
}

// UpdateBenefitDate sets the "benefit_date" field to the value that was provided on create.
func (u *UnsoldStatementUpsertOne) UpdateBenefitDate() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateBenefitDate()
	})
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (u *UnsoldStatementUpsertOne) ClearBenefitDate() *UnsoldStatementUpsertOne {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearBenefitDate()
	})
}

// Exec executes the query.
func (u *UnsoldStatementUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UnsoldStatementCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UnsoldStatementUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UnsoldStatementUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UnsoldStatementUpsertOne.ID is not supported by MySQL driver. Use UnsoldStatementUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UnsoldStatementUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UnsoldStatementCreateBulk is the builder for creating many UnsoldStatement entities in bulk.
type UnsoldStatementCreateBulk struct {
	config
	builders []*UnsoldStatementCreate
	conflict []sql.ConflictOption
}

// Save creates the UnsoldStatement entities in the database.
func (uscb *UnsoldStatementCreateBulk) Save(ctx context.Context) ([]*UnsoldStatement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UnsoldStatement, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UnsoldStatementMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UnsoldStatementCreateBulk) SaveX(ctx context.Context) []*UnsoldStatement {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UnsoldStatementCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UnsoldStatementCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UnsoldStatement.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UnsoldStatementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (uscb *UnsoldStatementCreateBulk) OnConflict(opts ...sql.ConflictOption) *UnsoldStatementUpsertBulk {
	uscb.conflict = opts
	return &UnsoldStatementUpsertBulk{
		create: uscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UnsoldStatement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uscb *UnsoldStatementCreateBulk) OnConflictColumns(columns ...string) *UnsoldStatementUpsertBulk {
	uscb.conflict = append(uscb.conflict, sql.ConflictColumns(columns...))
	return &UnsoldStatementUpsertBulk{
		create: uscb,
	}
}

// UnsoldStatementUpsertBulk is the builder for "upsert"-ing
// a bulk of UnsoldStatement nodes.
type UnsoldStatementUpsertBulk struct {
	create *UnsoldStatementCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UnsoldStatement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(unsoldstatement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UnsoldStatementUpsertBulk) UpdateNewValues() *UnsoldStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(unsoldstatement.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UnsoldStatement.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UnsoldStatementUpsertBulk) Ignore() *UnsoldStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UnsoldStatementUpsertBulk) DoNothing() *UnsoldStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UnsoldStatementCreateBulk.OnConflict
// documentation for more info.
func (u *UnsoldStatementUpsertBulk) Update(set func(*UnsoldStatementUpsert)) *UnsoldStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UnsoldStatementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UnsoldStatementUpsertBulk) SetCreatedAt(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UnsoldStatementUpsertBulk) AddCreatedAt(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateCreatedAt() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UnsoldStatementUpsertBulk) SetUpdatedAt(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UnsoldStatementUpsertBulk) AddUpdatedAt(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateUpdatedAt() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UnsoldStatementUpsertBulk) SetDeletedAt(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UnsoldStatementUpsertBulk) AddDeletedAt(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateDeletedAt() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *UnsoldStatementUpsertBulk) SetGoodID(v uuid.UUID) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateGoodID() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *UnsoldStatementUpsertBulk) ClearGoodID() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *UnsoldStatementUpsertBulk) SetCoinTypeID(v uuid.UUID) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateCoinTypeID() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *UnsoldStatementUpsertBulk) ClearCoinTypeID() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *UnsoldStatementUpsertBulk) SetAmount(v decimal.Decimal) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *UnsoldStatementUpsertBulk) AddAmount(v decimal.Decimal) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateAmount() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *UnsoldStatementUpsertBulk) ClearAmount() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearAmount()
	})
}

// SetBenefitDate sets the "benefit_date" field.
func (u *UnsoldStatementUpsertBulk) SetBenefitDate(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.SetBenefitDate(v)
	})
}

// AddBenefitDate adds v to the "benefit_date" field.
func (u *UnsoldStatementUpsertBulk) AddBenefitDate(v uint32) *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.AddBenefitDate(v)
	})
}

// UpdateBenefitDate sets the "benefit_date" field to the value that was provided on create.
func (u *UnsoldStatementUpsertBulk) UpdateBenefitDate() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.UpdateBenefitDate()
	})
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (u *UnsoldStatementUpsertBulk) ClearBenefitDate() *UnsoldStatementUpsertBulk {
	return u.Update(func(s *UnsoldStatementUpsert) {
		s.ClearBenefitDate()
	})
}

// Exec executes the query.
func (u *UnsoldStatementUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UnsoldStatementCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UnsoldStatementCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UnsoldStatementUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
