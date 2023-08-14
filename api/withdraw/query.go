package withdraw

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	withdraw1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/withdraw"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetWithdraws(ctx context.Context, in *npool.GetWithdrawsRequest) (
	*npool.GetWithdrawsResponse,
	error,
) {
	handler, err := withdraw1.NewHandler(
		ctx,
		withdraw1.WithConds(in.GetConds()),
		withdraw1.WithOffset(in.Offset),
		withdraw1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetWithdraws",
			"In", in,
			"Error", err,
		)
		return &npool.GetWithdrawsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetWithdraws(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetWithdraws",
			"In", in,
			"Error", err,
		)
		return &npool.GetWithdrawsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetWithdrawsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetWithdraw(ctx context.Context, in *npool.GetWithdrawRequest) (*npool.GetWithdrawResponse, error) {
	handler, err := withdraw1.NewHandler(
		ctx,
		withdraw1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetWithdraw",
			"In", in,
			"error", err,
		)
		return &npool.GetWithdrawResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetWithdraw",
			"In", in,
			"error", err,
		)
		return &npool.GetWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetWithdrawResponse{
		Info: info,
	}, nil
}
