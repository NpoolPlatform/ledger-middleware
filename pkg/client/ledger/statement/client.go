package statement

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
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

func GetStatement(ctx context.Context, id string) (*npool.Statement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetStatement(ctx, &npool.GetStatementRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get statement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get statement: %v", err)
	}
	return info.(*npool.Statement), nil
}

func GetStatementOnly(ctx context.Context, conds *npool.Conds) (*npool.Statement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetStatementOnly(ctx, &npool.GetStatementOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get statement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get statement: %v", err)
	}
	return info.(*npool.Statement), nil
}

func GetStatements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Statement, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetStatements(ctx, &npool.GetStatementsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get statements: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get statements: %v", err)
	}
	return infos.([]*npool.Statement), total, nil
}

func CreateStatement(ctx context.Context, in *npool.StatementReq) (*npool.Statement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateStatement(ctx, &npool.CreateStatementRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create withdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create withdraw: %v", err)
	}
	return info.(*npool.Statement), nil
}

//nolint
func CreateStatements(ctx context.Context, in []*npool.StatementReq) ([]*npool.Statement, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateStatements(ctx, &npool.CreateStatementsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create statements: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create statements: %v", err)
	}
	return infos.([]*npool.Statement), nil
}

//nolint
func UnCreateStatements(ctx context.Context, in []*npool.StatementReq) ([]*npool.Statement, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UnCreateStatements(ctx, &npool.UnCreateStatementsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail uncreate statements: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail uncreate statements: %v", err)
	}
	return infos.([]*npool.Statement), nil
}
