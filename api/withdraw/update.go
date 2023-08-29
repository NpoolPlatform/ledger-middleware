package withdraw

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	withdraw1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/withdraw"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateWithdraw(ctx context.Context, in *npool.UpdateWithdrawRequest) (*npool.UpdateWithdrawResponse, error) {
	req := in.GetInfo()
	handler, err := withdraw1.NewHandler(
		ctx,
		withdraw1.WithID(req.ID, true),
		withdraw1.WithPlatformTransactionID(req.PlatformTransactionID, false),
		withdraw1.WithChainTransactionID(req.ChainTransactionID, false),
		withdraw1.WithState(req.State, false),
		withdraw1.WithFeeAmount(req.FeeAmount, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.UpdateWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.UpdateWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateWithdrawResponse{
		Info: info,
	}, nil
}
