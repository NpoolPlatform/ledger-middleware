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
		lock1.WithAppID(&req.AppID),
		lock1.WithUserID(&req.UserID),
		lock1.WithCoinTypeID(&req.CoinTypeID),
		lock1.WithAmount(&req.Amount),
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

//nolint
func (s *Server) UnlockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (
	*npool.UnlockBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID),
		lock1.WithUserID(&req.UserID),
		lock1.WithCoinTypeID(&req.CoinTypeID),
		lock1.WithAmount(&req.Amount),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UnlockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UnlockBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UnlockBalance",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UnlockBalanceResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) SpendBalance(ctx context.Context, in *npool.SpendBalanceRequest) (
	*npool.SpendBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID),
		lock1.WithUserID(&req.UserID),
		lock1.WithCoinTypeID(&req.CoinTypeID),
		lock1.WithUnlocked(&req.Unlocked),
		lock1.WithOutcoming(&req.Outcoming),
		lock1.WithIOSubType(&req.IOSubType),
		lock1.WithIOExtra(&req.IOExtra),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SpendBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SpendBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UnlockBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SpendBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SpendBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.SpendBalanceResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) UnspendBalance(ctx context.Context, in *npool.UnspendBalanceRequest) (
	*npool.UnspendBalanceResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithAppID(&req.AppID),
		lock1.WithUserID(&req.UserID),
		lock1.WithCoinTypeID(&req.CoinTypeID),
		lock1.WithUnlocked(&req.Unlocked),
		lock1.WithOutcoming(&req.Outcoming),
		lock1.WithIOSubType(&req.IOSubType),
		lock1.WithIOExtra(&req.IOExtra),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UnspendBalance",
			"In", in,
			"Error", err,
		)
		return &npool.UnspendBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UnspendBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UnspendBalance",
			"In", in,
			"Error", err,
		)
		return &npool.UnspendBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnspendBalanceResponse{
		Info: info,
	}, nil
}
