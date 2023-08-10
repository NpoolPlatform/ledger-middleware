package ledger

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.LedgerSelect
	infos     []*npool.Ledger
	total     uint32
}

func (h *queryHandler) selectLedger(stm *ent.LedgerQuery) {
	h.stmSelect = stm.Select(
		entledger.FieldID,
		entledger.FieldAppID,
		entledger.FieldUserID,
		entledger.FieldCoinTypeID,
		entledger.FieldIncoming,
		entledger.FieldOutcoming,
		entledger.FieldLocked,
		entledger.FieldSpendable,
		entledger.FieldCreatedAt,
		entledger.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryLedger(cli *ent.Client) {
	h.selectLedger(
		cli.Ledger.
			Query().
			Where(
				entledger.ID(*h.ID),
				entledger.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryLedgers(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.Ledger.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectLedger(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		incoming, err := decimal.NewFromString(info.Incoming)
		if err != nil {
			info.Incoming = decimal.NewFromInt(0).String()
		} else {
			info.Incoming = incoming.String()
		}

		outcoming, err := decimal.NewFromString(info.Outcoming)
		if err != nil {
			info.Outcoming = decimal.NewFromInt(0).String()
		} else {
			info.Outcoming = outcoming.String()
		}

		locked, err := decimal.NewFromString(info.Locked)
		if err != nil {
			info.Locked = decimal.NewFromInt(0).String()
		} else {
			info.Locked = locked.String()
		}

		spendable, err := decimal.NewFromString(info.Spendable)
		if err != nil {
			info.Spendable = decimal.NewFromInt(0).String()
		} else {
			info.Spendable = spendable.String()
		}
	}
}

func (h *Handler) GetLedger(ctx context.Context) (*npool.Ledger, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Ledger{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryLedger(cli)
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

func (h *Handler) GetLedgers(ctx context.Context) ([]*npool.Ledger, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Ledger{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLedgers(ctx, cli); err != nil {
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

func (h *Handler) GetLedgerOnly(ctx context.Context) (*npool.Ledger, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLedgers(_ctx, cli); err != nil {
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