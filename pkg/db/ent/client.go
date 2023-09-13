// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledgerlock"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/profit"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// GoodLedger is the client for interacting with the GoodLedger builders.
	GoodLedger *GoodLedgerClient
	// GoodStatement is the client for interacting with the GoodStatement builders.
	GoodStatement *GoodStatementClient
	// Ledger is the client for interacting with the Ledger builders.
	Ledger *LedgerClient
	// LedgerLock is the client for interacting with the LedgerLock builders.
	LedgerLock *LedgerLockClient
	// Profit is the client for interacting with the Profit builders.
	Profit *ProfitClient
	// Statement is the client for interacting with the Statement builders.
	Statement *StatementClient
	// UnsoldStatement is the client for interacting with the UnsoldStatement builders.
	UnsoldStatement *UnsoldStatementClient
	// Withdraw is the client for interacting with the Withdraw builders.
	Withdraw *WithdrawClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.GoodLedger = NewGoodLedgerClient(c.config)
	c.GoodStatement = NewGoodStatementClient(c.config)
	c.Ledger = NewLedgerClient(c.config)
	c.LedgerLock = NewLedgerLockClient(c.config)
	c.Profit = NewProfitClient(c.config)
	c.Statement = NewStatementClient(c.config)
	c.UnsoldStatement = NewUnsoldStatementClient(c.config)
	c.Withdraw = NewWithdrawClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		GoodLedger:      NewGoodLedgerClient(cfg),
		GoodStatement:   NewGoodStatementClient(cfg),
		Ledger:          NewLedgerClient(cfg),
		LedgerLock:      NewLedgerLockClient(cfg),
		Profit:          NewProfitClient(cfg),
		Statement:       NewStatementClient(cfg),
		UnsoldStatement: NewUnsoldStatementClient(cfg),
		Withdraw:        NewWithdrawClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		GoodLedger:      NewGoodLedgerClient(cfg),
		GoodStatement:   NewGoodStatementClient(cfg),
		Ledger:          NewLedgerClient(cfg),
		LedgerLock:      NewLedgerLockClient(cfg),
		Profit:          NewProfitClient(cfg),
		Statement:       NewStatementClient(cfg),
		UnsoldStatement: NewUnsoldStatementClient(cfg),
		Withdraw:        NewWithdrawClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		GoodLedger.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.GoodLedger.Use(hooks...)
	c.GoodStatement.Use(hooks...)
	c.Ledger.Use(hooks...)
	c.LedgerLock.Use(hooks...)
	c.Profit.Use(hooks...)
	c.Statement.Use(hooks...)
	c.UnsoldStatement.Use(hooks...)
	c.Withdraw.Use(hooks...)
}

// GoodLedgerClient is a client for the GoodLedger schema.
type GoodLedgerClient struct {
	config
}

// NewGoodLedgerClient returns a client for the GoodLedger from the given config.
func NewGoodLedgerClient(c config) *GoodLedgerClient {
	return &GoodLedgerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `goodledger.Hooks(f(g(h())))`.
func (c *GoodLedgerClient) Use(hooks ...Hook) {
	c.hooks.GoodLedger = append(c.hooks.GoodLedger, hooks...)
}

// Create returns a builder for creating a GoodLedger entity.
func (c *GoodLedgerClient) Create() *GoodLedgerCreate {
	mutation := newGoodLedgerMutation(c.config, OpCreate)
	return &GoodLedgerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GoodLedger entities.
func (c *GoodLedgerClient) CreateBulk(builders ...*GoodLedgerCreate) *GoodLedgerCreateBulk {
	return &GoodLedgerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GoodLedger.
func (c *GoodLedgerClient) Update() *GoodLedgerUpdate {
	mutation := newGoodLedgerMutation(c.config, OpUpdate)
	return &GoodLedgerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GoodLedgerClient) UpdateOne(gl *GoodLedger) *GoodLedgerUpdateOne {
	mutation := newGoodLedgerMutation(c.config, OpUpdateOne, withGoodLedger(gl))
	return &GoodLedgerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GoodLedgerClient) UpdateOneID(id uuid.UUID) *GoodLedgerUpdateOne {
	mutation := newGoodLedgerMutation(c.config, OpUpdateOne, withGoodLedgerID(id))
	return &GoodLedgerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GoodLedger.
func (c *GoodLedgerClient) Delete() *GoodLedgerDelete {
	mutation := newGoodLedgerMutation(c.config, OpDelete)
	return &GoodLedgerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GoodLedgerClient) DeleteOne(gl *GoodLedger) *GoodLedgerDeleteOne {
	return c.DeleteOneID(gl.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GoodLedgerClient) DeleteOneID(id uuid.UUID) *GoodLedgerDeleteOne {
	builder := c.Delete().Where(goodledger.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GoodLedgerDeleteOne{builder}
}

// Query returns a query builder for GoodLedger.
func (c *GoodLedgerClient) Query() *GoodLedgerQuery {
	return &GoodLedgerQuery{
		config: c.config,
	}
}

// Get returns a GoodLedger entity by its id.
func (c *GoodLedgerClient) Get(ctx context.Context, id uuid.UUID) (*GoodLedger, error) {
	return c.Query().Where(goodledger.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GoodLedgerClient) GetX(ctx context.Context, id uuid.UUID) *GoodLedger {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GoodLedgerClient) Hooks() []Hook {
	hooks := c.hooks.GoodLedger
	return append(hooks[:len(hooks):len(hooks)], goodledger.Hooks[:]...)
}

// GoodStatementClient is a client for the GoodStatement schema.
type GoodStatementClient struct {
	config
}

// NewGoodStatementClient returns a client for the GoodStatement from the given config.
func NewGoodStatementClient(c config) *GoodStatementClient {
	return &GoodStatementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `goodstatement.Hooks(f(g(h())))`.
func (c *GoodStatementClient) Use(hooks ...Hook) {
	c.hooks.GoodStatement = append(c.hooks.GoodStatement, hooks...)
}

// Create returns a builder for creating a GoodStatement entity.
func (c *GoodStatementClient) Create() *GoodStatementCreate {
	mutation := newGoodStatementMutation(c.config, OpCreate)
	return &GoodStatementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GoodStatement entities.
func (c *GoodStatementClient) CreateBulk(builders ...*GoodStatementCreate) *GoodStatementCreateBulk {
	return &GoodStatementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GoodStatement.
func (c *GoodStatementClient) Update() *GoodStatementUpdate {
	mutation := newGoodStatementMutation(c.config, OpUpdate)
	return &GoodStatementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GoodStatementClient) UpdateOne(gs *GoodStatement) *GoodStatementUpdateOne {
	mutation := newGoodStatementMutation(c.config, OpUpdateOne, withGoodStatement(gs))
	return &GoodStatementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GoodStatementClient) UpdateOneID(id uuid.UUID) *GoodStatementUpdateOne {
	mutation := newGoodStatementMutation(c.config, OpUpdateOne, withGoodStatementID(id))
	return &GoodStatementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GoodStatement.
func (c *GoodStatementClient) Delete() *GoodStatementDelete {
	mutation := newGoodStatementMutation(c.config, OpDelete)
	return &GoodStatementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GoodStatementClient) DeleteOne(gs *GoodStatement) *GoodStatementDeleteOne {
	return c.DeleteOneID(gs.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GoodStatementClient) DeleteOneID(id uuid.UUID) *GoodStatementDeleteOne {
	builder := c.Delete().Where(goodstatement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GoodStatementDeleteOne{builder}
}

// Query returns a query builder for GoodStatement.
func (c *GoodStatementClient) Query() *GoodStatementQuery {
	return &GoodStatementQuery{
		config: c.config,
	}
}

// Get returns a GoodStatement entity by its id.
func (c *GoodStatementClient) Get(ctx context.Context, id uuid.UUID) (*GoodStatement, error) {
	return c.Query().Where(goodstatement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GoodStatementClient) GetX(ctx context.Context, id uuid.UUID) *GoodStatement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GoodStatementClient) Hooks() []Hook {
	hooks := c.hooks.GoodStatement
	return append(hooks[:len(hooks):len(hooks)], goodstatement.Hooks[:]...)
}

// LedgerClient is a client for the Ledger schema.
type LedgerClient struct {
	config
}

// NewLedgerClient returns a client for the Ledger from the given config.
func NewLedgerClient(c config) *LedgerClient {
	return &LedgerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ledger.Hooks(f(g(h())))`.
func (c *LedgerClient) Use(hooks ...Hook) {
	c.hooks.Ledger = append(c.hooks.Ledger, hooks...)
}

// Create returns a builder for creating a Ledger entity.
func (c *LedgerClient) Create() *LedgerCreate {
	mutation := newLedgerMutation(c.config, OpCreate)
	return &LedgerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ledger entities.
func (c *LedgerClient) CreateBulk(builders ...*LedgerCreate) *LedgerCreateBulk {
	return &LedgerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ledger.
func (c *LedgerClient) Update() *LedgerUpdate {
	mutation := newLedgerMutation(c.config, OpUpdate)
	return &LedgerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LedgerClient) UpdateOne(l *Ledger) *LedgerUpdateOne {
	mutation := newLedgerMutation(c.config, OpUpdateOne, withLedger(l))
	return &LedgerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LedgerClient) UpdateOneID(id uuid.UUID) *LedgerUpdateOne {
	mutation := newLedgerMutation(c.config, OpUpdateOne, withLedgerID(id))
	return &LedgerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ledger.
func (c *LedgerClient) Delete() *LedgerDelete {
	mutation := newLedgerMutation(c.config, OpDelete)
	return &LedgerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LedgerClient) DeleteOne(l *Ledger) *LedgerDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *LedgerClient) DeleteOneID(id uuid.UUID) *LedgerDeleteOne {
	builder := c.Delete().Where(ledger.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LedgerDeleteOne{builder}
}

// Query returns a query builder for Ledger.
func (c *LedgerClient) Query() *LedgerQuery {
	return &LedgerQuery{
		config: c.config,
	}
}

// Get returns a Ledger entity by its id.
func (c *LedgerClient) Get(ctx context.Context, id uuid.UUID) (*Ledger, error) {
	return c.Query().Where(ledger.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LedgerClient) GetX(ctx context.Context, id uuid.UUID) *Ledger {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LedgerClient) Hooks() []Hook {
	hooks := c.hooks.Ledger
	return append(hooks[:len(hooks):len(hooks)], ledger.Hooks[:]...)
}

// LedgerLockClient is a client for the LedgerLock schema.
type LedgerLockClient struct {
	config
}

// NewLedgerLockClient returns a client for the LedgerLock from the given config.
func NewLedgerLockClient(c config) *LedgerLockClient {
	return &LedgerLockClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ledgerlock.Hooks(f(g(h())))`.
func (c *LedgerLockClient) Use(hooks ...Hook) {
	c.hooks.LedgerLock = append(c.hooks.LedgerLock, hooks...)
}

// Create returns a builder for creating a LedgerLock entity.
func (c *LedgerLockClient) Create() *LedgerLockCreate {
	mutation := newLedgerLockMutation(c.config, OpCreate)
	return &LedgerLockCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LedgerLock entities.
func (c *LedgerLockClient) CreateBulk(builders ...*LedgerLockCreate) *LedgerLockCreateBulk {
	return &LedgerLockCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LedgerLock.
func (c *LedgerLockClient) Update() *LedgerLockUpdate {
	mutation := newLedgerLockMutation(c.config, OpUpdate)
	return &LedgerLockUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LedgerLockClient) UpdateOne(ll *LedgerLock) *LedgerLockUpdateOne {
	mutation := newLedgerLockMutation(c.config, OpUpdateOne, withLedgerLock(ll))
	return &LedgerLockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LedgerLockClient) UpdateOneID(id uuid.UUID) *LedgerLockUpdateOne {
	mutation := newLedgerLockMutation(c.config, OpUpdateOne, withLedgerLockID(id))
	return &LedgerLockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LedgerLock.
func (c *LedgerLockClient) Delete() *LedgerLockDelete {
	mutation := newLedgerLockMutation(c.config, OpDelete)
	return &LedgerLockDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LedgerLockClient) DeleteOne(ll *LedgerLock) *LedgerLockDeleteOne {
	return c.DeleteOneID(ll.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *LedgerLockClient) DeleteOneID(id uuid.UUID) *LedgerLockDeleteOne {
	builder := c.Delete().Where(ledgerlock.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LedgerLockDeleteOne{builder}
}

// Query returns a query builder for LedgerLock.
func (c *LedgerLockClient) Query() *LedgerLockQuery {
	return &LedgerLockQuery{
		config: c.config,
	}
}

// Get returns a LedgerLock entity by its id.
func (c *LedgerLockClient) Get(ctx context.Context, id uuid.UUID) (*LedgerLock, error) {
	return c.Query().Where(ledgerlock.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LedgerLockClient) GetX(ctx context.Context, id uuid.UUID) *LedgerLock {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LedgerLockClient) Hooks() []Hook {
	hooks := c.hooks.LedgerLock
	return append(hooks[:len(hooks):len(hooks)], ledgerlock.Hooks[:]...)
}

// ProfitClient is a client for the Profit schema.
type ProfitClient struct {
	config
}

// NewProfitClient returns a client for the Profit from the given config.
func NewProfitClient(c config) *ProfitClient {
	return &ProfitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profit.Hooks(f(g(h())))`.
func (c *ProfitClient) Use(hooks ...Hook) {
	c.hooks.Profit = append(c.hooks.Profit, hooks...)
}

// Create returns a builder for creating a Profit entity.
func (c *ProfitClient) Create() *ProfitCreate {
	mutation := newProfitMutation(c.config, OpCreate)
	return &ProfitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profit entities.
func (c *ProfitClient) CreateBulk(builders ...*ProfitCreate) *ProfitCreateBulk {
	return &ProfitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profit.
func (c *ProfitClient) Update() *ProfitUpdate {
	mutation := newProfitMutation(c.config, OpUpdate)
	return &ProfitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfitClient) UpdateOne(pr *Profit) *ProfitUpdateOne {
	mutation := newProfitMutation(c.config, OpUpdateOne, withProfit(pr))
	return &ProfitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfitClient) UpdateOneID(id uuid.UUID) *ProfitUpdateOne {
	mutation := newProfitMutation(c.config, OpUpdateOne, withProfitID(id))
	return &ProfitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profit.
func (c *ProfitClient) Delete() *ProfitDelete {
	mutation := newProfitMutation(c.config, OpDelete)
	return &ProfitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProfitClient) DeleteOne(pr *Profit) *ProfitDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ProfitClient) DeleteOneID(id uuid.UUID) *ProfitDeleteOne {
	builder := c.Delete().Where(profit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfitDeleteOne{builder}
}

// Query returns a query builder for Profit.
func (c *ProfitClient) Query() *ProfitQuery {
	return &ProfitQuery{
		config: c.config,
	}
}

// Get returns a Profit entity by its id.
func (c *ProfitClient) Get(ctx context.Context, id uuid.UUID) (*Profit, error) {
	return c.Query().Where(profit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfitClient) GetX(ctx context.Context, id uuid.UUID) *Profit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProfitClient) Hooks() []Hook {
	hooks := c.hooks.Profit
	return append(hooks[:len(hooks):len(hooks)], profit.Hooks[:]...)
}

// StatementClient is a client for the Statement schema.
type StatementClient struct {
	config
}

// NewStatementClient returns a client for the Statement from the given config.
func NewStatementClient(c config) *StatementClient {
	return &StatementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `statement.Hooks(f(g(h())))`.
func (c *StatementClient) Use(hooks ...Hook) {
	c.hooks.Statement = append(c.hooks.Statement, hooks...)
}

// Create returns a builder for creating a Statement entity.
func (c *StatementClient) Create() *StatementCreate {
	mutation := newStatementMutation(c.config, OpCreate)
	return &StatementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Statement entities.
func (c *StatementClient) CreateBulk(builders ...*StatementCreate) *StatementCreateBulk {
	return &StatementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Statement.
func (c *StatementClient) Update() *StatementUpdate {
	mutation := newStatementMutation(c.config, OpUpdate)
	return &StatementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatementClient) UpdateOne(s *Statement) *StatementUpdateOne {
	mutation := newStatementMutation(c.config, OpUpdateOne, withStatement(s))
	return &StatementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatementClient) UpdateOneID(id uuid.UUID) *StatementUpdateOne {
	mutation := newStatementMutation(c.config, OpUpdateOne, withStatementID(id))
	return &StatementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Statement.
func (c *StatementClient) Delete() *StatementDelete {
	mutation := newStatementMutation(c.config, OpDelete)
	return &StatementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatementClient) DeleteOne(s *Statement) *StatementDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *StatementClient) DeleteOneID(id uuid.UUID) *StatementDeleteOne {
	builder := c.Delete().Where(statement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatementDeleteOne{builder}
}

// Query returns a query builder for Statement.
func (c *StatementClient) Query() *StatementQuery {
	return &StatementQuery{
		config: c.config,
	}
}

// Get returns a Statement entity by its id.
func (c *StatementClient) Get(ctx context.Context, id uuid.UUID) (*Statement, error) {
	return c.Query().Where(statement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatementClient) GetX(ctx context.Context, id uuid.UUID) *Statement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StatementClient) Hooks() []Hook {
	hooks := c.hooks.Statement
	return append(hooks[:len(hooks):len(hooks)], statement.Hooks[:]...)
}

// UnsoldStatementClient is a client for the UnsoldStatement schema.
type UnsoldStatementClient struct {
	config
}

// NewUnsoldStatementClient returns a client for the UnsoldStatement from the given config.
func NewUnsoldStatementClient(c config) *UnsoldStatementClient {
	return &UnsoldStatementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `unsoldstatement.Hooks(f(g(h())))`.
func (c *UnsoldStatementClient) Use(hooks ...Hook) {
	c.hooks.UnsoldStatement = append(c.hooks.UnsoldStatement, hooks...)
}

// Create returns a builder for creating a UnsoldStatement entity.
func (c *UnsoldStatementClient) Create() *UnsoldStatementCreate {
	mutation := newUnsoldStatementMutation(c.config, OpCreate)
	return &UnsoldStatementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UnsoldStatement entities.
func (c *UnsoldStatementClient) CreateBulk(builders ...*UnsoldStatementCreate) *UnsoldStatementCreateBulk {
	return &UnsoldStatementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UnsoldStatement.
func (c *UnsoldStatementClient) Update() *UnsoldStatementUpdate {
	mutation := newUnsoldStatementMutation(c.config, OpUpdate)
	return &UnsoldStatementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UnsoldStatementClient) UpdateOne(us *UnsoldStatement) *UnsoldStatementUpdateOne {
	mutation := newUnsoldStatementMutation(c.config, OpUpdateOne, withUnsoldStatement(us))
	return &UnsoldStatementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UnsoldStatementClient) UpdateOneID(id uuid.UUID) *UnsoldStatementUpdateOne {
	mutation := newUnsoldStatementMutation(c.config, OpUpdateOne, withUnsoldStatementID(id))
	return &UnsoldStatementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UnsoldStatement.
func (c *UnsoldStatementClient) Delete() *UnsoldStatementDelete {
	mutation := newUnsoldStatementMutation(c.config, OpDelete)
	return &UnsoldStatementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UnsoldStatementClient) DeleteOne(us *UnsoldStatement) *UnsoldStatementDeleteOne {
	return c.DeleteOneID(us.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UnsoldStatementClient) DeleteOneID(id uuid.UUID) *UnsoldStatementDeleteOne {
	builder := c.Delete().Where(unsoldstatement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UnsoldStatementDeleteOne{builder}
}

// Query returns a query builder for UnsoldStatement.
func (c *UnsoldStatementClient) Query() *UnsoldStatementQuery {
	return &UnsoldStatementQuery{
		config: c.config,
	}
}

// Get returns a UnsoldStatement entity by its id.
func (c *UnsoldStatementClient) Get(ctx context.Context, id uuid.UUID) (*UnsoldStatement, error) {
	return c.Query().Where(unsoldstatement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UnsoldStatementClient) GetX(ctx context.Context, id uuid.UUID) *UnsoldStatement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UnsoldStatementClient) Hooks() []Hook {
	hooks := c.hooks.UnsoldStatement
	return append(hooks[:len(hooks):len(hooks)], unsoldstatement.Hooks[:]...)
}

// WithdrawClient is a client for the Withdraw schema.
type WithdrawClient struct {
	config
}

// NewWithdrawClient returns a client for the Withdraw from the given config.
func NewWithdrawClient(c config) *WithdrawClient {
	return &WithdrawClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `withdraw.Hooks(f(g(h())))`.
func (c *WithdrawClient) Use(hooks ...Hook) {
	c.hooks.Withdraw = append(c.hooks.Withdraw, hooks...)
}

// Create returns a builder for creating a Withdraw entity.
func (c *WithdrawClient) Create() *WithdrawCreate {
	mutation := newWithdrawMutation(c.config, OpCreate)
	return &WithdrawCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Withdraw entities.
func (c *WithdrawClient) CreateBulk(builders ...*WithdrawCreate) *WithdrawCreateBulk {
	return &WithdrawCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Withdraw.
func (c *WithdrawClient) Update() *WithdrawUpdate {
	mutation := newWithdrawMutation(c.config, OpUpdate)
	return &WithdrawUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WithdrawClient) UpdateOne(w *Withdraw) *WithdrawUpdateOne {
	mutation := newWithdrawMutation(c.config, OpUpdateOne, withWithdraw(w))
	return &WithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WithdrawClient) UpdateOneID(id uuid.UUID) *WithdrawUpdateOne {
	mutation := newWithdrawMutation(c.config, OpUpdateOne, withWithdrawID(id))
	return &WithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Withdraw.
func (c *WithdrawClient) Delete() *WithdrawDelete {
	mutation := newWithdrawMutation(c.config, OpDelete)
	return &WithdrawDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WithdrawClient) DeleteOne(w *Withdraw) *WithdrawDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *WithdrawClient) DeleteOneID(id uuid.UUID) *WithdrawDeleteOne {
	builder := c.Delete().Where(withdraw.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WithdrawDeleteOne{builder}
}

// Query returns a query builder for Withdraw.
func (c *WithdrawClient) Query() *WithdrawQuery {
	return &WithdrawQuery{
		config: c.config,
	}
}

// Get returns a Withdraw entity by its id.
func (c *WithdrawClient) Get(ctx context.Context, id uuid.UUID) (*Withdraw, error) {
	return c.Query().Where(withdraw.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WithdrawClient) GetX(ctx context.Context, id uuid.UUID) *Withdraw {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WithdrawClient) Hooks() []Hook {
	hooks := c.hooks.Withdraw
	return append(hooks[:len(hooks):len(hooks)], withdraw.Hooks[:]...)
}
