//nolint:dupl
package ledger

import (
	"context"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type settleHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *settleHandler) settleBalance(ctx context.Context) error {
	outcoming := h.lock.Amount
	locked := decimal.NewFromInt(0).Sub(outcoming)
	stm, err := ledgercrud.UpdateSetWithValidate(h.lop.ledger, &ledgercrud.Req{
		Locked:    &locked,
		Outcoming: &outcoming,
	})
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *settleHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	ioType := types.IOType_Outcoming
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			ID:         h.StatementID,
			AppID:      &h.lop.ledger.AppID,
			UserID:     &h.lop.ledger.UserID,
			CoinTypeID: &h.lop.ledger.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  h.IOSubType,
			IOExtra:    h.IOExtra,
			Amount:     &h.lock.Amount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SettleBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &settleHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockSettle.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.getLock(ctx); err != nil {
		return nil, err
	}
	handler.lop.ledgerID = &handler.lock.LedgerID

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lop.getLedger(ctx, tx); err != nil {
			return err
		}
		if err := handler.settleBalance(ctx); err != nil {
			return err
		}
		if err := handler.createStatement(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.ID = &handler.lop.ledger.ID
	return h.GetLedger(ctx)
}
