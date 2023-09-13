package profit

import (
	"context"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/profit"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/profit"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createProfit(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Profit.Create(),
			&h.Req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) CreateProfit(ctx context.Context) (*npool.Profit, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.createProfit(ctx); err != nil {
		return nil, err
	}

	return h.GetProfit(ctx)
}
