package lock

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*Handler
}

func (h *lockHandler) tryUpdateLedger(req ledgercrud.Req, ctx context.Context, tx *ent.Tx) (*ledgerpb.Ledger, error) {
	stm, err := ledgercrud.SetQueryConds(tx.Ledger.Query(), &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: req.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: req.CoinTypeID},
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

	handler := ledger1.Handler{
		Req: ledgercrud.Req{
			ID: &info.ID,
		},
	}
	return handler.GetLedger(ctx)
}

func (h *Handler) UnlockBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	// TODO: LockBalanceOut Can Only Be Called Once
	locked := decimal.RequireFromString(fmt.Sprintf("-%v", h.Amount.String()))
	spendable := h.Amount

	handler := &lockHandler{
		Handler: h,
	}

	key := fmt.Sprintf("ledger-lock-balance-out:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = handler.tryUpdateLedger(ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Spendable:  spendable,
		}, ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, err
}

func (h *Handler) LockBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	locked := h.Amount
	spendable := decimal.RequireFromString(fmt.Sprintf("-%v", h.Amount.String()))

	handler := &lockHandler{
		Handler: h,
	}
	key := fmt.Sprintf("ledger-lock-balance:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = handler.tryUpdateLedger(ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     locked,
			Spendable:  &spendable,
		}, ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, err
}

func (h *lockHandler) tryCreateStatement(ctx context.Context, tx *ent.Tx) error {
	ioType := basetypes.IOType_Outcoming

	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  h.IOSubType,
			Amount:     h.Outcoming,
			IOExtra:    h.IOExtra,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockHandler) tryDeleteStatement(ctx context.Context, tx *ent.Tx) error {
	ioType := basetypes.IOType_Outcoming

	stm, err := statementcrud.SetQueryConds(
		tx.Statement.Query(),
		&statementcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: ioType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *h.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("statement not found")
		}
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := statementcrud.UpdateSet(
		tx.Statement.UpdateOneID(info.ID),
		&statementcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *Handler) UnspendBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	if h.Unlocked.Cmp(decimal.NewFromInt(0)) == 0 && h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
		return nil, fmt.Errorf("nothing todo")
	}

	key := fmt.Sprintf("ledger-unspend-balance:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler := &lockHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err = handler.tryDeleteStatement(ctx, tx); err != nil {
			return err
		}

		// TODO:
		spendable := h.Unlocked.Sub(*h.Outcoming)
		unlocked := decimal.RequireFromString(h.Unlocked.String())

		info, err = handler.tryUpdateLedger(ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &unlocked,
			Outcoming:  h.Outcoming,
			Spendable:  &spendable,
		}, ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

//nolint
func (h *Handler) SpendBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	if h.Unlocked.Cmp(decimal.NewFromInt(0)) == 0 && h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
		return nil, fmt.Errorf("nothing todo")
	}

	key := fmt.Sprintf("ledger-spend-balance:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler := &lockHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err = handler.tryCreateStatement(ctx, tx); err != nil {
			return err
		}

		spendable := h.Unlocked.Sub(*h.Outcoming)
		unlocked := decimal.RequireFromString(h.Unlocked.String())

		info, err = handler.tryUpdateLedger(ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &unlocked,
			Outcoming:  h.Outcoming,
			Spendable:  &spendable,
		}, ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}