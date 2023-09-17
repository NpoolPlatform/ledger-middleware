package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistStatementConds(ctx context.Context, in *npool.ExistStatementCondsRequest) (
	*npool.ExistStatementCondsResponse,
	error,
) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistStatementConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistStatementCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistStatementConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistStatementConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistStatementCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStatementCondsResponse{
		Info: info,
	}, nil
}
