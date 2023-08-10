package ledger

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
)

func (h *Handler) DeleteLedger(ctx context.Context) (*npool.Ledger, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetLedger(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.Ledger.
			Update().
			Where(
				entledger.ID(*h.ID),
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
