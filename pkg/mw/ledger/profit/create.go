package profit

import (
	"context"
	"fmt"

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
		stm, err := crud.CreateSet(
			cli.Profit.Create(),
			&h.Req,
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) CreateProfit(ctx context.Context) (*npool.Profit, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}

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
