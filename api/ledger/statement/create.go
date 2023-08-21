package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) CreateStatements(ctx context.Context, in *npool.CreateStatementsRequest) (*npool.CreateStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CreateStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateStatementsResponse{
		Infos: infos,
	}, nil
}

//nolint
func (s *Server) UnCreateStatements(ctx context.Context, in *npool.UnCreateStatementsRequest) (*npool.UnCreateStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UnCreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.UnCreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.UnCreateStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UnCreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.UnCreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnCreateStatementsResponse{
		Infos: infos,
	}, nil
}
