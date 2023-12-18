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

func (s *Server) UpdateCouponWithdraw(ctx context.Context, in *npool.UpdateCouponWithdrawRequest) (*npool.UpdateCouponWithdrawResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCouponWithdraw",
			"In", in,
		)
		return &npool.UpdateCouponWithdrawResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := couponwithdraw1.NewHandler(
		ctx,
		couponwithdraw1.WithID(req.ID, true),
		couponwithdraw1.WithState(req.State, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCouponWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.UpdateCouponWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateCouponWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCouponWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.UpdateCouponWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateCouponWithdrawResponse{
		Info: info,
	}, nil
}
