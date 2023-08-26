package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteStatement(ctx context.Context, in *npool.DeleteStatementRequest) (*npool.DeleteStatementResponse, error) {
	req := in.GetInfo()
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteStatementResponse{
		Info: info,
	}, nil
}