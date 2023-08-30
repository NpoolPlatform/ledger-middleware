package statement 

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistGoodStatementConds(
	ctx context.Context,
	in *npool.ExistGoodStatementCondsRequest,
) (
	*npool.ExistGoodStatementCondsResponse,
	error,
) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistGoodStatementCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistGoodStatementConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistGoodStatementCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistGoodStatementCondsResponse{
		Info: info,
	}, nil
}
