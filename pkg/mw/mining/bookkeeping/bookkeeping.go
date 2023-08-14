package bookkeeping

import (
	"context"
	"fmt"

	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodstatement"
	unsoldstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	goodledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodledger"
	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodstatement"
	unsoldstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/unsoldstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

func (h *Handler) BookKeeping(ctx context.Context) error {
	if h.GoodID == nil {
		return fmt.Errorf("invalid good id")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid coin type id")
	}
	if h.UnsoldAmount == nil {
		return fmt.Errorf("invalid unsold amount")
	}
	if h.TotalAmount == nil {
		return fmt.Errorf("invalid total amount")
	}
	if h.TechniqueServiceFeeAmount == nil {
		return fmt.Errorf("invalid fee amount")
	}
	if h.BenefitDate == nil {
		return fmt.Errorf("invalid benefit date")
	}

	goodstatementHandler := &goodstatement1.Handler{
		Conds: &goodstatement.Conds{
			GoodID:      &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
			BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: h.BenefitDate},
		},
	}
	info, err := goodstatementHandler.GetGoodStatementOnly(ctx)
	if err != nil {
		return err
	}
	if info != nil {
		return fmt.Errorf("benefit exist, good id: %v, benefit date: %v", *h.GoodID, *h.BenefitDate)
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		goodstatementHandler := &goodstatement1.Handler{
			Req: goodstatement.Req{
				GoodID:      h.GoodID,
				CoinTypeID:  h.CoinTypeID,
				Amount:      h.TotalAmount,
				BenefitDate: h.BenefitDate,
			},
		}
		if _, err := goodstatementHandler.CreateGoodStatement(ctx); err != nil {
			return err
		}

		goodledgerHandler := &goodledger1.Handler{
			Conds: &goodledgercrud.Conds{
				GoodID:     &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
				CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
			},
			Req: goodledgercrud.Req{
				GoodID:     h.GoodID,
				CoinTypeID: h.CoinTypeID,
			},
		}

		goodledger, err := goodledgerHandler.GetGoodLedgerOnly(ctx)
		if err != nil {
			return err
		}
		goodledgerID := ""

		if goodledger == nil {
			_goodledger, err := goodledgerHandler.CreateGoodLedger(ctx)
			if err != nil {
				return err
			}
			goodledgerID = _goodledger.ID
		} else {
			goodledgerID = goodledger.ID
		}

		toPlatform := h.UnsoldAmount.Add(*h.TechniqueServiceFeeAmount)
		toUser := h.TotalAmount.Sub(toPlatform)

		_goodLedgerID, err := uuid.Parse(goodledgerID)
		if err != nil {
			return err
		}
		goodledgerHandler = &goodledger1.Handler{
			Req: goodledgercrud.Req{
				ID:         &_goodLedgerID,
				ToPlatform: &toPlatform,
				ToUser:     &toUser,
				Amount:     h.TotalAmount,
			},
		}
		if _, err := goodledgerHandler.UpdateGoodLedger(ctx); err != nil {
			return err
		}

		unsoldHandler := unsoldstatement1.Handler{
			Req: unsoldstatementcrud.Req{
				GoodID:      h.GoodID,
				CoinTypeID:  h.CoinTypeID,
				Amount:      h.TotalAmount,
				BenefitDate: h.BenefitDate,
			},
		}
		if _, err := unsoldHandler.CreateUnsoldStatement(ctx); err != nil {
			return err
		}
		return nil
	})
}
