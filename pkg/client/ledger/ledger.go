package ledger

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	generalpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/general"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v1/ledger"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetIntervalGenerals(
	ctx context.Context, appID, userID string, start, end uint32, limit, offset int32,
) (
	[]*generalpb.General, uint32, error,
) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetIntervalGenerals(ctx, &npool.GetIntervalGeneralsRequest{
			AppID:  appID,
			UserID: userID,
			Start:  start,
			End:    end,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*generalpb.General), total, nil
}
