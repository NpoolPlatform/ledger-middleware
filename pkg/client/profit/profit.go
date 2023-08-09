//nolint:dupl
package profit

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get profit connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateProfit(ctx context.Context, in *npool.ProfitReq) (*npool.Profit, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateProfit(ctx, &npool.CreateProfitRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create profit: %v", err)
	}
	return info.(*npool.Profit), nil
}

func CreateProfits(ctx context.Context, in []*npool.ProfitReq) ([]*npool.Profit, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateProfits(ctx, &npool.CreateProfitsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create profits: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create profits: %v", err)
	}
	return infos.([]*npool.Profit), nil
}

func AddProfit(ctx context.Context, in *npool.ProfitReq) (*npool.Profit, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.AddProfit(ctx, &npool.AddProfitRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update profit: %v", err)
	}
	return info.(*npool.Profit), nil
}

func GetProfit(ctx context.Context, id string) (*npool.Profit, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetProfit(ctx, &npool.GetProfitRequest{
			ID: id,
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
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetProfitOnly(ctx, &npool.GetProfitOnlyRequest{
			Conds: conds,
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

func GetProfits(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Profit, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
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

func ExistProfit(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistProfit(ctx, &npool.ExistProfitRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get profit: %v", err)
	}
	return infos.(bool), nil
}

func ExistProfitConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistProfitConds(ctx, &npool.ExistProfitCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get profit: %v", err)
	}
	return infos.(bool), nil
}

func CountProfits(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountProfits(ctx, &npool.CountProfitsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count profit: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteProfit(ctx context.Context, id string) (*npool.Profit, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteProfit(ctx, &npool.DeleteProfitRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete profit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete profit: %v", err)
	}
	return infos.(*npool.Profit), nil
}
