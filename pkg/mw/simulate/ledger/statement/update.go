package statement

import (
	"context"
	"fmt"
	"time"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/simulate/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/simulatestatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/statement"
)

func (h *Handler) UpdateStatement(ctx context.Context) (*npool.Statement, error) {
	if h.CashUsed == nil {
		return nil, fmt.Errorf("invalid cashused")
	}
	if !*h.CashUsed {
		return nil, fmt.Errorf("invalid cashused")
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			SimulateStatement.
			Query().
			Where(
				entstatement.ID(*h.ID),
				entstatement.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		now := uint32(time.Now().Unix())

		if _, err := crud.UpdateSet(
			info.Update(),
			&crud.Req{
				CashUsed:   h.CashUsed,
				CashUsedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStatement(ctx)
}
