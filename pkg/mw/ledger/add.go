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
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type addHandler struct {
	*Handler
	ledger    *ent.Ledger
	statement *ent.Statement
	rollback  *ent.Statement
}

func (h *addHandler) validate() error {
	if h.Spendable != nil && h.Locked != nil {
		return fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable != nil {
		if h.AppID == nil || h.UserID == nil || h.CoinTypeID == nil {
			return fmt.Errorf("invalid appid or userid or cointypeid")
		}
	}
	return nil
}

func (h *addHandler) getStatement(ctx context.Context, cli *ent.Client) error {
	info, err := cli.
		Statement.
		Query().
		Where(
			entstatement.ID(*h.StatementID),
			entstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	if info.IoType != types.IOType_Outcoming.String() {
		return fmt.Errorf("io type should be outcoming")
	}
	h.statement = info
	return nil
}

func (h *addHandler) getRollbackStatement(ctx context.Context) error {
	if h.Spendable != nil {
		return nil
	}
	if h.StatementID == nil {
		return fmt.Errorf("invalid statement id")
	}
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if err := h.getStatement(ctx, cli); err != nil {
			return err
		}
		info, err := cli.
			Statement.
			Query().
			Where(
				entstatement.IoType(types.IOType_Incoming.String()),
				entstatement.IoExtra(getStatementExtra(h.StatementID.String())),
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

func (h *addHandler) tryUnlock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}
	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(*h.AppID),
			entledger.UserID(*h.UserID),
			entledger.CoinTypeID(*h.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	h.ledger = info

	spendable := *h.Spendable
	locked := decimal.NewFromInt(0).Sub(*h.Spendable)

	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
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

func (h *addHandler) tryUnspend(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable != nil {
		return nil
	}

	ioType := types.IOType_Incoming
	ioSubType := types.IOSubType(types.IOSubType_value[h.statement.IoSubType])
	ioExtra := getStatementExtra(h.StatementID.String())
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:      &h.statement.AppID,
			UserID:     &h.statement.UserID,
			CoinTypeID: &h.statement.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			IOExtra:    &ioExtra,
			Amount:     &h.statement.Amount,
		},
	).Save(ctx); err != nil {
		return err
	}

	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(h.statement.AppID),
			entledger.UserID(h.statement.UserID),
			entledger.CoinTypeID(h.statement.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	h.ledger = info

	outcoming := decimal.NewFromInt(0).Sub(h.statement.Amount)
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Locked:    &h.statement.Amount,
			Outcoming: &outcoming,
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

func (h *Handler) AddBalance(ctx context.Context) (*ledgermwpb.Ledger, error) {
	handler := &addHandler{
		Handler: h,
	}
	if err := handler.validate(); err != nil {
		return nil, err
	}
	if err := handler.getRollbackStatement(ctx); err != nil {
		return nil, err
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

	h.ID = &handler.ledger.ID

	return h.GetLedger(ctx)
}
