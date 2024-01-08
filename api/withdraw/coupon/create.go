//nolint
package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	couponwithdraw1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/withdraw/coupon"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCouponWithdraw(ctx context.Context, in *npool.CreateCouponWithdrawRequest) (*npool.CreateCouponWithdrawResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCouponWithdraw",
			"In", in,
		)
		return &npool.CreateCouponWithdrawResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := couponwithdraw1.NewHandler(
		ctx,
		couponwithdraw1.WithEntID(req.EntID, false),
		couponwithdraw1.WithAppID(req.AppID, true),
		couponwithdraw1.WithUserID(req.UserID, true),
		couponwithdraw1.WithCoinTypeID(req.CoinTypeID, true),
		couponwithdraw1.WithAllocatedID(req.AllocatedID, true),
		couponwithdraw1.WithAmount(req.Amount, true),
		couponwithdraw1.WithReviewID(req.ReviewID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCouponWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateCouponWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateCouponWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCouponWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateCouponWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCouponWithdrawResponse{
		Info: info,
	}, nil
}
