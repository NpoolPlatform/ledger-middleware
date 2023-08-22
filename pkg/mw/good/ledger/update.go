package ledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger"
)

func (h *Handler) UpdateGoodLedger(ctx context.Context) (*npool.GoodLedger, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		line, err := cli.GoodLedger.Query().Where(entgoodledger.ID(*h.ID)).Only(_ctx)
		if err != nil {
			return err
		}
		entity, err := crud.UpdateSet(
			line,
			cli.GoodLedger.UpdateOneID(*h.ID),
			&crud.Req{
				ToPlatform: h.ToPlatform,
				ToUser:     h.ToUser,
				Amount:     h.Amount,
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

	return h.GetGoodLedger(ctx)
}
