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

func (s *Server) DeleteCouponWithdraw(
	ctx context.Context,
	in *npool.DeleteCouponWithdrawRequest,
) (*npool.DeleteCouponWithdrawResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCouponWithdraw",
			"In", in,
		)
		return &npool.DeleteCouponWithdrawResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := couponwithdraw1.NewHandler(
		ctx,
		couponwithdraw1.WithID(req.ID, false),
		couponwithdraw1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCouponWithdraw",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCouponWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteCouponWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCouponWithdraw",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCouponWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteCouponWithdrawResponse{
		Info: info,
	}, nil
}
