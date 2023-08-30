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

//nolint
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

func CreateGoodStatements(ctx context.Context, in []*npool.GoodStatementReq) ([]*npool.GoodStatement, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateGoodStatements(ctx, &npool.CreateGoodStatementsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create good statements: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create good statements: %v", err)
	}
	return infos.([]*npool.GoodStatement), nil
}

//nolint
func DeleteGoodStatement(ctx context.Context, req *npool.GoodStatementReq) (*npool.GoodStatement, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteGoodStatement(ctx, &npool.DeleteGoodStatementRequest{
			Info: req,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete goodstatement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete goodstatement: %v", err)
	}
	return info.(*npool.GoodStatement), nil
}

func DeleteGoodStatements(ctx context.Context, in []*npool.GoodStatementReq) ([]*npool.GoodStatement, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteGoodStatements(ctx, &npool.DeleteGoodStatementsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete good statements: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete good statements: %v", err)
	}
	return infos.([]*npool.GoodStatement), nil
}

func GetGoodStatementOnly(ctx context.Context, conds *npool.Conds) (*npool.GoodStatement, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetGoodStatements(ctx, &npool.GetGoodStatementsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodstatement only: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodstatement only: %v", err)
	}
	if len(infos.([]*npool.GoodStatement)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.GoodStatement)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.GoodStatement)[0], nil
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
			return nil, fmt.Errorf("fail get goodstatements: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get goodstatements: %v", err)
	}
	return infos.([]*npool.GoodStatement), total, nil
}

func ExistGoodStatementConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistGoodStatementConds(ctx, &npool.ExistGoodStatementCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist goodstatement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail exist goodstatement: %v", err)
	}
	return infos.(bool), nil
}
