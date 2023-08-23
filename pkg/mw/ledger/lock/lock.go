package lock

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ledgerpb "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*Handler
	info     *ledgermwpb.Ledger
	ledger   ent.Ledger
	rollback ent.Statement
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

func (h *lockHandler) tryGetStatement(req *statementcrud.Req, ctx context.Context, tx *ent.Tx) (*ent.Statement, error) {
	if req.IOSubType == nil {
		return nil, fmt.Errorf("invalid io sub type")
	}
	if req.IOExtra == nil {
		return nil, fmt.Errorf("invalid io extra")
	}

	conds := &statementcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		IOType:     &cruder.Cond{Op: cruder.EQ, Val: *req.IOType},
		IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *req.IOSubType},
		IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *req.IOExtra},
	}

	stm, err := statementcrud.SetQueryConds(
		tx.Statement.Query(),
		conds,
	)
	if err != nil {
		return nil, err
	}

	info, err := stm.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return info, nil
}

func (h *lockHandler) tryGetRolledBackStatement(origin *ent.Statement, ctx context.Context, tx *ent.Tx) (*ent.Statement, error) {
	ioType := ledgerpb.IOType_Incoming
	ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, origin.ID.String())
	info, err := h.tryGetStatement(&statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    &ioExtra,
	}, ctx, tx)
	if err != nil {
		return nil, err
	}
	h.rollback = *info
	return &h.rollback, nil
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

	stm, err := ledgercrud.UpdateSet(
		&h.ledger,
		tx.Ledger.UpdateOneID(h.ledger.ID),
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Spendable:  &spendable,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
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

	ioType := ledgerpb.IOType_Outcoming
	info, err := h.tryGetStatement(&statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	}, ctx, tx)
	if err != nil {
		return err
	}
	if info != nil {
		// try get rolled back statement
		info, err := h.tryGetRolledBackStatement(info, ctx, tx)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("statement already exist")
		}
	}

	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithChangeLedger(),
	)
	if err != nil {
		return err
	}

	handler.Req = statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
		Amount:     h.Locked,
	}
	if _, err := handler.CreateStatement(ctx); err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(*h.Locked)
	outcoming := *h.Locked

	stm, err := ledgercrud.UpdateSet(
		&h.ledger,
		tx.Ledger.UpdateOneID(h.ledger.ID),
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Outcoming:  &outcoming,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

// Lock & Spend
func (h *Handler) SubBalance(ctx context.Context) (info *ledgermwpb.Ledger, err error) {
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

	stm, err := ledgercrud.UpdateSet(
		&h.ledger,
		tx.Ledger.UpdateOneID(h.ledger.ID),
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Spendable:  &spendable,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
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

	ioType := ledgerpb.IOType_Outcoming
	info, err := h.tryGetStatement(&statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	}, ctx, tx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	// whether have been rolled back
	rolled, err := h.tryGetRolledBackStatement(info, ctx, tx)
	if err != nil {
		return err
	}
	if rolled != nil {
		return fmt.Errorf("rollback statement already exist")
	}

	// rollback
	ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, info.ID.String())
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

	stm, err := ledgercrud.UpdateSet(
		&h.ledger,
		tx.Ledger.UpdateOneID(h.ledger.ID),
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Outcoming:  &outcoming,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

// Unlock & Unspend
func (h *Handler) AddBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
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

func (h *lockHandler) getLedger(ctx context.Context, tx *ent.Tx) error {
	stm, err := ledgercrud.SetQueryConds(
		tx.Ledger.Query(),
		&ledgercrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		})
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("ledger not exist, AppID: %v, UserID: %v, CoinTypeID: %v", *h.AppID, *h.UserID, *h.CoinTypeID)
		}
		return err
	}
	h.ledger = *info
	return nil
}
