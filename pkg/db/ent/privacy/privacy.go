// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The GoodLedgerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodLedgerQueryRuleFunc func(context.Context, *ent.GoodLedgerQuery) error

// EvalQuery return f(ctx, q).
func (f GoodLedgerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodLedgerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodLedgerQuery", q)
}

// The GoodLedgerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodLedgerMutationRuleFunc func(context.Context, *ent.GoodLedgerMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodLedgerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodLedgerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodLedgerMutation", m)
}

// The GoodStatementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodStatementQueryRuleFunc func(context.Context, *ent.GoodStatementQuery) error

// EvalQuery return f(ctx, q).
func (f GoodStatementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodStatementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodStatementQuery", q)
}

// The GoodStatementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodStatementMutationRuleFunc func(context.Context, *ent.GoodStatementMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodStatementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodStatementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodStatementMutation", m)
}

// The LedgerQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LedgerQueryRuleFunc func(context.Context, *ent.LedgerQuery) error

// EvalQuery return f(ctx, q).
func (f LedgerQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LedgerQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LedgerQuery", q)
}

// The LedgerMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LedgerMutationRuleFunc func(context.Context, *ent.LedgerMutation) error

// EvalMutation calls f(ctx, m).
func (f LedgerMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LedgerMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LedgerMutation", m)
}

// The LedgerLockQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LedgerLockQueryRuleFunc func(context.Context, *ent.LedgerLockQuery) error

// EvalQuery return f(ctx, q).
func (f LedgerLockQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LedgerLockQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LedgerLockQuery", q)
}

// The LedgerLockMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LedgerLockMutationRuleFunc func(context.Context, *ent.LedgerLockMutation) error

// EvalMutation calls f(ctx, m).
func (f LedgerLockMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LedgerLockMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LedgerLockMutation", m)
}

// The ProfitQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProfitQueryRuleFunc func(context.Context, *ent.ProfitQuery) error

// EvalQuery return f(ctx, q).
func (f ProfitQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProfitQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProfitQuery", q)
}

// The ProfitMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProfitMutationRuleFunc func(context.Context, *ent.ProfitMutation) error

// EvalMutation calls f(ctx, m).
func (f ProfitMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProfitMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProfitMutation", m)
}

// The StatementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StatementQueryRuleFunc func(context.Context, *ent.StatementQuery) error

// EvalQuery return f(ctx, q).
func (f StatementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StatementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StatementQuery", q)
}

// The StatementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StatementMutationRuleFunc func(context.Context, *ent.StatementMutation) error

// EvalMutation calls f(ctx, m).
func (f StatementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StatementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StatementMutation", m)
}

// The UnsoldStatementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UnsoldStatementQueryRuleFunc func(context.Context, *ent.UnsoldStatementQuery) error

// EvalQuery return f(ctx, q).
func (f UnsoldStatementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UnsoldStatementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UnsoldStatementQuery", q)
}

// The UnsoldStatementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UnsoldStatementMutationRuleFunc func(context.Context, *ent.UnsoldStatementMutation) error

// EvalMutation calls f(ctx, m).
func (f UnsoldStatementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UnsoldStatementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UnsoldStatementMutation", m)
}

// The WithdrawQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type WithdrawQueryRuleFunc func(context.Context, *ent.WithdrawQuery) error

// EvalQuery return f(ctx, q).
func (f WithdrawQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.WithdrawQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.WithdrawQuery", q)
}

// The WithdrawMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type WithdrawMutationRuleFunc func(context.Context, *ent.WithdrawMutation) error

// EvalMutation calls f(ctx, m).
func (f WithdrawMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.WithdrawMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.WithdrawMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.GoodLedgerQuery:
		return q.Filter(), nil
	case *ent.GoodStatementQuery:
		return q.Filter(), nil
	case *ent.LedgerQuery:
		return q.Filter(), nil
	case *ent.LedgerLockQuery:
		return q.Filter(), nil
	case *ent.ProfitQuery:
		return q.Filter(), nil
	case *ent.StatementQuery:
		return q.Filter(), nil
	case *ent.UnsoldStatementQuery:
		return q.Filter(), nil
	case *ent.WithdrawQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.GoodLedgerMutation:
		return m.Filter(), nil
	case *ent.GoodStatementMutation:
		return m.Filter(), nil
	case *ent.LedgerMutation:
		return m.Filter(), nil
	case *ent.LedgerLockMutation:
		return m.Filter(), nil
	case *ent.ProfitMutation:
		return m.Filter(), nil
	case *ent.StatementMutation:
		return m.Filter(), nil
	case *ent.UnsoldStatementMutation:
		return m.Filter(), nil
	case *ent.WithdrawMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
