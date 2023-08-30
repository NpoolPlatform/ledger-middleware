//nolint
package statement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteStatement(ctx context.Context, in *npool.DeleteStatementRequest) (*npool.DeleteStatementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"In", in,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteStatement(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteStatementResponse{
		Info: info,
	}, nil
}

//nolint
func (s *Server) DeleteStatements(ctx context.Context, in *npool.DeleteStatementsRequest) (*npool.DeleteStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(in.GetInfos(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.DeleteStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.DeleteStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteStatementsResponse{
		Infos: infos,
	}, nil
}
