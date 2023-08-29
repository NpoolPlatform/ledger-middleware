package statement

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger"
	goodstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/statement"
	unsoldcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/unsold"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type rollbackHandler struct {
	*Handler
	statementsMap map[string]*npool.GoodStatement
}

func (h *rollbackHandler) tryUpdateGoodLedger(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, _ := h.statementsMap[req.ID.String()] //nolint
	stm, err := goodledgercrud.SetQueryConds(
		tx.GoodLedger.Query(),
		&goodledgercrud.Conds{
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: uuid.MustParse(statement.GoodID)},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: uuid.MustParse(statement.CoinTypeID)},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	amount := decimal.RequireFromString(fmt.Sprintf("-%v", statement.Amount))
	toUser := decimal.RequireFromString(fmt.Sprintf("-%v", statement.ToUser))
	toPlatform := decimal.RequireFromString(fmt.Sprintf("-%v", statement.ToPlatform))

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

func (h *rollbackHandler) tryDeleteGoodStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	key := fmt.Sprintf("%v:%v",
		commonpb.Prefix_PrefixDeleteGoodStatement,
		*req.ID,
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()
	now := uint32(time.Now().Unix())
	if _, err := goodstatementcrud.UpdateSet(
		tx.GoodStatement.UpdateOneID(*req.ID),
		&goodstatementcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *rollbackHandler) tryDeleteUnsoldStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		UnsoldStatement.
		Query().
		Where(
			unsoldstatement.StatementID(*req.ID),
			unsoldstatement.DeletedAt(0),
		).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := unsoldcrud.UpdateSet(
		tx.UnsoldStatement.UpdateOneID(info.ID),
		&unsoldcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *rollbackHandler) tryGetAllGoodStatements(ctx context.Context) error {
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return fmt.Errorf("invalid good statement id")
		}
		ids = append(ids, *req.ID)
	}

	h.Conds = &goodstatementcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Limit = int32(len(ids))
	infos, _, err := h.GetGoodStatements(ctx)
	if err != nil {
		return err
	}

	h.statementsMap = map[string]*npool.GoodStatement{}
	for _, info := range infos {
		h.statementsMap[info.ID] = info
	}
	return nil
}

func (h *Handler) DeleteGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	handler := &rollbackHandler{
		Handler: h,
	}
	if err := handler.tryGetAllGoodStatements(ctx); err != nil {
		return nil, err
	}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				_, ok := handler.statementsMap[req.ID.String()]
				if !ok {
					return fmt.Errorf("good statement not found %v", req.ID.String())
				}
				if err := handler.tryDeleteGoodStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryDeleteUnsoldStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryUpdateGoodLedger(req, ctx, tx); err != nil {
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

	infos := []*npool.GoodStatement{}
	for _, value := range handler.statementsMap {
		infos = append(infos, value)
	}
	return infos, nil
}

func (h *Handler) DeleteGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	h.Reqs = append(h.Reqs, &h.Req)

	infos, err := h.DeleteGoodStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos[0], nil
}
