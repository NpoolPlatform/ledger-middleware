package statement

import (
	"context"
	"fmt"
	"time"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	profitcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/profit"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) updateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		Statement.
		Query().
		Where(
			entstatement.ID(*req.ID),
			entstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	if statement.IoSubType != basetypes.IOSubType_MiningBenefit.String() {
		return nil
	}

	stm, err := profitcrud.SetQueryConds(
		tx.Profit.Query(),
		&profitcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: statement.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: statement.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: statement.CoinTypeID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	amount := decimal.NewFromInt(0).Sub(statement.Amount)
	stm1, err := profitcrud.UpdateSetWithValidate(
		info,
		&profitcrud.Req{
			Incoming: &amount,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) updateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		Statement.
		Query().
		Where(
			entstatement.ID(*req.ID),
			entstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	stm, err := ledgercrud.SetQueryConds(
		tx.Ledger.Query(),
		&ledgercrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: statement.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: statement.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: statement.CoinTypeID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	incoming := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	ioType := basetypes.IOType(basetypes.IOType_value[statement.IoType])
	switch ioType {
	case basetypes.IOType_Incoming:
		incoming = decimal.NewFromInt(0).Sub(statement.Amount)
	case basetypes.IOType_Outcoming:
		outcoming = decimal.NewFromInt(0).Sub(statement.Amount)
	default:
		return fmt.Errorf("invalid io type %v", statement.IoType)
	}
	spendable := incoming.Sub(outcoming)

	stm1, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming:  &incoming,
			Outcoming: &outcoming,
			Spendable: &spendable,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *deleteHandler) deleteStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Statement.
		Query().
		Where(
			entstatement.ID(*req.ID),
			entstatement.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	now := uint32(time.Now().Unix())
	if _, err := crud.UpdateSet(
		info.Update(),
		&crud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteStatements(ctx context.Context) ([]*npool.Statement, error) {
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return nil, fmt.Errorf("invalid statement id")
		}
		ids = append(ids, *req.ID)
	}

	h.Conds = &crud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Limit = int32(len(ids))
	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) != len(h.Reqs) {
		if len(h.Reqs) > 0 && h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		if h.Rollback == nil {
			return nil, fmt.Errorf("statement not found")
		}
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				if err := handler.updateProfit(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.updateLedger(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.deleteStatement(req, ctx, tx); err != nil {
					return err
				}
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (h *Handler) DeleteStatement(ctx context.Context) (*npool.Statement, error) {
	info, err := h.GetStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		if h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, fmt.Errorf("statement not found")
	}

	h.Reqs = []*crud.Req{&h.Req}
	if _, err := h.DeleteStatements(ctx); err != nil {
		return nil, err
	}

	return info, nil
}
