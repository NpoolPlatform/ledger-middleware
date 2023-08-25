package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RollbackStatement(ctx context.Context, in *npool.RollbackStatementRequest) (*npool.RollbackStatementResponse, error) {
	req := in.GetInfo()
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID),
		statement1.WithAppID(req.AppID),
		statement1.WithUserID(req.UserID),
		statement1.WithCoinTypeID(req.CoinTypeID),
		statement1.WithIOType(req.IOType),
		statement1.WithIOSubType(req.IOSubType),
		statement1.WithAmount(req.Amount),
		statement1.WithIOExtra(req.IOExtra),
		statement1.WithCreatedAt(*req.CreatedAt),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.RollbackStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.RollbackStatementResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) RollbackStatements(ctx context.Context, in *npool.RollbackStatementsRequest) (*npool.RollbackStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.RollbackStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.RollbackStatementsResponse{
		Infos: infos,
	}, nil
}
