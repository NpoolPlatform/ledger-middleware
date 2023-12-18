package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.CouponWithdraw{
		EntID:      uuid.NewString(),
		AppID:      uuid.NewString(),
		UserID:     uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		CouponID:   uuid.Nil.String(),
		Amount:     "999.999999999",
		State:      types.WithdrawState_Reviewing,
		StateStr:   types.WithdrawState_Reviewing.String(),
		ReviewID:   uuid.NewString(),
	}
)

func createCouponWithdraw(t *testing.T) {
	info, err := CreateCouponWithdraw(context.Background(), &npool.CouponWithdrawReq{
		EntID:      &ret.EntID,
		AppID:      &ret.AppID,
		UserID:     &ret.UserID,
		CoinTypeID: &ret.CoinTypeID,
		CouponID:   &ret.CouponID,
		Amount:     &ret.Amount,
		ReviewID:   &ret.ReviewID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func updateCouponWithdraw(t *testing.T) {
	ret.State = types.WithdrawState_Approved
	ret.StateStr = types.WithdrawState_Approved.String()

	info, err := UpdateCouponWithdraw(context.Background(), &npool.CouponWithdrawReq{
		ID:    &ret.ID,
		State: &ret.State,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getCouponWithdraw(t *testing.T) {
	info, err := GetCouponWithdraw(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getCouponWithdraws(t *testing.T) {
	infos, _, err := GetCouponWithdraws(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		ReviewID:   &commonpb.StringVal{Op: cruder.EQ, Value: ret.ReviewID},
		CouponID:   &commonpb.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		State:      &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.State)},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteCouponWithdraw(t *testing.T) {
	info, err := DeleteCouponWithdraw(context.Background(), &npool.CouponWithdrawReq{EntID: &ret.EntID})
	assert.NotNil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCouponWithdraw", createCouponWithdraw)
	t.Run("updateCouponWithdraw", updateCouponWithdraw)
	t.Run("getCouponWithdraw", getCouponWithdraw)
	t.Run("getCouponWithdraws", getCouponWithdraws)
	t.Run("deleteCouponWithdraw", deleteCouponWithdraw)
}
