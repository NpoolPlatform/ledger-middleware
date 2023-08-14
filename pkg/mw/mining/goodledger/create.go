package goodledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"
	"github.com/google/uuid"
)

func (h *Handler) CreateGoodLedger(ctx context.Context) (*npool.GoodLedger, error) {
	if h.GoodID == nil {
		return nil, fmt.Errorf("invalid good id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}

	h.Conds = &crud.Conds{
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
	}
	if _, err := h.GetGoodLedgerOnly(ctx); err != nil {
		return nil, err
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
