package profit

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entprofit "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/simulateprofit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/profit"
)

func (h *Handler) DeleteProfit(ctx context.Context) (*npool.Profit, error) {
	info, err := h.GetProfit(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid id %v", *h.ID)
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.SimulateProfit.
			Update().
			Where(
				entprofit.ID(*h.ID),
			).
			SetDeletedAt(now).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
