package goodstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/goodstatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/goodstatement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetGoodStatements(ctx context.Context, in *npool.GetGoodStatementsRequest) (
	*npool.GetGoodStatementsResponse,
	error,
) {
	handler, err := goodstatement1.NewHandler(
		ctx,
		goodstatement1.WithConds(in.GetConds()),
		goodstatement1.WithOffset(in.Offset),
		goodstatement1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetGoodStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodStatementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
