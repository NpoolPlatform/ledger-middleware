package ledger

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	ledgerpb "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type addHandler struct {
	*Handler
	ledger   *ent.Ledger
	origin   *ent.Statement
	rollback *ent.Statement
}

func (h *addHandler) getLedger(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CoinTypeID(*h.CoinTypeID),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	h.ledger = info
	return nil
}

func (h *addHandler) tryUnlock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}

	spendable := *h.Spendable
	locked := decimal.NewFromInt(0).Sub(*h.Spendable)

	stm, err := ledgercrud.UpdateSet(
		h.ledger,
		tx.Ledger.UpdateOneID(h.ledger.ID),
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Spendable:  &spendable,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *addHandler) getStatement(ctx context.Context, tx *ent.Tx) (*ent.Statement, error) {
	if h.IOSubType == nil {
		return nil, fmt.Errorf("invalid io sub type")
	}
	if h.IOExtra == nil {
		return nil, fmt.Errorf("invalid io extra")
	}

	// get statement
	ioType := ledgerpb.IOType_Outcoming
	origin, err := tx.
		Statement.
		Query().
		Where(
			entstatement.AppID(*h.AppID),
			entstatement.UserID(*h.UserID),
			entstatement.CoinTypeID(*h.CoinTypeID),
			entstatement.IoType(ioType.String()),
			entstatement.IoSubType(h.IOSubType.String()),
			entstatement.IoExtra(*h.IOExtra),
		).
		Order(ent.Desc(entstatement.FieldUpdatedAt)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("statement not exist")
		}
		return nil, err
	}
	h.origin = origin
	return origin, nil
}

func (h *addHandler) getRollbackStatement(ctx context.Context, tx *ent.Tx) error {
	// get statement
	origin, err := h.getStatement(ctx, tx)
	if err != nil {
		return err
	}

	// get rollback statement
	ioType := ledgerpb.IOType_Incoming
	ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, origin.ID.String())
	info, err := tx.
		Statement.
		Query().
		Where(
			entstatement.AppID(*h.AppID),
			entstatement.UserID(*h.UserID),
			entstatement.CoinTypeID(*h.CoinTypeID),
			entstatement.IoType(ioType.String()),
			entstatement.IoSubType(h.IOSubType.String()),
			entstatement.IoExtra(ioExtra),
		).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}
	h.rollback = info
	return nil
}

func (h *addHandler) tryUnspend(ctx context.Context, tx *ent.Tx) error {
	if h.Locked == nil {
		return nil
	}
	// already rollback
	if h.rollback != nil {
		return fmt.Errorf("statement already rolled back")
	}

	// rollback
	ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, h.origin.ID.String())
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithChangeLedger(),
	)
	if err != nil {
		return err
	}

	ioType := ledgerpb.IOType_Incoming
	handler.Req = statementcrud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    &ioExtra,
		Amount:     h.Locked,
	}
	if _, err := handler.CreateStatement(ctx); err != nil {
		return err
	}

	locked := *h.Locked
	outcoming := decimal.NewFromInt(0).Sub(*h.Locked)

	stm, err := ledgercrud.UpdateSet(
		h.ledger,
		tx.Ledger.UpdateOneID(h.ledger.ID),
		&ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Outcoming:  &outcoming,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

// Unlock & Unspend
func (h *Handler) AddBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	if err := h.validate(); err != nil {
		return nil, err
	}

	handler := &addHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.getLedger(ctx, tx); err != nil {
			return err
		}
		if err := handler.tryUnlock(ctx, tx); err != nil {
			return err
		}
		if err := handler.getRollbackStatement(ctx, tx); err != nil {
			return err
		}
		if err := handler.tryUnspend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	ledgerID := handler.ledger.ID
	h.ID = &ledgerID

	return h.GetLedger(ctx)
}
