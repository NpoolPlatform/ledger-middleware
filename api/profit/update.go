package profit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	profit1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddProfit(ctx context.Context, in *npool.AddProfitRequest) (*npool.AddProfitResponse, error) {
	req := in.GetInfo()
	handler, err := profit1.NewHandler(
		ctx,
		profit1.WithID(req.ID),
		profit1.WithAppID(req.AppID),
		profit1.WithUserID(req.UserID),
		profit1.WithCoinTypeID(req.CoinTypeID),
		profit1.WithIncoming(req.Incoming),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddProfit",
			"Req", in,
			"Error", err,
		)
		return &npool.AddProfitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateProfit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddProfit",
			"Req", in,
			"Error", err,
		)
		return &npool.AddProfitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddProfitResponse{
		Info: info,
	}, nil
}
