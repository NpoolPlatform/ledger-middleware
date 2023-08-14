package mining

import (
	"context"

	bookkeeping1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/bookkeeping"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/bookkeeping"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (
	*npool.BookKeepingResponse,
	error,
) {
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithGoodID(&in.GoodID),
		bookkeeping1.WithCoinTypeID(&in.CoinTypeID),
		bookkeeping1.WithTotalAmount(&in.TotalAmount),
		bookkeeping1.WithUnsoldAmount(&in.UnsoldAmount),
		bookkeeping1.WithTechniqueServiceFeeAmount(&in.TechniqueServiceFeeAmount),
		bookkeeping1.WithBenefitDate(&in.BenefitDate),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"BookKeeping",
			"In", in,
			"Error", err,
		)
		return &npool.BookKeepingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.BookKeeping(ctx); err != nil {
		logger.Sugar().Errorw(
			"BookKeeping",
			"In", in,
			"Error", err,
		)
		return &npool.BookKeepingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}
