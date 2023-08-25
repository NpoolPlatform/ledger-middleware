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
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type subHandler struct {
	*Handler
	ledger    *ent.Ledger
	statement *ent.Statement
}

func (h *subHandler) getLedger(ctx context.Context) error {
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

func (h *subHandler) getRollbackStatement(ctx context.Context, cli *ent.Client) error {
	if _, err := cli.
		Statement.
		Query().
		Where(
			entstatement.AppID(*h.AppID),
			entstatement.UserID(*h.UserID),
			entstatement.CoinTypeID(*h.CoinTypeID),
			entstatement.IoType(types.IOType_Incoming.String()),
			entstatement.IoSubType(h.IOSubType.String()),
			entstatement.IoExtra(getStatementExtra(h.StatementID.String())),
			entstatement.DeletedAt(0),
		).
		Only(ctx); err != nil {
		return err
	}
	return nil
}

func (h *subHandler) getStatement(ctx context.Context) error {
	if h.Locked == nil {
		return nil
	}
	if h.StatementID != nil {
		return fmt.Errorf("invalid statement id")
	}

	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Statement.
			Query().
			Where(
				entstatement.ID(*h.StatementID),
				entstatement.IoType(types.IOType_Outcoming.String()),
				entstatement.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		h.statement = info

		if err := h.getRollbackStatement(ctx, cli); err != nil {
			return err
		}
		h.statement = nil
		return nil
	})
}

func (h *subHandler) tryLock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}

	spendable := decimal.NewFromInt(0).Sub(*h.Spendable)
	locked := *h.Spendable

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

func (h *subHandler) trySpend(ctx context.Context, tx *ent.Tx) error {
	if h.Locked == nil {
		return nil
	}

	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithChangeLedger(false),
	)
	if err != nil {
		return err
	}

	ioType := types.IOType_Outcoming
	handler.Req = statementcrud.Req{
		ID:         h.StatementID,
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
		Amount:     h.Locked,
	}
	if _, err := handler.CreateStatement(ctx); err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(*h.Locked)
	outcoming := *h.Locked

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

// Lock & Spend
func (h *Handler) SubBalance(ctx context.Context) (info *ledgermwpb.Ledger, err error) {
	if err := h.validate(); err != nil {
		return nil, err
	}

	handler := &subHandler{
		Handler: h,
	}
	if err := handler.getLedger(ctx); err != nil {
		return nil, err
	}
	if err := handler.getStatement(ctx); err != nil {
		return nil, err
	}
	if handler.Locked != nil && handler.statement != nil {
		return nil, fmt.Errorf("statement already exist")
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.tryLock(ctx, tx); err != nil {
			return err
		}
		if err := handler.trySpend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.ID = &handler.ledger.ID
	return h.GetLedger(ctx)
}
