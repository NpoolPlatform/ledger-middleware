package coupon

import (
	"context"
	"fmt"

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"

	statementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw/coupon"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entcouponwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/couponwithdraw"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
)

type updateHandler struct {
	*Handler
	couponwithdraw *ent.CouponWithdraw
}

func (h *updateHandler) checkCouponWithdraw(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			CouponWithdraw.
			Query().
			Where(
				entcouponwithdraw.ID(*h.ID),
				entcouponwithdraw.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.couponwithdraw = info
		return nil
	})
}

func (h *updateHandler) updateLedger(ctx context.Context, tx *ent.Tx) error {
	switch {
	case h.State == nil:
		fallthrough //nolint
	case h.State.String() == h.couponwithdraw.State:
		fallthrough //nolint
	case *h.State != types.WithdrawState_Approved:
		return nil
	}

	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(h.couponwithdraw.AppID),
			entledger.UserID(h.couponwithdraw.UserID),
			entledger.CoinTypeID(h.couponwithdraw.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	incoming := h.couponwithdraw.Amount
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming: &incoming,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateCouponWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.UpdateSet(
		tx.CouponWithdraw.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	if h.State == nil {
		return nil
	}
	if h.State.String() != types.WithdrawState_Approved.String() {
		return nil
	}

	ioExtra := fmt.Sprintf(
		`{"AppID":"%v","UserID":"%v","CouponWithdrawID":"%v","CouponID":"%v"}`,
		h.couponwithdraw.AppID,
		h.couponwithdraw.UserID,
		h.couponwithdraw.EntID,
		h.couponwithdraw.CouponID.String(),
	)
	ioType := types.IOType_Incoming
	ioSubType := types.IOSubType_RandomCouponCash
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:      &h.couponwithdraw.AppID,
			UserID:     &h.couponwithdraw.UserID,
			CoinTypeID: &h.couponwithdraw.CoinTypeID,
			Amount:     &h.couponwithdraw.Amount,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			IOExtra:    &ioExtra,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.checkCouponWithdraw(ctx); err != nil {
		return nil, err
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCouponWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLedger(ctx, tx); err != nil {
			return err
		}
		if err := handler.createStatement(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCouponWithdraw(ctx)
}
