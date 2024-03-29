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
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/simulateledger"
)

// SimulateLedgerQuery is the builder for querying SimulateLedger entities.
type SimulateLedgerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SimulateLedger
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SimulateLedgerQuery builder.
func (slq *SimulateLedgerQuery) Where(ps ...predicate.SimulateLedger) *SimulateLedgerQuery {
	slq.predicates = append(slq.predicates, ps...)
	return slq
}

// Limit adds a limit step to the query.
func (slq *SimulateLedgerQuery) Limit(limit int) *SimulateLedgerQuery {
	slq.limit = &limit
	return slq
}

// Offset adds an offset step to the query.
func (slq *SimulateLedgerQuery) Offset(offset int) *SimulateLedgerQuery {
	slq.offset = &offset
	return slq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (slq *SimulateLedgerQuery) Unique(unique bool) *SimulateLedgerQuery {
	slq.unique = &unique
	return slq
}

// Order adds an order step to the query.
func (slq *SimulateLedgerQuery) Order(o ...OrderFunc) *SimulateLedgerQuery {
	slq.order = append(slq.order, o...)
	return slq
}

// First returns the first SimulateLedger entity from the query.
// Returns a *NotFoundError when no SimulateLedger was found.
func (slq *SimulateLedgerQuery) First(ctx context.Context) (*SimulateLedger, error) {
	nodes, err := slq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{simulateledger.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (slq *SimulateLedgerQuery) FirstX(ctx context.Context) *SimulateLedger {
	node, err := slq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SimulateLedger ID from the query.
// Returns a *NotFoundError when no SimulateLedger ID was found.
func (slq *SimulateLedgerQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = slq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{simulateledger.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (slq *SimulateLedgerQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := slq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SimulateLedger entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SimulateLedger entity is found.
// Returns a *NotFoundError when no SimulateLedger entities are found.
func (slq *SimulateLedgerQuery) Only(ctx context.Context) (*SimulateLedger, error) {
	nodes, err := slq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{simulateledger.Label}
	default:
		return nil, &NotSingularError{simulateledger.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (slq *SimulateLedgerQuery) OnlyX(ctx context.Context) *SimulateLedger {
	node, err := slq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SimulateLedger ID in the query.
// Returns a *NotSingularError when more than one SimulateLedger ID is found.
// Returns a *NotFoundError when no entities are found.
func (slq *SimulateLedgerQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = slq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{simulateledger.Label}
	default:
		err = &NotSingularError{simulateledger.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (slq *SimulateLedgerQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := slq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SimulateLedgers.
func (slq *SimulateLedgerQuery) All(ctx context.Context) ([]*SimulateLedger, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return slq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (slq *SimulateLedgerQuery) AllX(ctx context.Context) []*SimulateLedger {
	nodes, err := slq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SimulateLedger IDs.
func (slq *SimulateLedgerQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := slq.Select(simulateledger.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (slq *SimulateLedgerQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := slq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (slq *SimulateLedgerQuery) Count(ctx context.Context) (int, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return slq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (slq *SimulateLedgerQuery) CountX(ctx context.Context) int {
	count, err := slq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (slq *SimulateLedgerQuery) Exist(ctx context.Context) (bool, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return slq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (slq *SimulateLedgerQuery) ExistX(ctx context.Context) bool {
	exist, err := slq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SimulateLedgerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (slq *SimulateLedgerQuery) Clone() *SimulateLedgerQuery {
	if slq == nil {
		return nil
	}
	return &SimulateLedgerQuery{
		config:     slq.config,
		limit:      slq.limit,
		offset:     slq.offset,
		order:      append([]OrderFunc{}, slq.order...),
		predicates: append([]predicate.SimulateLedger{}, slq.predicates...),
		// clone intermediate query.
		sql:    slq.sql.Clone(),
		path:   slq.path,
		unique: slq.unique,
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
//	client.SimulateLedger.Query().
//		GroupBy(simulateledger.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (slq *SimulateLedgerQuery) GroupBy(field string, fields ...string) *SimulateLedgerGroupBy {
	grbuild := &SimulateLedgerGroupBy{config: slq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := slq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return slq.sqlQuery(ctx), nil
	}
	grbuild.label = simulateledger.Label
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
//	client.SimulateLedger.Query().
//		Select(simulateledger.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (slq *SimulateLedgerQuery) Select(fields ...string) *SimulateLedgerSelect {
	slq.fields = append(slq.fields, fields...)
	selbuild := &SimulateLedgerSelect{SimulateLedgerQuery: slq}
	selbuild.label = simulateledger.Label
	selbuild.flds, selbuild.scan = &slq.fields, selbuild.Scan
	return selbuild
}

func (slq *SimulateLedgerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range slq.fields {
		if !simulateledger.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if slq.path != nil {
		prev, err := slq.path(ctx)
		if err != nil {
			return err
		}
		slq.sql = prev
	}
	if simulateledger.Policy == nil {
		return errors.New("ent: uninitialized simulateledger.Policy (forgotten import ent/runtime?)")
	}
	if err := simulateledger.Policy.EvalQuery(ctx, slq); err != nil {
		return err
	}
	return nil
}

func (slq *SimulateLedgerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SimulateLedger, error) {
	var (
		nodes = []*SimulateLedger{}
		_spec = slq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SimulateLedger).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SimulateLedger{config: slq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(slq.modifiers) > 0 {
		_spec.Modifiers = slq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, slq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (slq *SimulateLedgerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := slq.querySpec()
	if len(slq.modifiers) > 0 {
		_spec.Modifiers = slq.modifiers
	}
	_spec.Node.Columns = slq.fields
	if len(slq.fields) > 0 {
		_spec.Unique = slq.unique != nil && *slq.unique
	}
	return sqlgraph.CountNodes(ctx, slq.driver, _spec)
}

func (slq *SimulateLedgerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := slq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (slq *SimulateLedgerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   simulateledger.Table,
			Columns: simulateledger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: simulateledger.FieldID,
			},
		},
		From:   slq.sql,
		Unique: true,
	}
	if unique := slq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := slq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, simulateledger.FieldID)
		for i := range fields {
			if fields[i] != simulateledger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := slq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := slq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := slq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := slq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (slq *SimulateLedgerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(slq.driver.Dialect())
	t1 := builder.Table(simulateledger.Table)
	columns := slq.fields
	if len(columns) == 0 {
		columns = simulateledger.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if slq.sql != nil {
		selector = slq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if slq.unique != nil && *slq.unique {
		selector.Distinct()
	}
	for _, m := range slq.modifiers {
		m(selector)
	}
	for _, p := range slq.predicates {
		p(selector)
	}
	for _, p := range slq.order {
		p(selector)
	}
	if offset := slq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := slq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (slq *SimulateLedgerQuery) ForUpdate(opts ...sql.LockOption) *SimulateLedgerQuery {
	if slq.driver.Dialect() == dialect.Postgres {
		slq.Unique(false)
	}
	slq.modifiers = append(slq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return slq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (slq *SimulateLedgerQuery) ForShare(opts ...sql.LockOption) *SimulateLedgerQuery {
	if slq.driver.Dialect() == dialect.Postgres {
		slq.Unique(false)
	}
	slq.modifiers = append(slq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return slq
}

// SimulateLedgerGroupBy is the group-by builder for SimulateLedger entities.
type SimulateLedgerGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (slgb *SimulateLedgerGroupBy) Aggregate(fns ...AggregateFunc) *SimulateLedgerGroupBy {
	slgb.fns = append(slgb.fns, fns...)
	return slgb
}

// Scan applies the group-by query and scans the result into the given value.
func (slgb *SimulateLedgerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := slgb.path(ctx)
	if err != nil {
		return err
	}
	slgb.sql = query
	return slgb.sqlScan(ctx, v)
}

func (slgb *SimulateLedgerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range slgb.fields {
		if !simulateledger.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := slgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := slgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (slgb *SimulateLedgerGroupBy) sqlQuery() *sql.Selector {
	selector := slgb.sql.Select()
	aggregation := make([]string, 0, len(slgb.fns))
	for _, fn := range slgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(slgb.fields)+len(slgb.fns))
		for _, f := range slgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(slgb.fields...)...)
}

// SimulateLedgerSelect is the builder for selecting fields of SimulateLedger entities.
type SimulateLedgerSelect struct {
	*SimulateLedgerQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sls *SimulateLedgerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sls.prepareQuery(ctx); err != nil {
		return err
	}
	sls.sql = sls.SimulateLedgerQuery.sqlQuery(ctx)
	return sls.sqlScan(ctx, v)
}

func (sls *SimulateLedgerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sls.sql.Query()
	if err := sls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
