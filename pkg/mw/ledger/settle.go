package ledger

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type settleHandler struct {
	*lockopHandler
	lop *ledgeropHandler
}

func (h *settleHandler) settleBalances(ctx context.Context) error {
	for _, lock := range h.locks {
		outcoming := lock.Amount
		locked := decimal.NewFromInt(0).Sub(outcoming)

		ledger := h.lop.ledger(lock.LedgerID)
		if ledger == nil {
			return fmt.Errorf("invalid ledger")
		}

		stm, err := ledgercrud.UpdateSetWithValidate(ledger, &ledgercrud.Req{
			Locked:    &locked,
			Outcoming: &outcoming,
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

func (h *settleHandler) createStatements(ctx context.Context, tx *ent.Tx) error {
	for i, lock := range h.locks {
		if err := func() error {
			ledger := h.lop.ledger(lock.LedgerID)
			if ledger == nil {
				return fmt.Errorf("invalid ledger")
			}

			key := statement1.LockKey(ledger.AppID, ledger.UserID, ledger.CoinTypeID, *h.IOExtra)
			if err := redis2.TryLock(key, 0); err != nil {
				return err
			}
			defer func() {
				_ = redis2.Unlock(key)
			}()
			ioType := types.IOType_Outcoming
			stm, err := statementcrud.SetQueryConds(tx.Statement.Query(), &statementcrud.Conds{
				AppID:      &cruder.Cond{Op: cruder.EQ, Val: ledger.AppID},
				UserID:     &cruder.Cond{Op: cruder.EQ, Val: ledger.UserID},
				CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: ledger.CoinTypeID},
				IOType:     &cruder.Cond{Op: cruder.EQ, Val: ioType},
				IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *h.IOSubType},
				IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *h.IOExtra},
			})
			if err != nil {
				return err
			}
			exist, err := stm.Exist(ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("statement already exist")
			}

			if _, err := statementcrud.CreateSet(tx.Statement.Create(), &statementcrud.Req{
				EntID:      &h.StatementIDs[i],
				AppID:      &ledger.AppID,
				UserID:     &ledger.UserID,
				CoinTypeID: &ledger.CoinTypeID,
				IOType:     &ioType,
				IOSubType:  h.IOSubType,
				IOExtra:    h.IOExtra,
				Amount:     &lock.Amount,
			}).Save(ctx); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return err
		}
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

	if err := handler.getLocks(ctx); err != nil {
		return nil, err
	}

	handler.lop.ledgerIDs = []uuid.UUID{handler.locks[0].LedgerID}
	h.StatementIDs = []uuid.UUID{*h.StatementID}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		if err := handler.settleBalances(ctx); err != nil {
			return err
		}
		if err := handler.createStatements(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.ID = &handler.lop.ledgers[0].ID
	return h.GetLedger(ctx)
}

func (h *Handler) SettleBalances(ctx context.Context) ([]*ledgermwpb.Ledger, error) {
	handler := &settleHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.LedgerLockState_LedgerLockSettle.Enum(),
		},
		lop: &ledgeropHandler{
			Handler: h,
		},
	}

	if err := handler.getLocks(ctx); err != nil {
		return nil, err
	}
	for _, lock := range handler.locks {
		handler.lop.ledgerIDs = append(handler.lop.ledgerIDs, lock.LedgerID)
	}
	if len(h.StatementIDs) != len(handler.locks) {
		return nil, fmt.Errorf("mismatched statementids")
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.lop.getLedgers(ctx, tx); err != nil {
			return err
		}
		if err := handler.settleBalances(ctx); err != nil {
			return err
		}
		if err := handler.createStatements(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
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
