//nolint:dupl
package detail

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get detail connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateDetail(ctx context.Context, in *npool.DetailReq) (*mgrpb.Detail, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateDetail(ctx, &npool.CreateDetailRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create detail: %v", err)
	}
	return info.(*mgrpb.Detail), nil
}

func GetDetailOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.Detail, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDetailOnly(ctx, &npool.GetDetailOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get detail: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get detail: %v", err)
	}
	return info.(*mgrpb.Detail), nil
}

func GetDetails(ctx context.Context, conds *mgrpb.Conds, offset, limit uint32) ([]*mgrpb.Detail, uint32, error) {
	var total = uint32(0)
	rows, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDetails(ctx, &npool.GetDetailsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get detail: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, total, fmt.Errorf("fail get detail: %v", err)
	}
	return rows.([]*mgrpb.Detail), total, nil
}
