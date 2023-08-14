package unsoldstatement

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entunsoldstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/unsoldstatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.UnsoldStatementSelect
	infos     []*npool.UnsoldStatement
	total     uint32
}

func (h *queryHandler) selectUnsoldStatement(stm *ent.UnsoldStatementQuery) {
	h.stmSelect = stm.Select(
		entunsoldstatement.FieldID,
		entunsoldstatement.FieldGoodID,
		entunsoldstatement.FieldCoinTypeID,
		entunsoldstatement.FieldAmount,
		entunsoldstatement.FieldBenefitDate,
		entunsoldstatement.FieldCreatedAt,
		entunsoldstatement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryUnsoldStatement(cli *ent.Client) {
	h.selectUnsoldStatement(
		cli.UnsoldStatement.
			Query().
			Where(
				entunsoldstatement.ID(*h.ID),
				entunsoldstatement.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryUnsoldStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.UnsoldStatement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectUnsoldStatement(stm)
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

func (h *Handler) GetUnsoldStatement(ctx context.Context) (*npool.UnsoldStatement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.UnsoldStatement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryUnsoldStatement(cli)
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

func (h *Handler) GetUnsoldStatementOnly(ctx context.Context) (*npool.UnsoldStatement, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUnsoldStatements(_ctx, cli); err != nil {
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
