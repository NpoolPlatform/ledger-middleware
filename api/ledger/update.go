package ledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddLedger(ctx context.Context, in *npool.AddLedgerRequest) (*npool.AddLedgerResponse, error) {
	req := in.GetInfo()
	handler, err := ledger1.NewHandler(
		ctx,
		ledger1.WithID(req.ID),
		ledger1.WithCoinTypeID(req.CoinTypeID),
		ledger1.WithIncoming(req.Incoming),
		ledger1.WithOutcoming(req.Outcoming),
		ledger1.WithLocked(req.Locked),
		ledger1.WithSpendable(req.Spendable),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.AddLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateLedger(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.AddLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddLedgerResponse{
		Info: info,
	}, nil
}
