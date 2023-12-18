package coupon

import (
	"context"
	"time"

	couponwithdrawcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw/coupon"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteWithdraw(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := couponwithdrawcrud.UpdateSet(
			cli.CouponWithdraw.UpdateOneID(*h.ID),
			&couponwithdrawcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) DeleteCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	info, err := h.GetCouponWithdraw(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	handler := &deleteHandler{
		Handler: h,
	}
	if err := handler.deleteWithdraw(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
