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

type unlockHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *unlockHandler) unlockBalance(ctx context.Context, tx *ent.Tx) error {
	spendable := h.lock.Amount
	locked := decimal.NewFromInt(0).Sub(spendable)
	stm, err := ledgercrud.UpdateSetWithValidate(h.lop.ledger, &ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     &locked,
		Spendable:  &spendable,
	})
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UnlockBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &unlockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockRollback.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.lop.getLedger(ctx); err != nil {
		if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, err
	}
	if err := handler.getLock(ctx); err != nil {
		if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, err
	}

	if h.Rollback == nil || !*h.Rollback {
		handler.state = types.LedgerLockState_LedgerLockCanceled.Enum()
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.unlockBalance(ctx, tx); err != nil {
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
