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

func (s *Server) UnlockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (*npool.UnlockBalanceResponse, error) {
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithLockID(&in.LockID, true),
		lock1.WithRollback(&in.Rollback, true),
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
