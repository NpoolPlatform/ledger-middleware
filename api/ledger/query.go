package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLedgers(ctx context.Context, in *npool.GetLedgersRequest) (
	*npool.GetLedgersResponse,
	error,
) {
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithConds(in.GetConds()),
		ledger1.WithOffset(in.Offset),
		ledger1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLedgers",
			"In", in,
			"Error", err,
		)
		return &npool.GetLedgersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetLedgers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLedgers",
			"In", in,
			"Error", err,
		)
		return &npool.GetLedgersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLedgersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetLedger(ctx context.Context, in *npool.GetLedgerRequest) (*npool.GetLedgerResponse, error) {
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLedger",
			"In", in,
			"error", err,
		)
		return &npool.GetLedgerResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetLedger(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLedger",
			"In", in,
			"error", err,
		)
		return &npool.GetLedgerResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLedgerResponse{
		Info: info,
	}, nil
}
