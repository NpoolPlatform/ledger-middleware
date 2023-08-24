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
	rollback *ent.Statement
}

func (h *addHandler) getLedger(ctx context.Context) error {
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Ledger.
			Query().
			Where(
				entledger.AppID(*h.AppID),
				entledger.UserID(*h.UserID),
				entledger.CoinTypeID(*h.CoinTypeID),
				entledger.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.ledger = info
		return nil
	})
	return err
}

func (h *addHandler) tryUnlock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}

	spendable := *h.Spendable
	locked := decimal.NewFromInt(0).Sub(*h.Spendable)

	stm, err := ledgercrud.UpdateSetWithValidate(
		h.ledger,
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

func (h *addHandler) getStatement(ctx context.Context, cli *ent.Client) error {
	if h.IOSubType == nil {
		return fmt.Errorf("invalid io sub type")
	}
	if h.IOExtra == nil {
		return fmt.Errorf("invalid io extra")
	}

	// get statement
	ioType := ledgerpb.IOType_Outcoming
	statement, err := cli.
		Statement.
		Query().
		Where(
			entstatement.ID(*h.StatementID),
			entstatement.IoType(ioType.String()),
			entstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("statement not exist")
		}
		return err
	}

	if statement.Amount.Cmp(*h.Locked) != 0 {
		return fmt.Errorf("rollback amount(%v) not match spend amount(%v)", statement.Amount.String(), h.Locked.String())
	}
	return nil
}

func (h *addHandler) getRollbackStatement(ctx context.Context) error {
	if h.Locked == nil {
		return nil
	}
	if h.StatementID == nil {
		return fmt.Errorf("invalid statement id")
	}
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		// get statement
		if err := h.getStatement(ctx, cli); err != nil {
			return err
		}
		// get rollback statement
		ioType := ledgerpb.IOType_Incoming
		ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, h.StatementID.String())
		info, err := cli.
			Statement.
			Query().
			Where(
				entstatement.AppID(*h.AppID),
				entstatement.UserID(*h.UserID),
				entstatement.CoinTypeID(*h.CoinTypeID),
				entstatement.IoType(ioType.String()),
				entstatement.IoSubType(h.IOSubType.String()),
				entstatement.IoExtra(ioExtra),
				entstatement.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		h.rollback = info
		return nil
	})
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
	ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, h.StatementID.String())
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

	stm, err := ledgercrud.UpdateSetWithValidate(
		h.ledger,
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

	if err := handler.getLedger(ctx); err != nil {
		return nil, err
	}
	if err := handler.getRollbackStatement(ctx); err != nil {
		return nil, err
	}
	if h.Locked != nil && handler.rollback != nil {
		return nil, fmt.Errorf("statement already rolled back")
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.tryUnlock(ctx, tx); err != nil {
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
