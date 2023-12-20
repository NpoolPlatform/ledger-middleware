package coupon

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw/coupon"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ledgertypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"github.com/google/uuid"
)

func (h *Handler) CreateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	h.Conds = &crud.Conds{
		AppID:    &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:   &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		CouponID: &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		State:    &cruder.Cond{Op: cruder.EQ, Val: ledgertypes.WithdrawState_Reviewing},
	}
	exist, err := h.ExistCouponWithdrawConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coupon withdraw already exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
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
