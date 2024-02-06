// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/couponwithdraw"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledgerlock"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/profit"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/schema"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	couponwithdrawMixin := schema.CouponWithdraw{}.Mixin()
	couponwithdraw.Policy = privacy.NewPolicies(couponwithdrawMixin[0], schema.CouponWithdraw{})
	couponwithdraw.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := couponwithdraw.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	couponwithdrawMixinFields0 := couponwithdrawMixin[0].Fields()
	_ = couponwithdrawMixinFields0
	couponwithdrawMixinFields1 := couponwithdrawMixin[1].Fields()
	_ = couponwithdrawMixinFields1
	couponwithdrawFields := schema.CouponWithdraw{}.Fields()
	_ = couponwithdrawFields
	// couponwithdrawDescCreatedAt is the schema descriptor for created_at field.
	couponwithdrawDescCreatedAt := couponwithdrawMixinFields0[0].Descriptor()
	// couponwithdraw.DefaultCreatedAt holds the default value on creation for the created_at field.
	couponwithdraw.DefaultCreatedAt = couponwithdrawDescCreatedAt.Default.(func() uint32)
	// couponwithdrawDescUpdatedAt is the schema descriptor for updated_at field.
	couponwithdrawDescUpdatedAt := couponwithdrawMixinFields0[1].Descriptor()
	// couponwithdraw.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	couponwithdraw.DefaultUpdatedAt = couponwithdrawDescUpdatedAt.Default.(func() uint32)
	// couponwithdraw.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	couponwithdraw.UpdateDefaultUpdatedAt = couponwithdrawDescUpdatedAt.UpdateDefault.(func() uint32)
	// couponwithdrawDescDeletedAt is the schema descriptor for deleted_at field.
	couponwithdrawDescDeletedAt := couponwithdrawMixinFields0[2].Descriptor()
	// couponwithdraw.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	couponwithdraw.DefaultDeletedAt = couponwithdrawDescDeletedAt.Default.(func() uint32)
	// couponwithdrawDescEntID is the schema descriptor for ent_id field.
	couponwithdrawDescEntID := couponwithdrawMixinFields1[1].Descriptor()
	// couponwithdraw.DefaultEntID holds the default value on creation for the ent_id field.
	couponwithdraw.DefaultEntID = couponwithdrawDescEntID.Default.(func() uuid.UUID)
	// couponwithdrawDescAppID is the schema descriptor for app_id field.
	couponwithdrawDescAppID := couponwithdrawFields[0].Descriptor()
	// couponwithdraw.DefaultAppID holds the default value on creation for the app_id field.
	couponwithdraw.DefaultAppID = couponwithdrawDescAppID.Default.(func() uuid.UUID)
	// couponwithdrawDescUserID is the schema descriptor for user_id field.
	couponwithdrawDescUserID := couponwithdrawFields[1].Descriptor()
	// couponwithdraw.DefaultUserID holds the default value on creation for the user_id field.
	couponwithdraw.DefaultUserID = couponwithdrawDescUserID.Default.(func() uuid.UUID)
	// couponwithdrawDescCoinTypeID is the schema descriptor for coin_type_id field.
	couponwithdrawDescCoinTypeID := couponwithdrawFields[2].Descriptor()
	// couponwithdraw.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	couponwithdraw.DefaultCoinTypeID = couponwithdrawDescCoinTypeID.Default.(func() uuid.UUID)
	// couponwithdrawDescAllocatedID is the schema descriptor for allocated_id field.
	couponwithdrawDescAllocatedID := couponwithdrawFields[3].Descriptor()
	// couponwithdraw.DefaultAllocatedID holds the default value on creation for the allocated_id field.
	couponwithdraw.DefaultAllocatedID = couponwithdrawDescAllocatedID.Default.(func() uuid.UUID)
	// couponwithdrawDescState is the schema descriptor for state field.
	couponwithdrawDescState := couponwithdrawFields[4].Descriptor()
	// couponwithdraw.DefaultState holds the default value on creation for the state field.
	couponwithdraw.DefaultState = couponwithdrawDescState.Default.(string)
	// couponwithdrawDescReviewID is the schema descriptor for review_id field.
	couponwithdrawDescReviewID := couponwithdrawFields[6].Descriptor()
	// couponwithdraw.DefaultReviewID holds the default value on creation for the review_id field.
	couponwithdraw.DefaultReviewID = couponwithdrawDescReviewID.Default.(func() uuid.UUID)
	goodledgerMixin := schema.GoodLedger{}.Mixin()
	goodledger.Policy = privacy.NewPolicies(goodledgerMixin[0], schema.GoodLedger{})
	goodledger.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := goodledger.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	goodledgerMixinFields0 := goodledgerMixin[0].Fields()
	_ = goodledgerMixinFields0
	goodledgerMixinFields1 := goodledgerMixin[1].Fields()
	_ = goodledgerMixinFields1
	goodledgerFields := schema.GoodLedger{}.Fields()
	_ = goodledgerFields
	// goodledgerDescCreatedAt is the schema descriptor for created_at field.
	goodledgerDescCreatedAt := goodledgerMixinFields0[0].Descriptor()
	// goodledger.DefaultCreatedAt holds the default value on creation for the created_at field.
	goodledger.DefaultCreatedAt = goodledgerDescCreatedAt.Default.(func() uint32)
	// goodledgerDescUpdatedAt is the schema descriptor for updated_at field.
	goodledgerDescUpdatedAt := goodledgerMixinFields0[1].Descriptor()
	// goodledger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	goodledger.DefaultUpdatedAt = goodledgerDescUpdatedAt.Default.(func() uint32)
	// goodledger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	goodledger.UpdateDefaultUpdatedAt = goodledgerDescUpdatedAt.UpdateDefault.(func() uint32)
	// goodledgerDescDeletedAt is the schema descriptor for deleted_at field.
	goodledgerDescDeletedAt := goodledgerMixinFields0[2].Descriptor()
	// goodledger.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	goodledger.DefaultDeletedAt = goodledgerDescDeletedAt.Default.(func() uint32)
	// goodledgerDescEntID is the schema descriptor for ent_id field.
	goodledgerDescEntID := goodledgerMixinFields1[1].Descriptor()
	// goodledger.DefaultEntID holds the default value on creation for the ent_id field.
	goodledger.DefaultEntID = goodledgerDescEntID.Default.(func() uuid.UUID)
	// goodledgerDescGoodID is the schema descriptor for good_id field.
	goodledgerDescGoodID := goodledgerFields[0].Descriptor()
	// goodledger.DefaultGoodID holds the default value on creation for the good_id field.
	goodledger.DefaultGoodID = goodledgerDescGoodID.Default.(func() uuid.UUID)
	// goodledgerDescCoinTypeID is the schema descriptor for coin_type_id field.
	goodledgerDescCoinTypeID := goodledgerFields[1].Descriptor()
	// goodledger.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	goodledger.DefaultCoinTypeID = goodledgerDescCoinTypeID.Default.(func() uuid.UUID)
	goodstatementMixin := schema.GoodStatement{}.Mixin()
	goodstatement.Policy = privacy.NewPolicies(goodstatementMixin[0], schema.GoodStatement{})
	goodstatement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := goodstatement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	goodstatementMixinFields0 := goodstatementMixin[0].Fields()
	_ = goodstatementMixinFields0
	goodstatementMixinFields1 := goodstatementMixin[1].Fields()
	_ = goodstatementMixinFields1
	goodstatementFields := schema.GoodStatement{}.Fields()
	_ = goodstatementFields
	// goodstatementDescCreatedAt is the schema descriptor for created_at field.
	goodstatementDescCreatedAt := goodstatementMixinFields0[0].Descriptor()
	// goodstatement.DefaultCreatedAt holds the default value on creation for the created_at field.
	goodstatement.DefaultCreatedAt = goodstatementDescCreatedAt.Default.(func() uint32)
	// goodstatementDescUpdatedAt is the schema descriptor for updated_at field.
	goodstatementDescUpdatedAt := goodstatementMixinFields0[1].Descriptor()
	// goodstatement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	goodstatement.DefaultUpdatedAt = goodstatementDescUpdatedAt.Default.(func() uint32)
	// goodstatement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	goodstatement.UpdateDefaultUpdatedAt = goodstatementDescUpdatedAt.UpdateDefault.(func() uint32)
	// goodstatementDescDeletedAt is the schema descriptor for deleted_at field.
	goodstatementDescDeletedAt := goodstatementMixinFields0[2].Descriptor()
	// goodstatement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	goodstatement.DefaultDeletedAt = goodstatementDescDeletedAt.Default.(func() uint32)
	// goodstatementDescEntID is the schema descriptor for ent_id field.
	goodstatementDescEntID := goodstatementMixinFields1[1].Descriptor()
	// goodstatement.DefaultEntID holds the default value on creation for the ent_id field.
	goodstatement.DefaultEntID = goodstatementDescEntID.Default.(func() uuid.UUID)
	// goodstatementDescGoodID is the schema descriptor for good_id field.
	goodstatementDescGoodID := goodstatementFields[0].Descriptor()
	// goodstatement.DefaultGoodID holds the default value on creation for the good_id field.
	goodstatement.DefaultGoodID = goodstatementDescGoodID.Default.(func() uuid.UUID)
	// goodstatementDescCoinTypeID is the schema descriptor for coin_type_id field.
	goodstatementDescCoinTypeID := goodstatementFields[1].Descriptor()
	// goodstatement.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	goodstatement.DefaultCoinTypeID = goodstatementDescCoinTypeID.Default.(func() uuid.UUID)
	// goodstatementDescBenefitDate is the schema descriptor for benefit_date field.
	goodstatementDescBenefitDate := goodstatementFields[6].Descriptor()
	// goodstatement.DefaultBenefitDate holds the default value on creation for the benefit_date field.
	goodstatement.DefaultBenefitDate = goodstatementDescBenefitDate.Default.(uint32)
	ledgerMixin := schema.Ledger{}.Mixin()
	ledger.Policy = privacy.NewPolicies(ledgerMixin[0], schema.Ledger{})
	ledger.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := ledger.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	ledgerMixinFields0 := ledgerMixin[0].Fields()
	_ = ledgerMixinFields0
	ledgerMixinFields1 := ledgerMixin[1].Fields()
	_ = ledgerMixinFields1
	ledgerFields := schema.Ledger{}.Fields()
	_ = ledgerFields
	// ledgerDescCreatedAt is the schema descriptor for created_at field.
	ledgerDescCreatedAt := ledgerMixinFields0[0].Descriptor()
	// ledger.DefaultCreatedAt holds the default value on creation for the created_at field.
	ledger.DefaultCreatedAt = ledgerDescCreatedAt.Default.(func() uint32)
	// ledgerDescUpdatedAt is the schema descriptor for updated_at field.
	ledgerDescUpdatedAt := ledgerMixinFields0[1].Descriptor()
	// ledger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ledger.DefaultUpdatedAt = ledgerDescUpdatedAt.Default.(func() uint32)
	// ledger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ledger.UpdateDefaultUpdatedAt = ledgerDescUpdatedAt.UpdateDefault.(func() uint32)
	// ledgerDescDeletedAt is the schema descriptor for deleted_at field.
	ledgerDescDeletedAt := ledgerMixinFields0[2].Descriptor()
	// ledger.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	ledger.DefaultDeletedAt = ledgerDescDeletedAt.Default.(func() uint32)
	// ledgerDescEntID is the schema descriptor for ent_id field.
	ledgerDescEntID := ledgerMixinFields1[1].Descriptor()
	// ledger.DefaultEntID holds the default value on creation for the ent_id field.
	ledger.DefaultEntID = ledgerDescEntID.Default.(func() uuid.UUID)
	// ledgerDescAppID is the schema descriptor for app_id field.
	ledgerDescAppID := ledgerFields[0].Descriptor()
	// ledger.DefaultAppID holds the default value on creation for the app_id field.
	ledger.DefaultAppID = ledgerDescAppID.Default.(func() uuid.UUID)
	// ledgerDescUserID is the schema descriptor for user_id field.
	ledgerDescUserID := ledgerFields[1].Descriptor()
	// ledger.DefaultUserID holds the default value on creation for the user_id field.
	ledger.DefaultUserID = ledgerDescUserID.Default.(func() uuid.UUID)
	// ledgerDescCoinTypeID is the schema descriptor for coin_type_id field.
	ledgerDescCoinTypeID := ledgerFields[2].Descriptor()
	// ledger.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	ledger.DefaultCoinTypeID = ledgerDescCoinTypeID.Default.(func() uuid.UUID)
	ledgerlockMixin := schema.LedgerLock{}.Mixin()
	ledgerlock.Policy = privacy.NewPolicies(ledgerlockMixin[0], schema.LedgerLock{})
	ledgerlock.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := ledgerlock.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	ledgerlockMixinFields0 := ledgerlockMixin[0].Fields()
	_ = ledgerlockMixinFields0
	ledgerlockMixinFields1 := ledgerlockMixin[1].Fields()
	_ = ledgerlockMixinFields1
	ledgerlockFields := schema.LedgerLock{}.Fields()
	_ = ledgerlockFields
	// ledgerlockDescCreatedAt is the schema descriptor for created_at field.
	ledgerlockDescCreatedAt := ledgerlockMixinFields0[0].Descriptor()
	// ledgerlock.DefaultCreatedAt holds the default value on creation for the created_at field.
	ledgerlock.DefaultCreatedAt = ledgerlockDescCreatedAt.Default.(func() uint32)
	// ledgerlockDescUpdatedAt is the schema descriptor for updated_at field.
	ledgerlockDescUpdatedAt := ledgerlockMixinFields0[1].Descriptor()
	// ledgerlock.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ledgerlock.DefaultUpdatedAt = ledgerlockDescUpdatedAt.Default.(func() uint32)
	// ledgerlock.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ledgerlock.UpdateDefaultUpdatedAt = ledgerlockDescUpdatedAt.UpdateDefault.(func() uint32)
	// ledgerlockDescDeletedAt is the schema descriptor for deleted_at field.
	ledgerlockDescDeletedAt := ledgerlockMixinFields0[2].Descriptor()
	// ledgerlock.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	ledgerlock.DefaultDeletedAt = ledgerlockDescDeletedAt.Default.(func() uint32)
	// ledgerlockDescEntID is the schema descriptor for ent_id field.
	ledgerlockDescEntID := ledgerlockMixinFields1[1].Descriptor()
	// ledgerlock.DefaultEntID holds the default value on creation for the ent_id field.
	ledgerlock.DefaultEntID = ledgerlockDescEntID.Default.(func() uuid.UUID)
	// ledgerlockDescLedgerID is the schema descriptor for ledger_id field.
	ledgerlockDescLedgerID := ledgerlockFields[0].Descriptor()
	// ledgerlock.DefaultLedgerID holds the default value on creation for the ledger_id field.
	ledgerlock.DefaultLedgerID = ledgerlockDescLedgerID.Default.(func() uuid.UUID)
	// ledgerlockDescStatementID is the schema descriptor for statement_id field.
	ledgerlockDescStatementID := ledgerlockFields[1].Descriptor()
	// ledgerlock.DefaultStatementID holds the default value on creation for the statement_id field.
	ledgerlock.DefaultStatementID = ledgerlockDescStatementID.Default.(func() uuid.UUID)
	// ledgerlockDescLockState is the schema descriptor for lock_state field.
	ledgerlockDescLockState := ledgerlockFields[3].Descriptor()
	// ledgerlock.DefaultLockState holds the default value on creation for the lock_state field.
	ledgerlock.DefaultLockState = ledgerlockDescLockState.Default.(string)
	// ledgerlockDescExLockID is the schema descriptor for ex_lock_id field.
	ledgerlockDescExLockID := ledgerlockFields[4].Descriptor()
	// ledgerlock.DefaultExLockID holds the default value on creation for the ex_lock_id field.
	ledgerlock.DefaultExLockID = ledgerlockDescExLockID.Default.(func() uuid.UUID)
	profitMixin := schema.Profit{}.Mixin()
	profit.Policy = privacy.NewPolicies(profitMixin[0], schema.Profit{})
	profit.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := profit.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	profitMixinFields0 := profitMixin[0].Fields()
	_ = profitMixinFields0
	profitMixinFields1 := profitMixin[1].Fields()
	_ = profitMixinFields1
	profitFields := schema.Profit{}.Fields()
	_ = profitFields
	// profitDescCreatedAt is the schema descriptor for created_at field.
	profitDescCreatedAt := profitMixinFields0[0].Descriptor()
	// profit.DefaultCreatedAt holds the default value on creation for the created_at field.
	profit.DefaultCreatedAt = profitDescCreatedAt.Default.(func() uint32)
	// profitDescUpdatedAt is the schema descriptor for updated_at field.
	profitDescUpdatedAt := profitMixinFields0[1].Descriptor()
	// profit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profit.DefaultUpdatedAt = profitDescUpdatedAt.Default.(func() uint32)
	// profit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	profit.UpdateDefaultUpdatedAt = profitDescUpdatedAt.UpdateDefault.(func() uint32)
	// profitDescDeletedAt is the schema descriptor for deleted_at field.
	profitDescDeletedAt := profitMixinFields0[2].Descriptor()
	// profit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	profit.DefaultDeletedAt = profitDescDeletedAt.Default.(func() uint32)
	// profitDescEntID is the schema descriptor for ent_id field.
	profitDescEntID := profitMixinFields1[1].Descriptor()
	// profit.DefaultEntID holds the default value on creation for the ent_id field.
	profit.DefaultEntID = profitDescEntID.Default.(func() uuid.UUID)
	// profitDescAppID is the schema descriptor for app_id field.
	profitDescAppID := profitFields[0].Descriptor()
	// profit.DefaultAppID holds the default value on creation for the app_id field.
	profit.DefaultAppID = profitDescAppID.Default.(func() uuid.UUID)
	// profitDescUserID is the schema descriptor for user_id field.
	profitDescUserID := profitFields[1].Descriptor()
	// profit.DefaultUserID holds the default value on creation for the user_id field.
	profit.DefaultUserID = profitDescUserID.Default.(func() uuid.UUID)
	// profitDescCoinTypeID is the schema descriptor for coin_type_id field.
	profitDescCoinTypeID := profitFields[2].Descriptor()
	// profit.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	profit.DefaultCoinTypeID = profitDescCoinTypeID.Default.(func() uuid.UUID)
	statementMixin := schema.Statement{}.Mixin()
	statement.Policy = privacy.NewPolicies(statementMixin[0], schema.Statement{})
	statement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := statement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	statementMixinFields0 := statementMixin[0].Fields()
	_ = statementMixinFields0
	statementMixinFields1 := statementMixin[1].Fields()
	_ = statementMixinFields1
	statementFields := schema.Statement{}.Fields()
	_ = statementFields
	// statementDescCreatedAt is the schema descriptor for created_at field.
	statementDescCreatedAt := statementMixinFields0[0].Descriptor()
	// statement.DefaultCreatedAt holds the default value on creation for the created_at field.
	statement.DefaultCreatedAt = statementDescCreatedAt.Default.(func() uint32)
	// statementDescUpdatedAt is the schema descriptor for updated_at field.
	statementDescUpdatedAt := statementMixinFields0[1].Descriptor()
	// statement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	statement.DefaultUpdatedAt = statementDescUpdatedAt.Default.(func() uint32)
	// statement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	statement.UpdateDefaultUpdatedAt = statementDescUpdatedAt.UpdateDefault.(func() uint32)
	// statementDescDeletedAt is the schema descriptor for deleted_at field.
	statementDescDeletedAt := statementMixinFields0[2].Descriptor()
	// statement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	statement.DefaultDeletedAt = statementDescDeletedAt.Default.(func() uint32)
	// statementDescEntID is the schema descriptor for ent_id field.
	statementDescEntID := statementMixinFields1[1].Descriptor()
	// statement.DefaultEntID holds the default value on creation for the ent_id field.
	statement.DefaultEntID = statementDescEntID.Default.(func() uuid.UUID)
	// statementDescAppID is the schema descriptor for app_id field.
	statementDescAppID := statementFields[0].Descriptor()
	// statement.DefaultAppID holds the default value on creation for the app_id field.
	statement.DefaultAppID = statementDescAppID.Default.(func() uuid.UUID)
	// statementDescUserID is the schema descriptor for user_id field.
	statementDescUserID := statementFields[1].Descriptor()
	// statement.DefaultUserID holds the default value on creation for the user_id field.
	statement.DefaultUserID = statementDescUserID.Default.(func() uuid.UUID)
	// statementDescCoinTypeID is the schema descriptor for coin_type_id field.
	statementDescCoinTypeID := statementFields[2].Descriptor()
	// statement.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	statement.DefaultCoinTypeID = statementDescCoinTypeID.Default.(func() uuid.UUID)
	// statementDescIoType is the schema descriptor for io_type field.
	statementDescIoType := statementFields[3].Descriptor()
	// statement.DefaultIoType holds the default value on creation for the io_type field.
	statement.DefaultIoType = statementDescIoType.Default.(string)
	// statementDescIoSubType is the schema descriptor for io_sub_type field.
	statementDescIoSubType := statementFields[4].Descriptor()
	// statement.DefaultIoSubType holds the default value on creation for the io_sub_type field.
	statement.DefaultIoSubType = statementDescIoSubType.Default.(string)
	// statementDescIoExtra is the schema descriptor for io_extra field.
	statementDescIoExtra := statementFields[6].Descriptor()
	// statement.DefaultIoExtra holds the default value on creation for the io_extra field.
	statement.DefaultIoExtra = statementDescIoExtra.Default.(string)
	// statement.IoExtraValidator is a validator for the "io_extra" field. It is called by the builders before save.
	statement.IoExtraValidator = statementDescIoExtra.Validators[0].(func(string) error)
	// statementDescIoExtraV1 is the schema descriptor for io_extra_v1 field.
	statementDescIoExtraV1 := statementFields[7].Descriptor()
	// statement.DefaultIoExtraV1 holds the default value on creation for the io_extra_v1 field.
	statement.DefaultIoExtraV1 = statementDescIoExtraV1.Default.(string)
	// statement.IoExtraV1Validator is a validator for the "io_extra_v1" field. It is called by the builders before save.
	statement.IoExtraV1Validator = statementDescIoExtraV1.Validators[0].(func(string) error)
	unsoldstatementMixin := schema.UnsoldStatement{}.Mixin()
	unsoldstatement.Policy = privacy.NewPolicies(unsoldstatementMixin[0], schema.UnsoldStatement{})
	unsoldstatement.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := unsoldstatement.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	unsoldstatementMixinFields0 := unsoldstatementMixin[0].Fields()
	_ = unsoldstatementMixinFields0
	unsoldstatementMixinFields1 := unsoldstatementMixin[1].Fields()
	_ = unsoldstatementMixinFields1
	unsoldstatementFields := schema.UnsoldStatement{}.Fields()
	_ = unsoldstatementFields
	// unsoldstatementDescCreatedAt is the schema descriptor for created_at field.
	unsoldstatementDescCreatedAt := unsoldstatementMixinFields0[0].Descriptor()
	// unsoldstatement.DefaultCreatedAt holds the default value on creation for the created_at field.
	unsoldstatement.DefaultCreatedAt = unsoldstatementDescCreatedAt.Default.(func() uint32)
	// unsoldstatementDescUpdatedAt is the schema descriptor for updated_at field.
	unsoldstatementDescUpdatedAt := unsoldstatementMixinFields0[1].Descriptor()
	// unsoldstatement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	unsoldstatement.DefaultUpdatedAt = unsoldstatementDescUpdatedAt.Default.(func() uint32)
	// unsoldstatement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	unsoldstatement.UpdateDefaultUpdatedAt = unsoldstatementDescUpdatedAt.UpdateDefault.(func() uint32)
	// unsoldstatementDescDeletedAt is the schema descriptor for deleted_at field.
	unsoldstatementDescDeletedAt := unsoldstatementMixinFields0[2].Descriptor()
	// unsoldstatement.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	unsoldstatement.DefaultDeletedAt = unsoldstatementDescDeletedAt.Default.(func() uint32)
	// unsoldstatementDescEntID is the schema descriptor for ent_id field.
	unsoldstatementDescEntID := unsoldstatementMixinFields1[1].Descriptor()
	// unsoldstatement.DefaultEntID holds the default value on creation for the ent_id field.
	unsoldstatement.DefaultEntID = unsoldstatementDescEntID.Default.(func() uuid.UUID)
	// unsoldstatementDescGoodID is the schema descriptor for good_id field.
	unsoldstatementDescGoodID := unsoldstatementFields[0].Descriptor()
	// unsoldstatement.DefaultGoodID holds the default value on creation for the good_id field.
	unsoldstatement.DefaultGoodID = unsoldstatementDescGoodID.Default.(func() uuid.UUID)
	// unsoldstatementDescCoinTypeID is the schema descriptor for coin_type_id field.
	unsoldstatementDescCoinTypeID := unsoldstatementFields[1].Descriptor()
	// unsoldstatement.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	unsoldstatement.DefaultCoinTypeID = unsoldstatementDescCoinTypeID.Default.(func() uuid.UUID)
	// unsoldstatementDescBenefitDate is the schema descriptor for benefit_date field.
	unsoldstatementDescBenefitDate := unsoldstatementFields[3].Descriptor()
	// unsoldstatement.DefaultBenefitDate holds the default value on creation for the benefit_date field.
	unsoldstatement.DefaultBenefitDate = unsoldstatementDescBenefitDate.Default.(uint32)
	// unsoldstatementDescStatementID is the schema descriptor for statement_id field.
	unsoldstatementDescStatementID := unsoldstatementFields[4].Descriptor()
	// unsoldstatement.DefaultStatementID holds the default value on creation for the statement_id field.
	unsoldstatement.DefaultStatementID = unsoldstatementDescStatementID.Default.(func() uuid.UUID)
	withdrawMixin := schema.Withdraw{}.Mixin()
	withdraw.Policy = privacy.NewPolicies(withdrawMixin[0], schema.Withdraw{})
	withdraw.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := withdraw.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	withdrawMixinFields0 := withdrawMixin[0].Fields()
	_ = withdrawMixinFields0
	withdrawMixinFields1 := withdrawMixin[1].Fields()
	_ = withdrawMixinFields1
	withdrawFields := schema.Withdraw{}.Fields()
	_ = withdrawFields
	// withdrawDescCreatedAt is the schema descriptor for created_at field.
	withdrawDescCreatedAt := withdrawMixinFields0[0].Descriptor()
	// withdraw.DefaultCreatedAt holds the default value on creation for the created_at field.
	withdraw.DefaultCreatedAt = withdrawDescCreatedAt.Default.(func() uint32)
	// withdrawDescUpdatedAt is the schema descriptor for updated_at field.
	withdrawDescUpdatedAt := withdrawMixinFields0[1].Descriptor()
	// withdraw.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	withdraw.DefaultUpdatedAt = withdrawDescUpdatedAt.Default.(func() uint32)
	// withdraw.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	withdraw.UpdateDefaultUpdatedAt = withdrawDescUpdatedAt.UpdateDefault.(func() uint32)
	// withdrawDescDeletedAt is the schema descriptor for deleted_at field.
	withdrawDescDeletedAt := withdrawMixinFields0[2].Descriptor()
	// withdraw.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	withdraw.DefaultDeletedAt = withdrawDescDeletedAt.Default.(func() uint32)
	// withdrawDescEntID is the schema descriptor for ent_id field.
	withdrawDescEntID := withdrawMixinFields1[1].Descriptor()
	// withdraw.DefaultEntID holds the default value on creation for the ent_id field.
	withdraw.DefaultEntID = withdrawDescEntID.Default.(func() uuid.UUID)
	// withdrawDescAppID is the schema descriptor for app_id field.
	withdrawDescAppID := withdrawFields[0].Descriptor()
	// withdraw.DefaultAppID holds the default value on creation for the app_id field.
	withdraw.DefaultAppID = withdrawDescAppID.Default.(func() uuid.UUID)
	// withdrawDescUserID is the schema descriptor for user_id field.
	withdrawDescUserID := withdrawFields[1].Descriptor()
	// withdraw.DefaultUserID holds the default value on creation for the user_id field.
	withdraw.DefaultUserID = withdrawDescUserID.Default.(func() uuid.UUID)
	// withdrawDescCoinTypeID is the schema descriptor for coin_type_id field.
	withdrawDescCoinTypeID := withdrawFields[2].Descriptor()
	// withdraw.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	withdraw.DefaultCoinTypeID = withdrawDescCoinTypeID.Default.(func() uuid.UUID)
	// withdrawDescAccountID is the schema descriptor for account_id field.
	withdrawDescAccountID := withdrawFields[3].Descriptor()
	// withdraw.DefaultAccountID holds the default value on creation for the account_id field.
	withdraw.DefaultAccountID = withdrawDescAccountID.Default.(func() uuid.UUID)
	// withdrawDescAddress is the schema descriptor for address field.
	withdrawDescAddress := withdrawFields[4].Descriptor()
	// withdraw.DefaultAddress holds the default value on creation for the address field.
	withdraw.DefaultAddress = withdrawDescAddress.Default.(string)
	// withdrawDescPlatformTransactionID is the schema descriptor for platform_transaction_id field.
	withdrawDescPlatformTransactionID := withdrawFields[5].Descriptor()
	// withdraw.DefaultPlatformTransactionID holds the default value on creation for the platform_transaction_id field.
	withdraw.DefaultPlatformTransactionID = withdrawDescPlatformTransactionID.Default.(func() uuid.UUID)
	// withdrawDescChainTransactionID is the schema descriptor for chain_transaction_id field.
	withdrawDescChainTransactionID := withdrawFields[6].Descriptor()
	// withdraw.DefaultChainTransactionID holds the default value on creation for the chain_transaction_id field.
	withdraw.DefaultChainTransactionID = withdrawDescChainTransactionID.Default.(string)
	// withdrawDescState is the schema descriptor for state field.
	withdrawDescState := withdrawFields[7].Descriptor()
	// withdraw.DefaultState holds the default value on creation for the state field.
	withdraw.DefaultState = withdrawDescState.Default.(string)
	// withdrawDescReviewID is the schema descriptor for review_id field.
	withdrawDescReviewID := withdrawFields[9].Descriptor()
	// withdraw.DefaultReviewID holds the default value on creation for the review_id field.
	withdraw.DefaultReviewID = withdrawDescReviewID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)
