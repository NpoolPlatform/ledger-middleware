package statement

import (
	"context"
	"crypto/sha256"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	profitcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/profit"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	ledgerhandler "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	profithandler "github.com/NpoolPlatform/ledger-middleware/pkg/mw/profit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type bookkeepingHandler struct {
	*Handler
	Unlocked  *decimal.Decimal
	Outcoming *decimal.Decimal
}

func statementKey(in *crud.Req) string {
	extra := sha256.Sum256([]byte(*in.IOExtra))
	return fmt.Sprintf("ledger-statement:%v:%v:%v:%v:%v:%v:%v",
		*in.AppID,
		*in.UserID,
		*in.CoinTypeID,
		in.IOType.String(),
		in.IOSubType.String(),
		*in.IOExtra,
		extra,
	)
}

func (h *bookkeepingHandler) tryCreateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) (string, error) {
	key := fmt.Sprintf("ledger-ledger:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)

	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ledger1 := &ledgerhandler.Handler{
		Req: ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
		},
	}
	ledger1.Conds = &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
	}

	ledger, err := ledger1.GetLedgerOnly(ctx)
	if err != nil {
		return "", err
	}
	if ledger != nil {
		return ledger.ID, nil
	}

	info, err := ledger1.CreateLedger(ctx)
	if err != nil {
		return "", err
	}
	return info.ID, nil
}

func (h *bookkeepingHandler) tryCreateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) (string, error) {
	key := fmt.Sprintf("ledger-profit:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)

	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	profit1 := &profithandler.Handler{
		Req: profitcrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
		},
	}
	profit1.Conds = &profitcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
	}

	profit, err := profit1.GetProfitOnly(ctx)
	if err != nil {
		return "", err
	}
	if profit != nil {
		return profit.ID, nil
	}

	info, err := profit1.CreateProfit(ctx)
	if err != nil {
		return "", err
	}

	return info.ID, nil
}

func (h *bookkeepingHandler) tryBookKeepingV2(req *crud.Req, ledgerID, profitID string, ctx context.Context, tx *ent.Tx) error {

	return nil
}

func (h *Handler) BookKeepingV2(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid coin type id")
	}

	handler := &bookkeepingHandler{
		Handler: h,
	}

	for _, req := range h.Reqs {
		err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			ledgerID, err := handler.tryCreateLedger(req, _ctx, tx)
			if err != nil {
				return err
			}
			profitID, err := handler.tryCreateProfit(req, _ctx, tx)
			if err != nil {
				return err
			}
			if err := handler.tryBookKeepingV2(req, ledgerID, profitID, _ctx, tx); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *bookkeepingHandler) LockBalance(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid coin type id")
	}
	if h.Amount == nil {
		return fmt.Errorf("invalid amount in lock balance")
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		_ledgerID, err := h.tryCreateLedger(&h.Req, ctx, tx)
		if err != nil {
			return err
		}
		ledgerID, err := uuid.Parse(_ledgerID)
		if err != nil {
			return err
		}

		spendable, err := decimal.NewFromString(fmt.Sprintf("-%v", h.Amount.String()))
		if err != nil {
			return err
		}

		ledger1 := &ledgerhandler.Handler{
			Req: ledgercrud.Req{
				ID:         &ledgerID,
				AppID:      h.AppID,
				UserID:     h.UserID,
				CoinTypeID: h.CoinTypeID,
				Locked:     h.Amount,
				Spendable:  &spendable,
			},
		}
		if _, err := ledger1.UpdateLedger(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *bookkeepingHandler) UnlockBalance(ctx context.Context) error {
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

	key := statementKey(&h.Req)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ioType := basetypes.IOType_Outcoming
	h.IOType = &ioType

	h.Conds = &crud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		IOType:     &cruder.Cond{Op: cruder.EQ, Val: h.IOType},
		IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: h.IOSubType},
		IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
	}
	exist, err := h.ExistStatementConds(ctx)
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

	ledger1 := &ledgerhandler.Handler{
		Req: ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &_unlocked,
			Spendable:  &spendable,
			Outcoming:  h.Outcoming,
		},
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := h.tryCreateLedger(&h.Req, ctx, tx); err != nil {
			return err
		}

		if _, err := ledger1.UpdateLedger(ctx); err != nil {
			return err
		}

		if h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
			return nil
		}

		if _, err := h.CreateStatement(ctx); err != nil {
			return err
		}
		return nil
	})
}
