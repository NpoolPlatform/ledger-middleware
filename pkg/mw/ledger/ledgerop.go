package ledger

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"

	"github.com/google/uuid"
)

type ledgeropHandler struct {
	*Handler
	ledgers   []*ent.Ledger
	ledgerIDs []uuid.UUID
}

func (h *ledgeropHandler) getLedgers(ctx context.Context, tx *ent.Tx) error {
	stm := tx.Ledger.Query()
	if len(h.ledgerIDs) > 0 {
		stm.Where(
			entledger.EntIDIn(h.ledgerIDs...),
			entledger.DeletedAt(0),
		)
	} else if len(h.Balances) > 0 {
		coinTypeIDs := []uuid.UUID{}
		for _, balance := range h.Balances {
			coinTypeIDs = append(coinTypeIDs, balance.CoinTypeID)
		}
		stm.Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CoinTypeIDIn(coinTypeIDs...),
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
	ledgers, err := stm.ForUpdate().All(ctx)
	if err != nil {
		return err
	}
	if len(ledgers) == 0 {
		return fmt.Errorf("invalid ledgers")
	}
	h.ledgers = ledgers
	return nil
}

func (h *ledgeropHandler) coinLedger(coinTypeID uuid.UUID) *ent.Ledger {
	for _, ledger := range h.ledgers {
		if ledger.CoinTypeID == coinTypeID {
			return ledger
		}
	}
	return nil
}
