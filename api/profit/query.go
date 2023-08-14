package profit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	profit1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetProfits(ctx context.Context, in *npool.GetProfitsRequest) (
	*npool.GetProfitsResponse,
	error,
) {
	handler, err := profit1.NewHandler(
		ctx,
		profit1.WithConds(in.GetConds()),
		profit1.WithOffset(in.Offset),
		profit1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetProfits",
			"In", in,
			"Error", err,
		)
		return &npool.GetProfitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetProfits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetProfits",
			"In", in,
			"Error", err,
		)
		return &npool.GetProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetProfitsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetProfit(ctx context.Context, in *npool.GetProfitRequest) (*npool.GetProfitResponse, error) {
	handler, err := profit1.NewHandler(
		ctx,
		profit1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetProfit",
			"In", in,
			"error", err,
		)
		return &npool.GetProfitResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetProfit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetProfit",
			"In", in,
			"error", err,
		)
		return &npool.GetProfitResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetProfitResponse{
		Info: info,
	}, nil
}
