package coupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	couponwithdraw1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/withdraw/coupon"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCouponWithdraws(ctx context.Context, in *npool.GetCouponWithdrawsRequest) (
	*npool.GetCouponWithdrawsResponse,
	error,
) {
	handler, err := couponwithdraw1.NewHandler(
		ctx,
		couponwithdraw1.WithConds(in.GetConds()),
		couponwithdraw1.WithOffset(in.Offset),
		couponwithdraw1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponWithdraws",
			"In", in,
			"Error", err,
		)
		return &npool.GetCouponWithdrawsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCouponWithdraws(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponWithdraws",
			"In", in,
			"Error", err,
		)
		return &npool.GetCouponWithdrawsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponWithdrawsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCouponWithdraw(ctx context.Context, in *npool.GetCouponWithdrawRequest) (*npool.GetCouponWithdrawResponse, error) {
	handler, err := couponwithdraw1.NewHandler(
		ctx,
		couponwithdraw1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponWithdraw",
			"In", in,
			"error", err,
		)
		return &npool.GetCouponWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCouponWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCouponWithdraw",
			"In", in,
			"error", err,
		)
		return &npool.GetCouponWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCouponWithdrawResponse{
		Info: info,
	}, nil
}
