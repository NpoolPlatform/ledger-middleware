//nolint
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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateWithdraw",
			"In", in,
		)
		return &npool.CreateWithdrawResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := withdraw1.NewHandler(
		ctx,
		withdraw1.WithEntID(req.EntID, false),
		withdraw1.WithAppID(req.AppID, true),
		withdraw1.WithUserID(req.UserID, true),
		withdraw1.WithCoinTypeID(req.CoinTypeID, true),
		withdraw1.WithAccountID(req.AccountID, true),
		withdraw1.WithAmount(req.Amount, true),
		withdraw1.WithAddress(req.Address, true),
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
