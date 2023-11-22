// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledgerlock"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/profit"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 8)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   goodledger.Table,
			Columns: goodledger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodledger.FieldID,
			},
		},
		Type: "GoodLedger",
		Fields: map[string]*sqlgraph.FieldSpec{
			goodledger.FieldCreatedAt:  {Type: field.TypeUint32, Column: goodledger.FieldCreatedAt},
			goodledger.FieldUpdatedAt:  {Type: field.TypeUint32, Column: goodledger.FieldUpdatedAt},
			goodledger.FieldDeletedAt:  {Type: field.TypeUint32, Column: goodledger.FieldDeletedAt},
			goodledger.FieldEntID:      {Type: field.TypeUUID, Column: goodledger.FieldEntID},
			goodledger.FieldGoodID:     {Type: field.TypeUUID, Column: goodledger.FieldGoodID},
			goodledger.FieldCoinTypeID: {Type: field.TypeUUID, Column: goodledger.FieldCoinTypeID},
			goodledger.FieldAmount:     {Type: field.TypeFloat64, Column: goodledger.FieldAmount},
			goodledger.FieldToPlatform: {Type: field.TypeFloat64, Column: goodledger.FieldToPlatform},
			goodledger.FieldToUser:     {Type: field.TypeFloat64, Column: goodledger.FieldToUser},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   goodstatement.Table,
			Columns: goodstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodstatement.FieldID,
			},
		},
		Type: "GoodStatement",
		Fields: map[string]*sqlgraph.FieldSpec{
			goodstatement.FieldCreatedAt:                 {Type: field.TypeUint32, Column: goodstatement.FieldCreatedAt},
			goodstatement.FieldUpdatedAt:                 {Type: field.TypeUint32, Column: goodstatement.FieldUpdatedAt},
			goodstatement.FieldDeletedAt:                 {Type: field.TypeUint32, Column: goodstatement.FieldDeletedAt},
			goodstatement.FieldEntID:                     {Type: field.TypeUUID, Column: goodstatement.FieldEntID},
			goodstatement.FieldGoodID:                    {Type: field.TypeUUID, Column: goodstatement.FieldGoodID},
			goodstatement.FieldCoinTypeID:                {Type: field.TypeUUID, Column: goodstatement.FieldCoinTypeID},
			goodstatement.FieldAmount:                    {Type: field.TypeFloat64, Column: goodstatement.FieldAmount},
			goodstatement.FieldToPlatform:                {Type: field.TypeFloat64, Column: goodstatement.FieldToPlatform},
			goodstatement.FieldToUser:                    {Type: field.TypeFloat64, Column: goodstatement.FieldToUser},
			goodstatement.FieldTechniqueServiceFeeAmount: {Type: field.TypeFloat64, Column: goodstatement.FieldTechniqueServiceFeeAmount},
			goodstatement.FieldBenefitDate:               {Type: field.TypeUint32, Column: goodstatement.FieldBenefitDate},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   ledger.Table,
			Columns: ledger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: ledger.FieldID,
			},
		},
		Type: "Ledger",
		Fields: map[string]*sqlgraph.FieldSpec{
			ledger.FieldCreatedAt:  {Type: field.TypeUint32, Column: ledger.FieldCreatedAt},
			ledger.FieldUpdatedAt:  {Type: field.TypeUint32, Column: ledger.FieldUpdatedAt},
			ledger.FieldDeletedAt:  {Type: field.TypeUint32, Column: ledger.FieldDeletedAt},
			ledger.FieldEntID:      {Type: field.TypeUUID, Column: ledger.FieldEntID},
			ledger.FieldAppID:      {Type: field.TypeUUID, Column: ledger.FieldAppID},
			ledger.FieldUserID:     {Type: field.TypeUUID, Column: ledger.FieldUserID},
			ledger.FieldCoinTypeID: {Type: field.TypeUUID, Column: ledger.FieldCoinTypeID},
			ledger.FieldIncoming:   {Type: field.TypeFloat64, Column: ledger.FieldIncoming},
			ledger.FieldLocked:     {Type: field.TypeFloat64, Column: ledger.FieldLocked},
			ledger.FieldOutcoming:  {Type: field.TypeFloat64, Column: ledger.FieldOutcoming},
			ledger.FieldSpendable:  {Type: field.TypeFloat64, Column: ledger.FieldSpendable},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   ledgerlock.Table,
			Columns: ledgerlock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: ledgerlock.FieldID,
			},
		},
		Type: "LedgerLock",
		Fields: map[string]*sqlgraph.FieldSpec{
			ledgerlock.FieldCreatedAt:   {Type: field.TypeUint32, Column: ledgerlock.FieldCreatedAt},
			ledgerlock.FieldUpdatedAt:   {Type: field.TypeUint32, Column: ledgerlock.FieldUpdatedAt},
			ledgerlock.FieldDeletedAt:   {Type: field.TypeUint32, Column: ledgerlock.FieldDeletedAt},
			ledgerlock.FieldEntID:       {Type: field.TypeUUID, Column: ledgerlock.FieldEntID},
			ledgerlock.FieldLedgerID:    {Type: field.TypeUUID, Column: ledgerlock.FieldLedgerID},
			ledgerlock.FieldStatementID: {Type: field.TypeUUID, Column: ledgerlock.FieldStatementID},
			ledgerlock.FieldAmount:      {Type: field.TypeFloat64, Column: ledgerlock.FieldAmount},
			ledgerlock.FieldLockState:   {Type: field.TypeString, Column: ledgerlock.FieldLockState},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   profit.Table,
			Columns: profit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: profit.FieldID,
			},
		},
		Type: "Profit",
		Fields: map[string]*sqlgraph.FieldSpec{
			profit.FieldCreatedAt:  {Type: field.TypeUint32, Column: profit.FieldCreatedAt},
			profit.FieldUpdatedAt:  {Type: field.TypeUint32, Column: profit.FieldUpdatedAt},
			profit.FieldDeletedAt:  {Type: field.TypeUint32, Column: profit.FieldDeletedAt},
			profit.FieldEntID:      {Type: field.TypeUUID, Column: profit.FieldEntID},
			profit.FieldAppID:      {Type: field.TypeUUID, Column: profit.FieldAppID},
			profit.FieldUserID:     {Type: field.TypeUUID, Column: profit.FieldUserID},
			profit.FieldCoinTypeID: {Type: field.TypeUUID, Column: profit.FieldCoinTypeID},
			profit.FieldIncoming:   {Type: field.TypeFloat64, Column: profit.FieldIncoming},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   statement.Table,
			Columns: statement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: statement.FieldID,
			},
		},
		Type: "Statement",
		Fields: map[string]*sqlgraph.FieldSpec{
			statement.FieldCreatedAt:  {Type: field.TypeUint32, Column: statement.FieldCreatedAt},
			statement.FieldUpdatedAt:  {Type: field.TypeUint32, Column: statement.FieldUpdatedAt},
			statement.FieldDeletedAt:  {Type: field.TypeUint32, Column: statement.FieldDeletedAt},
			statement.FieldEntID:      {Type: field.TypeUUID, Column: statement.FieldEntID},
			statement.FieldAppID:      {Type: field.TypeUUID, Column: statement.FieldAppID},
			statement.FieldUserID:     {Type: field.TypeUUID, Column: statement.FieldUserID},
			statement.FieldCoinTypeID: {Type: field.TypeUUID, Column: statement.FieldCoinTypeID},
			statement.FieldIoType:     {Type: field.TypeString, Column: statement.FieldIoType},
			statement.FieldIoSubType:  {Type: field.TypeString, Column: statement.FieldIoSubType},
			statement.FieldAmount:     {Type: field.TypeFloat64, Column: statement.FieldAmount},
			statement.FieldIoExtra:    {Type: field.TypeString, Column: statement.FieldIoExtra},
			statement.FieldIoExtraV1:  {Type: field.TypeString, Column: statement.FieldIoExtraV1},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   unsoldstatement.Table,
			Columns: unsoldstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: unsoldstatement.FieldID,
			},
		},
		Type: "UnsoldStatement",
		Fields: map[string]*sqlgraph.FieldSpec{
			unsoldstatement.FieldCreatedAt:   {Type: field.TypeUint32, Column: unsoldstatement.FieldCreatedAt},
			unsoldstatement.FieldUpdatedAt:   {Type: field.TypeUint32, Column: unsoldstatement.FieldUpdatedAt},
			unsoldstatement.FieldDeletedAt:   {Type: field.TypeUint32, Column: unsoldstatement.FieldDeletedAt},
			unsoldstatement.FieldEntID:       {Type: field.TypeUUID, Column: unsoldstatement.FieldEntID},
			unsoldstatement.FieldGoodID:      {Type: field.TypeUUID, Column: unsoldstatement.FieldGoodID},
			unsoldstatement.FieldCoinTypeID:  {Type: field.TypeUUID, Column: unsoldstatement.FieldCoinTypeID},
			unsoldstatement.FieldAmount:      {Type: field.TypeFloat64, Column: unsoldstatement.FieldAmount},
			unsoldstatement.FieldBenefitDate: {Type: field.TypeUint32, Column: unsoldstatement.FieldBenefitDate},
			unsoldstatement.FieldStatementID: {Type: field.TypeUUID, Column: unsoldstatement.FieldStatementID},
		},
	}
	graph.Nodes[7] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   withdraw.Table,
			Columns: withdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: withdraw.FieldID,
			},
		},
		Type: "Withdraw",
		Fields: map[string]*sqlgraph.FieldSpec{
			withdraw.FieldCreatedAt:             {Type: field.TypeUint32, Column: withdraw.FieldCreatedAt},
			withdraw.FieldUpdatedAt:             {Type: field.TypeUint32, Column: withdraw.FieldUpdatedAt},
			withdraw.FieldDeletedAt:             {Type: field.TypeUint32, Column: withdraw.FieldDeletedAt},
			withdraw.FieldEntID:                 {Type: field.TypeUUID, Column: withdraw.FieldEntID},
			withdraw.FieldAppID:                 {Type: field.TypeUUID, Column: withdraw.FieldAppID},
			withdraw.FieldUserID:                {Type: field.TypeUUID, Column: withdraw.FieldUserID},
			withdraw.FieldCoinTypeID:            {Type: field.TypeUUID, Column: withdraw.FieldCoinTypeID},
			withdraw.FieldAccountID:             {Type: field.TypeUUID, Column: withdraw.FieldAccountID},
			withdraw.FieldAddress:               {Type: field.TypeString, Column: withdraw.FieldAddress},
			withdraw.FieldPlatformTransactionID: {Type: field.TypeUUID, Column: withdraw.FieldPlatformTransactionID},
			withdraw.FieldChainTransactionID:    {Type: field.TypeString, Column: withdraw.FieldChainTransactionID},
			withdraw.FieldState:                 {Type: field.TypeString, Column: withdraw.FieldState},
			withdraw.FieldAmount:                {Type: field.TypeFloat64, Column: withdraw.FieldAmount},
			withdraw.FieldReviewID:              {Type: field.TypeUUID, Column: withdraw.FieldReviewID},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (glq *GoodLedgerQuery) addPredicate(pred func(s *sql.Selector)) {
	glq.predicates = append(glq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GoodLedgerQuery builder.
func (glq *GoodLedgerQuery) Filter() *GoodLedgerFilter {
	return &GoodLedgerFilter{config: glq.config, predicateAdder: glq}
}

// addPredicate implements the predicateAdder interface.
func (m *GoodLedgerMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GoodLedgerMutation builder.
func (m *GoodLedgerMutation) Filter() *GoodLedgerFilter {
	return &GoodLedgerFilter{config: m.config, predicateAdder: m}
}

// GoodLedgerFilter provides a generic filtering capability at runtime for GoodLedgerQuery.
type GoodLedgerFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GoodLedgerFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *GoodLedgerFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(goodledger.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *GoodLedgerFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(goodledger.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *GoodLedgerFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(goodledger.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *GoodLedgerFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(goodledger.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *GoodLedgerFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(goodledger.FieldEntID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *GoodLedgerFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(goodledger.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *GoodLedgerFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(goodledger.FieldCoinTypeID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *GoodLedgerFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(goodledger.FieldAmount))
}

// WhereToPlatform applies the entql float64 predicate on the to_platform field.
func (f *GoodLedgerFilter) WhereToPlatform(p entql.Float64P) {
	f.Where(p.Field(goodledger.FieldToPlatform))
}

// WhereToUser applies the entql float64 predicate on the to_user field.
func (f *GoodLedgerFilter) WhereToUser(p entql.Float64P) {
	f.Where(p.Field(goodledger.FieldToUser))
}

// addPredicate implements the predicateAdder interface.
func (gsq *GoodStatementQuery) addPredicate(pred func(s *sql.Selector)) {
	gsq.predicates = append(gsq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GoodStatementQuery builder.
func (gsq *GoodStatementQuery) Filter() *GoodStatementFilter {
	return &GoodStatementFilter{config: gsq.config, predicateAdder: gsq}
}

// addPredicate implements the predicateAdder interface.
func (m *GoodStatementMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GoodStatementMutation builder.
func (m *GoodStatementMutation) Filter() *GoodStatementFilter {
	return &GoodStatementFilter{config: m.config, predicateAdder: m}
}

// GoodStatementFilter provides a generic filtering capability at runtime for GoodStatementQuery.
type GoodStatementFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GoodStatementFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *GoodStatementFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(goodstatement.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *GoodStatementFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(goodstatement.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *GoodStatementFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(goodstatement.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *GoodStatementFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(goodstatement.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *GoodStatementFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(goodstatement.FieldEntID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *GoodStatementFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(goodstatement.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *GoodStatementFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(goodstatement.FieldCoinTypeID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *GoodStatementFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(goodstatement.FieldAmount))
}

// WhereToPlatform applies the entql float64 predicate on the to_platform field.
func (f *GoodStatementFilter) WhereToPlatform(p entql.Float64P) {
	f.Where(p.Field(goodstatement.FieldToPlatform))
}

// WhereToUser applies the entql float64 predicate on the to_user field.
func (f *GoodStatementFilter) WhereToUser(p entql.Float64P) {
	f.Where(p.Field(goodstatement.FieldToUser))
}

// WhereTechniqueServiceFeeAmount applies the entql float64 predicate on the technique_service_fee_amount field.
func (f *GoodStatementFilter) WhereTechniqueServiceFeeAmount(p entql.Float64P) {
	f.Where(p.Field(goodstatement.FieldTechniqueServiceFeeAmount))
}

// WhereBenefitDate applies the entql uint32 predicate on the benefit_date field.
func (f *GoodStatementFilter) WhereBenefitDate(p entql.Uint32P) {
	f.Where(p.Field(goodstatement.FieldBenefitDate))
}

// addPredicate implements the predicateAdder interface.
func (lq *LedgerQuery) addPredicate(pred func(s *sql.Selector)) {
	lq.predicates = append(lq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the LedgerQuery builder.
func (lq *LedgerQuery) Filter() *LedgerFilter {
	return &LedgerFilter{config: lq.config, predicateAdder: lq}
}

// addPredicate implements the predicateAdder interface.
func (m *LedgerMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the LedgerMutation builder.
func (m *LedgerMutation) Filter() *LedgerFilter {
	return &LedgerFilter{config: m.config, predicateAdder: m}
}

// LedgerFilter provides a generic filtering capability at runtime for LedgerQuery.
type LedgerFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *LedgerFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *LedgerFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(ledger.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *LedgerFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(ledger.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *LedgerFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(ledger.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *LedgerFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(ledger.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *LedgerFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(ledger.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *LedgerFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(ledger.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *LedgerFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(ledger.FieldUserID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *LedgerFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(ledger.FieldCoinTypeID))
}

// WhereIncoming applies the entql float64 predicate on the incoming field.
func (f *LedgerFilter) WhereIncoming(p entql.Float64P) {
	f.Where(p.Field(ledger.FieldIncoming))
}

// WhereLocked applies the entql float64 predicate on the locked field.
func (f *LedgerFilter) WhereLocked(p entql.Float64P) {
	f.Where(p.Field(ledger.FieldLocked))
}

// WhereOutcoming applies the entql float64 predicate on the outcoming field.
func (f *LedgerFilter) WhereOutcoming(p entql.Float64P) {
	f.Where(p.Field(ledger.FieldOutcoming))
}

// WhereSpendable applies the entql float64 predicate on the spendable field.
func (f *LedgerFilter) WhereSpendable(p entql.Float64P) {
	f.Where(p.Field(ledger.FieldSpendable))
}

// addPredicate implements the predicateAdder interface.
func (llq *LedgerLockQuery) addPredicate(pred func(s *sql.Selector)) {
	llq.predicates = append(llq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the LedgerLockQuery builder.
func (llq *LedgerLockQuery) Filter() *LedgerLockFilter {
	return &LedgerLockFilter{config: llq.config, predicateAdder: llq}
}

// addPredicate implements the predicateAdder interface.
func (m *LedgerLockMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the LedgerLockMutation builder.
func (m *LedgerLockMutation) Filter() *LedgerLockFilter {
	return &LedgerLockFilter{config: m.config, predicateAdder: m}
}

// LedgerLockFilter provides a generic filtering capability at runtime for LedgerLockQuery.
type LedgerLockFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *LedgerLockFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *LedgerLockFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(ledgerlock.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *LedgerLockFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(ledgerlock.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *LedgerLockFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(ledgerlock.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *LedgerLockFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(ledgerlock.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *LedgerLockFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(ledgerlock.FieldEntID))
}

// WhereLedgerID applies the entql [16]byte predicate on the ledger_id field.
func (f *LedgerLockFilter) WhereLedgerID(p entql.ValueP) {
	f.Where(p.Field(ledgerlock.FieldLedgerID))
}

// WhereStatementID applies the entql [16]byte predicate on the statement_id field.
func (f *LedgerLockFilter) WhereStatementID(p entql.ValueP) {
	f.Where(p.Field(ledgerlock.FieldStatementID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *LedgerLockFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(ledgerlock.FieldAmount))
}

// WhereLockState applies the entql string predicate on the lock_state field.
func (f *LedgerLockFilter) WhereLockState(p entql.StringP) {
	f.Where(p.Field(ledgerlock.FieldLockState))
}

// addPredicate implements the predicateAdder interface.
func (pq *ProfitQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ProfitQuery builder.
func (pq *ProfitQuery) Filter() *ProfitFilter {
	return &ProfitFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *ProfitMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ProfitMutation builder.
func (m *ProfitMutation) Filter() *ProfitFilter {
	return &ProfitFilter{config: m.config, predicateAdder: m}
}

// ProfitFilter provides a generic filtering capability at runtime for ProfitQuery.
type ProfitFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ProfitFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *ProfitFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(profit.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ProfitFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(profit.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ProfitFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(profit.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ProfitFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(profit.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *ProfitFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(profit.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ProfitFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(profit.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *ProfitFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(profit.FieldUserID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *ProfitFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(profit.FieldCoinTypeID))
}

// WhereIncoming applies the entql float64 predicate on the incoming field.
func (f *ProfitFilter) WhereIncoming(p entql.Float64P) {
	f.Where(p.Field(profit.FieldIncoming))
}

// addPredicate implements the predicateAdder interface.
func (sq *StatementQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the StatementQuery builder.
func (sq *StatementQuery) Filter() *StatementFilter {
	return &StatementFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *StatementMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the StatementMutation builder.
func (m *StatementMutation) Filter() *StatementFilter {
	return &StatementFilter{config: m.config, predicateAdder: m}
}

// StatementFilter provides a generic filtering capability at runtime for StatementQuery.
type StatementFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *StatementFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *StatementFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(statement.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *StatementFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(statement.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *StatementFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(statement.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *StatementFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(statement.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *StatementFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(statement.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *StatementFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(statement.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *StatementFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(statement.FieldUserID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *StatementFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(statement.FieldCoinTypeID))
}

// WhereIoType applies the entql string predicate on the io_type field.
func (f *StatementFilter) WhereIoType(p entql.StringP) {
	f.Where(p.Field(statement.FieldIoType))
}

// WhereIoSubType applies the entql string predicate on the io_sub_type field.
func (f *StatementFilter) WhereIoSubType(p entql.StringP) {
	f.Where(p.Field(statement.FieldIoSubType))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *StatementFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(statement.FieldAmount))
}

// WhereIoExtra applies the entql string predicate on the io_extra field.
func (f *StatementFilter) WhereIoExtra(p entql.StringP) {
	f.Where(p.Field(statement.FieldIoExtra))
}

// WhereIoExtraV1 applies the entql string predicate on the io_extra_v1 field.
func (f *StatementFilter) WhereIoExtraV1(p entql.StringP) {
	f.Where(p.Field(statement.FieldIoExtraV1))
}

// addPredicate implements the predicateAdder interface.
func (usq *UnsoldStatementQuery) addPredicate(pred func(s *sql.Selector)) {
	usq.predicates = append(usq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UnsoldStatementQuery builder.
func (usq *UnsoldStatementQuery) Filter() *UnsoldStatementFilter {
	return &UnsoldStatementFilter{config: usq.config, predicateAdder: usq}
}

// addPredicate implements the predicateAdder interface.
func (m *UnsoldStatementMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UnsoldStatementMutation builder.
func (m *UnsoldStatementMutation) Filter() *UnsoldStatementFilter {
	return &UnsoldStatementFilter{config: m.config, predicateAdder: m}
}

// UnsoldStatementFilter provides a generic filtering capability at runtime for UnsoldStatementQuery.
type UnsoldStatementFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UnsoldStatementFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *UnsoldStatementFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(unsoldstatement.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *UnsoldStatementFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(unsoldstatement.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *UnsoldStatementFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(unsoldstatement.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *UnsoldStatementFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(unsoldstatement.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *UnsoldStatementFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(unsoldstatement.FieldEntID))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *UnsoldStatementFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(unsoldstatement.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *UnsoldStatementFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(unsoldstatement.FieldCoinTypeID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *UnsoldStatementFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(unsoldstatement.FieldAmount))
}

// WhereBenefitDate applies the entql uint32 predicate on the benefit_date field.
func (f *UnsoldStatementFilter) WhereBenefitDate(p entql.Uint32P) {
	f.Where(p.Field(unsoldstatement.FieldBenefitDate))
}

// WhereStatementID applies the entql [16]byte predicate on the statement_id field.
func (f *UnsoldStatementFilter) WhereStatementID(p entql.ValueP) {
	f.Where(p.Field(unsoldstatement.FieldStatementID))
}

// addPredicate implements the predicateAdder interface.
func (wq *WithdrawQuery) addPredicate(pred func(s *sql.Selector)) {
	wq.predicates = append(wq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the WithdrawQuery builder.
func (wq *WithdrawQuery) Filter() *WithdrawFilter {
	return &WithdrawFilter{config: wq.config, predicateAdder: wq}
}

// addPredicate implements the predicateAdder interface.
func (m *WithdrawMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the WithdrawMutation builder.
func (m *WithdrawMutation) Filter() *WithdrawFilter {
	return &WithdrawFilter{config: m.config, predicateAdder: m}
}

// WithdrawFilter provides a generic filtering capability at runtime for WithdrawQuery.
type WithdrawFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *WithdrawFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[7].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *WithdrawFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(withdraw.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *WithdrawFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(withdraw.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *WithdrawFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(withdraw.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *WithdrawFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(withdraw.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *WithdrawFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *WithdrawFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *WithdrawFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldUserID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *WithdrawFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldCoinTypeID))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *WithdrawFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldAccountID))
}

// WhereAddress applies the entql string predicate on the address field.
func (f *WithdrawFilter) WhereAddress(p entql.StringP) {
	f.Where(p.Field(withdraw.FieldAddress))
}

// WherePlatformTransactionID applies the entql [16]byte predicate on the platform_transaction_id field.
func (f *WithdrawFilter) WherePlatformTransactionID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldPlatformTransactionID))
}

// WhereChainTransactionID applies the entql string predicate on the chain_transaction_id field.
func (f *WithdrawFilter) WhereChainTransactionID(p entql.StringP) {
	f.Where(p.Field(withdraw.FieldChainTransactionID))
}

// WhereState applies the entql string predicate on the state field.
func (f *WithdrawFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(withdraw.FieldState))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *WithdrawFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(withdraw.FieldAmount))
}

// WhereReviewID applies the entql [16]byte predicate on the review_id field.
func (f *WithdrawFilter) WhereReviewID(p entql.ValueP) {
	f.Where(p.Field(withdraw.FieldReviewID))
}
