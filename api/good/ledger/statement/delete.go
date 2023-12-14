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
func (s *Server) DeleteGoodStatement(ctx context.Context, in *npool.DeleteGoodStatementRequest) (
	*npool.DeleteGoodStatementResponse,
	error,
) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteGoodStatement",
			"In", in,
		)
		return &npool.DeleteGoodStatementResponse{}, status.Error(codes.Aborted, "invalid info")
	}

	handler, err := statement1.NewHandler(
		ctx,
		goodstatement1.WithID(req.ID, false),
		goodstatement1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteGoodStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteGoodStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteGoodStatementResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) DeleteGoodStatements(ctx context.Context, in *npool.DeleteGoodStatementsRequest) (*npool.DeleteGoodStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		goodstatement1.WithReqs(in.GetInfos(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.DeleteGoodStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteGoodStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteGoodStatementsResponse{
		Infos: infos,
	}, nil
}
