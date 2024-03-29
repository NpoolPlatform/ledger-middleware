package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/simulate/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetStatement(ctx context.Context, in *npool.GetStatementRequest) (*npool.GetStatementResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatement",
			"In", in,
			"error", err,
		)
		return &npool.GetStatementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatement",
			"In", in,
			"error", err,
		)
		return &npool.GetStatementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStatementResponse{
		Info: info,
	}, nil
}

func (s *Server) GetStatements(ctx context.Context, in *npool.GetStatementsRequest) (
	*npool.GetStatementsResponse,
	error,
) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithConds(in.GetConds()),
		statement1.WithOffset(in.Offset),
		statement1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetStatementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetStatements",
			"In", in,
			"Error", err,
		)
		return &npool.GetStatementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStatementsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
