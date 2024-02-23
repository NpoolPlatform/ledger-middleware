package statement

import (
	"context"

	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/simulate/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/statement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func (s *Server) UpdateStatement(ctx context.Context, in *npool.UpdateStatementRequest) (*npool.UpdateStatementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateStatement",
			"In", in,
		)
		return &npool.UpdateStatementResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID, true),
		statement1.WithCashUsed(req.CashUsed, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateStatement",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateStatement",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateStatementResponse{
		Info: info,
	}, nil
}
