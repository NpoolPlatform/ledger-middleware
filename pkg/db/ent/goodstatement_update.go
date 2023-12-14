// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodStatementUpdate is the builder for updating GoodStatement entities.
type GoodStatementUpdate struct {
	config
	hooks    []Hook
	mutation *GoodStatementMutation
}

// Where appends a list predicates to the GoodStatementUpdate builder.
func (gsu *GoodStatementUpdate) Where(ps ...predicate.GoodStatement) *GoodStatementUpdate {
	gsu.mutation.Where(ps...)
	return gsu
}

// SetCreatedAt sets the "created_at" field.
func (gsu *GoodStatementUpdate) SetCreatedAt(u uint32) *GoodStatementUpdate {
	gsu.mutation.ResetCreatedAt()
	gsu.mutation.SetCreatedAt(u)
	return gsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableCreatedAt(u *uint32) *GoodStatementUpdate {
	if u != nil {
		gsu.SetCreatedAt(*u)
	}
	return gsu
}

// AddCreatedAt adds u to the "created_at" field.
func (gsu *GoodStatementUpdate) AddCreatedAt(u int32) *GoodStatementUpdate {
	gsu.mutation.AddCreatedAt(u)
	return gsu
}

// SetUpdatedAt sets the "updated_at" field.
func (gsu *GoodStatementUpdate) SetUpdatedAt(u uint32) *GoodStatementUpdate {
	gsu.mutation.ResetUpdatedAt()
	gsu.mutation.SetUpdatedAt(u)
	return gsu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gsu *GoodStatementUpdate) AddUpdatedAt(u int32) *GoodStatementUpdate {
	gsu.mutation.AddUpdatedAt(u)
	return gsu
}

// SetDeletedAt sets the "deleted_at" field.
func (gsu *GoodStatementUpdate) SetDeletedAt(u uint32) *GoodStatementUpdate {
	gsu.mutation.ResetDeletedAt()
	gsu.mutation.SetDeletedAt(u)
	return gsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableDeletedAt(u *uint32) *GoodStatementUpdate {
	if u != nil {
		gsu.SetDeletedAt(*u)
	}
	return gsu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gsu *GoodStatementUpdate) AddDeletedAt(u int32) *GoodStatementUpdate {
	gsu.mutation.AddDeletedAt(u)
	return gsu
}

// SetEntID sets the "ent_id" field.
func (gsu *GoodStatementUpdate) SetEntID(u uuid.UUID) *GoodStatementUpdate {
	gsu.mutation.SetEntID(u)
	return gsu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableEntID(u *uuid.UUID) *GoodStatementUpdate {
	if u != nil {
		gsu.SetEntID(*u)
	}
	return gsu
}

// SetGoodID sets the "good_id" field.
func (gsu *GoodStatementUpdate) SetGoodID(u uuid.UUID) *GoodStatementUpdate {
	gsu.mutation.SetGoodID(u)
	return gsu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableGoodID(u *uuid.UUID) *GoodStatementUpdate {
	if u != nil {
		gsu.SetGoodID(*u)
	}
	return gsu
}

// ClearGoodID clears the value of the "good_id" field.
func (gsu *GoodStatementUpdate) ClearGoodID() *GoodStatementUpdate {
	gsu.mutation.ClearGoodID()
	return gsu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gsu *GoodStatementUpdate) SetCoinTypeID(u uuid.UUID) *GoodStatementUpdate {
	gsu.mutation.SetCoinTypeID(u)
	return gsu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableCoinTypeID(u *uuid.UUID) *GoodStatementUpdate {
	if u != nil {
		gsu.SetCoinTypeID(*u)
	}
	return gsu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gsu *GoodStatementUpdate) ClearCoinTypeID() *GoodStatementUpdate {
	gsu.mutation.ClearCoinTypeID()
	return gsu
}

// SetAmount sets the "amount" field.
func (gsu *GoodStatementUpdate) SetAmount(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.ResetAmount()
	gsu.mutation.SetAmount(d)
	return gsu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableAmount(d *decimal.Decimal) *GoodStatementUpdate {
	if d != nil {
		gsu.SetAmount(*d)
	}
	return gsu
}

// AddAmount adds d to the "amount" field.
func (gsu *GoodStatementUpdate) AddAmount(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.AddAmount(d)
	return gsu
}

// ClearAmount clears the value of the "amount" field.
func (gsu *GoodStatementUpdate) ClearAmount() *GoodStatementUpdate {
	gsu.mutation.ClearAmount()
	return gsu
}

// SetToPlatform sets the "to_platform" field.
func (gsu *GoodStatementUpdate) SetToPlatform(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.ResetToPlatform()
	gsu.mutation.SetToPlatform(d)
	return gsu
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableToPlatform(d *decimal.Decimal) *GoodStatementUpdate {
	if d != nil {
		gsu.SetToPlatform(*d)
	}
	return gsu
}

// AddToPlatform adds d to the "to_platform" field.
func (gsu *GoodStatementUpdate) AddToPlatform(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.AddToPlatform(d)
	return gsu
}

// ClearToPlatform clears the value of the "to_platform" field.
func (gsu *GoodStatementUpdate) ClearToPlatform() *GoodStatementUpdate {
	gsu.mutation.ClearToPlatform()
	return gsu
}

// SetToUser sets the "to_user" field.
func (gsu *GoodStatementUpdate) SetToUser(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.ResetToUser()
	gsu.mutation.SetToUser(d)
	return gsu
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableToUser(d *decimal.Decimal) *GoodStatementUpdate {
	if d != nil {
		gsu.SetToUser(*d)
	}
	return gsu
}

// AddToUser adds d to the "to_user" field.
func (gsu *GoodStatementUpdate) AddToUser(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.AddToUser(d)
	return gsu
}

// ClearToUser clears the value of the "to_user" field.
func (gsu *GoodStatementUpdate) ClearToUser() *GoodStatementUpdate {
	gsu.mutation.ClearToUser()
	return gsu
}

// SetTechniqueServiceFeeAmount sets the "technique_service_fee_amount" field.
func (gsu *GoodStatementUpdate) SetTechniqueServiceFeeAmount(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.ResetTechniqueServiceFeeAmount()
	gsu.mutation.SetTechniqueServiceFeeAmount(d)
	return gsu
}

// SetNillableTechniqueServiceFeeAmount sets the "technique_service_fee_amount" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableTechniqueServiceFeeAmount(d *decimal.Decimal) *GoodStatementUpdate {
	if d != nil {
		gsu.SetTechniqueServiceFeeAmount(*d)
	}
	return gsu
}

// AddTechniqueServiceFeeAmount adds d to the "technique_service_fee_amount" field.
func (gsu *GoodStatementUpdate) AddTechniqueServiceFeeAmount(d decimal.Decimal) *GoodStatementUpdate {
	gsu.mutation.AddTechniqueServiceFeeAmount(d)
	return gsu
}

// ClearTechniqueServiceFeeAmount clears the value of the "technique_service_fee_amount" field.
func (gsu *GoodStatementUpdate) ClearTechniqueServiceFeeAmount() *GoodStatementUpdate {
	gsu.mutation.ClearTechniqueServiceFeeAmount()
	return gsu
}

// SetBenefitDate sets the "benefit_date" field.
func (gsu *GoodStatementUpdate) SetBenefitDate(u uint32) *GoodStatementUpdate {
	gsu.mutation.ResetBenefitDate()
	gsu.mutation.SetBenefitDate(u)
	return gsu
}

// SetNillableBenefitDate sets the "benefit_date" field if the given value is not nil.
func (gsu *GoodStatementUpdate) SetNillableBenefitDate(u *uint32) *GoodStatementUpdate {
	if u != nil {
		gsu.SetBenefitDate(*u)
	}
	return gsu
}

// AddBenefitDate adds u to the "benefit_date" field.
func (gsu *GoodStatementUpdate) AddBenefitDate(u int32) *GoodStatementUpdate {
	gsu.mutation.AddBenefitDate(u)
	return gsu
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (gsu *GoodStatementUpdate) ClearBenefitDate() *GoodStatementUpdate {
	gsu.mutation.ClearBenefitDate()
	return gsu
}

// Mutation returns the GoodStatementMutation object of the builder.
func (gsu *GoodStatementUpdate) Mutation() *GoodStatementMutation {
	return gsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gsu *GoodStatementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := gsu.defaults(); err != nil {
		return 0, err
	}
	if len(gsu.hooks) == 0 {
		affected, err = gsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodStatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gsu.mutation = mutation
			affected, err = gsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gsu.hooks) - 1; i >= 0; i-- {
			if gsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gsu *GoodStatementUpdate) SaveX(ctx context.Context) int {
	affected, err := gsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gsu *GoodStatementUpdate) Exec(ctx context.Context) error {
	_, err := gsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsu *GoodStatementUpdate) ExecX(ctx context.Context) {
	if err := gsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsu *GoodStatementUpdate) defaults() error {
	if _, ok := gsu.mutation.UpdatedAt(); !ok {
		if goodstatement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodstatement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodstatement.UpdateDefaultUpdatedAt()
		gsu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (gsu *GoodStatementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodstatement.Table,
			Columns: goodstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodstatement.FieldID,
			},
		},
	}
	if ps := gsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldCreatedAt,
		})
	}
	if value, ok := gsu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldCreatedAt,
		})
	}
	if value, ok := gsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldUpdatedAt,
		})
	}
	if value, ok := gsu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldUpdatedAt,
		})
	}
	if value, ok := gsu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldDeletedAt,
		})
	}
	if value, ok := gsu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldDeletedAt,
		})
	}
	if value, ok := gsu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodstatement.FieldEntID,
		})
	}
	if value, ok := gsu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodstatement.FieldGoodID,
		})
	}
	if gsu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodstatement.FieldGoodID,
		})
	}
	if value, ok := gsu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodstatement.FieldCoinTypeID,
		})
	}
	if gsu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodstatement.FieldCoinTypeID,
		})
	}
	if value, ok := gsu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldAmount,
		})
	}
	if value, ok := gsu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldAmount,
		})
	}
	if gsu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldAmount,
		})
	}
	if value, ok := gsu.mutation.ToPlatform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToPlatform,
		})
	}
	if value, ok := gsu.mutation.AddedToPlatform(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToPlatform,
		})
	}
	if gsu.mutation.ToPlatformCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldToPlatform,
		})
	}
	if value, ok := gsu.mutation.ToUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToUser,
		})
	}
	if value, ok := gsu.mutation.AddedToUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToUser,
		})
	}
	if gsu.mutation.ToUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldToUser,
		})
	}
	if value, ok := gsu.mutation.TechniqueServiceFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldTechniqueServiceFeeAmount,
		})
	}
	if value, ok := gsu.mutation.AddedTechniqueServiceFeeAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldTechniqueServiceFeeAmount,
		})
	}
	if gsu.mutation.TechniqueServiceFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldTechniqueServiceFeeAmount,
		})
	}
	if value, ok := gsu.mutation.BenefitDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldBenefitDate,
		})
	}
	if value, ok := gsu.mutation.AddedBenefitDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldBenefitDate,
		})
	}
	if gsu.mutation.BenefitDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: goodstatement.FieldBenefitDate,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodstatement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoodStatementUpdateOne is the builder for updating a single GoodStatement entity.
type GoodStatementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodStatementMutation
}

// SetCreatedAt sets the "created_at" field.
func (gsuo *GoodStatementUpdateOne) SetCreatedAt(u uint32) *GoodStatementUpdateOne {
	gsuo.mutation.ResetCreatedAt()
	gsuo.mutation.SetCreatedAt(u)
	return gsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableCreatedAt(u *uint32) *GoodStatementUpdateOne {
	if u != nil {
		gsuo.SetCreatedAt(*u)
	}
	return gsuo
}

// AddCreatedAt adds u to the "created_at" field.
func (gsuo *GoodStatementUpdateOne) AddCreatedAt(u int32) *GoodStatementUpdateOne {
	gsuo.mutation.AddCreatedAt(u)
	return gsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (gsuo *GoodStatementUpdateOne) SetUpdatedAt(u uint32) *GoodStatementUpdateOne {
	gsuo.mutation.ResetUpdatedAt()
	gsuo.mutation.SetUpdatedAt(u)
	return gsuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gsuo *GoodStatementUpdateOne) AddUpdatedAt(u int32) *GoodStatementUpdateOne {
	gsuo.mutation.AddUpdatedAt(u)
	return gsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gsuo *GoodStatementUpdateOne) SetDeletedAt(u uint32) *GoodStatementUpdateOne {
	gsuo.mutation.ResetDeletedAt()
	gsuo.mutation.SetDeletedAt(u)
	return gsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableDeletedAt(u *uint32) *GoodStatementUpdateOne {
	if u != nil {
		gsuo.SetDeletedAt(*u)
	}
	return gsuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gsuo *GoodStatementUpdateOne) AddDeletedAt(u int32) *GoodStatementUpdateOne {
	gsuo.mutation.AddDeletedAt(u)
	return gsuo
}

// SetEntID sets the "ent_id" field.
func (gsuo *GoodStatementUpdateOne) SetEntID(u uuid.UUID) *GoodStatementUpdateOne {
	gsuo.mutation.SetEntID(u)
	return gsuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableEntID(u *uuid.UUID) *GoodStatementUpdateOne {
	if u != nil {
		gsuo.SetEntID(*u)
	}
	return gsuo
}

// SetGoodID sets the "good_id" field.
func (gsuo *GoodStatementUpdateOne) SetGoodID(u uuid.UUID) *GoodStatementUpdateOne {
	gsuo.mutation.SetGoodID(u)
	return gsuo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableGoodID(u *uuid.UUID) *GoodStatementUpdateOne {
	if u != nil {
		gsuo.SetGoodID(*u)
	}
	return gsuo
}

// ClearGoodID clears the value of the "good_id" field.
func (gsuo *GoodStatementUpdateOne) ClearGoodID() *GoodStatementUpdateOne {
	gsuo.mutation.ClearGoodID()
	return gsuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gsuo *GoodStatementUpdateOne) SetCoinTypeID(u uuid.UUID) *GoodStatementUpdateOne {
	gsuo.mutation.SetCoinTypeID(u)
	return gsuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *GoodStatementUpdateOne {
	if u != nil {
		gsuo.SetCoinTypeID(*u)
	}
	return gsuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gsuo *GoodStatementUpdateOne) ClearCoinTypeID() *GoodStatementUpdateOne {
	gsuo.mutation.ClearCoinTypeID()
	return gsuo
}

// SetAmount sets the "amount" field.
func (gsuo *GoodStatementUpdateOne) SetAmount(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.ResetAmount()
	gsuo.mutation.SetAmount(d)
	return gsuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableAmount(d *decimal.Decimal) *GoodStatementUpdateOne {
	if d != nil {
		gsuo.SetAmount(*d)
	}
	return gsuo
}

// AddAmount adds d to the "amount" field.
func (gsuo *GoodStatementUpdateOne) AddAmount(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.AddAmount(d)
	return gsuo
}

// ClearAmount clears the value of the "amount" field.
func (gsuo *GoodStatementUpdateOne) ClearAmount() *GoodStatementUpdateOne {
	gsuo.mutation.ClearAmount()
	return gsuo
}

// SetToPlatform sets the "to_platform" field.
func (gsuo *GoodStatementUpdateOne) SetToPlatform(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.ResetToPlatform()
	gsuo.mutation.SetToPlatform(d)
	return gsuo
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableToPlatform(d *decimal.Decimal) *GoodStatementUpdateOne {
	if d != nil {
		gsuo.SetToPlatform(*d)
	}
	return gsuo
}

// AddToPlatform adds d to the "to_platform" field.
func (gsuo *GoodStatementUpdateOne) AddToPlatform(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.AddToPlatform(d)
	return gsuo
}

// ClearToPlatform clears the value of the "to_platform" field.
func (gsuo *GoodStatementUpdateOne) ClearToPlatform() *GoodStatementUpdateOne {
	gsuo.mutation.ClearToPlatform()
	return gsuo
}

// SetToUser sets the "to_user" field.
func (gsuo *GoodStatementUpdateOne) SetToUser(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.ResetToUser()
	gsuo.mutation.SetToUser(d)
	return gsuo
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableToUser(d *decimal.Decimal) *GoodStatementUpdateOne {
	if d != nil {
		gsuo.SetToUser(*d)
	}
	return gsuo
}

// AddToUser adds d to the "to_user" field.
func (gsuo *GoodStatementUpdateOne) AddToUser(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.AddToUser(d)
	return gsuo
}

// ClearToUser clears the value of the "to_user" field.
func (gsuo *GoodStatementUpdateOne) ClearToUser() *GoodStatementUpdateOne {
	gsuo.mutation.ClearToUser()
	return gsuo
}

// SetTechniqueServiceFeeAmount sets the "technique_service_fee_amount" field.
func (gsuo *GoodStatementUpdateOne) SetTechniqueServiceFeeAmount(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.ResetTechniqueServiceFeeAmount()
	gsuo.mutation.SetTechniqueServiceFeeAmount(d)
	return gsuo
}

// SetNillableTechniqueServiceFeeAmount sets the "technique_service_fee_amount" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableTechniqueServiceFeeAmount(d *decimal.Decimal) *GoodStatementUpdateOne {
	if d != nil {
		gsuo.SetTechniqueServiceFeeAmount(*d)
	}
	return gsuo
}

// AddTechniqueServiceFeeAmount adds d to the "technique_service_fee_amount" field.
func (gsuo *GoodStatementUpdateOne) AddTechniqueServiceFeeAmount(d decimal.Decimal) *GoodStatementUpdateOne {
	gsuo.mutation.AddTechniqueServiceFeeAmount(d)
	return gsuo
}

// ClearTechniqueServiceFeeAmount clears the value of the "technique_service_fee_amount" field.
func (gsuo *GoodStatementUpdateOne) ClearTechniqueServiceFeeAmount() *GoodStatementUpdateOne {
	gsuo.mutation.ClearTechniqueServiceFeeAmount()
	return gsuo
}

// SetBenefitDate sets the "benefit_date" field.
func (gsuo *GoodStatementUpdateOne) SetBenefitDate(u uint32) *GoodStatementUpdateOne {
	gsuo.mutation.ResetBenefitDate()
	gsuo.mutation.SetBenefitDate(u)
	return gsuo
}

// SetNillableBenefitDate sets the "benefit_date" field if the given value is not nil.
func (gsuo *GoodStatementUpdateOne) SetNillableBenefitDate(u *uint32) *GoodStatementUpdateOne {
	if u != nil {
		gsuo.SetBenefitDate(*u)
	}
	return gsuo
}

// AddBenefitDate adds u to the "benefit_date" field.
func (gsuo *GoodStatementUpdateOne) AddBenefitDate(u int32) *GoodStatementUpdateOne {
	gsuo.mutation.AddBenefitDate(u)
	return gsuo
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (gsuo *GoodStatementUpdateOne) ClearBenefitDate() *GoodStatementUpdateOne {
	gsuo.mutation.ClearBenefitDate()
	return gsuo
}

// Mutation returns the GoodStatementMutation object of the builder.
func (gsuo *GoodStatementUpdateOne) Mutation() *GoodStatementMutation {
	return gsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gsuo *GoodStatementUpdateOne) Select(field string, fields ...string) *GoodStatementUpdateOne {
	gsuo.fields = append([]string{field}, fields...)
	return gsuo
}

// Save executes the query and returns the updated GoodStatement entity.
func (gsuo *GoodStatementUpdateOne) Save(ctx context.Context) (*GoodStatement, error) {
	var (
		err  error
		node *GoodStatement
	)
	if err := gsuo.defaults(); err != nil {
		return nil, err
	}
	if len(gsuo.hooks) == 0 {
		node, err = gsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodStatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gsuo.mutation = mutation
			node, err = gsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gsuo.hooks) - 1; i >= 0; i-- {
			if gsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodStatement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodStatementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gsuo *GoodStatementUpdateOne) SaveX(ctx context.Context) *GoodStatement {
	node, err := gsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gsuo *GoodStatementUpdateOne) Exec(ctx context.Context) error {
	_, err := gsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsuo *GoodStatementUpdateOne) ExecX(ctx context.Context) {
	if err := gsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsuo *GoodStatementUpdateOne) defaults() error {
	if _, ok := gsuo.mutation.UpdatedAt(); !ok {
		if goodstatement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodstatement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodstatement.UpdateDefaultUpdatedAt()
		gsuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (gsuo *GoodStatementUpdateOne) sqlSave(ctx context.Context) (_node *GoodStatement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodstatement.Table,
			Columns: goodstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodstatement.FieldID,
			},
		},
	}
	id, ok := gsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodStatement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodstatement.FieldID)
		for _, f := range fields {
			if !goodstatement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodstatement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldCreatedAt,
		})
	}
	if value, ok := gsuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldCreatedAt,
		})
	}
	if value, ok := gsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldUpdatedAt,
		})
	}
	if value, ok := gsuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldUpdatedAt,
		})
	}
	if value, ok := gsuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldDeletedAt,
		})
	}
	if value, ok := gsuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldDeletedAt,
		})
	}
	if value, ok := gsuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodstatement.FieldEntID,
		})
	}
	if value, ok := gsuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodstatement.FieldGoodID,
		})
	}
	if gsuo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodstatement.FieldGoodID,
		})
	}
	if value, ok := gsuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodstatement.FieldCoinTypeID,
		})
	}
	if gsuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodstatement.FieldCoinTypeID,
		})
	}
	if value, ok := gsuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldAmount,
		})
	}
	if value, ok := gsuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldAmount,
		})
	}
	if gsuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldAmount,
		})
	}
	if value, ok := gsuo.mutation.ToPlatform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToPlatform,
		})
	}
	if value, ok := gsuo.mutation.AddedToPlatform(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToPlatform,
		})
	}
	if gsuo.mutation.ToPlatformCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldToPlatform,
		})
	}
	if value, ok := gsuo.mutation.ToUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToUser,
		})
	}
	if value, ok := gsuo.mutation.AddedToUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldToUser,
		})
	}
	if gsuo.mutation.ToUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldToUser,
		})
	}
	if value, ok := gsuo.mutation.TechniqueServiceFeeAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldTechniqueServiceFeeAmount,
		})
	}
	if value, ok := gsuo.mutation.AddedTechniqueServiceFeeAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: goodstatement.FieldTechniqueServiceFeeAmount,
		})
	}
	if gsuo.mutation.TechniqueServiceFeeAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: goodstatement.FieldTechniqueServiceFeeAmount,
		})
	}
	if value, ok := gsuo.mutation.BenefitDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldBenefitDate,
		})
	}
	if value, ok := gsuo.mutation.AddedBenefitDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodstatement.FieldBenefitDate,
		})
	}
	if gsuo.mutation.BenefitDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: goodstatement.FieldBenefitDate,
		})
	}
	_node = &GoodStatement{config: gsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodstatement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
