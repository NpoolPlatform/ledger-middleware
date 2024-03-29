//nolint
package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/simulate/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateStatement(ctx context.Context, in *npool.CreateStatementRequest) (*npool.CreateStatementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithEntID(req.EntID, false),
		statement1.WithAppID(req.AppID, true),
		statement1.WithUserID(req.UserID, true),
		statement1.WithCoinTypeID(req.CoinTypeID, true),
		statement1.WithIOType(req.IOType, true),
		statement1.WithIOSubType(req.IOSubType, true),
		statement1.WithAmount(req.Amount, true),
		statement1.WithIOExtra(req.IOExtra, true),
		statement1.WithCreatedAt(req.CreatedAt, false),
		statement1.WithSendCoupon(req.SendCoupon, false),
		statement1.WithCashable(req.Cashable, false),
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
		statement1.WithReqs(in.GetInfos(), true),
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
