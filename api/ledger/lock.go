//nolint:dupl
package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	lock1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LockBalance(ctx context.Context, in *npool.LockBalanceRequest) (*npool.LockBalanceResponse, error) {
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&in.AppID, true),
		lock1.WithUserID(&in.UserID, true),
		lock1.WithCoinTypeID(&in.CoinTypeID, true),
		lock1.WithLocked(&in.Amount, true),
		lock1.WithLockID(&in.LockID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"LockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.LockBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"LockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LockBalanceResponse{
		Info: info,
	}, nil
}

func (s *Server) LockBalances(ctx context.Context, in *npool.LockBalancesRequest) (*npool.LockBalancesResponse, error) {
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&in.AppID, true),
		lock1.WithUserID(&in.UserID, true),
		lock1.WithLockID(&in.LockID, true),
		lock1.WithBalances(in.Balances, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"LockBalances",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalancesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.LockBalances(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"LockBalances",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalancesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LockBalancesResponse{
		Infos: infos,
	}, nil
}
