package goodledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/goodledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/goodledger"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.GoodLedgerSelect
	infos     []*npool.GoodLedger
	total     uint32
}

func (h *queryHandler) selectGoodLedger(stm *ent.GoodLedgerQuery) {
	h.stmSelect = stm.Select(
		entgoodledger.FieldID,
		entgoodledger.FieldGoodID,
		entgoodledger.FieldCoinTypeID,
		entgoodledger.FieldAmount,
		entgoodledger.FieldToPlatform,
		entgoodledger.FieldToUser,
		entgoodledger.FieldCreatedAt,
		entgoodledger.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryGoodLedger(cli *ent.Client) {
	h.selectGoodLedger(
		cli.GoodLedger.
			Query().
			Where(
				entgoodledger.ID(*h.ID),
				entgoodledger.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryGoodLedgers(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.GoodLedger.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectGoodLedger(stm)
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

func (h *Handler) GetGoodLedger(ctx context.Context) (*npool.GoodLedger, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.GoodLedger{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryGoodLedger(cli)
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

func (h *Handler) GetGoodLedgers(ctx context.Context) ([]*npool.GoodLedger, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.GoodLedger{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodLedgers(ctx, cli); err != nil {
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
