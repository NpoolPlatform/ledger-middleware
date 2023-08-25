//nolint
package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodStatement(ctx context.Context, in *npool.CreateGoodStatementRequest) (
	*npool.CreateGoodStatementResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := goodstatement1.NewHandler(
		ctx,
		goodstatement1.WithID(req.ID, false),
		goodstatement1.WithUnsoldStatementID(req.UnsoldStatementID, false),
		goodstatement1.WithGoodID(req.GoodID, true),
		goodstatement1.WithCoinTypeID(req.CoinTypeID, true),
		goodstatement1.WithTotalAmount(req.TotalAmount, true),
		goodstatement1.WithUnsoldAmount(req.UnsoldAmount, true),
		goodstatement1.WithTechniqueServiceFeeAmount(req.TechniqueServiceFeeAmount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodStatementResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateGoodStatements(ctx context.Context, in *npool.CreateGoodStatementsRequest) (
	*npool.CreateGoodStatementsResponse,
	error,
) {
	handler, err := goodstatement1.NewHandler(
		ctx,
		goodstatement1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CreateGoodStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodStatementsResponse{
		Infos: infos,
	}, nil
}
