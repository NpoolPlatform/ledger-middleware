package bookkeeping

import (
	"context"
	"fmt"

	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodledger"
	goodstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodstatement"
	unsoldstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	goodledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodledger"
	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodstatement"
	unsoldstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/unsoldstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) tryCreateUnsoldStatement(ctx context.Context) error {
	handler := unsoldstatement1.Handler{
		Req: unsoldstatementcrud.Req{
			GoodID:      h.GoodID,
			CoinTypeID:  h.CoinTypeID,
			Amount:      h.TotalAmount,
			BenefitDate: h.BenefitDate,
		},
	}
	if _, err := handler.CreateUnsoldStatement(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateGoodStatement(ctx context.Context) error {
	handler := &goodstatement1.Handler{
		Req: goodstatementcrud.Req{
			GoodID:      h.GoodID,
			CoinTypeID:  h.CoinTypeID,
			Amount:      h.TotalAmount,
			BenefitDate: h.BenefitDate,
		},
	}
	if _, err := handler.CreateGoodStatement(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateOrAddGoodLedger(ctx context.Context, tx *ent.Tx) error {
	handler := &goodledger1.Handler{
		Conds: &goodledgercrud.Conds{
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		},
	}

	info, err := handler.GetGoodLedgerOnly(ctx)
	if err != nil {
		return err
	}

	toPlatform := h.UnsoldAmount.Add(*h.TechniqueServiceFeeAmount)
	toUser := h.TotalAmount.Sub(toPlatform)

	// create
	if info == nil {
		if _, err := goodledgercrud.CreateSet(
			tx.GoodLedger.Create(),
			&goodledgercrud.Req{
				GoodID:     h.GoodID,
				CoinTypeID: h.CoinTypeID,
				ToPlatform: &toPlatform,
				ToUser:     &toUser,
				Amount:     h.TotalAmount,
			},
		).Save(ctx); err != nil {
			return nil
		}
		return nil
	}

	// update
	id, err := uuid.Parse(info.ID)
	if err != nil {
		return err
	}
	handler = &goodledger1.Handler{
		Req: goodledgercrud.Req{
			ID:         &id,
			ToPlatform: &toPlatform,
			ToUser:     &toUser,
			Amount:     h.TotalAmount,
		},
	}
	if _, err := handler.UpdateGoodLedger(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) BookKeeping(ctx context.Context) error {
	goodstatementHandler := &goodstatement1.Handler{
		Conds: &goodstatementcrud.Conds{
			GoodID:      &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
			BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: h.BenefitDate},
		},
	}
	exist, err := goodstatementHandler.ExistGoodStatementConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("benefit exist, good id: %v, benefit date: %v", *h.GoodID, *h.BenefitDate)
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		handler := &createHandler{
			Handler: h,
		}

		if err := handler.tryCreateGoodStatement(ctx); err != nil {
			return err
		}
		if err := handler.tryCreateOrAddGoodLedger(ctx, tx); err != nil {
			return err
		}
		if err := handler.tryCreateUnsoldStatement(ctx); err != nil {
			return err
		}
		return nil
	})
}
