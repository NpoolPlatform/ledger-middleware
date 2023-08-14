package statement

import (
	"context"

	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetIntervalStatements(ctx context.Context, in *npool.GetIntervalStatementsResponse) (
	*npool.GetIntervalStatementsResponse,
	error,
) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithAppID(&in.AppID),
		statement1.WithUserID(&in.UserID),
		statement1.WithStart(in.Start),
		statement1.WithEnd(in.End),
		statement1.WithOffset(in.Offset),
		statement1.WithLimit(in.Limit),
	)
	if err != nil {
		return &npool.GetIntervalStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetStatements(ctx)
	if err != nil {
		return &npool.GetIntervalStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetIntervalStatementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
