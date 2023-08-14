package goodstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/goodstatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/goodstatement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodStatement(ctx context.Context, in *npool.CreateGoodStatementRequest) (*npool.CreateGoodStatementResponse, error) {
	req := in.GetInfo()
	handler, err := goodstatement1.NewHandler(
		ctx,
		goodstatement1.WithID(req.ID),
		goodstatement1.WithGoodID(req.GoodID),
		goodstatement1.WithCoinTypeID(req.CoinTypeID),
		goodstatement1.WithAmount(req.Amount),
		goodstatement1.WithBenefitDate(req.BenefitDate),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodStatementResponse{
		Info: info,
	}, nil
}
