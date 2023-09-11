package ledger

import (
	"context"
	"fmt"
	"time"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	ledgerlockcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/lock"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	entledgerlock "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledgerlock"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type subHandler struct {
	*Handler
	ledger    *ent.Ledger
	statement *ent.Statement
}

func (h *subHandler) validate() error {
	if h.Spendable != nil && h.Locked != nil {
		return fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable == nil && h.Locked == nil {
		return fmt.Errorf("spendable or locked needed")
	}
	if h.Spendable != nil {
		if h.AppID == nil || h.UserID == nil || h.CoinTypeID == nil {
			return fmt.Errorf("invalid appid or userid or cointypeid")
		}
	}
	return nil
}

func (h *subHandler) getStatement(ctx context.Context) error {
	if h.Locked == nil {
		return nil
	}
	if h.IOSubType == nil {
		return fmt.Errorf("invalid io sub type")
	}
	if h.IOExtra == nil {
		return fmt.Errorf("invalid io extra")
	}
	if h.StatementID == nil {
		return fmt.Errorf("invalid statement id")
	}
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Statement.
			Query().
			Where(
				entstatement.ID(*h.StatementID),
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
		return nil
	})
}

func (h *subHandler) createLedgerLock(ctx context.Context, tx *ent.Tx) error {
	if _, err := ledgerlockcrud.CreateSet(
		tx.LedgerLock.Create(),
		&ledgerlockcrud.Req{
			ID:     h.LockID,
			Amount: h.Spendable,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *subHandler) deleteLedgerLock(ctx context.Context, tx *ent.Tx) (bool, error) {
	lock, err := tx.
		LedgerLock.
		Query().
		Where(
			entledgerlock.ID(*h.LockID),
			entledgerlock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	if h.Locked.Cmp(lock.Amount) != 0 {
		return false, fmt.Errorf("invalid amount")
	}

	now := uint32(time.Now().Unix())
	if _, err := ledgerlockcrud.UpdateSet(lock.Update(), &ledgerlockcrud.Req{
		DeletedAt: &now,
	}).Save(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (h *subHandler) tryLock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}

	if err := h.createLedgerLock(ctx, tx); err != nil {
		return err
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

	spendable := decimal.NewFromInt(0).Sub(*h.Spendable)
	locked := *h.Spendable

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

func (h *subHandler) trySpend(ctx context.Context, tx *ent.Tx) error {
	if h.Locked == nil {
		return nil
	}

	if deleted, err := h.deleteLedgerLock(ctx, tx); err != nil || !deleted {
		return err
	}

	ioType := types.IOType_Outcoming
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			ID:         h.StatementID,
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  h.IOSubType,
			IOExtra:    h.IOExtra,
			Amount:     h.Locked,
		},
	).Save(ctx); err != nil {
		return err
	}

	locked := decimal.NewFromInt(0).Sub(*h.Locked)
	outcoming := *h.Locked

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

func (h *Handler) SubBalance(ctx context.Context) (info *ledgermwpb.Ledger, err error) {
	handler := &subHandler{
		Handler: h,
	}
	if err := handler.validate(); err != nil {
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
