package withdraw

import (
	"context"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createWithdraw(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Withdraw.Create(),
			&h.Req,
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) CreateWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.createWithdraw(ctx); err != nil {
		return nil, err
	}

	return h.GetWithdraw(ctx)
}
