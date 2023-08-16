//nolint:dupl
package withdraw

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
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

func CreateWithdraw(ctx context.Context, in *npool.WithdrawReq) (*npool.Withdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateWithdraw(ctx, &npool.CreateWithdrawRequest{
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
	return info.(*npool.Withdraw), nil
}

func CreateWithdraws(ctx context.Context, in []*npool.WithdrawReq) ([]*npool.Withdraw, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateWithdraws(ctx, &npool.CreateWithdrawsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create withdraws: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create withdraws: %v", err)
	}
	return infos.([]*npool.Withdraw), nil
}

func UpdateWithdraw(ctx context.Context, in *npool.WithdrawReq) (*npool.Withdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateWithdraw(ctx, &npool.UpdateWithdrawRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update withdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update withdraw: %v", err)
	}
	return info.(*npool.Withdraw), nil
}

func GetWithdraw(ctx context.Context, id string) (*npool.Withdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetWithdraw(ctx, &npool.GetWithdrawRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get withdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get withdraw: %v", err)
	}
	return info.(*npool.Withdraw), nil
}

func GetWithdrawOnly(ctx context.Context, conds *npool.Conds) (*npool.Withdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetWithdrawOnly(ctx, &npool.GetWithdrawOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get withdraw only: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get withdraw only: %v", err)
	}
	return info.(*npool.Withdraw), nil
}

func GetWithdraws(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Withdraw, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetWithdraws(ctx, &npool.GetWithdrawsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get withdraws: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get withdraws: %v", err)
	}
	return infos.([]*npool.Withdraw), total, nil
}

func ExistWithdraw(ctx context.Context, id string) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistWithdraw(ctx, &npool.ExistWithdrawRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get withdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get withdraw: %v", err)
	}
	return infos.(bool), nil
}

func ExistWithdrawConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistWithdrawConds(ctx, &npool.ExistWithdrawCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get withdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get withdraw: %v", err)
	}
	return infos.(bool), nil
}