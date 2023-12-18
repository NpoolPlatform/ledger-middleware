package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
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
		CouponID:   uuid.NewString(),
		Amount:     "999.999999999",
		State:      types.WithdrawState_Reviewing,
		StateStr:   types.WithdrawState_Reviewing.String(),
		ReviewID:   uuid.Nil.String(),
	}
)

func createCouponWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithCouponID(&ret.CouponID, true),
		WithAmount(&ret.Amount, true),
		WithReviewID(&ret.ReviewID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCouponWithdraw(context.Background())
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
	ret.ReviewID = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithState(&ret.State, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCouponWithdraw(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getCouponWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCouponWithdraw(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getCouponWithdraws(t *testing.T) {
	conds := &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		CouponID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		ReviewID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.ReviewID},
		State:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.State)},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCouponWithdraws(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteCouponWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCouponWithdraw(context.Background())
	assert.NotNil(t, err)
	assert.Nil(t, info)
}

func TestCouponWithdraw(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createCouponWithdraw", createCouponWithdraw)
	t.Run("updateCouponWithdraw", updateCouponWithdraw)
	t.Run("getCouponWithdraw", getCouponWithdraw)
	t.Run("getCouponWithdraws", getCouponWithdraws)
	t.Run("deleteCouponWithdraw", deleteCouponWithdraw)
}
