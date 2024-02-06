package ledger

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"

	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *lockHandler) lockBalance(ctx context.Context) error {
	spendable := decimal.NewFromInt(0).Sub(*h.Locked)
	stm, err := ledgercrud.UpdateSetWithValidate(h.lop.ledgers[0], &ledgercrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		Locked:     h.Locked,
		Spendable:  &spendable,
	})
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) LockBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &lockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.lockopHandler.getLocks(ctx); err != nil {
		return nil, err
	}
	if len(handler.lockopHandler.locks) > 0 {
		return nil, fmt.Errorf("invalid lockid")
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		h.EntID = &handler.lop.ledgers[0].EntID
		if err := handler.lockBalance(ctx); err != nil {
			return err
		}
		if err := handler.createLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetLedger(ctx)
}

func (h *lockHandler) lockBalances(ctx context.Context) error {
	for _, balance := range h.Balances {
		ledger := h.lop.coinLedger(balance.CoinTypeID)
		if ledger == nil {
			return fmt.Errorf("invalid ledger")
		}
		spendable := decimal.NewFromInt(0).Sub(balance.Amount)
		stm, err := ledgercrud.UpdateSetWithValidate(ledger, &ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: &balance.CoinTypeID,
			Locked:     &balance.Amount,
			Spendable:  &spendable,
		})
		if err != nil {
			return err
		}
		if _, err := stm.Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) LockBalances(ctx context.Context) ([]*ledgermwpb.Ledger, error) {
	handler := &lockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.lockopHandler.getLocks(ctx); err != nil {
		return nil, err
	}
	if len(handler.lockopHandler.locks) > 0 {
		return nil, fmt.Errorf("invalid lockid")
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		for _, balance := range h.Balances {
			ledger := handler.lop.coinLedger(balance.CoinTypeID)
			if ledger == nil {
				return fmt.Errorf("invalid ledger")
			}
			balance.LedgerID = ledger.EntID
		}
		if err := handler.lockBalances(ctx); err != nil {
			return err
		}
		if err := handler.createLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &ledgercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: handler.lop.ledgerIDs},
	}
	h.Offset = 0
	h.Limit = int32(len(handler.lop.ledgerIDs))
	infos, _, err := h.GetLedgers(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
