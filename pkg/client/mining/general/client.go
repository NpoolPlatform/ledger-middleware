//nolint:dupl
package general

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get general connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateGeneral(ctx context.Context, in *mgrpb.GeneralReq) (*mgrpb.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateGeneral(ctx, &npool.CreateGeneralRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create general: %v", err)
	}
	return info.(*mgrpb.General), nil
}

func GetGeneralOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetGeneralOnly(ctx, &npool.GetGeneralOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get general: %v", err)
	}
	return info.(*mgrpb.General), nil
}

func AddGeneral(ctx context.Context, in *mgrpb.GeneralReq) (*mgrpb.General, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.AddGeneral(ctx, &npool.AddGeneralRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add general: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail add general: %v", err)
	}
	return info.(*mgrpb.General), nil
}
