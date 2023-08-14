package goodstatement

import (
	"context"
	"fmt"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodstatement"
	"github.com/google/uuid"
)

func (h *Handler) CreateGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	if h.GoodID == nil {
		return nil, fmt.Errorf("invalid good id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}

	now := uint32(time.Now().Unix())
	seconds := *h.BenefitIntervalHours * timedef.SecondsPerHour
	timestamp := now / seconds * seconds

	h.Conds = &crud.Conds{
		GoodID:      &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
		CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: timestamp},
	}

	info, err := h.GetGoodStatementOnly(ctx)
	if err != nil {
		return nil, err
	}
	if info != nil {
		id, err := uuid.Parse(info.ID)
		if err != nil {
			return nil, err
		}
		h.ID = &id
		h.GetGoodStatement(ctx)
		return h.GetGoodStatement(ctx)
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.GoodStatement.Create(),
			&h.Req,
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return h.GetGoodStatement(ctx)
}
