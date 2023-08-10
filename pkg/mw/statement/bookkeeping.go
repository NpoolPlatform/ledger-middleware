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
)

type bookkeepingHandler struct {
	*Handler
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
