package coupon

import (
	"context"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw/coupon"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"github.com/google/uuid"
)

func (h *Handler) CreateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.CouponWithdraw.Create(),
			&h.Req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCouponWithdraw(ctx)
}
