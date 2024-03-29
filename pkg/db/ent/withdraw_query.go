// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
)

// WithdrawQuery is the builder for querying Withdraw entities.
type WithdrawQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Withdraw
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithdrawQuery builder.
func (wq *WithdrawQuery) Where(ps ...predicate.Withdraw) *WithdrawQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WithdrawQuery) Limit(limit int) *WithdrawQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WithdrawQuery) Offset(offset int) *WithdrawQuery {
	wq.offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WithdrawQuery) Unique(unique bool) *WithdrawQuery {
	wq.unique = &unique
	return wq
}

// Order adds an order step to the query.
func (wq *WithdrawQuery) Order(o ...OrderFunc) *WithdrawQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// First returns the first Withdraw entity from the query.
// Returns a *NotFoundError when no Withdraw was found.
func (wq *WithdrawQuery) First(ctx context.Context) (*Withdraw, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withdraw.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WithdrawQuery) FirstX(ctx context.Context) *Withdraw {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Withdraw ID from the query.
// Returns a *NotFoundError when no Withdraw ID was found.
func (wq *WithdrawQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withdraw.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WithdrawQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Withdraw entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Withdraw entity is found.
// Returns a *NotFoundError when no Withdraw entities are found.
func (wq *WithdrawQuery) Only(ctx context.Context) (*Withdraw, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withdraw.Label}
	default:
		return nil, &NotSingularError{withdraw.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WithdrawQuery) OnlyX(ctx context.Context) *Withdraw {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Withdraw ID in the query.
// Returns a *NotSingularError when more than one Withdraw ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WithdrawQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withdraw.Label}
	default:
		err = &NotSingularError{withdraw.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WithdrawQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Withdraws.
func (wq *WithdrawQuery) All(ctx context.Context) ([]*Withdraw, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WithdrawQuery) AllX(ctx context.Context) []*Withdraw {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Withdraw IDs.
func (wq *WithdrawQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := wq.Select(withdraw.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WithdrawQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WithdrawQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WithdrawQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WithdrawQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WithdrawQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithdrawQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WithdrawQuery) Clone() *WithdrawQuery {
	if wq == nil {
		return nil
	}
	return &WithdrawQuery{
		config:     wq.config,
		limit:      wq.limit,
		offset:     wq.offset,
		order:      append([]OrderFunc{}, wq.order...),
		predicates: append([]predicate.Withdraw{}, wq.predicates...),
		// clone intermediate query.
		sql:    wq.sql.Clone(),
		path:   wq.path,
		unique: wq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Withdraw.Query().
//		GroupBy(withdraw.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *WithdrawQuery) GroupBy(field string, fields ...string) *WithdrawGroupBy {
	grbuild := &WithdrawGroupBy{config: wq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(ctx), nil
	}
	grbuild.label = withdraw.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.Withdraw.Query().
//		Select(withdraw.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (wq *WithdrawQuery) Select(fields ...string) *WithdrawSelect {
	wq.fields = append(wq.fields, fields...)
	selbuild := &WithdrawSelect{WithdrawQuery: wq}
	selbuild.label = withdraw.Label
	selbuild.flds, selbuild.scan = &wq.fields, selbuild.Scan
	return selbuild
}

func (wq *WithdrawQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wq.fields {
		if !withdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	if withdraw.Policy == nil {
		return errors.New("ent: uninitialized withdraw.Policy (forgotten import ent/runtime?)")
	}
	if err := withdraw.Policy.EvalQuery(ctx, wq); err != nil {
		return err
	}
	return nil
}

func (wq *WithdrawQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Withdraw, error) {
	var (
		nodes = []*Withdraw{}
		_spec = wq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Withdraw).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Withdraw{config: wq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wq *WithdrawQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.fields
	if len(wq.fields) > 0 {
		_spec.Unique = wq.unique != nil && *wq.unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WithdrawQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wq *WithdrawQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withdraw.Table,
			Columns: withdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: withdraw.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if unique := wq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withdraw.FieldID)
		for i := range fields {
			if fields[i] != withdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WithdrawQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(withdraw.Table)
	columns := wq.fields
	if len(columns) == 0 {
		columns = withdraw.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.unique != nil && *wq.unique {
		selector.Distinct()
	}
	for _, m := range wq.modifiers {
		m(selector)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (wq *WithdrawQuery) ForUpdate(opts ...sql.LockOption) *WithdrawQuery {
	if wq.driver.Dialect() == dialect.Postgres {
		wq.Unique(false)
	}
	wq.modifiers = append(wq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return wq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (wq *WithdrawQuery) ForShare(opts ...sql.LockOption) *WithdrawQuery {
	if wq.driver.Dialect() == dialect.Postgres {
		wq.Unique(false)
	}
	wq.modifiers = append(wq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return wq
}

// WithdrawGroupBy is the group-by builder for Withdraw entities.
type WithdrawGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WithdrawGroupBy) Aggregate(fns ...AggregateFunc) *WithdrawGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wgb *WithdrawGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

func (wgb *WithdrawGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wgb.fields {
		if !withdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WithdrawGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql.Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
		for _, f := range wgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wgb.fields...)...)
}

// WithdrawSelect is the builder for selecting fields of Withdraw entities.
type WithdrawSelect struct {
	*WithdrawQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WithdrawSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	ws.sql = ws.WithdrawQuery.sqlQuery(ctx)
	return ws.sqlScan(ctx, v)
}

func (ws *WithdrawSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ws.sql.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
