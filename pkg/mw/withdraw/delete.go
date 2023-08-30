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
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	withdraw *npool.Withdraw
}

func (h *deleteHandler) unlockBalance(ctx context.Context, tx *ent.Tx) error {
	if h.withdraw.StateStr != types.WithdrawState_Reviewing.String() {
		return nil
	}

	appID := uuid.MustParse(h.withdraw.AppID)
	userID := uuid.MustParse(h.withdraw.UserID)
	coinTypeID := uuid.MustParse(h.withdraw.CoinTypeID)
	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(appID),
			entledger.UserID(userID),
			entledger.CoinTypeID(coinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	spendable := decimal.RequireFromString(h.withdraw.Amount)
	locked := decimal.NewFromInt(0).Sub(spendable)
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
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
		return nil, nil
	}
	if info.State == types.WithdrawState_Transferring {
		return nil, fmt.Errorf("withdraw in transferring state can not be delete")
	}

	handler := &deleteHandler{
		Handler: h,
	}
	handler.withdraw = info

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
