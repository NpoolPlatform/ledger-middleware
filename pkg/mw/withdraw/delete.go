package withdraw

import (
	"context"
	"time"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	entwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	withdraw *ent.Withdraw
}

func (h *deleteHandler) tryGetWithdraw(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Withdraw.
		Query().
		Where(
			entwithdraw.ID(*h.ID),
			entwithdraw.DeletedAt(0),
		).
		First(ctx)
	if err != nil {
		return err
	}
	h.withdraw = info
	return nil
}

func (h *deleteHandler) tryUnlockBalance(ctx context.Context, tx *ent.Tx) error {
	if h.withdraw.State != types.WithdrawState_Transferring.String() {
		return nil
	}
	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(h.withdraw.AppID),
			entledger.UserID(h.withdraw.UserID),
			entledger.CoinTypeID(h.withdraw.CoinTypeID),
			entledger.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(h.withdraw.Amount)
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			AppID:      &h.withdraw.AppID,
			UserID:     &h.withdraw.UserID,
			CoinTypeID: &h.withdraw.CoinTypeID,
			Locked:     &locked,
			Spendable:  &h.withdraw.Amount,
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

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.tryGetWithdraw(ctx, tx); err != nil {
			return err
		}
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
