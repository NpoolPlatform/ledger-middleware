package profit

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/profit"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entprofit "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.ProfitSelect
	infos     []*npool.Profit
	total     uint32
}

func (h *queryHandler) selectProfit(stm *ent.ProfitQuery) {
	h.stmSelect = stm.Select(
		entprofit.FieldID,
		entprofit.FieldAppID,
		entprofit.FieldUserID,
		entprofit.FieldCoinTypeID,
		entprofit.FieldIncoming,
		entprofit.FieldCreatedAt,
		entprofit.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryProfit(cli *ent.Client) {
	h.selectProfit(
		cli.Profit.
			Query().
			Where(
				entprofit.ID(*h.ID),
				entprofit.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryProfits(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.Profit.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectProfit(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		incoming := decimal.NewFromInt(0).String()
		if _incoming, err := decimal.NewFromString(info.Incoming); err == nil {
			incoming = _incoming.String()
		}
		info.Incoming = incoming
	}
}

func (h *Handler) GetProfit(ctx context.Context) (*npool.Profit, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Profit{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryProfit(cli)
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

func (h *Handler) GetProfits(ctx context.Context) ([]*npool.Profit, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Profit{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryProfits(ctx, cli); err != nil {
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
