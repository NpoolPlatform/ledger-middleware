package withdraw

import (
	"context"
	"time"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	withdraw *npool.Withdraw
}

func (h *deleteHandler) tryUnlockBalance(ctx context.Context, tx *ent.Tx) error {
	if h.withdraw.StateStr != types.WithdrawState_Reviewing.String() {
		return nil
	}
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithAppID(&h.withdraw.AppID, true),
		ledger1.WithUserID(&h.withdraw.UserID, true),
		ledger1.WithCoinTypeID(&h.withdraw.CoinTypeID, true),
	)
	if err != nil {
		return err
	}

	info, err := handler.TryGetLedgerOnly(ctx, tx)
	if err != nil {
		return err
	}

	spendable := decimal.RequireFromString(h.withdraw.Amount)
	locked := decimal.NewFromInt(0).Sub(spendable)
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			AppID:      handler.AppID,
			UserID:     handler.UserID,
			CoinTypeID: handler.CoinTypeID,
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

func (h *deleteHandler) tryDeleteWithdraw(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := tx.
		Withdraw.
		Update().
		Where(
			entwithdraw.ID(*h.ID),
		).
		SetDeletedAt(now).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	info, err := h.GetWithdraw(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &deleteHandler{
		Handler: h,
	}
	handler.withdraw = info

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.tryDeleteWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.tryUnlockBalance(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}
