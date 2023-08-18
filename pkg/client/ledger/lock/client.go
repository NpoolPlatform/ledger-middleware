package lock

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/lock"
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
func LockBalance(ctx context.Context, in *npool.BalanceReq) (*ledgerpb.Ledger, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.LockBalance(ctx, &npool.LockBalanceRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail lock balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail lock balance: %v", err)
	}
	return info.(*ledgerpb.Ledger), nil
}

//nolint
func UnlockBalance(ctx context.Context, in *npool.BalanceReq) (*ledgerpb.Ledger, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UnlockBalance(ctx, &npool.UnlockBalanceRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail unlock balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail unlock balance: %v", err)
	}
	return info.(*ledgerpb.Ledger), nil
}

//nolint
func SpendBalance(ctx context.Context, in *npool.SpendBalanceReq) ([]*ledgerpb.Ledger, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.SpendBalance(ctx, &npool.SpendBalanceRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail spend balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail spend balance: %v", err)
	}
	return infos.([]*ledgerpb.Ledger), nil
}

//nolint
func UnspendBalance(ctx context.Context, in *npool.SpendBalanceReq) ([]*ledgerpb.Ledger, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UnspendBalance(ctx, &npool.UnspendBalanceRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail unspend balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail unspend balance: %v", err)
	}
	return infos.([]*ledgerpb.Ledger), nil
}
