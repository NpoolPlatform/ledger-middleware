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
	unsold1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/unsold"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	ledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/unsold"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type rollbackHandler struct {
	*Handler
	statementsMap       map[string]*npool.GoodStatement
	unsoldstatementsMap map[string]*ledgermwpb.UnsoldStatement
}

func (h *rollbackHandler) tryUpdateGoodLedger(req *Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := goodledgercrud.SetQueryConds(tx.GoodLedger.Query(), &goodledgercrud.Conds{
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: *req.GoodID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
	})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	toPlatform := h.UnsoldAmount.Add(*h.TechniqueServiceFeeAmount)
	toUser := h.TotalAmount.Sub(toPlatform)
	if h.TotalAmount.Cmp(toPlatform.Add(toUser)) != 0 {
		return fmt.Errorf("TotalAmount(%v) != ToPlatform(%v) + ToUser(%v)", h.TotalAmount.String(), toPlatform.String(), toUser.String())
	}
	_amount := decimal.RequireFromString(fmt.Sprintf("-%v", req.TotalAmount.String()))
	_toUser := decimal.RequireFromString(fmt.Sprintf("-%v", toUser.String()))
	_toPlatform := decimal.RequireFromString(fmt.Sprintf("-%v", toPlatform.String()))

	stm1, err := goodledgercrud.UpdateSetWithValidate(
		info,
		&goodledgercrud.Req{
			Amount:     &_amount,
			ToUser:     &_toUser,
			ToPlatform: &_toPlatform,
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
func (h *rollbackHandler) tryDeleteGoodStatement(req *Req, ctx context.Context, tx *ent.Tx) error {
	statement, _ := h.statementsMap[req.ID.String()]
	if statement.Amount != req.TotalAmount.String() {
		return fmt.Errorf("total amount not matched")
	}

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

func (h *rollbackHandler) tryDeleteUnsoldStatement(req *Req, ctx context.Context, tx *ent.Tx) error {
	unsold, ok := h.unsoldstatementsMap[req.UnsoldStatementID.String()]
	if !ok {
		return nil
	}
	if unsold.Amount != req.UnsoldAmount.String() {
		return fmt.Errorf("unsold amount not matched")
	}

	now := uint32(time.Now().Unix())
	if _, err := unsoldcrud.UpdateSet(
		tx.UnsoldStatement.UpdateOneID(*req.UnsoldStatementID),
		&unsoldcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *rollbackHandler) tryGetAllUnsoldGoodStatements(ctx context.Context) error {
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.UnsoldStatementID == nil {
			return fmt.Errorf("invalid unsold good statement id")
		}
		ids = append(ids, *req.UnsoldStatementID)
	}

	handler, err := unsold1.NewHandler(ctx)
	if err != nil {
		return err
	}
	handler.Conds = &unsoldcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}

	infos, _, err := handler.GetUnsoldStatements(ctx)
	if err != nil {
		return err
	}

	for _, info := range infos {
		h.unsoldstatementsMap[info.ID] = info
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
	infos, _, err := h.GetGoodStatements(ctx)
	if err != nil {
		return err
	}

	for _, info := range infos {
		h.statementsMap[info.ID] = info
	}
	return nil
}

func (h *Handler) RollbackGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	handler := &rollbackHandler{
		Handler: h,
	}
	if err := handler.tryGetAllGoodStatements(ctx); err != nil {
		return nil, err
	}
	if err := handler.tryGetAllUnsoldGoodStatements(ctx); err != nil {
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

func (h *Handler) RollbackGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	h.Reqs = append(h.Reqs, h.Req)

	infos, err := h.RollbackGoodStatements(ctx)
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