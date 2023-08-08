//nolint:dupl
package unsold

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get unsold connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateUnsold(ctx context.Context, in *npool.UnsoldReq) (*mgrpb.Unsold, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateUnsold(ctx, &npool.CreateUnsoldRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create unsold: %v", err)
	}
	return info.(*mgrpb.Unsold), nil
}

func GetUnsoldOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.Unsold, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUnsoldOnly(ctx, &npool.GetUnsoldOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get unsold: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get unsold: %v", err)
	}
	return info.(*mgrpb.Unsold), nil
}
