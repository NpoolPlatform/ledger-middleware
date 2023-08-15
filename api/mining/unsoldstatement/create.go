package unsoldstatement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	unsoldstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/unsoldstatement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUnsoldStatement(ctx context.Context, in *npool.CreateUnsoldStatementRequest) (
	*npool.CreateUnsoldStatementResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := unsoldstatement1.NewHandler(
		ctx,
		unsoldstatement1.WithID(req.ID),
		unsoldstatement1.WithGoodID(req.GoodID),
		unsoldstatement1.WithCoinTypeID(req.CoinTypeID),
		unsoldstatement1.WithAmount(req.Amount),
		unsoldstatement1.WithBenefitDate(req.BenefitIntervalHours),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUnsoldStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateUnsoldStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateUnsoldStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUnsoldStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateUnsoldStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUnsoldStatementResponse{
		Info: info,
	}, nil
}
