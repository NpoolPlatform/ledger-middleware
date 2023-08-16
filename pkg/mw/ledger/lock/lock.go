package lock

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*Handler
}

func (h *Handler) LockBalanceOut(ctx context.Context) (info *ledgerpb.Ledger, err error) {
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

//nolint
func (h *Handler) UnlockBalanceOut(ctx context.Context) error {

	if h.Unlocked.Cmp(decimal.NewFromInt(0)) == 0 && h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
		return fmt.Errorf("nothing todo")
	}

	ioType := basetypes.IOType_Outcoming
	h.IOType = &ioType

	key := statementKey(&crud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     h.IOType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	})
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler := &statement1.Handler{
		Conds: &crud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: h.IOType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: h.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
		},
	}
	statement, err := handler.GetStatementOnly(ctx)
	if err != nil {
		return err
	}

	if statement != nil {
		statementID, err := uuid.Parse(statement.ID)
		if err != nil {
			return err
		}
		h.ID = &statementID

		//TODO:
		spendable := h.Unlocked.Sub(*h.Outcoming)

		_outcoming, err := decimal.NewFromString(fmt.Sprintf("-%v", *h.Outcoming))
		if err != nil {
			return err
		}

		return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
			ledger1 := &ledger1.Handler{
				Req: ledgercrud.Req{
					AppID:      h.AppID,
					UserID:     h.UserID,
					CoinTypeID: h.CoinTypeID,
					Locked:     h.Unlocked,
					Spendable:  &spendable,
					Outcoming:  &_outcoming,
				},
			}
			if _, err := ledger1.UpdateLedger(ctx); err != nil {
				return err
			}

			if h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
				return nil
			}

			if _, err := handler.DeleteStatement(ctx); err != nil {
				return err
			}

			return nil
		})
	}
	return nil
}

//nolint
func (h *Handler) UnlockBalance(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid coin type id")
	}
	if h.Unlocked == nil {
		return fmt.Errorf("invalid unlocked")
	}
	if h.Outcoming == nil {
		return fmt.Errorf("invalid outcoming")
	}
	if h.IOExtra == nil {
		return fmt.Errorf("invalid extra")
	}
	if h.Unlocked.Cmp(decimal.NewFromInt(0)) == 0 && h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
		return fmt.Errorf("nothing todo")
	}

	key := statementKey(&crud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     h.IOType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	})
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ioType := basetypes.IOType_Outcoming
	h.IOType = &ioType

	handler := &statement1.Handler{
		Conds: &crud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: h.IOType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: h.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
		},
	}
	exist, err := handler.ExistStatementConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("statement already exist, app id %v, user id %v, coin type id %v", *h.AppID, *h.UserID, *h.CoinTypeID)
	}

	spendable := h.Unlocked.Sub(*h.Outcoming)
	h.Amount = h.Outcoming
	_unlocked, err := decimal.NewFromString(fmt.Sprintf("-%v", *h.Unlocked))
	if err != nil {
		return err
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		bookkeeping1 := &bookkeepingHandler{
			Handler: h,
		}
		if _, err := bookkeeping1.tryCreateLedger(&crud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			IOType:     h.IOType,
			IOSubType:  h.IOSubType,
			IOExtra:    h.IOExtra,
		}, ctx, tx); err != nil {
			return err
		}

		ledger1 := &ledger1.Handler{
			Req: ledgercrud.Req{
				AppID:      h.AppID,
				UserID:     h.UserID,
				CoinTypeID: h.CoinTypeID,
				Locked:     &_unlocked,
				Spendable:  &spendable,
				Outcoming:  h.Outcoming,
			},
		}
		if _, err := ledger1.UpdateLedger(ctx); err != nil {
			return err
		}

		if h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
			return nil
		}

		if _, err := handler.CreateStatement(ctx); err != nil {
			return err
		}
		return nil
	})
}
