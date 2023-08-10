package statement

import (
	"context"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
)

type bookkeepingHandler struct {
	*Handler
}

func (h *bookkeepingHandler) tryCreateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) (string, error) {

	return "", nil
}

func (h *bookkeepingHandler) tryCreateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) (string, error) {
	return "", nil
}

func (h *bookkeepingHandler) tryBookKeepingV2(req *crud.Req, ledgerID, profitID string, ctx context.Context, tx *ent.Tx) error {
	return nil
}

func (h *Handler) BookKeepingV2(ctx context.Context) error {
	handler := &bookkeepingHandler{
		Handler: h,
	}

	for _, req := range h.Reqs {
		err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			ledgerID, err := handler.tryCreateLedger(req, _ctx, tx)
			if err != nil {
				return err
			}
			profitID, err := handler.tryCreateProfit(req, _ctx, tx)
			if err != nil {
				return err
			}
			if err := handler.tryBookKeepingV2(req, ledgerID, profitID, _ctx, tx); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
