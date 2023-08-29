package withdraw

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	ledger     *ent.Ledger
	withdraw   *npool.Withdraw
	needUpdate bool
}

func (h *updateHandler) checkWithdrawState(ctx context.Context) error {
	info, err := h.GetWithdraw(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("withdraw not found")
	}
	if h.State == nil {
		return nil
	}
	if info.StateStr == types.WithdrawState_Reviewing.String() {
		if h.State.String() == types.WithdrawState_Rejected.String() {
			h.needUpdate = true
		}
	}
	if info.StateStr == types.WithdrawState_Transferring.String() {
		if h.State.String() == types.WithdrawState_TransactionFail.String() {
			h.needUpdate = true
		}
		if h.State.String() == types.WithdrawState_Successful.String() {
			h.needUpdate = true
		}
	}
	h.withdraw = info
	return nil
}

func (h *updateHandler) updateLedger(ctx context.Context, tx *ent.Tx) error {
	if !h.needUpdate {
		return nil
	}

	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(uuid.MustParse(h.withdraw.AppID)),
			entledger.UserID(uuid.MustParse(h.withdraw.UserID)),
			entledger.CoinTypeID(uuid.MustParse(h.withdraw.CoinTypeID)),
			entledger.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	spendable := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *h.State {
	case types.WithdrawState_Successful:
		outcoming = decimal.RequireFromString(h.withdraw.Amount)
	case types.WithdrawState_TransactionFail:
		fallthrough
	case types.WithdrawState_Rejected:
		spendable = decimal.RequireFromString(h.withdraw.Amount)
	default:
		return nil
	}

	locked := decimal.RequireFromString(fmt.Sprintf("-%v", h.withdraw.Amount))
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
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
