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
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type rollbackHandler struct {
	*Handler
	statementsMap map[string]*npool.Statement
}

func (h *rollbackHandler) tryGetAllStatements(ctx context.Context) error {
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return fmt.Errorf("invalid statement id")
		}
		ids = append(ids, *req.ID)
	}

	h.Conds = &crud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Limit = int32(len(ids))

	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return err
	}

	h.statementsMap = map[string]*npool.Statement{}
	for _, info := range infos {
		h.statementsMap[info.ID] = info
	}
	return nil
}

func (h *rollbackHandler) tryUpdateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if *req.IOSubType != basetypes.IOSubType_MiningBenefit {
		return nil
	}

	stm, err := profitcrud.SetQueryConds(
		tx.Profit.Query(),
		&profitcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("profit not found")
	}

	amount := decimal.RequireFromString(fmt.Sprintf("-%v", req.Amount))
	stm1, err := profitcrud.UpdateSetWithValidate(
		info,
		&profitcrud.Req{
			AppID:      req.AppID,
			UserID:     req.UserID,
			CoinTypeID: req.CoinTypeID,
			Incoming:   &amount,
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

func (h *rollbackHandler) tryUpdateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := ledgercrud.SetQueryConds(
		tx.Ledger.Query(),
		&ledgercrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("ledger not found")
	}

	incoming := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *req.IOType {
	case basetypes.IOType_Incoming:
		incoming = decimal.RequireFromString(fmt.Sprintf("-%v", req.Amount.String()))
	case basetypes.IOType_Outcoming:
		outcoming = *req.Amount
	default:
		return fmt.Errorf("invalid io type %v", *req.IOType)
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

func (h *rollbackHandler) tryDeleteStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	statement, _ := h.statementsMap[req.ID.String()] //nolint
	if statement.Amount != req.Amount.String() {
		return fmt.Errorf("amount not matched")
	}
	now := uint32(time.Now().Unix())
	if _, err := crud.UpdateSet(
		tx.Statement.UpdateOneID(*req.ID),
		&crud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) RollbackStatements(ctx context.Context) ([]*npool.Statement, error) {
	handler := &rollbackHandler{
		Handler: h,
	}
	if err := handler.tryGetAllStatements(ctx); err != nil {
		return nil, err
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				_, ok := handler.statementsMap[req.ID.String()]
				if !ok {
					return fmt.Errorf("statement not found %v", req.ID.String())
				}
				if err := handler.tryDeleteStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryUpdateProfit(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryUpdateLedger(req, ctx, tx); err != nil {
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

	infos := []*npool.Statement{}
	for _, value := range handler.statementsMap {
		infos = append(infos, value)
	}
	return infos, nil
}

func (h *Handler) RollbackStatement(ctx context.Context) (*npool.Statement, error) {
	h.Reqs = append(h.Reqs, &h.Req)
	infos, err := h.RollbackStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("to many statements")
	}

	return infos[0], nil
}
