package ledger

import (
	"context"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *lockHandler) lockBalance(ctx context.Context) error {
	spendable := decimal.NewFromInt(0).Sub(*h.Locked)
	stm, err := ledgercrud.UpdateSetWithValidate(h.lop.ledger, &ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     h.Locked,
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

func (h *Handler) LockBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &lockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lop.getLedger(ctx, tx); err != nil {
			return err
		}
		h.ID = &handler.lop.ledger.ID
		if err := handler.lockBalance(ctx); err != nil {
			return err
		}
		if err := handler.createLock(ctx, tx); err != nil {
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
