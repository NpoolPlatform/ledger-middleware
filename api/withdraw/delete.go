package withdraw

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	withdraw1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/withdraw"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteWithdraw(ctx context.Context, in *npool.DeleteWithdrawRequest) (*npool.DeleteWithdrawResponse, error) {
	req := in.GetInfo()
	handler, err := withdraw1.NewHandler(
		ctx,
		withdraw1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteWithdraw",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteWithdraw(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteWithdraw",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteWithdrawResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteWithdrawResponse{
		Info: info,
	}, nil
}
