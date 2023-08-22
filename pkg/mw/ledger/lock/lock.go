package lock

import (
	"context"
	"fmt"
	"time"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*Handler
}

func (h *lockHandler) setConds() *statementcrud.Conds {
	conds := &statementcrud.Conds{}
	if h.AppID != nil {
		conds.AppID = &cruder.Cond{Op: cruder.EQ, Val: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &cruder.Cond{Op: cruder.EQ, Val: *h.UserID}
	}
	if h.CoinTypeID != nil {
		conds.CoinTypeID = &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID}
	}
	if h.IOSubType != nil {
		conds.IOSubType = &cruder.Cond{Op: cruder.EQ, Val: *h.IOSubType}
	}
	if h.IOExtra != nil {
		conds.IOExtra = &cruder.Cond{Op: cruder.LIKE, Val: *h.IOExtra}
	}
	ioType := basetypes.IOType_Outcoming
	conds.IOType = &cruder.Cond{Op: cruder.EQ, Val: ioType}
	return conds
}

func (h *lockHandler) tryCreateStatement(ctx context.Context, tx *ent.Tx) error {
	ioType := basetypes.IOType_Outcoming

	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  h.IOSubType,
			Amount:     h.Outcoming,
			IOExtra:    h.IOExtra,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockHandler) tryDeleteStatement(ctx context.Context, tx *ent.Tx) error {
	stm, err := statementcrud.SetQueryConds(
		tx.Statement.Query(),
		h.setConds(),
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := statementcrud.UpdateSet(
		tx.Statement.UpdateOneID(info.ID),
		&statementcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockHandler) tryGetStatement(ctx context.Context, tx *ent.Tx) (bool, error) {
	stm, err := statementcrud.SetQueryConds(
		tx.Statement.Query(),
		h.setConds(),
	)
	if err != nil {
		return false, err
	}
	exist, err := stm.Exist(ctx)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *lockHandler) tryUpdateLedger(req ledgercrud.Req, ctx context.Context, tx *ent.Tx) (*ledgerpb.Ledger, error) {
	stm, err := ledgercrud.SetQueryConds(tx.Ledger.Query(), &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
	})
	if err != nil {
		return nil, err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("ledger not exist, AppID: %v, UserID: %v, CoinTypeID: %v", *req.AppID, *req.UserID, *req.CoinTypeID)
		}
		return nil, err
	}

	// update
	old, err := tx.Ledger.Get(ctx, info.ID)
	if err != nil {
		return nil, err
	}
	if old == nil {
		return nil, fmt.Errorf("ledger not exist, id %v", info.ID)
	}

	stm1, err := ledgercrud.UpdateSet(
		old,
		tx.Ledger.UpdateOneID(info.ID),
		&ledgercrud.Req{
			Outcoming: req.Outcoming,
			Spendable: req.Spendable,
			Locked:    req.Locked,
		},
	)
	if err != nil {
		return nil, err
	}
	info1, err := stm1.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ledgerpb.Ledger{
		ID:         info.ID.String(),
		AppID:      info.AppID.String(),
		UserID:     info.UserID.String(),
		CoinTypeID: info.CoinTypeID.String(),
		Outcoming:  info1.Outcoming.String(),
		Incoming:   info1.Incoming.String(),
		Spendable:  info1.Spendable.String(),
		Locked:     info1.Locked.String(),
	}, nil
}

// Lock & Spend
func (h *Handler) SubBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	if h.Spendable != nil && h.Locked != nil {
		return nil, fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable == nil && h.Locked == nil {
		return nil, fmt.Errorf("spendable or locked needed")
	}

	handler := &lockHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		locked := decimal.NewFromInt(0)
		spendable := decimal.NewFromInt(0)
		outcoming := decimal.NewFromInt(0)

		if h.Locked != nil { // lock
			locked = *h.Locked
			spendable = decimal.NewFromInt(0).Sub(*h.Locked)
		}

		if h.Spendable != nil { // spend
			locked = decimal.NewFromInt(0).Sub(*h.Spendable)
			outcoming = *h.Spendable
		}

		if h.Spendable != nil {
			if h.IOSubType == nil {
				return fmt.Errorf("invalid io sub type")
			}
			if h.IOExtra != nil {
				return fmt.Errorf("invalid io extra")
			}

			exist, err := handler.tryGetStatement(ctx, tx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("statement already exist")
			}

			if err := handler.tryCreateStatement(ctx, tx); err != nil {
				return err
			}
		}
		info, err = handler.tryUpdateLedger(ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Spendable:  &spendable,
			Outcoming:  &outcoming,
		}, ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, err
}

// Unlock & Unspend
func (h *Handler) AddBalance(ctx context.Context) (info *ledgerpb.Ledger, err error) {
	if h.Spendable != nil && h.Locked != nil {
		return nil, fmt.Errorf("spendable & locked is not allowed")
	}
	if h.Spendable == nil && h.Locked == nil {
		return nil, fmt.Errorf("spendable or locked needed")
	}

	handler := &lockHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		locked := decimal.NewFromInt(0)
		spendable := decimal.NewFromInt(0)
		outcoming := decimal.NewFromInt(0)

		if h.Locked != nil { // unlock
			locked = decimal.NewFromInt(0).Sub(*h.Locked)
			spendable = *h.Locked
		}
		if h.Spendable != nil { // unspend
			locked = *h.Spendable
			spendable = decimal.NewFromInt(0).Sub(*h.Spendable)
		}

		if h.Spendable != nil {
			if h.IOSubType == nil {
				return fmt.Errorf("invalid io sub type")
			}
			if h.IOExtra != nil {
				return fmt.Errorf("invalid io extra")
			}
			exist, err := handler.tryGetStatement(ctx, tx)
			if err != nil {
				return err
			}
			if exist {
				return nil
			}
			if err := handler.tryDeleteStatement(ctx, tx); err != nil {
				return err
			}
		}

		info, err = handler.tryUpdateLedger(ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &locked,
			Spendable:  &spendable,
			Outcoming:  &outcoming,
		}, ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, err
}
