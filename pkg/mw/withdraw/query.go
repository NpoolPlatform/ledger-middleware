package withdraw

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.WithdrawSelect
	infos     []*npool.Withdraw
	total     uint32
}

func (h *queryHandler) selectWithdraw(stm *ent.WithdrawQuery) {
	h.stmSelect = stm.Select(
		entwithdraw.FieldID,
		entwithdraw.FieldAppID,
		entwithdraw.FieldUserID,
		entwithdraw.FieldCoinTypeID,
		entwithdraw.FieldAccountID,
		entwithdraw.FieldChainTransactionID,
		entwithdraw.FieldPlatformTransactionID,
		entwithdraw.FieldAddress,
		entwithdraw.FieldAmount,
		entwithdraw.FieldCreatedAt,
		entwithdraw.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryWithdraw(cli *ent.Client) {
	h.selectWithdraw(
		cli.Withdraw.
			Query().
			Where(
				entwithdraw.ID(*h.ID),
				entwithdraw.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryWithdraws(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.Withdraw.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectWithdraw(stm)
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
		info.State = basetypes.WithdrawState(basetypes.WithdrawState_value[info.StateStr])
	}
}

func (h *Handler) GetWithdraw(ctx context.Context) (*npool.Withdraw, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Withdraw{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryWithdraw(cli)
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

func (h *Handler) GetWithdraws(ctx context.Context) ([]*npool.Withdraw, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Withdraw{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryWithdraws(ctx, cli); err != nil {
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

func (h *Handler) GetWithdrawOnly(ctx context.Context) (*npool.Withdraw, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryWithdraws(_ctx, cli); err != nil {
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
