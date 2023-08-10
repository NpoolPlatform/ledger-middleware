package unsoldstatement

import (
	"context"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/unsoldstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/unsoldstatement"
	"github.com/google/uuid"
)

func (h *Handler) CreateUnsoldStatement(ctx context.Context) (*npool.UnsoldStatement, error) {
	now := uint32(time.Now().Unix())
	seconds := h.BenefitIntervalHours * timedef.SecondsPerHour
	timestamp := now / seconds * seconds

	h.Conds = &crud.Conds{
		GoodID:      &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
		CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: timestamp},
	}
	if _, err := h.GetUnsoldStatementOnly(ctx); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.UnsoldStatement.Create(),
			&crud.Req{
				ID:          h.ID,
				GoodID:      h.GoodID,
				CoinTypeID:  h.CoinTypeID,
				Amount:      h.Amount,
				BenefitDate: &timestamp,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return h.GetUnsoldStatement(ctx)
}
