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
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/mininggeneral"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// MiningGeneralCreate is the builder for creating a MiningGeneral entity.
type MiningGeneralCreate struct {
	config
	mutation *MiningGeneralMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mgc *MiningGeneralCreate) SetCreatedAt(u uint32) *MiningGeneralCreate {
	mgc.mutation.SetCreatedAt(u)
	return mgc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableCreatedAt(u *uint32) *MiningGeneralCreate {
	if u != nil {
		mgc.SetCreatedAt(*u)
	}
	return mgc
}

// SetUpdatedAt sets the "updated_at" field.
func (mgc *MiningGeneralCreate) SetUpdatedAt(u uint32) *MiningGeneralCreate {
	mgc.mutation.SetUpdatedAt(u)
	return mgc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableUpdatedAt(u *uint32) *MiningGeneralCreate {
	if u != nil {
		mgc.SetUpdatedAt(*u)
	}
	return mgc
}

// SetDeletedAt sets the "deleted_at" field.
func (mgc *MiningGeneralCreate) SetDeletedAt(u uint32) *MiningGeneralCreate {
	mgc.mutation.SetDeletedAt(u)
	return mgc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableDeletedAt(u *uint32) *MiningGeneralCreate {
	if u != nil {
		mgc.SetDeletedAt(*u)
	}
	return mgc
}

// SetGoodID sets the "good_id" field.
func (mgc *MiningGeneralCreate) SetGoodID(u uuid.UUID) *MiningGeneralCreate {
	mgc.mutation.SetGoodID(u)
	return mgc
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableGoodID(u *uuid.UUID) *MiningGeneralCreate {
	if u != nil {
		mgc.SetGoodID(*u)
	}
	return mgc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (mgc *MiningGeneralCreate) SetCoinTypeID(u uuid.UUID) *MiningGeneralCreate {
	mgc.mutation.SetCoinTypeID(u)
	return mgc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableCoinTypeID(u *uuid.UUID) *MiningGeneralCreate {
	if u != nil {
		mgc.SetCoinTypeID(*u)
	}
	return mgc
}

// SetAmount sets the "amount" field.
func (mgc *MiningGeneralCreate) SetAmount(d decimal.Decimal) *MiningGeneralCreate {
	mgc.mutation.SetAmount(d)
	return mgc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableAmount(d *decimal.Decimal) *MiningGeneralCreate {
	if d != nil {
		mgc.SetAmount(*d)
	}
	return mgc
}

// SetToPlatform sets the "to_platform" field.
func (mgc *MiningGeneralCreate) SetToPlatform(d decimal.Decimal) *MiningGeneralCreate {
	mgc.mutation.SetToPlatform(d)
	return mgc
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableToPlatform(d *decimal.Decimal) *MiningGeneralCreate {
	if d != nil {
		mgc.SetToPlatform(*d)
	}
	return mgc
}

// SetToUser sets the "to_user" field.
func (mgc *MiningGeneralCreate) SetToUser(d decimal.Decimal) *MiningGeneralCreate {
	mgc.mutation.SetToUser(d)
	return mgc
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableToUser(d *decimal.Decimal) *MiningGeneralCreate {
	if d != nil {
		mgc.SetToUser(*d)
	}
	return mgc
}

// SetID sets the "id" field.
func (mgc *MiningGeneralCreate) SetID(u uuid.UUID) *MiningGeneralCreate {
	mgc.mutation.SetID(u)
	return mgc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mgc *MiningGeneralCreate) SetNillableID(u *uuid.UUID) *MiningGeneralCreate {
	if u != nil {
		mgc.SetID(*u)
	}
	return mgc
}

// Mutation returns the MiningGeneralMutation object of the builder.
func (mgc *MiningGeneralCreate) Mutation() *MiningGeneralMutation {
	return mgc.mutation
}

// Save creates the MiningGeneral in the database.
func (mgc *MiningGeneralCreate) Save(ctx context.Context) (*MiningGeneral, error) {
	var (
		err  error
		node *MiningGeneral
	)
	if err := mgc.defaults(); err != nil {
		return nil, err
	}
	if len(mgc.hooks) == 0 {
		if err = mgc.check(); err != nil {
			return nil, err
		}
		node, err = mgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiningGeneralMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mgc.check(); err != nil {
				return nil, err
			}
			mgc.mutation = mutation
			if node, err = mgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mgc.hooks) - 1; i >= 0; i-- {
			if mgc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mgc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mgc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MiningGeneral)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MiningGeneralMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mgc *MiningGeneralCreate) SaveX(ctx context.Context) *MiningGeneral {
	v, err := mgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mgc *MiningGeneralCreate) Exec(ctx context.Context) error {
	_, err := mgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgc *MiningGeneralCreate) ExecX(ctx context.Context) {
	if err := mgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mgc *MiningGeneralCreate) defaults() error {
	if _, ok := mgc.mutation.CreatedAt(); !ok {
		if mininggeneral.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := mininggeneral.DefaultCreatedAt()
		mgc.mutation.SetCreatedAt(v)
	}
	if _, ok := mgc.mutation.UpdatedAt(); !ok {
		if mininggeneral.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := mininggeneral.DefaultUpdatedAt()
		mgc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mgc.mutation.DeletedAt(); !ok {
		if mininggeneral.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := mininggeneral.DefaultDeletedAt()
		mgc.mutation.SetDeletedAt(v)
	}
	if _, ok := mgc.mutation.GoodID(); !ok {
		if mininggeneral.DefaultGoodID == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.DefaultGoodID (forgotten import ent/runtime?)")
		}
		v := mininggeneral.DefaultGoodID()
		mgc.mutation.SetGoodID(v)
	}
	if _, ok := mgc.mutation.CoinTypeID(); !ok {
		if mininggeneral.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := mininggeneral.DefaultCoinTypeID()
		mgc.mutation.SetCoinTypeID(v)
	}
	if _, ok := mgc.mutation.ID(); !ok {
		if mininggeneral.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized mininggeneral.DefaultID (forgotten import ent/runtime?)")
		}
		v := mininggeneral.DefaultID()
		mgc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mgc *MiningGeneralCreate) check() error {
	if _, ok := mgc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MiningGeneral.created_at"`)}
	}
	if _, ok := mgc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MiningGeneral.updated_at"`)}
	}
	if _, ok := mgc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "MiningGeneral.deleted_at"`)}
	}
	return nil
}

func (mgc *MiningGeneralCreate) sqlSave(ctx context.Context) (*MiningGeneral, error) {
	_node, _spec := mgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mgc.driver, _spec); err != nil {
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

func (mgc *MiningGeneralCreate) createSpec() (*MiningGeneral, *sqlgraph.CreateSpec) {
	var (
		_node = &MiningGeneral{config: mgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mininggeneral.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mininggeneral.FieldID,
			},
		}
	)
	_spec.OnConflict = mgc.conflict
	if id, ok := mgc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mgc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mgc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mgc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mininggeneral.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := mgc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mininggeneral.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := mgc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: mininggeneral.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := mgc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := mgc.mutation.ToPlatform(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToPlatform,
		})
		_node.ToPlatform = value
	}
	if value, ok := mgc.mutation.ToUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mininggeneral.FieldToUser,
		})
		_node.ToUser = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MiningGeneral.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MiningGeneralUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (mgc *MiningGeneralCreate) OnConflict(opts ...sql.ConflictOption) *MiningGeneralUpsertOne {
	mgc.conflict = opts
	return &MiningGeneralUpsertOne{
		create: mgc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MiningGeneral.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mgc *MiningGeneralCreate) OnConflictColumns(columns ...string) *MiningGeneralUpsertOne {
	mgc.conflict = append(mgc.conflict, sql.ConflictColumns(columns...))
	return &MiningGeneralUpsertOne{
		create: mgc,
	}
}

type (
	// MiningGeneralUpsertOne is the builder for "upsert"-ing
	//  one MiningGeneral node.
	MiningGeneralUpsertOne struct {
		create *MiningGeneralCreate
	}

	// MiningGeneralUpsert is the "OnConflict" setter.
	MiningGeneralUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MiningGeneralUpsert) SetCreatedAt(v uint32) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateCreatedAt() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *MiningGeneralUpsert) AddCreatedAt(v uint32) *MiningGeneralUpsert {
	u.Add(mininggeneral.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MiningGeneralUpsert) SetUpdatedAt(v uint32) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateUpdatedAt() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *MiningGeneralUpsert) AddUpdatedAt(v uint32) *MiningGeneralUpsert {
	u.Add(mininggeneral.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MiningGeneralUpsert) SetDeletedAt(v uint32) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateDeletedAt() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MiningGeneralUpsert) AddDeletedAt(v uint32) *MiningGeneralUpsert {
	u.Add(mininggeneral.FieldDeletedAt, v)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *MiningGeneralUpsert) SetGoodID(v uuid.UUID) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateGoodID() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *MiningGeneralUpsert) ClearGoodID() *MiningGeneralUpsert {
	u.SetNull(mininggeneral.FieldGoodID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *MiningGeneralUpsert) SetCoinTypeID(v uuid.UUID) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateCoinTypeID() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *MiningGeneralUpsert) ClearCoinTypeID() *MiningGeneralUpsert {
	u.SetNull(mininggeneral.FieldCoinTypeID)
	return u
}

// SetAmount sets the "amount" field.
func (u *MiningGeneralUpsert) SetAmount(v decimal.Decimal) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateAmount() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *MiningGeneralUpsert) AddAmount(v decimal.Decimal) *MiningGeneralUpsert {
	u.Add(mininggeneral.FieldAmount, v)
	return u
}

// ClearAmount clears the value of the "amount" field.
func (u *MiningGeneralUpsert) ClearAmount() *MiningGeneralUpsert {
	u.SetNull(mininggeneral.FieldAmount)
	return u
}

// SetToPlatform sets the "to_platform" field.
func (u *MiningGeneralUpsert) SetToPlatform(v decimal.Decimal) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldToPlatform, v)
	return u
}

// UpdateToPlatform sets the "to_platform" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateToPlatform() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldToPlatform)
	return u
}

// AddToPlatform adds v to the "to_platform" field.
func (u *MiningGeneralUpsert) AddToPlatform(v decimal.Decimal) *MiningGeneralUpsert {
	u.Add(mininggeneral.FieldToPlatform, v)
	return u
}

// ClearToPlatform clears the value of the "to_platform" field.
func (u *MiningGeneralUpsert) ClearToPlatform() *MiningGeneralUpsert {
	u.SetNull(mininggeneral.FieldToPlatform)
	return u
}

// SetToUser sets the "to_user" field.
func (u *MiningGeneralUpsert) SetToUser(v decimal.Decimal) *MiningGeneralUpsert {
	u.Set(mininggeneral.FieldToUser, v)
	return u
}

// UpdateToUser sets the "to_user" field to the value that was provided on create.
func (u *MiningGeneralUpsert) UpdateToUser() *MiningGeneralUpsert {
	u.SetExcluded(mininggeneral.FieldToUser)
	return u
}

// AddToUser adds v to the "to_user" field.
func (u *MiningGeneralUpsert) AddToUser(v decimal.Decimal) *MiningGeneralUpsert {
	u.Add(mininggeneral.FieldToUser, v)
	return u
}

// ClearToUser clears the value of the "to_user" field.
func (u *MiningGeneralUpsert) ClearToUser() *MiningGeneralUpsert {
	u.SetNull(mininggeneral.FieldToUser)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MiningGeneral.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mininggeneral.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MiningGeneralUpsertOne) UpdateNewValues() *MiningGeneralUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mininggeneral.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.MiningGeneral.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MiningGeneralUpsertOne) Ignore() *MiningGeneralUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MiningGeneralUpsertOne) DoNothing() *MiningGeneralUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MiningGeneralCreate.OnConflict
// documentation for more info.
func (u *MiningGeneralUpsertOne) Update(set func(*MiningGeneralUpsert)) *MiningGeneralUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MiningGeneralUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MiningGeneralUpsertOne) SetCreatedAt(v uint32) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *MiningGeneralUpsertOne) AddCreatedAt(v uint32) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateCreatedAt() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MiningGeneralUpsertOne) SetUpdatedAt(v uint32) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *MiningGeneralUpsertOne) AddUpdatedAt(v uint32) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateUpdatedAt() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MiningGeneralUpsertOne) SetDeletedAt(v uint32) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MiningGeneralUpsertOne) AddDeletedAt(v uint32) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateDeletedAt() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *MiningGeneralUpsertOne) SetGoodID(v uuid.UUID) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateGoodID() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *MiningGeneralUpsertOne) ClearGoodID() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *MiningGeneralUpsertOne) SetCoinTypeID(v uuid.UUID) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateCoinTypeID() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *MiningGeneralUpsertOne) ClearCoinTypeID() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *MiningGeneralUpsertOne) SetAmount(v decimal.Decimal) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *MiningGeneralUpsertOne) AddAmount(v decimal.Decimal) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateAmount() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *MiningGeneralUpsertOne) ClearAmount() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearAmount()
	})
}

// SetToPlatform sets the "to_platform" field.
func (u *MiningGeneralUpsertOne) SetToPlatform(v decimal.Decimal) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetToPlatform(v)
	})
}

// AddToPlatform adds v to the "to_platform" field.
func (u *MiningGeneralUpsertOne) AddToPlatform(v decimal.Decimal) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddToPlatform(v)
	})
}

// UpdateToPlatform sets the "to_platform" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateToPlatform() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateToPlatform()
	})
}

// ClearToPlatform clears the value of the "to_platform" field.
func (u *MiningGeneralUpsertOne) ClearToPlatform() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearToPlatform()
	})
}

// SetToUser sets the "to_user" field.
func (u *MiningGeneralUpsertOne) SetToUser(v decimal.Decimal) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetToUser(v)
	})
}

// AddToUser adds v to the "to_user" field.
func (u *MiningGeneralUpsertOne) AddToUser(v decimal.Decimal) *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddToUser(v)
	})
}

// UpdateToUser sets the "to_user" field to the value that was provided on create.
func (u *MiningGeneralUpsertOne) UpdateToUser() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateToUser()
	})
}

// ClearToUser clears the value of the "to_user" field.
func (u *MiningGeneralUpsertOne) ClearToUser() *MiningGeneralUpsertOne {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearToUser()
	})
}

// Exec executes the query.
func (u *MiningGeneralUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MiningGeneralCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MiningGeneralUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MiningGeneralUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MiningGeneralUpsertOne.ID is not supported by MySQL driver. Use MiningGeneralUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MiningGeneralUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MiningGeneralCreateBulk is the builder for creating many MiningGeneral entities in bulk.
type MiningGeneralCreateBulk struct {
	config
	builders []*MiningGeneralCreate
	conflict []sql.ConflictOption
}

// Save creates the MiningGeneral entities in the database.
func (mgcb *MiningGeneralCreateBulk) Save(ctx context.Context) ([]*MiningGeneral, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mgcb.builders))
	nodes := make([]*MiningGeneral, len(mgcb.builders))
	mutators := make([]Mutator, len(mgcb.builders))
	for i := range mgcb.builders {
		func(i int, root context.Context) {
			builder := mgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MiningGeneralMutation)
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
					_, err = mutators[i+1].Mutate(root, mgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mgcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mgcb *MiningGeneralCreateBulk) SaveX(ctx context.Context) []*MiningGeneral {
	v, err := mgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mgcb *MiningGeneralCreateBulk) Exec(ctx context.Context) error {
	_, err := mgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgcb *MiningGeneralCreateBulk) ExecX(ctx context.Context) {
	if err := mgcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MiningGeneral.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MiningGeneralUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (mgcb *MiningGeneralCreateBulk) OnConflict(opts ...sql.ConflictOption) *MiningGeneralUpsertBulk {
	mgcb.conflict = opts
	return &MiningGeneralUpsertBulk{
		create: mgcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MiningGeneral.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mgcb *MiningGeneralCreateBulk) OnConflictColumns(columns ...string) *MiningGeneralUpsertBulk {
	mgcb.conflict = append(mgcb.conflict, sql.ConflictColumns(columns...))
	return &MiningGeneralUpsertBulk{
		create: mgcb,
	}
}

// MiningGeneralUpsertBulk is the builder for "upsert"-ing
// a bulk of MiningGeneral nodes.
type MiningGeneralUpsertBulk struct {
	create *MiningGeneralCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MiningGeneral.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mininggeneral.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MiningGeneralUpsertBulk) UpdateNewValues() *MiningGeneralUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mininggeneral.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MiningGeneral.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MiningGeneralUpsertBulk) Ignore() *MiningGeneralUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MiningGeneralUpsertBulk) DoNothing() *MiningGeneralUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MiningGeneralCreateBulk.OnConflict
// documentation for more info.
func (u *MiningGeneralUpsertBulk) Update(set func(*MiningGeneralUpsert)) *MiningGeneralUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MiningGeneralUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MiningGeneralUpsertBulk) SetCreatedAt(v uint32) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *MiningGeneralUpsertBulk) AddCreatedAt(v uint32) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateCreatedAt() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MiningGeneralUpsertBulk) SetUpdatedAt(v uint32) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *MiningGeneralUpsertBulk) AddUpdatedAt(v uint32) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateUpdatedAt() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MiningGeneralUpsertBulk) SetDeletedAt(v uint32) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MiningGeneralUpsertBulk) AddDeletedAt(v uint32) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateDeletedAt() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *MiningGeneralUpsertBulk) SetGoodID(v uuid.UUID) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateGoodID() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *MiningGeneralUpsertBulk) ClearGoodID() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *MiningGeneralUpsertBulk) SetCoinTypeID(v uuid.UUID) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateCoinTypeID() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *MiningGeneralUpsertBulk) ClearCoinTypeID() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *MiningGeneralUpsertBulk) SetAmount(v decimal.Decimal) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *MiningGeneralUpsertBulk) AddAmount(v decimal.Decimal) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateAmount() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *MiningGeneralUpsertBulk) ClearAmount() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearAmount()
	})
}

// SetToPlatform sets the "to_platform" field.
func (u *MiningGeneralUpsertBulk) SetToPlatform(v decimal.Decimal) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetToPlatform(v)
	})
}

// AddToPlatform adds v to the "to_platform" field.
func (u *MiningGeneralUpsertBulk) AddToPlatform(v decimal.Decimal) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddToPlatform(v)
	})
}

// UpdateToPlatform sets the "to_platform" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateToPlatform() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateToPlatform()
	})
}

// ClearToPlatform clears the value of the "to_platform" field.
func (u *MiningGeneralUpsertBulk) ClearToPlatform() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearToPlatform()
	})
}

// SetToUser sets the "to_user" field.
func (u *MiningGeneralUpsertBulk) SetToUser(v decimal.Decimal) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.SetToUser(v)
	})
}

// AddToUser adds v to the "to_user" field.
func (u *MiningGeneralUpsertBulk) AddToUser(v decimal.Decimal) *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.AddToUser(v)
	})
}

// UpdateToUser sets the "to_user" field to the value that was provided on create.
func (u *MiningGeneralUpsertBulk) UpdateToUser() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.UpdateToUser()
	})
}

// ClearToUser clears the value of the "to_user" field.
func (u *MiningGeneralUpsertBulk) ClearToUser() *MiningGeneralUpsertBulk {
	return u.Update(func(s *MiningGeneralUpsert) {
		s.ClearToUser()
	})
}

// Exec executes the query.
func (u *MiningGeneralUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MiningGeneralCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MiningGeneralCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MiningGeneralUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
