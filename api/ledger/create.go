package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v1/ledger"

	"github.com/NpoolPlatform/ledger-manager/api/detail"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/ledger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if err := detail.Validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := ledger1.BookKeeping(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, status.Error(codes.Unimplemented, "NOT IMEPLEMENTED")
}

func (s *Server) UnlockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (*npool.UnlockBalanceResponse, error) {
	return &npool.UnlockBalanceResponse{}, status.Error(codes.Unimplemented, "NOT IMEPLEMENTED")
}
