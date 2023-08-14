package ledger

import (
	"context"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v1/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetIntervalLedgers(ctx context.Context, in *npool.GetIntervalLedgersRequest) (
	*npool.GetIntervalLedgersResponse,
	error,
) {
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithAppID(&in.AppID),
		ledger1.WithUserID(&in.UserID),
		ledger1.WithStart(in.Start),
		ledger1.WithEnd(in.End),
		ledger1.WithOffset(in.Offset),
		ledger1.WithLimit(in.Limit),
	)
	if err != nil {
		return &npool.GetIntervalLedgersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetLedgers(ctx)
	if err != nil {
		return &npool.GetIntervalLedgersResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetIntervalLedgersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
