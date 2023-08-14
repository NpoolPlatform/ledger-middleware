package withdraw

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	withdraw1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/withdraw"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateWithdraw(ctx context.Context, in *npool.CreateWithdrawRequest) (*npool.CreateWithdrawResponse, error) {
	req := in.GetInfo()
	handler, err := withdraw1.NewHandler(
		ctx,
		withdraw1.WithID(req.ID),
		withdraw1.WithAppID(req.AppID),
		withdraw1.WithUserID(req.UserID),
		withdraw1.WithCoinTypeID(req.CoinTypeID),
		withdraw1.WithAccountID(req.AccountID),
		withdraw1.WithAmount(req.Amount),
		withdraw1.WithAddress(req.Address),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateWithdraw",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateWithdrawResponse{
		Info: info,
	}, nil
}
