package goodstatement

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodstatement"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.GoodStatementSelect
	infos     []*npool.GoodStatement
	total     uint32
}

func (h *queryHandler) selectGoodStatement(stm *ent.GoodStatementQuery) {
	h.stmSelect = stm.Select(
		entgoodstatement.FieldID,
		entgoodstatement.FieldGoodID,
		entgoodstatement.FieldCoinTypeID,
		entgoodstatement.FieldAmount,
		entgoodstatement.FieldBenefitDate,
		entgoodstatement.FieldCreatedAt,
		entgoodstatement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryGoodStatement(cli *ent.Client) {
	h.selectGoodStatement(
		cli.GoodStatement.
			Query().
			Where(
				entgoodstatement.ID(*h.ID),
				entgoodstatement.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryGoodStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.GoodStatement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectGoodStatement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount := decimal.NewFromInt(0).String()
		if _amount, err := decimal.NewFromString(info.Amount); err == nil {
			amount = _amount.String()
		}
		info.Amount = amount
	}
}

func (h *Handler) GetGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.GoodStatement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryGoodStatement(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetGoodStatements(ctx context.Context) ([]*npool.GoodStatement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.GoodStatement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodStatements(ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}


func (h *Handler) GetGoodStatementOnly(ctx context.Context) (*npool.GoodStatement, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodStatements(_ctx, cli); err != nil {
			return err
		}

		_, err := handler.stmSelect.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("to many record")
	}

	return handler.infos[0], nil
}