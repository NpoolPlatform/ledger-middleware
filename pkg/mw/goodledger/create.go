package goodledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/goodledger"
	"github.com/google/uuid"
)

func (h *Handler) CreateGoodLedger(ctx context.Context) (*npool.GoodLedger, error) {
	if h.GoodID == nil {
		return nil, fmt.Errorf("invalid good id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	if h.Amount == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	if h.ToPlatform == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	if h.ToUser == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.GoodLedger.Create(),
			&h.Req,
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return h.GetGoodLedger(ctx)
}
