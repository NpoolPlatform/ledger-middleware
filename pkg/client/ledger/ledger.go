//nolint:dupl
package ledger

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	detailpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/detail"
	generalpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/general"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v1/ledger"

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

func GetIntervalGenerals(
	ctx context.Context, appID, userID string, start, end uint32, limit, offset int32,
) (
	[]*generalpb.General, uint32, error,
) {
	var total uint32
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetIntervalGenerals(_ctx, &npool.GetIntervalGeneralsRequest{
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
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*generalpb.General), total, nil
}

func GetIntervalDetails(
	ctx context.Context, appID, userID string, start, end uint32, limit, offset int32,
) (
	[]*detailpb.Detail, uint32, error,
) {
	var total uint32
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetIntervalDetails(_ctx, &npool.GetIntervalDetailsRequest{
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
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*detailpb.Detail), total, nil
}

func GetIntervalProfits(
	ctx context.Context, appID, userID string, start, end uint32, limit, offset int32,
) (
	[]*detailpb.Detail, uint32, error,
) {
	var total uint32
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetIntervalProfits(_ctx, &npool.GetIntervalProfitsRequest{
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
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*detailpb.Detail), total, nil
}

func BookKeeping(ctx context.Context, in *detailpb.DetailReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.BookKeeping(_ctx, &npool.BookKeepingRequest{
			Info: in,
		})
		return nil, err
	})
	return err
}

func LockBalance(ctx context.Context, appID, userID, coinTypeID string, amount decimal.Decimal) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.LockBalance(_ctx, &npool.LockBalanceRequest{
			AppID:      appID,
			UserID:     userID,
			CoinTypeID: coinTypeID,
			Amount:     amount.String(),
		})
		return nil, err
	})
	return err
}

func UnlockBalance(
	ctx context.Context,
	appID, userID, coinTypeID string,
	ioSubType detailpb.IOSubType,
	unlocked, outcoming decimal.Decimal,
	ioExtra string,
) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.UnlockBalance(_ctx, &npool.UnlockBalanceRequest{
			AppID:      appID,
			UserID:     userID,
			CoinTypeID: coinTypeID,
			IOSubType:  ioSubType,
			Unlocked:   unlocked.String(),
			Outcoming:  outcoming.String(),
			IOExtra:    ioExtra,
		})
		return nil, err
	})
	return err
}
