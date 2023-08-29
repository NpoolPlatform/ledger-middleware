package goodstatement

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func CreateGoodStatement(ctx context.Context, req *npool.GoodStatementReq) (*npool.GoodStatement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateGoodStatement(ctx, &npool.CreateGoodStatementRequest{
			Info: req,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodstatement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodstatement: %v", err)
	}
	return info.(*npool.GoodStatement), nil
}

func GetGoodStatementOnly(ctx context.Context, conds *npool.Conds) (*npool.GoodStatement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetGoodStatementOnly(ctx, &npool.GetGoodStatementOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodstatement only: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodstatement only: %v", err)
	}
	return info.(*npool.GoodStatement), nil
}

func GetGoodStatements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.GoodStatement, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetGoodStatements(ctx, &npool.GetGoodStatementsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get details: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get details: %v", err)
	}
	return infos.([]*npool.GoodStatement), total, nil
}
