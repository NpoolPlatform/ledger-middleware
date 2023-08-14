package goodledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/goodledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodLedger(ctx context.Context, in *npool.CreateGoodLedgerRequest) (*npool.CreateGoodLedgerResponse, error) {
	req := in.GetInfo()
	handler, err := goodledger1.NewHandler(
		ctx,
		goodledger1.WithID(req.ID),
		goodledger1.WithGoodID(req.GoodID),
		goodledger1.WithCoinTypeID(req.CoinTypeID),
		goodledger1.WithAmount(req.Amount),
		goodledger1.WithToPlatform(req.ToPlatform),
		goodledger1.WithToUser(req.ToUser),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodLedger(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodLedgerResponse{
		Info: info,
	}, nil
}
