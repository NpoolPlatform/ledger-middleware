package withdraw

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	ledger   *ent.Ledger
	withdraw *npool.Withdraw
}

func (h *updateHandler) checkWithdrawState(ctx context.Context) error {
	if h.State == nil {
		return nil
	}
	info, err := h.GetWithdraw(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("withdraw not found")
	}
	if info.State == types.WithdrawState_TransactionFail {
		return fmt.Errorf("current withdraw have already failed")
	}
	if h.State.String() == info.StateStr {
		return fmt.Errorf("current state not need to update")
	}
	return nil
}

func (h *updateHandler) tryGetLedger(ctx context.Context) error {
	if h.State == nil {
		return nil
	}
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := ledgercrud.SetQueryConds(
			cli.Ledger.Query(),
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
			return err
		}
		h.ledger = info
		return nil
	})
}

func (h *updateHandler) updateLedger(ctx context.Context, tx *ent.Tx) error {
	if h.State == nil {
		return nil
	}
	spendable := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *h.State {
	case types.WithdrawState_Successful:
		outcoming = decimal.RequireFromString(h.withdraw.Amount)
	case types.WithdrawState_TransactionFail:
		spendable = decimal.RequireFromString(h.withdraw.Amount)
	default:
		return nil
	}

	locked := decimal.RequireFromString(fmt.Sprintf("-%v", h.withdraw.Amount))
	stm, err := ledgercrud.UpdateSetWithValidate(
		h.ledger,
		&ledgercrud.Req{
			Spendable: &spendable,
			Outcoming: &outcoming,
			Locked:    &locked,
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

func (h *updateHandler) updateWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.UpdateSet(
		tx.Withdraw.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.checkWithdrawState(ctx); err != nil {
		return nil, err
	}
	if err := handler.tryGetLedger(ctx); err != nil {
		return nil, err
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLedger(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetWithdraw(ctx)
}
