package ledger

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
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

type subHandler struct {
	*Handler
	ledger *ent.Ledger
	origin *ent.Statement
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
	ioType := ledgerpb.IOType_Incoming
	ioExtra := fmt.Sprintf(`{"StatementID": "%v", "Rollback": "true"}`, h.origin.ID.String())
	if _, err := cli.
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
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("statement already exist")
		}
		return err
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

	type extra struct {
		SubID string
	}
	e := extra{}
	if err := json.Unmarshal([]byte(*h.IOExtra), &e); err != nil {
		logger.Sugar().Errorf("need sub id in extra")
		return err
	}

	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		ioType := ledgerpb.IOType_Outcoming
		origin, err := cli.
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
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		h.origin = origin

		if err := h.getRollbackStatement(ctx, cli); err != nil {
			return err
		}
		// can create statement, set h.origin = nil
		h.origin = nil
		return nil
	})
}

func (h *subHandler) tryLock(ctx context.Context, tx *ent.Tx) error {
	if h.Spendable == nil {
		return nil
	}

	spendable := decimal.NewFromInt(0).Sub(*h.Spendable)
	locked := *h.Spendable

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

func (h *subHandler) trySpend(ctx context.Context, tx *ent.Tx) error {
	if h.Locked == nil {
		return nil
	}
	if h.origin != nil {
		return fmt.Errorf("statement already exist")
	}

	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithChangeLedger(),
	)
	if err != nil {
		return err
	}

	ioType := ledgerpb.IOType_Outcoming
	handler.Req = statementcrud.Req{
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

	ledgerID := handler.ledger.ID
	h.ID = &ledgerID

	return h.GetLedger(ctx)
}
