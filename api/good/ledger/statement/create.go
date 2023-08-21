package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodStatements(ctx context.Context, in *npool.CreateGoodStatementsRequest) (
	*npool.CreateGoodStatementsResponse,
	error,
) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CreateGoodStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodStatementsResponse{
		Infos: infos,
	}, nil
}
