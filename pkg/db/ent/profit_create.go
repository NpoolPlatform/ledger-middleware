// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/profit"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProfitCreate is the builder for creating a Profit entity.
type ProfitCreate struct {
	config
	mutation *ProfitMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProfitCreate) SetCreatedAt(u uint32) *ProfitCreate {
	pc.mutation.SetCreatedAt(u)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableCreatedAt(u *uint32) *ProfitCreate {
	if u != nil {
		pc.SetCreatedAt(*u)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProfitCreate) SetUpdatedAt(u uint32) *ProfitCreate {
	pc.mutation.SetUpdatedAt(u)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableUpdatedAt(u *uint32) *ProfitCreate {
	if u != nil {
		pc.SetUpdatedAt(*u)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *ProfitCreate) SetDeletedAt(u uint32) *ProfitCreate {
	pc.mutation.SetDeletedAt(u)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableDeletedAt(u *uint32) *ProfitCreate {
	if u != nil {
		pc.SetDeletedAt(*u)
	}
	return pc
}

// SetEntID sets the "ent_id" field.
func (pc *ProfitCreate) SetEntID(u uuid.UUID) *ProfitCreate {
	pc.mutation.SetEntID(u)
	return pc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableEntID(u *uuid.UUID) *ProfitCreate {
	if u != nil {
		pc.SetEntID(*u)
	}
	return pc
}

// SetAppID sets the "app_id" field.
func (pc *ProfitCreate) SetAppID(u uuid.UUID) *ProfitCreate {
	pc.mutation.SetAppID(u)
	return pc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableAppID(u *uuid.UUID) *ProfitCreate {
	if u != nil {
		pc.SetAppID(*u)
	}
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *ProfitCreate) SetUserID(u uuid.UUID) *ProfitCreate {
	pc.mutation.SetUserID(u)
	return pc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableUserID(u *uuid.UUID) *ProfitCreate {
	if u != nil {
		pc.SetUserID(*u)
	}
	return pc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pc *ProfitCreate) SetCoinTypeID(u uuid.UUID) *ProfitCreate {
	pc.mutation.SetCoinTypeID(u)
	return pc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableCoinTypeID(u *uuid.UUID) *ProfitCreate {
	if u != nil {
		pc.SetCoinTypeID(*u)
	}
	return pc
}

// SetIncoming sets the "incoming" field.
func (pc *ProfitCreate) SetIncoming(d decimal.Decimal) *ProfitCreate {
	pc.mutation.SetIncoming(d)
	return pc
}

// SetNillableIncoming sets the "incoming" field if the given value is not nil.
func (pc *ProfitCreate) SetNillableIncoming(d *decimal.Decimal) *ProfitCreate {
	if d != nil {
		pc.SetIncoming(*d)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProfitCreate) SetID(u uint32) *ProfitCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the ProfitMutation object of the builder.
func (pc *ProfitCreate) Mutation() *ProfitMutation {
	return pc.mutation
}

// Save creates the Profit in the database.
func (pc *ProfitCreate) Save(ctx context.Context) (*Profit, error) {
	var (
		err  error
		node *Profit
	)
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Profit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProfitMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfitCreate) SaveX(ctx context.Context) *Profit {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfitCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfitCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfitCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if profit.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := profit.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if profit.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := profit.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.DeletedAt(); !ok {
		if profit.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := profit.DefaultDeletedAt()
		pc.mutation.SetDeletedAt(v)
	}
	if _, ok := pc.mutation.EntID(); !ok {
		if profit.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := profit.DefaultEntID()
		pc.mutation.SetEntID(v)
	}
	if _, ok := pc.mutation.AppID(); !ok {
		if profit.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := profit.DefaultAppID()
		pc.mutation.SetAppID(v)
	}
	if _, ok := pc.mutation.UserID(); !ok {
		if profit.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := profit.DefaultUserID()
		pc.mutation.SetUserID(v)
	}
	if _, ok := pc.mutation.CoinTypeID(); !ok {
		if profit.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized profit.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := profit.DefaultCoinTypeID()
		pc.mutation.SetCoinTypeID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfitCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Profit.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Profit.updated_at"`)}
	}
	if _, ok := pc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Profit.deleted_at"`)}
	}
	if _, ok := pc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Profit.ent_id"`)}
	}
	return nil
}

func (pc *ProfitCreate) sqlSave(ctx context.Context) (*Profit, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (pc *ProfitCreate) createSpec() (*Profit, *sqlgraph.CreateSpec) {
	var (
		_node = &Profit{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: profit.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profit.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profit.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profit.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: profit.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := pc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: profit.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := pc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: profit.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := pc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: profit.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := pc.mutation.Incoming(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: profit.FieldIncoming,
		})
		_node.Incoming = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Profit.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfitUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pc *ProfitCreate) OnConflict(opts ...sql.ConflictOption) *ProfitUpsertOne {
	pc.conflict = opts
	return &ProfitUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Profit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pc *ProfitCreate) OnConflictColumns(columns ...string) *ProfitUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProfitUpsertOne{
		create: pc,
	}
}

type (
	// ProfitUpsertOne is the builder for "upsert"-ing
	//  one Profit node.
	ProfitUpsertOne struct {
		create *ProfitCreate
	}

	// ProfitUpsert is the "OnConflict" setter.
	ProfitUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ProfitUpsert) SetCreatedAt(v uint32) *ProfitUpsert {
	u.Set(profit.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateCreatedAt() *ProfitUpsert {
	u.SetExcluded(profit.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ProfitUpsert) AddCreatedAt(v uint32) *ProfitUpsert {
	u.Add(profit.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfitUpsert) SetUpdatedAt(v uint32) *ProfitUpsert {
	u.Set(profit.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateUpdatedAt() *ProfitUpsert {
	u.SetExcluded(profit.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ProfitUpsert) AddUpdatedAt(v uint32) *ProfitUpsert {
	u.Add(profit.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ProfitUpsert) SetDeletedAt(v uint32) *ProfitUpsert {
	u.Set(profit.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateDeletedAt() *ProfitUpsert {
	u.SetExcluded(profit.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ProfitUpsert) AddDeletedAt(v uint32) *ProfitUpsert {
	u.Add(profit.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *ProfitUpsert) SetEntID(v uuid.UUID) *ProfitUpsert {
	u.Set(profit.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateEntID() *ProfitUpsert {
	u.SetExcluded(profit.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *ProfitUpsert) SetAppID(v uuid.UUID) *ProfitUpsert {
	u.Set(profit.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateAppID() *ProfitUpsert {
	u.SetExcluded(profit.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *ProfitUpsert) ClearAppID() *ProfitUpsert {
	u.SetNull(profit.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ProfitUpsert) SetUserID(v uuid.UUID) *ProfitUpsert {
	u.Set(profit.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateUserID() *ProfitUpsert {
	u.SetExcluded(profit.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *ProfitUpsert) ClearUserID() *ProfitUpsert {
	u.SetNull(profit.FieldUserID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *ProfitUpsert) SetCoinTypeID(v uuid.UUID) *ProfitUpsert {
	u.Set(profit.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateCoinTypeID() *ProfitUpsert {
	u.SetExcluded(profit.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *ProfitUpsert) ClearCoinTypeID() *ProfitUpsert {
	u.SetNull(profit.FieldCoinTypeID)
	return u
}

// SetIncoming sets the "incoming" field.
func (u *ProfitUpsert) SetIncoming(v decimal.Decimal) *ProfitUpsert {
	u.Set(profit.FieldIncoming, v)
	return u
}

// UpdateIncoming sets the "incoming" field to the value that was provided on create.
func (u *ProfitUpsert) UpdateIncoming() *ProfitUpsert {
	u.SetExcluded(profit.FieldIncoming)
	return u
}

// AddIncoming adds v to the "incoming" field.
func (u *ProfitUpsert) AddIncoming(v decimal.Decimal) *ProfitUpsert {
	u.Add(profit.FieldIncoming, v)
	return u
}

// ClearIncoming clears the value of the "incoming" field.
func (u *ProfitUpsert) ClearIncoming() *ProfitUpsert {
	u.SetNull(profit.FieldIncoming)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Profit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(profit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ProfitUpsertOne) UpdateNewValues() *ProfitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(profit.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Profit.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ProfitUpsertOne) Ignore() *ProfitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfitUpsertOne) DoNothing() *ProfitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfitCreate.OnConflict
// documentation for more info.
func (u *ProfitUpsertOne) Update(set func(*ProfitUpsert)) *ProfitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfitUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfitUpsertOne) SetCreatedAt(v uint32) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ProfitUpsertOne) AddCreatedAt(v uint32) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateCreatedAt() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfitUpsertOne) SetUpdatedAt(v uint32) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ProfitUpsertOne) AddUpdatedAt(v uint32) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateUpdatedAt() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ProfitUpsertOne) SetDeletedAt(v uint32) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ProfitUpsertOne) AddDeletedAt(v uint32) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateDeletedAt() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *ProfitUpsertOne) SetEntID(v uuid.UUID) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateEntID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *ProfitUpsertOne) SetAppID(v uuid.UUID) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateAppID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *ProfitUpsertOne) ClearAppID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ProfitUpsertOne) SetUserID(v uuid.UUID) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateUserID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *ProfitUpsertOne) ClearUserID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *ProfitUpsertOne) SetCoinTypeID(v uuid.UUID) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateCoinTypeID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *ProfitUpsertOne) ClearCoinTypeID() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetIncoming sets the "incoming" field.
func (u *ProfitUpsertOne) SetIncoming(v decimal.Decimal) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.SetIncoming(v)
	})
}

// AddIncoming adds v to the "incoming" field.
func (u *ProfitUpsertOne) AddIncoming(v decimal.Decimal) *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.AddIncoming(v)
	})
}

// UpdateIncoming sets the "incoming" field to the value that was provided on create.
func (u *ProfitUpsertOne) UpdateIncoming() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateIncoming()
	})
}

// ClearIncoming clears the value of the "incoming" field.
func (u *ProfitUpsertOne) ClearIncoming() *ProfitUpsertOne {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearIncoming()
	})
}

// Exec executes the query.
func (u *ProfitUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfitCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfitUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProfitUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProfitUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProfitCreateBulk is the builder for creating many Profit entities in bulk.
type ProfitCreateBulk struct {
	config
	builders []*ProfitCreate
	conflict []sql.ConflictOption
}

// Save creates the Profit entities in the database.
func (pcb *ProfitCreateBulk) Save(ctx context.Context) ([]*Profit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profit, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfitMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfitCreateBulk) SaveX(ctx context.Context) []*Profit {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfitCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfitCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Profit.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfitUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pcb *ProfitCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProfitUpsertBulk {
	pcb.conflict = opts
	return &ProfitUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Profit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pcb *ProfitCreateBulk) OnConflictColumns(columns ...string) *ProfitUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProfitUpsertBulk{
		create: pcb,
	}
}

// ProfitUpsertBulk is the builder for "upsert"-ing
// a bulk of Profit nodes.
type ProfitUpsertBulk struct {
	create *ProfitCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Profit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(profit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ProfitUpsertBulk) UpdateNewValues() *ProfitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(profit.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Profit.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ProfitUpsertBulk) Ignore() *ProfitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfitUpsertBulk) DoNothing() *ProfitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfitCreateBulk.OnConflict
// documentation for more info.
func (u *ProfitUpsertBulk) Update(set func(*ProfitUpsert)) *ProfitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfitUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfitUpsertBulk) SetCreatedAt(v uint32) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ProfitUpsertBulk) AddCreatedAt(v uint32) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateCreatedAt() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfitUpsertBulk) SetUpdatedAt(v uint32) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ProfitUpsertBulk) AddUpdatedAt(v uint32) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateUpdatedAt() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ProfitUpsertBulk) SetDeletedAt(v uint32) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ProfitUpsertBulk) AddDeletedAt(v uint32) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateDeletedAt() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *ProfitUpsertBulk) SetEntID(v uuid.UUID) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateEntID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *ProfitUpsertBulk) SetAppID(v uuid.UUID) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateAppID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *ProfitUpsertBulk) ClearAppID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ProfitUpsertBulk) SetUserID(v uuid.UUID) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateUserID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *ProfitUpsertBulk) ClearUserID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearUserID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *ProfitUpsertBulk) SetCoinTypeID(v uuid.UUID) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateCoinTypeID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *ProfitUpsertBulk) ClearCoinTypeID() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetIncoming sets the "incoming" field.
func (u *ProfitUpsertBulk) SetIncoming(v decimal.Decimal) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.SetIncoming(v)
	})
}

// AddIncoming adds v to the "incoming" field.
func (u *ProfitUpsertBulk) AddIncoming(v decimal.Decimal) *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.AddIncoming(v)
	})
}

// UpdateIncoming sets the "incoming" field to the value that was provided on create.
func (u *ProfitUpsertBulk) UpdateIncoming() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.UpdateIncoming()
	})
}

// ClearIncoming clears the value of the "incoming" field.
func (u *ProfitUpsertBulk) ClearIncoming() *ProfitUpsertBulk {
	return u.Update(func(s *ProfitUpsert) {
		s.ClearIncoming()
	})
}

// Exec executes the query.
func (u *ProfitUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProfitCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfitCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfitUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
