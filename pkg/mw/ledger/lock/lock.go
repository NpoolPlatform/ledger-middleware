package lock

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*Handler
	info *ledgerpb.Ledger
}

func (h *Handler) validate() error {
	if h.Spendable != nil && h.Locked != nil {
		return fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable == nil && h.Locked == nil {
		return fmt.Errorf("spendable or locked needed")
	}
	return nil
}

func (h *lockHandler) setConds() *statementcrud.Conds {
	conds := &statementcrud.Conds{}
	if h.AppID != nil {
		conds.AppID = &cruder.Cond{Op: cruder.EQ, Val: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &cruder.Cond{Op: cruder.EQ, Val: *h.UserID}
	}
	if h.CoinTypeID != nil {
		conds.CoinTypeID = &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID}
	}
	if h.IOSubType != nil {
		conds.IOSubType = &cruder.Cond{Op: cruder.EQ, Val: *h.IOSubType}
	}
	if h.IOExtra != nil {
		conds.IOExtra = &cruder.Cond{Op: cruder.LIKE, Val: *h.IOExtra}
	}
	ioType := basetypes.IOType_Outcoming
	conds.IOType = &cruder.Cond{Op: cruder.EQ, Val: ioType}
	return conds
}

func (h *lockHandler) tryCreateStatement(req *statementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithChangeLedger(),
	)
	if err != nil {
		return err
	}

	handler.Req = *req
	if _, err := handler.CreateStatement(ctx); err != nil {
		return err
	}

	return nil
}

func (h *lockHandler) tryGetStatement(ctx context.Context, tx *ent.Tx) (*ent.Statement, error) {
	if h.IOSubType == nil {
		return nil, fmt.Errorf("invalid io sub type")
	}
	if h.IOExtra == nil {
		return nil, fmt.Errorf("invalid io extra")
	}

	stm, err := statementcrud.SetQueryConds(
		tx.Statement.Query(),
		h.setConds(),
	)
	if err != nil {
		return nil, err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return info, nil
}

func (h *lockHandler) tryUpdateLedger(req ledgercrud.Req, ctx context.Context, tx *ent.Tx) (*ledgerpb.Ledger, error) {
	stm, err := ledgercrud.SetQueryConds(tx.Ledger.Query(), &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
	})
	if err != nil {
		return nil, err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("ledger not exist, AppID: %v, UserID: %v, CoinTypeID: %v", *req.AppID, *req.UserID, *req.CoinTypeID)
		}
		return nil, err
	}

	// update
	old, err := tx.Ledger.Get(ctx, info.ID)
	if err != nil {
		return nil, err
	}
	if old == nil {
		return nil, fmt.Errorf("ledger not exist, id %v", info.ID)
	}

	stm1, err := ledgercrud.UpdateSet(
		old,
		tx.Ledger.UpdateOneID(info.ID),
		&ledgercrud.Req{
			Outcoming: req.Outcoming,
			Spendable: req.Spendable,
			Locked:    req.Locked,
		},
	)
	if err != nil {
		return nil, err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return nil, err
	}

	ledgerID := old.ID.String()
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithID(&ledgerID),
	)
	if err != nil {
		return nil, err
	}

	h.info, err = handler.GetLedger(ctx)
	if err != nil {
		return nil, err
	}

	return h.info, nil
}

func (h *lockHandler) tryLock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}
	if h.Spendable.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("spendable less than equal 0")
	}

	spendable := decimal.NewFromInt(0).Sub(*h.Spendable)
	locked := *h.Spendable

	if _, err := h.tryUpdateLedger(ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     &locked,
		Spendable:  &spendable,
	}, ctx, tx); err != nil {
		return err
	}
	return nil
}

func (h *lockHandler) trySpend(ctx context.Context, tx *ent.Tx) error {
	if h.Locked == nil {
		return nil
	}
	if h.Locked.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("locked less than equal 0")
	}

	info, err := h.tryGetStatement(ctx, tx)
	if err != nil {
		return err
	}
	if info != nil {
		return fmt.Errorf("statement already exist")
	}

	ioType := basetypes.IOType_Outcoming
	if err := h.tryCreateStatement(&statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
		Amount:     h.Locked,
	}, ctx, tx); err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(*h.Locked)
	outcoming := *h.Locked

	if _, err = h.tryUpdateLedger(ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     &locked,
		Outcoming:  &outcoming,
	}, ctx, tx); err != nil {
		return err
	}
	return nil
}

// Lock & Spend
func (h *Handler) SubBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	if err := h.validate(); err != nil {
		return nil, err
	}

	handler := &lockHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.tryLock(ctx, tx); err != nil {
			return err
		}
		if err := handler.trySpend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return handler.info, err
}

func (h *lockHandler) tryUnlock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}
	if h.Spendable.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("spendable less than equal 0")
	}

	spendable := *h.Spendable
	locked := decimal.NewFromInt(0).Sub(*h.Spendable)

	if _, err := h.tryUpdateLedger(ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     &locked,
		Spendable:  &spendable,
	}, ctx, tx); err != nil {
		return err
	}
	return nil
}

func (h *lockHandler) tryUnspend(ctx context.Context, tx *ent.Tx) error {
	if h.Locked == nil {
		return nil
	}
	if h.Locked.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("locked less than equal 0")
	}

	info, err := h.tryGetStatement(ctx, tx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	// rollback
	ioType := basetypes.IOType_Outcoming
	ioExtra := fmt.Sprintf(`{"GeneralID": "%v", "Rollback": "true"}`, info.ID.String())
	if err := h.tryCreateStatement(&statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    &ioExtra,
		Amount:     h.Locked,
	}, ctx, tx); err != nil {
		return err
	}

	locked := *h.Locked
	outcoming := decimal.NewFromInt(0).Sub(*h.Locked)

	if _, err = h.tryUpdateLedger(ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     &locked,
		Outcoming:  &outcoming,
	}, ctx, tx); err != nil {
		return err
	}
	return nil
}

// Unlock & Unspend
func (h *Handler) AddBalance(ctx context.Context) (*ledgerpb.Ledger, error) {
	if err := h.validate(); err != nil {
		return nil, err
	}

	handler := &lockHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.tryUnlock(ctx, tx); err != nil {
			return err
		}
		if err := handler.tryUnspend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return handler.info, err
}
