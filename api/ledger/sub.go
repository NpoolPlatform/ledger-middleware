package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	lock1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) SubBalance(ctx context.Context, in *npool.SubBalanceRequest) (
	*npool.SubBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(req.AppID),
		lock1.WithUserID(req.UserID),
		lock1.WithCoinTypeID(req.CoinTypeID),
		lock1.WithSpendable(req.Spendable),
		lock1.WithLocked(req.Locked),
		lock1.WithIOSubType(req.IOSubType),
		lock1.WithIOExtra(req.IOExtra),
		lock1.WithStatementID(req.StatementID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SubBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SubBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.SubBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SubBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SubBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.SubBalanceResponse{
		Info: info,
	}, nil
}
