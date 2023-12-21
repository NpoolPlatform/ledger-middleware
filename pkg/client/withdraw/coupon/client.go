//nolint:dupl
package coupon

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/servicename"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
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

func CreateCouponWithdraw(ctx context.Context, in *npool.CouponWithdrawReq) (*npool.CouponWithdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCouponWithdraw(ctx, &npool.CreateCouponWithdrawRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create couponwithdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create couponwithdraw: %v", err)
	}
	return info.(*npool.CouponWithdraw), nil
}

func UpdateCouponWithdraw(ctx context.Context, in *npool.CouponWithdrawReq) (*npool.CouponWithdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateCouponWithdraw(ctx, &npool.UpdateCouponWithdrawRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update couponwithdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update withdraw: %v", err)
	}
	return info.(*npool.CouponWithdraw), nil
}

func GetCouponWithdraw(ctx context.Context, id string) (*npool.CouponWithdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCouponWithdraw(ctx, &npool.GetCouponWithdrawRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get couponwithdraw: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get withdraw: %v", err)
	}
	return info.(*npool.CouponWithdraw), nil
}

func GetCouponWithdrawOnly(ctx context.Context, conds *npool.Conds) (*npool.CouponWithdraw, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCouponWithdraws(ctx, &npool.GetCouponWithdrawsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get couponwithdraw only: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get couponwithdraw only: %v", err)
	}
	if len(infos.([]*npool.CouponWithdraw)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.CouponWithdraw)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.CouponWithdraw)[0], nil
}

func GetCouponWithdraws(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.CouponWithdraw, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCouponWithdraws(ctx, &npool.GetCouponWithdrawsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get couponwithdraws: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get couponwithdraws: %v", err)
	}
	return infos.([]*npool.CouponWithdraw), total, nil
}

func DeleteCouponWithdraw(ctx context.Context, in *npool.CouponWithdrawReq) (*npool.CouponWithdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteCouponWithdraw(ctx, &npool.DeleteCouponWithdrawRequest{
			Info: &npool.CouponWithdrawReq{
				ID:    in.ID,
				EntID: in.EntID,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.CouponWithdraw), nil
}
