package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) CreateStatement(ctx context.Context, in *npool.CreateStatementRequest) (*npool.CreateStatementResponse, error) {
	req := in.GetInfo()
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithAppID(req.AppID),
		statement1.WithUserID(req.UserID),
		statement1.WithCoinTypeID(req.CoinTypeID),
		statement1.WithIOType(req.IOType),
		statement1.WithIOSubType(req.IOSubType),
		statement1.WithAmount(req.Amount),
		statement1.WithIOExtra(req.IOExtra),
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
func (s *Server) UnCreateStatement(ctx context.Context, in *npool.UnCreateStatementRequest) (*npool.UnCreateStatementResponse, error) {
	req := in.GetInfo()
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithAppID(req.AppID),
		statement1.WithUserID(req.UserID),
		statement1.WithCoinTypeID(req.CoinTypeID),
		statement1.WithIOType(req.IOType),
		statement1.WithIOSubType(req.IOSubType),
		statement1.WithAmount(req.Amount),
		statement1.WithIOExtra(req.IOExtra),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.UnCreateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.UnCreateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnCreateStatementResponse{
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

func (s *Server) UnCreateStatements(ctx context.Context, in *npool.UnCreateStatementsRequest) (*npool.UnCreateStatementsResponse, error) {
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
		return &npool.UnCreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CreateStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.UnCreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnCreateStatementsResponse{
		Infos: infos,
	}, nil
}
