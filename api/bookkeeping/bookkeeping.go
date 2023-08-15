package bookkeeping

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	bookkeeping1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/bookkeeping"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/bookkeeping"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (
	*npool.BookKeepingResponse,
	error,
) {
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"BookKeeping",
			"In", in,
			"Error", err,
		)
		return &npool.BookKeepingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.BookKeeping(ctx); err != nil {
		logger.Sugar().Errorw(
			"BookKeeping",
			"In", in,
			"Error", err,
		)
		return &npool.BookKeepingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}

//nolint
func (s *Server) BookKeepingOut(ctx context.Context, in *npool.BookKeepingRequest) (
	*npool.BookKeepingResponse,
	error,
) {
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"BookKeepingOut",
			"In", in,
			"Error", err,
		)
		return &npool.BookKeepingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.BookKeepingOut(ctx); err != nil {
		logger.Sugar().Errorw(
			"BookKeepingOut",
			"In", in,
			"Error", err,
		)
		return &npool.BookKeepingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}

//nolint
func (s *Server) LockBalance(ctx context.Context, in *npool.LockBalanceRequest) (
	*npool.LockBalanceResponse,
	error,
) {
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithAppID(&in.AppID),
		bookkeeping1.WithUserID(&in.UserID),
		bookkeeping1.WithCoinTypeID(&in.CoinTypeID),
		bookkeeping1.WithAmount(&in.Amount),
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
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithAppID(&in.AppID),
		bookkeeping1.WithUserID(&in.UserID),
		bookkeeping1.WithCoinTypeID(&in.CoinTypeID),
		bookkeeping1.WithAmount(&in.Amount),
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
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithAppID(&in.AppID),
		bookkeeping1.WithUserID(&in.UserID),
		bookkeeping1.WithCoinTypeID(&in.CoinTypeID),
		bookkeeping1.WithUnlocked(&in.Unlocked),
		bookkeeping1.WithOutcoming(&in.Outcoming),
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
	handler, err := bookkeeping1.NewHandler(
		ctx,
		bookkeeping1.WithAppID(&in.AppID),
		bookkeeping1.WithUserID(&in.UserID),
		bookkeeping1.WithCoinTypeID(&in.CoinTypeID),
		bookkeeping1.WithUnlocked(&in.Unlocked),
		bookkeeping1.WithOutcoming(&in.Outcoming),
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
