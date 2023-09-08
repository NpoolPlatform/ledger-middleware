//nolint
package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	lock1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddBalance(ctx context.Context, in *npool.AddBalanceRequest) (
	*npool.AddBalanceResponse,
	error,
) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"AddBalance",
			"In", in,
		)
		return &npool.AddBalanceResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(req.AppID, false),
		lock1.WithUserID(req.UserID, false),
		lock1.WithCoinTypeID(req.CoinTypeID, false),
		lock1.WithLocked(req.Locked, false),
		lock1.WithLockID(req.LockID, false),
		lock1.WithSpendable(req.Spendable, false),
		lock1.WithIOSubType(req.IOSubType, false),
		lock1.WithIOExtra(req.IOExtra, false),
		lock1.WithStatementID(req.StatementID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddBalance",
			"In", in,
			"Error", err,
		)
		return &npool.AddBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.AddBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddBalance",
			"In", in,
			"Error", err,
		)
		return &npool.AddBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddBalanceResponse{
		Info: info,
	}, nil
}
