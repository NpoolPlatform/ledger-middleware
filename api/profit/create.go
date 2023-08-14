package profit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	profit1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateProfit(ctx context.Context, in *npool.CreateProfitRequest) (*npool.CreateProfitResponse, error) {
	req := in.GetInfo()
	handler, err := profit1.NewHandler(
		ctx,
		profit1.WithID(req.ID),
		profit1.WithAppID(req.AppID),
		profit1.WithUserID(req.UserID),
		profit1.WithCoinTypeID(req.CoinTypeID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateProfit",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateProfitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateProfit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateProfit",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateProfitResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateProfitResponse{
		Info: info,
	}, nil
}
