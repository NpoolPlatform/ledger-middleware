package goodledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"
)

func (h *Handler) UpdateGoodLedger(ctx context.Context) (*npool.GoodLedger, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.GoodLedger.UpdateOneID(*h.ID),
			&crud.Req{
				ToPlatform: h.ToPlatform,
				ToUser:     h.ToUser,
				Amount:     h.Amount,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGoodLedger(ctx)
}
