//nolint:dupl
package bookkeeping

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/bookkeeping"
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

func BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.BookKeeping(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail create bookkeeping: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create bookkeeping: %v", err)
	}
	return info.(*npool.BookKeepingResponse), nil
}

func LockBalance(ctx context.Context, in *npool.LockBalanceRequest) (*npool.LockBalanceResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.LockBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail lock balance: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail lock balance: %v", err)
	}
	return info.(*npool.LockBalanceResponse), nil
}

func LockBalanceOut(ctx context.Context, in *npool.LockBalanceRequest) (*npool.LockBalanceResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.LockBalanceOut(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail lock balance out: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail lock balance out: %v", err)
	}
	return info.(*npool.LockBalanceResponse), nil
}

func UnlockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (*npool.UnlockBalanceResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UnlockBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail update bookkeeping: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update bookkeeping: %v", err)
	}
	return info.(*npool.UnlockBalanceResponse), nil
}

func UnlockBalanceOut(ctx context.Context, in *npool.UnlockBalanceRequest) (*npool.UnlockBalanceResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UnlockBalanceOut(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail update bookkeeping: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update bookkeeping: %v", err)
	}
	return info.(*npool.UnlockBalanceResponse), nil
}
