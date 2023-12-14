//nolint:dupl
package ledger

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
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

func GetLedger(ctx context.Context, id string) (*npool.Ledger, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetLedger(ctx, &npool.GetLedgerRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ledger: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ledger: %v", err)
	}
	return info.(*npool.Ledger), nil
}

func GetLedgerOnly(ctx context.Context, conds *npool.Conds) (*npool.Ledger, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetLedgers(ctx, &npool.GetLedgersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ledger only: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get ledger only: %v", err)
	}
	if len(infos.([]*npool.Ledger)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Ledger)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.Ledger)[0], nil
}

func GetLedgers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Ledger, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetLedgers(ctx, &npool.GetLedgersRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get ledgers: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get ledgers: %v", err)
	}
	return infos.([]*npool.Ledger), total, nil
}

func LockBalance(ctx context.Context, in *npool.LockBalanceRequest) (*npool.Ledger, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.LockBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail lock balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail lock balance: %v", err)
	}
	return info.(*npool.Ledger), nil
}

func UnlockBalance(ctx context.Context, in *npool.UnlockBalanceRequest) (*npool.Ledger, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UnlockBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail unlock balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail unlock balance: %v", err)
	}
	return info.(*npool.Ledger), nil
}

func SettleBalance(ctx context.Context, in *npool.SettleBalanceRequest) (*npool.Ledger, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.SettleBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail settle balance: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail settle balance: %v", err)
	}
	return info.(*npool.Ledger), nil
}
