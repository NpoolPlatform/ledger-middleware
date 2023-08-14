package goodledger

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddGoodLedger(ctx context.Context, in *npool.AddGoodLedgerRequest) (*npool.AddGoodLedgerResponse, error) {
	req := in.GetInfo()
	handler, err := goodledger1.NewHandler(
		ctx,
		goodledger1.WithID(req.ID),
		goodledger1.WithToPlatform(req.ToPlatform),
		goodledger1.WithToUser(req.ToUser),
		goodledger1.WithAmount(req.Amount),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddGoodLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.AddGoodLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGoodLedger(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddGoodLedger",
			"Req", in,
			"Error", err,
		)
		return &npool.AddGoodLedgerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddGoodLedgerResponse{
		Info: info,
	}, nil
}
