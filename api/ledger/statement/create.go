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

func (s *Server) CreateStatement(ctx context.Context, in *npool.CreateStatementRequest) (*npool.CreateStatementResponse, error) {
	req := in.GetInfo()
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID, false),
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
			"CreateStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateStatementResponse{
		Info: info,
	}, nil
}

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
