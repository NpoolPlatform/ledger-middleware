package ledger

import (
	"context"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
)

type ledgeropHandler struct {
	*Handler
	ledger *ent.Ledger
}

func (h *ledgeropHandler) getLedger(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		ledger, err := cli.
			Ledger.
			Query().
			Where(
				entledger.AppID(*h.AppID),
				entledger.UserID(*h.UserID),
				entledger.CoinTypeID(*h.CoinTypeID),
				entledger.DeletedAt(0),
			).
			Only(_ctx)
		if err != nil {
			return err
		}
		h.ledger = ledger
		return nil
	})
}
