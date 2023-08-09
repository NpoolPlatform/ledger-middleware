package statement

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.StatementSelect
	infos     []*npool.Statement
	total     uint32
}

func (h *queryHandler) selectStatement(stm *ent.StatementQuery) {
	h.stmSelect = stm.Select(
		entstatement.FieldID,
		entstatement.FieldAppID,
		entstatement.FieldUserID,
		entstatement.FieldCoinTypeID,
		entstatement.FieldIoType,
		entstatement.FieldIoSubType,
		entstatement.FieldAmount,
		entstatement.FieldFromCoinTypeID,
		entstatement.FieldCoinUsdCurrency,
		entstatement.FieldIoExtra,
		entstatement.FieldCreatedAt,
		entstatement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryStatement(cli *ent.Client) {
	h.selectStatement(
		cli.Statement.
			Query().
			Where(
				entstatement.ID(*h.ID),
				entstatement.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.Statement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectStatement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		currency, err := decimal.NewFromString(info.CoinUSDCurrency)
		if err != nil {
			info.CoinUSDCurrency = decimal.NewFromInt(0).String()
		} else {
			info.CoinUSDCurrency = currency.String()
		}

		amount, err := decimal.NewFromString(info.Amount)
		if err != nil {
			info.Amount = decimal.NewFromInt(0).String()
		} else {
			info.Amount = amount.String()
		}

		info.IOType = basetypes.IOType(basetypes.IOType_value[info.IOTypeStr])
		info.IOSubType = basetypes.IOSubType(basetypes.IOSubType_value[info.IOSubTypeStr])
	}
}

func (h *Handler) GetStatement(ctx context.Context) (*npool.Statement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryStatement(cli)
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

func (h *Handler) GetStatements(ctx context.Context) ([]*npool.Statement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatements(ctx, cli); err != nil {
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
