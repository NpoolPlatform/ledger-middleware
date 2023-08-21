package unsoldstatement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entunsoldstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/unsold"
)

func (h *Handler) DeleteUnsoldStatement(ctx context.Context) (*npool.UnsoldStatement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetUnsoldStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("id not exist %v", *h.ID)
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.UnsoldStatement.
			Update().
			Where(
				entunsoldstatement.ID(*h.ID),
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
