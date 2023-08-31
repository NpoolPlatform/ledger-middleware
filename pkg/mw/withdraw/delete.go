package withdraw

import (
	"context"
	"fmt"
	"time"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	withdrawcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
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
}

func (h *deleteHandler) unlockBalance(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Withdraw.
		Query().
		Where(
			entwithdraw.ID(*h.ID),
			entwithdraw.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	ledgerInfo, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(info.AppID),
			entledger.UserID(info.UserID),
			entledger.CoinTypeID(info.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(info.Amount)
	stm, err := ledgercrud.UpdateSetWithValidate(
		ledgerInfo,
		&ledgercrud.Req{
			AppID:      &info.AppID,
			UserID:     &info.UserID,
			CoinTypeID: &info.CoinTypeID,
			Locked:     &locked,
			Spendable:  &info.Amount,
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

func (h *deleteHandler) deleteWithdraw(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Withdraw.
		Query().
		Where(
			entwithdraw.ID(*h.ID),
			entwithdraw.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := withdrawcrud.UpdateSet(
		info.Update(),
		&withdrawcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
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
		return nil, fmt.Errorf("withdraw not found")
	}
	if info.State != types.WithdrawState_Reviewing {
		return nil, fmt.Errorf("withdraw only in reviewing state can be deleted")
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.unlockBalance(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}
