package ledger

import (
	"context"
	"errors"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	detailmgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/ledger"

	errno "github.com/NpoolPlatform/ledger-middleware/pkg/errno"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if err := detail.Validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := ledger1.BookKeeping(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		if errors.Is(err, errno.ErrAlreadyExists) {
			return &npool.BookKeepingResponse{}, status.Error(codes.AlreadyExists, err.Error())
		}
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}

func (s *Server) LockBalance(ctx context.Context, in *npool.LockBalanceRequest) (*npool.LockBalanceResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("LockBalance", "AppID", in.GetAppID(), "error", err)
		return &npool.LockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("LockBalance", "UserID", in.GetUserID(), "error", err)
		return &npool.LockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("LockBalance", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.LockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	locked, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		logger.Sugar().Errorw("LockBalance", "Amount", in.GetAmount(), "error", err)
		return &npool.LockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = ledger1.LockBalance(ctx, in.GetAppID(), in.GetUserID(), in.GetCoinTypeID(), locked)
	if err != nil {
		logger.Sugar().Errorw("LockBalance", "error", err)
		return &npool.LockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.LockBalanceResponse{}, nil
}

func (s *Server) UnlockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (*npool.UnlockBalanceResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("UnlockBalance", "AppID", in.GetAppID(), "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("UnlockBalance", "UserID", in.GetUserID(), "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("UnlockBalance", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	unlocked, err := decimal.NewFromString(in.GetUnlocked())
	if err != nil {
		logger.Sugar().Errorw("UnlockBalance", "Unlocked", in.GetUnlocked(), "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	outcoming, err := decimal.NewFromString(in.GetOutcoming())
	if err != nil {
		logger.Sugar().Errorw("UnlockBalance", "Outcoming", in.GetOutcoming(), "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	switch in.GetIOSubType() {
	case detailmgrpb.IOSubType_Payment:
		// TODO: match extra pattern
	case detailmgrpb.IOSubType_Withdrawal:
		// TODO: match extra pattern
	default:
		logger.Sugar().Errorw("UnlockBalance", "IOSubType", in.GetIOSubType(), "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = ledger1.UnlockBalance(
		ctx,
		in.GetAppID(), in.GetUserID(), in.GetCoinTypeID(),
		in.GetIOSubType(),
		unlocked, outcoming,
		in.GetIOExtra(),
	)
	if err != nil {
		logger.Sugar().Errorw("UnlockBalance", "error", err)
		return &npool.UnlockBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UnlockBalanceResponse{}, nil
}
