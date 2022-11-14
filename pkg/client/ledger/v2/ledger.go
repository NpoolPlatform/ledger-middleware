package v2

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	detailpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/detail"
	generalpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/general"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withClient(ctx context.Context, handler handler) (cruder.Any, error) {
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

func BookKeeping(ctx context.Context, infos []*detailpb.DetailReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.BookKeeping(_ctx, &npool.BookKeepingRequest{
			Infos: infos,
		})
		return nil, err
	})
	return err
}

func GetGeneralOnly(ctx context.Context, conds *generalpb.Conds) (*generalpb.General, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		info, err := cli.GetGeneralOnly(_ctx, &npool.GetGeneralOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*generalpb.General), err
}
