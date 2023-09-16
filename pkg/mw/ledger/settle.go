package ledger

import (
	"context"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type settleHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *settleHandler) settleBalance(ctx context.Context, tx *ent.Tx) error {
	outcoming := h.lock.Amount
	locked := decimal.NewFromInt(0).Sub(outcoming)
	stm, err := ledgercrud.UpdateSetWithValidate(h.lop.ledger, &ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     &locked,
		Outcoming:  &outcoming,
	})
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SettleBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &settleHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockSettle.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.lop.getLedger(ctx); err != nil {
		return nil, err
	}
	if err := handler.getLock(ctx); err != nil {
		return nil, err
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.settleBalance(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.ID = &handler.lop.ledger.ID
	return h.GetLedger(ctx)
}
