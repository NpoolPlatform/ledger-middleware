package unsoldstatement

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/unsoldstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"
	"github.com/google/uuid"
)

func (h *Handler) CreateUnsoldStatement(ctx context.Context) (*npool.UnsoldStatement, error) {
	if h.GoodID == nil {
		return nil, fmt.Errorf("invalid good id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	if h.Amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}
	if h.BenefitDate == nil {
		return nil, fmt.Errorf("invalid benefit date")
	}

	h.Conds = &crud.Conds{
		GoodID:      &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
		CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: h.BenefitDate},
	}
	exist, err := h.ExistUnsoldStatementConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("unsold statement exist, goodid(%v), cointypeid(%v), benefitdate(%v)", *h.GoodID, *h.CoinTypeID, *h.BenefitDate)
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
				BenefitDate: h.BenefitDate,
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
