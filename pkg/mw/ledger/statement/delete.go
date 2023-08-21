package statement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
)

func (h *Handler) DeleteStatement(ctx context.Context) (*npool.Statement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid id %v", *h.ID)
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.Statement.
			Update().
			Where(
				entstatement.ID(*h.ID),
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
