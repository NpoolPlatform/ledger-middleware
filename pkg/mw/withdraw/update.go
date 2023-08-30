package withdraw

import (
	"context"
	"fmt"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"
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
	withdraw     *npool.Withdraw
	updateLedger bool
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
	if info.StateStr == types.WithdrawState_Rejected.String() ||
		info.StateStr == types.WithdrawState_TransactionFail.String() ||
		info.StateStr == types.WithdrawState_Successful.String() {
		return fmt.Errorf("current withdraw state(%v) can not be update", info.StateStr)
	}
	if h.State.String() != types.WithdrawState_Transferring.String() &&
		h.State.String() != types.WithdrawState_Rejected.String() {
		return fmt.Errorf("can not update withdraw state from %v to %v", info.StateStr, h.State.String())
	}
	if info.StateStr == types.WithdrawState_Reviewing.String() || info.StateStr == types.WithdrawState_Transferring.String() {
		if h.State.String() == types.WithdrawState_Rejected.String() {
			h.updateLedger = true
		}
	}
	if info.StateStr == types.WithdrawState_Transferring.String() {
		if h.State.String() == types.WithdrawState_TransactionFail.String() ||
			h.State.String() == types.WithdrawState_Successful.String() {
			h.updateLedger = true
		}
	}
	h.withdraw = info
	return nil
}

func (h *updateHandler) tryUpdateLedger(ctx context.Context, tx *ent.Tx) error {
	if !h.updateLedger {
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
		ForUpdate().
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
		fallthrough //nolint
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

func (h *updateHandler) tryCreateStatement(ctx context.Context, tx *ent.Tx) error {
	if !h.updateLedger {
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

	platformTransactionID := uuid.UUID{}
	if h.withdraw.PlatformTransactionID != uuid1.InvalidUUIDStr {
		platformTransactionID = uuid.MustParse(h.withdraw.PlatformTransactionID)
	}
	if h.PlatformTransactionID != nil {
		platformTransactionID = *h.PlatformTransactionID
	}
	if platformTransactionID.String() == uuid1.InvalidUUIDStr {
		return fmt.Errorf("invalid platform transaction id, %v", platformTransactionID.String())
	}

	ioExtra := fmt.Sprintf(`{"WithdrawID":"%v","TransactionID":"%v","CID":"%v","TransactionFee":"%v","AccountID":"%v"}`,
		h.withdraw.ID, platformTransactionID.String(), *h.ChainTransactionID, h.FeeAmount.String(), h.withdraw.AccountID,
	)
	amount := decimal.RequireFromString(h.withdraw.Amount)
	if _, err := tx.Statement.Create().
		SetAppID(uuid.MustParse(h.withdraw.AppID)).
		SetUserID(uuid.MustParse(h.withdraw.UserID)).
		SetCoinTypeID(uuid.MustParse(h.withdraw.CoinTypeID)).
		SetIoType(types.IOType_Outcoming.String()).
		SetIoSubType(types.IOSubType_Withdrawal.String()).
		SetIoExtra(ioExtra).
		SetAmount(amount).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	handler := &updateHandler{
		Handler:      h,
		updateLedger: false,
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
		if err := handler.tryCreateStatement(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetWithdraw(ctx)
}
