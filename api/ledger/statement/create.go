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

func (s *Server) CreateStatement(ctx context.Context, in *npool.CreateStatementRequest) (*npool.CreateStatementResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"In", in,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}

	reqs := []*npool.StatementReq{req}
	resp, err := s.CreateStatements(ctx, &npool.CreateStatementsRequest{
		Infos: reqs,
	})
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatement",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if len(resp.Infos) == 0 {
		return &npool.CreateStatementResponse{}, nil
	}
	if len(resp.Infos) > 1 {
		logger.Sugar().Errorw(
			"CreateStatement",
			"Req", in,
			"Error", "too many record",
		)
		return &npool.CreateStatementResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateStatementResponse{
		Info: resp.Infos[0],
	}, nil
}

//nolint
func (s *Server) CreateStatements(ctx context.Context, in *npool.CreateStatementsRequest) (*npool.CreateStatementsResponse, error) {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CreateStatements(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateStatements",
			"Req", in,
			"Error", err,
		)
		return &npool.CreateStatementsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateStatementsResponse{
		Infos: infos,
	}, nil
}
