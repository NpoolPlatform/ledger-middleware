package profit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	profit1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/simulate/ledger/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/profit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistProfitConds(ctx context.Context, in *npool.ExistProfitCondsRequest) (
	*npool.ExistProfitCondsResponse,
	error,
) {
	handler, err := profit1.NewHandler(
		ctx,
		profit1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistProfitConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistProfitCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistProfitConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistProfitConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistProfitCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistProfitCondsResponse{
		Info: info,
	}, nil
}
