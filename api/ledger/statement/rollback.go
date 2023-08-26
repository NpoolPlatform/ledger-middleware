//nolint
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
		statement1.WithID(req.ID, true),
		statement1.WithAppID(req.AppID, true),
		statement1.WithUserID(req.UserID, true),
		statement1.WithCoinTypeID(req.CoinTypeID, true),
		statement1.WithIOType(req.IOType, true),
		statement1.WithIOSubType(req.IOSubType, true),
		statement1.WithAmount(req.Amount, true),
		statement1.WithIOExtra(req.IOExtra, true),
		statement1.WithCreatedAt(*req.CreatedAt, false),
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
