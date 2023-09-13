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
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodStatementQuery is the builder for querying GoodStatement entities.
type GoodStatementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodStatement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodStatementQuery builder.
func (gsq *GoodStatementQuery) Where(ps ...predicate.GoodStatement) *GoodStatementQuery {
	gsq.predicates = append(gsq.predicates, ps...)
	return gsq
}

// Limit adds a limit step to the query.
func (gsq *GoodStatementQuery) Limit(limit int) *GoodStatementQuery {
	gsq.limit = &limit
	return gsq
}

// Offset adds an offset step to the query.
func (gsq *GoodStatementQuery) Offset(offset int) *GoodStatementQuery {
	gsq.offset = &offset
	return gsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gsq *GoodStatementQuery) Unique(unique bool) *GoodStatementQuery {
	gsq.unique = &unique
	return gsq
}

// Order adds an order step to the query.
func (gsq *GoodStatementQuery) Order(o ...OrderFunc) *GoodStatementQuery {
	gsq.order = append(gsq.order, o...)
	return gsq
}

// First returns the first GoodStatement entity from the query.
// Returns a *NotFoundError when no GoodStatement was found.
func (gsq *GoodStatementQuery) First(ctx context.Context) (*GoodStatement, error) {
	nodes, err := gsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodstatement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gsq *GoodStatementQuery) FirstX(ctx context.Context) *GoodStatement {
	node, err := gsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodStatement ID from the query.
// Returns a *NotFoundError when no GoodStatement ID was found.
func (gsq *GoodStatementQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodstatement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gsq *GoodStatementQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodStatement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodStatement entity is found.
// Returns a *NotFoundError when no GoodStatement entities are found.
func (gsq *GoodStatementQuery) Only(ctx context.Context) (*GoodStatement, error) {
	nodes, err := gsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodstatement.Label}
	default:
		return nil, &NotSingularError{goodstatement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gsq *GoodStatementQuery) OnlyX(ctx context.Context) *GoodStatement {
	node, err := gsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodStatement ID in the query.
// Returns a *NotSingularError when more than one GoodStatement ID is found.
// Returns a *NotFoundError when no entities are found.
func (gsq *GoodStatementQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodstatement.Label}
	default:
		err = &NotSingularError{goodstatement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gsq *GoodStatementQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodStatements.
func (gsq *GoodStatementQuery) All(ctx context.Context) ([]*GoodStatement, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gsq *GoodStatementQuery) AllX(ctx context.Context) []*GoodStatement {
	nodes, err := gsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodStatement IDs.
func (gsq *GoodStatementQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gsq.Select(goodstatement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gsq *GoodStatementQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gsq *GoodStatementQuery) Count(ctx context.Context) (int, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gsq *GoodStatementQuery) CountX(ctx context.Context) int {
	count, err := gsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gsq *GoodStatementQuery) Exist(ctx context.Context) (bool, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gsq *GoodStatementQuery) ExistX(ctx context.Context) bool {
	exist, err := gsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodStatementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gsq *GoodStatementQuery) Clone() *GoodStatementQuery {
	if gsq == nil {
		return nil
	}
	return &GoodStatementQuery{
		config:     gsq.config,
		limit:      gsq.limit,
		offset:     gsq.offset,
		order:      append([]OrderFunc{}, gsq.order...),
		predicates: append([]predicate.GoodStatement{}, gsq.predicates...),
		// clone intermediate query.
		sql:    gsq.sql.Clone(),
		path:   gsq.path,
		unique: gsq.unique,
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
//	client.GoodStatement.Query().
//		GroupBy(goodstatement.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gsq *GoodStatementQuery) GroupBy(field string, fields ...string) *GoodStatementGroupBy {
	grbuild := &GoodStatementGroupBy{config: gsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gsq.sqlQuery(ctx), nil
	}
	grbuild.label = goodstatement.Label
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
//	client.GoodStatement.Query().
//		Select(goodstatement.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (gsq *GoodStatementQuery) Select(fields ...string) *GoodStatementSelect {
	gsq.fields = append(gsq.fields, fields...)
	selbuild := &GoodStatementSelect{GoodStatementQuery: gsq}
	selbuild.label = goodstatement.Label
	selbuild.flds, selbuild.scan = &gsq.fields, selbuild.Scan
	return selbuild
}

func (gsq *GoodStatementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gsq.fields {
		if !goodstatement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gsq.path != nil {
		prev, err := gsq.path(ctx)
		if err != nil {
			return err
		}
		gsq.sql = prev
	}
	if goodstatement.Policy == nil {
		return errors.New("ent: uninitialized goodstatement.Policy (forgotten import ent/runtime?)")
	}
	if err := goodstatement.Policy.EvalQuery(ctx, gsq); err != nil {
		return err
	}
	return nil
}

func (gsq *GoodStatementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodStatement, error) {
	var (
		nodes = []*GoodStatement{}
		_spec = gsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GoodStatement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GoodStatement{config: gsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gsq *GoodStatementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gsq.querySpec()
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	_spec.Node.Columns = gsq.fields
	if len(gsq.fields) > 0 {
		_spec.Unique = gsq.unique != nil && *gsq.unique
	}
	return sqlgraph.CountNodes(ctx, gsq.driver, _spec)
}

func (gsq *GoodStatementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gsq *GoodStatementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodstatement.Table,
			Columns: goodstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodstatement.FieldID,
			},
		},
		From:   gsq.sql,
		Unique: true,
	}
	if unique := gsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodstatement.FieldID)
		for i := range fields {
			if fields[i] != goodstatement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gsq *GoodStatementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gsq.driver.Dialect())
	t1 := builder.Table(goodstatement.Table)
	columns := gsq.fields
	if len(columns) == 0 {
		columns = goodstatement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gsq.sql != nil {
		selector = gsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gsq.unique != nil && *gsq.unique {
		selector.Distinct()
	}
	for _, m := range gsq.modifiers {
		m(selector)
	}
	for _, p := range gsq.predicates {
		p(selector)
	}
	for _, p := range gsq.order {
		p(selector)
	}
	if offset := gsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gsq *GoodStatementQuery) ForUpdate(opts ...sql.LockOption) *GoodStatementQuery {
	if gsq.driver.Dialect() == dialect.Postgres {
		gsq.Unique(false)
	}
	gsq.modifiers = append(gsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gsq *GoodStatementQuery) ForShare(opts ...sql.LockOption) *GoodStatementQuery {
	if gsq.driver.Dialect() == dialect.Postgres {
		gsq.Unique(false)
	}
	gsq.modifiers = append(gsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gsq
}

// GoodStatementGroupBy is the group-by builder for GoodStatement entities.
type GoodStatementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gsgb *GoodStatementGroupBy) Aggregate(fns ...AggregateFunc) *GoodStatementGroupBy {
	gsgb.fns = append(gsgb.fns, fns...)
	return gsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gsgb *GoodStatementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gsgb.path(ctx)
	if err != nil {
		return err
	}
	gsgb.sql = query
	return gsgb.sqlScan(ctx, v)
}

func (gsgb *GoodStatementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gsgb.fields {
		if !goodstatement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gsgb *GoodStatementGroupBy) sqlQuery() *sql.Selector {
	selector := gsgb.sql.Select()
	aggregation := make([]string, 0, len(gsgb.fns))
	for _, fn := range gsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gsgb.fields)+len(gsgb.fns))
		for _, f := range gsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gsgb.fields...)...)
}

// GoodStatementSelect is the builder for selecting fields of GoodStatement entities.
type GoodStatementSelect struct {
	*GoodStatementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gss *GoodStatementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gss.prepareQuery(ctx); err != nil {
		return err
	}
	gss.sql = gss.GoodStatementQuery.sqlQuery(ctx)
	return gss.sqlScan(ctx, v)
}

func (gss *GoodStatementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gss.sql.Query()
	if err := gss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
