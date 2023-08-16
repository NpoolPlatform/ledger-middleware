package lock

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	lock1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/lock"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/lock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) LockBalance(ctx context.Context, in *npool.LockBalanceRequest) (
	*npool.LockBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID, true),
		lock1.WithUserID(&req.UserID, true),
		lock1.WithCoinTypeID(&req.CoinTypeID, true),
		lock1.WithAmount(&req.Amount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"LockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.LockBalance(ctx); err != nil {
		logger.Sugar().Errorw(
			"LockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LockBalanceResponse{}, nil
}

//nolint
func (s *Server) LockBalanceOut(ctx context.Context, in *npool.LockBalanceRequest) (
	*npool.LockBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID, true),
		lock1.WithUserID(&req.UserID, true),
		lock1.WithCoinTypeID(&req.CoinTypeID, true),
		lock1.WithAmount(&req.Amount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"LockBalanceOut",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.LockBalanceOut(ctx); err != nil {
		logger.Sugar().Errorw(
			"LockBalanceOut",
			"In", in,
			"Error", err,
		)
		return &npool.LockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.LockBalanceResponse{}, nil
}

//nolint
func (s *Server) UnLockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (
	*npool.UnlockBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID, true),
		lock1.WithUserID(&req.UserID, true),
		lock1.WithCoinTypeID(&req.CoinTypeID, true),
		lock1.WithUnlocked(&req.Unlocked, true),
		lock1.WithOutcoming(&req.Outcoming, true),
		lock1.WithIOSubType(req.IOSubType, false),
		lock1.WithIOExtra(&req.IOExtra, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UnLockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UnlockBalance(ctx); err != nil {
		logger.Sugar().Errorw(
			"UnLockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UnlockBalanceResponse{}, nil
}

//nolint
func (s *Server) UnLockBalanceOut(ctx context.Context, in *npool.UnlockBalanceRequest) (
	*npool.UnlockBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID, true),
		lock1.WithUserID(&req.UserID, true),
		lock1.WithCoinTypeID(&req.CoinTypeID, true),
		lock1.WithUnlocked(&req.Unlocked, true),
		lock1.WithOutcoming(&req.Outcoming, true),
		lock1.WithIOSubType(req.IOSubType, false),
		lock1.WithIOExtra(&req.IOExtra, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UnlockBalanceOut",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UnlockBalanceOut(ctx); err != nil {
		logger.Sugar().Errorw(
			"UnlockBalanceOut",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnlockBalanceResponse{}, nil
}
