package withdraw

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	entwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	withdraw *ent.Withdraw
}

func (h *updateHandler) checkWithdrawState(ctx context.Context) error {
    err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
        info, err := cli.
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
        h.withdraw = info
        return nil
    })
    if err != nil {
        return err
    }

    if h.State == nil {
        return nil
    }
    state := types.WithdrawState(types.WithdrawState_value[h.withdraw.State])
    switch state {
    case types.WithdrawState_Rejected:
        fallthrough //nolint
    case types.WithdrawState_TransactionFail:
        fallthrough //nolint
    case types.WithdrawState_Successful:
        return fmt.Errorf("current withdraw state(%v) can not be update", h.withdraw.State)
}

    switch state {
    case types.WithdrawState_Reviewing:
        switch *h.State {
        case types.WithdrawState_Rejected:
        case types.WithdrawState_Transferring:
        default:
            return fmt.Errorf("can not update withdraw state from %v to %v", h.withdraw.State, h.State.String())
    }
    case types.WithdrawState_Transferring:
        if h.PlatformTransactionID == nil && h.withdraw.PlatformTransactionID.String() == uuid.Nil.String() {
            return fmt.Errorf("invalid platform transaction id")
        }
        switch *h.State {
        case types.WithdrawState_TransactionFail:
        case types.WithdrawState_Successful:
        default:
            return fmt.Errorf("can not update withdraw state from %v to %v", h.withdraw.State, h.State.String())
    }
        return nil
}
    if *h.State == types.WithdrawState_Transferring {
        if h.PlatformTransactionID == nil && h.withdraw.PlatformTransactionID.String() == uuid.Nil.String() {
            return fmt.Errorf("invalid platform transaction id")
        }
    }
    return nil
}

func (h *updateHandler) tryUpdateLedger(ctx context.Context, tx *ent.Tx) error {
	if h.State == nil {
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
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	spendable := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *h.State {
	case types.WithdrawState_Successful:
		outcoming = h.withdraw.Amount
	case types.WithdrawState_TransactionFail:
		fallthrough //nolint
	case types.WithdrawState_Rejected:
		spendable = h.withdraw.Amount
	default:
		return nil
	}

	locked := decimal.NewFromInt(0).Sub(h.withdraw.Amount)
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
	if h.PlatformTransactionID != nil && h.withdraw.PlatformTransactionID.String() != uuid.Nil.String() {
		return fmt.Errorf("current platform transaction id can not be updated")
	}
	if _, err := crud.UpdateSet(
		tx.Withdraw.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	if h.State == nil {
		return nil
	}
	if h.State.String() != types.WithdrawState_Successful.String() {
		return nil
	}
	if h.ChainTransactionID == nil {
		return fmt.Errorf("invalid chain transaction id")
	}
	if h.FeeAmount == nil {
		return fmt.Errorf("invalid fee amount")
	}

	ioExtra := fmt.Sprintf(`{"WithdrawID":"%v","TransactionID":"%v","CID":"%v","TransactionFee":"%v","AccountID":"%v"}`,
		h.withdraw.ID, h.withdraw.PlatformTransactionID.String(), *h.ChainTransactionID, h.FeeAmount.String(), h.withdraw.AccountID,
	)

	ioType := types.IOType_Outcoming
	ioSubType := types.IOSubType_Withdrawal
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:      &h.withdraw.AppID,
			UserID:     &h.withdraw.UserID,
			CoinTypeID: &h.withdraw.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			IOExtra:    &ioExtra,
			Amount:     &h.withdraw.Amount,
		},
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
		if err := handler.tryUpdateLedger(ctx, tx); err != nil {
			return err
		}
		if err := handler.createStatement(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetWithdraw(ctx)
}
