package ledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createLedger(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Ledger.Create(),
			&h.Req,
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) CreateLedger(ctx context.Context) (*npool.Ledger, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	// if h.Incoming == nil {
	// 	return nil, fmt.Errorf("invalid incoming")
	// }
	// if h.Outcoming == nil {
	// 	return nil, fmt.Errorf("invalid outcoming")
	// }
	// if h.Spendable == nil {
	// 	return nil, fmt.Errorf("invalid spendable")
	// }
	// if h.Locked == nil {
	// 	return nil, fmt.Errorf("invalid locked")
	// }

	h.Conds = &crud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
	}

	exist, err := h.ExistStatementConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("ledger already exist, cointypeid %v, userid %v", h.CoinTypeID, h.UserID)
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.createLedger(ctx); err != nil {
		return nil, err
	}

	return h.GetLedger(ctx)
}
