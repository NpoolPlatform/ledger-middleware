package profit

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/profit"
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

func GetProfit(ctx context.Context, id string) (*npool.Profit, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetProfit(ctx, &npool.GetProfitRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get profit: %v", err)
	}
	return info.(*npool.Profit), nil
}

func GetProfitOnly(ctx context.Context, conds *npool.Conds) (*npool.Profit, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetProfits(ctx, &npool.GetProfitsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get profit only: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get profit only: %v", err)
	}
	if len(infos.([]*npool.Profit)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Profit)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.Profit)[0], nil
}

func GetProfits(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Profit, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetProfits(ctx, &npool.GetProfitsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get profits: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get profits: %v", err)
	}
	return infos.([]*npool.Profit), total, nil
}

func ExistProfitConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistProfitConds(ctx, &npool.ExistProfitCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return false, fmt.Errorf("fail get profits: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get profits: %v", err)
	}
	return info.(bool), nil
}
