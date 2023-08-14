package ledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
)

func (h *Handler) UpdateLedger(ctx context.Context) (*npool.Ledger, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		line, err := cli.Ledger.Query().Where(entledger.ID(*h.ID)).Only(_ctx)
		if err != nil {
			return err
		}

		entity, err := crud.UpdateSet(
			line,
			cli.Ledger.UpdateOneID(*h.ID),
			&crud.Req{
				Incoming:  h.Incoming,
				Outcoming: h.Outcoming,
				Spendable: h.Spendable,
				Locked:    h.Locked,
			},
		)
		if err != nil {
			return err
		}
		if _, err := entity.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLedger(ctx)
}