package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLedger(ctx context.Context, in *npool.CreateLedgerRequest) (*npool.CreateLedgerResponse, error) {
	req := in.GetInfo()
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithID(req.ID),
		ledger1.WithAppID(req.AppID),
		ledger1.WithUserID(req.UserID),
		ledger1.WithCoinTypeID(req.CoinTypeID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLedger(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLedgerResponse{
		Info: info,
	}, nil
}
