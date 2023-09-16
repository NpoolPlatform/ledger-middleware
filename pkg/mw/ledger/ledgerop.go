package ledger

import (
	"context"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"

	"github.com/google/uuid"
)

type ledgeropHandler struct {
	*Handler
	ledger   *ent.Ledger
	ledgerID *uuid.UUID
}

func (h *ledgeropHandler) getLedger(ctx context.Context, tx *ent.Tx) error {
	stm := tx.Ledger.Query()
	if h.ledgerID != nil {
		stm.Where(
			entledger.ID(*h.ledgerID),
			entledger.DeletedAt(0),
		)
	} else {
		stm.Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CoinTypeID(*h.CoinTypeID),
			entledger.DeletedAt(0),
		)
	}
	ledger, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		return err
	}
	h.ledger = ledger
	return nil
}
