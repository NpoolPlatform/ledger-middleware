package withdraw

import (
	"context"
	"fmt"

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
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	if h.AccountID == nil {
		return nil, fmt.Errorf("invalid account id")
	}
	if h.Address == nil {
		return nil, fmt.Errorf("invalid address")
	}
	if h.Amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

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
