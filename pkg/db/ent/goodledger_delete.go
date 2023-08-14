// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
)

// GoodLedgerDelete is the builder for deleting a GoodLedger entity.
type GoodLedgerDelete struct {
	config
	hooks    []Hook
	mutation *GoodLedgerMutation
}

// Where appends a list predicates to the GoodLedgerDelete builder.
func (gld *GoodLedgerDelete) Where(ps ...predicate.GoodLedger) *GoodLedgerDelete {
	gld.mutation.Where(ps...)
	return gld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gld *GoodLedgerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gld.hooks) == 0 {
		affected, err = gld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodLedgerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gld.mutation = mutation
			affected, err = gld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gld.hooks) - 1; i >= 0; i-- {
			if gld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gld *GoodLedgerDelete) ExecX(ctx context.Context) int {
	n, err := gld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gld *GoodLedgerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodledger.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodledger.FieldID,
			},
		},
	}
	if ps := gld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GoodLedgerDeleteOne is the builder for deleting a single GoodLedger entity.
type GoodLedgerDeleteOne struct {
	gld *GoodLedgerDelete
}

// Exec executes the deletion query.
func (gldo *GoodLedgerDeleteOne) Exec(ctx context.Context) error {
	n, err := gldo.gld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodledger.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gldo *GoodLedgerDeleteOne) ExecX(ctx context.Context) {
	gldo.gld.ExecX(ctx)
}