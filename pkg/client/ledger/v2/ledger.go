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

var timeout = 1200 * time.Second

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
		res, err := cli.GetGeneralOnly(_ctx, &npool.GetGeneralOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*generalpb.General), err
}

func AddGeneral(ctx context.Context, in *generalpb.GeneralReq) (*generalpb.General, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		res, err := cli.AddGeneral(_ctx, &npool.AddGeneralRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*generalpb.General), err
}

func GetDetails(ctx context.Context, conds *detailpb.Conds, offset, limit uint32) ([]*detailpb.Detail, uint32, error) {
	var total uint32
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		res, err := cli.GetDetails(_ctx, &npool.GetDetailsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = res.GetTotal()
		return res.Infos, err
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*detailpb.Detail), total, err
}
