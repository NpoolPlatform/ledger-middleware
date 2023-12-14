package ledger

import (
	"context"
	"fmt"

	ledgerlockcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/lock"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledgerlock "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledgerlock"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
)

type lockopHandler struct {
	*Handler
	lock  *ent.LedgerLock
	state *types.LedgerLockState
}

func (h *lockopHandler) getLock(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		lock, err := cli.
			LedgerLock.
			Query().
			Where(
				entledgerlock.EntID(*h.LockID),
				entledgerlock.DeletedAt(0),
			).
			Only(_ctx)
		if err != nil {
			return err
		}
		h.lock = lock
		return nil
	})
}

func (h *lockopHandler) createLock(ctx context.Context, tx *ent.Tx) error {
	if _, err := ledgerlockcrud.CreateSet(
		tx.LedgerLock.Create(),
		&ledgerlockcrud.Req{
			EntID:    h.LockID,
			LedgerID: h.EntID,
			Amount:   h.Locked,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockopHandler) updateLock(ctx context.Context, tx *ent.Tx) error {
	switch h.lock.LockState {
	case types.LedgerLockState_LedgerLockLocked.String():
		switch *h.state {
		case types.LedgerLockState_LedgerLockSettle:
		case types.LedgerLockState_LedgerLockRollback:
		case types.LedgerLockState_LedgerLockCanceled:
		default:
			return fmt.Errorf("invalid ledgerlockstate")
		}
	default:
		return fmt.Errorf("invalid ledgerlockstate")
	}

	stm := tx.
		LedgerLock.
		UpdateOneID(h.lock.ID).
		SetLockState(h.state.String())
	if *h.state == types.LedgerLockState_LedgerLockSettle {
		stm.SetStatementID(*h.StatementID)
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}
