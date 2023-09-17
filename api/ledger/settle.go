package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	lock1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SettleBalance(ctx context.Context, in *npool.SettleBalanceRequest) (*npool.SettleBalanceResponse, error) {
	handler, err := lock1.NewHandler(
		ctx,
		lock1.WithLockID(&in.LockID, true),
		lock1.WithIOSubType(&in.IOSubType, true),
		lock1.WithIOExtra(&in.IOExtra, true),
		lock1.WithStatementID(&in.StatementID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SettleBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SettleBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.SettleBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SettleBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SettleBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.SettleBalanceResponse{
		Info: info,
	}, nil
}
