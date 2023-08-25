package profit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	profit1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/profit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Server) DeleteProfit(ctx context.Context, in *npool.DeleteProfitRequest) (*npool.DeleteProfitResponse, error) {
	req := in.GetInfo()
	handler, err := profit1.NewHandler(
		ctx,
		profit1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteProfit",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteProfitResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteProfit(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteProfit",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteProfitResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteProfitResponse{
		Info: info,
	}, nil
}