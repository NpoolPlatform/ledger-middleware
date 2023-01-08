package bookkeeping

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/bookkeeping"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get bookkeeping connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
