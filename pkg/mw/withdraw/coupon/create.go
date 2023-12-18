package coupon

import (
	"context"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw/coupon"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCouponWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.CreateSet(
		tx.CouponWithdraw.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.createCouponWithdraw(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCouponWithdraw(ctx)
}
