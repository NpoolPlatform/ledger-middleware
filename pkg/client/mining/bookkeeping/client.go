package bookkeeping

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/bookkeeping"
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

func BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.BookKeeping(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail bookkeeping: %v", err)
		}
		return nil, err
	})
	if err != nil {
		return fmt.Errorf("fail bookkeeping: %v", err)
	}
	return nil
}
