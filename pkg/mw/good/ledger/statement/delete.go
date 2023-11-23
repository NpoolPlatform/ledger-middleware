package statement

import (
	"context"
	"fmt"
	"time"

	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger"
	goodstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/statement"
	unsoldcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/unsold"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) updateGoodLedger(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, err := tx.
		GoodStatement.
		Query().
		Where(
			entgoodstatement.ID(*req.ID),
			entgoodstatement.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	stm, err := goodledgercrud.SetQueryConds(
		tx.GoodLedger.Query(),
		&goodledgercrud.Conds{
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: statement.GoodID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: statement.CoinTypeID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	amount := decimal.NewFromInt(0).Sub(statement.Amount)
	toUser := decimal.NewFromInt(0).Sub(statement.ToUser)
	toPlatform := decimal.NewFromInt(0).Sub(statement.ToPlatform)

	stm1, err := goodledgercrud.UpdateSetWithValidate(
		info,
		&goodledgercrud.Req{
			Amount:     &amount,
			ToUser:     &toUser,
			ToPlatform: &toPlatform,
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

//nolint
func (h *deleteHandler) deleteGoodStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		GoodStatement.
		Query().
		Where(
			entgoodstatement.ID(*req.ID),
			entgoodstatement.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	now := uint32(time.Now().Unix())
	if _, err := goodstatementcrud.UpdateSet(
		info.Update(),
		&goodstatementcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *deleteHandler) deleteUnsoldStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		UnsoldStatement.
		Query().
		Where(
			unsoldstatement.StatementID(*req.EntID),
			unsoldstatement.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	now := uint32(time.Now().Unix())
	if _, err := unsoldcrud.UpdateSet(
		info.Update(),
		&unsoldcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	handler := &deleteHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		ids = append(ids, *req.EntID)
	}
	h.Conds = &goodstatementcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Limit = int32(len(ids))
	infos, _, err := h.GetGoodStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) != len(h.Reqs) {
		if h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, fmt.Errorf("good statement not found")
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				if err := handler.deleteUnsoldStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.updateGoodLedger(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.deleteGoodStatement(req, ctx, tx); err != nil {
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

func (h *Handler) DeleteGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	info, err := h.GetGoodStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		if h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, fmt.Errorf("statement not found")
	}
	if h.ID == nil {
		h.ID = &info.ID
	}
	if h.EntID == nil {
		id, err := uuid.Parse(info.EntID)
		if err != nil {
			return nil, err
		}
		h.EntID = &id
	}
	h.Reqs = []*goodstatementcrud.Req{&h.Req}

	if _, err := h.DeleteGoodStatements(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
