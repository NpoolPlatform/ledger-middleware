//nolint
package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/statement"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint
func (s *Server) RollbackGoodStatement(ctx context.Context, in *npool.RollbackGoodStatementRequest) (
	*npool.RollbackGoodStatementResponse,
	error,
) {
	req := in.GetInfo()
	handler, err := statement1.NewHandler(
		ctx,
		goodstatement1.WithID(req.ID, true),
		goodstatement1.WithUnsoldStatementID(req.UnsoldStatementID, true),
		goodstatement1.WithGoodID(req.GoodID, true),
		goodstatement1.WithCoinTypeID(req.CoinTypeID, true),
		goodstatement1.WithTotalAmount(req.TotalAmount, true),
		goodstatement1.WithUnsoldAmount(req.UnsoldAmount, true),
		goodstatement1.WithTechniqueServiceFeeAmount(req.TechniqueServiceFeeAmount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.RollbackGoodStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.RollbackGoodStatementResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) RollbackGoodStatements(ctx context.Context, in *npool.RollbackGoodStatementsRequest) (*npool.RollbackGoodStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		goodstatement1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.RollbackGoodStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"RollbackGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.RollbackGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.RollbackGoodStatementsResponse{
		Infos: infos,
	}, nil
}
