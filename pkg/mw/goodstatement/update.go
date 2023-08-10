package goodstatement

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/goodstatement"
)

func (h *Handler) UpdateGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.GoodStatement.UpdateOneID(*h.ID),
			&crud.Req{
				GoodID:     h.GoodID,
				CoinTypeID: h.CoinTypeID,
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

	return h.GetGoodStatement(ctx)
}
