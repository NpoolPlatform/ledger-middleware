package withdraw

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

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
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
	ret = npool.Withdraw{
		ID:                    uuid.NewString(),
		AppID:                 uuid.NewString(),
		UserID:                uuid.NewString(),
		CoinTypeID:            uuid.NewString(),
		AccountID:             uuid.NewString(),
		Address:               uuid.NewString(),
		Amount:                "999.999999999",
		State:                 basetypes.WithdrawState_Reviewing,
		StateStr:              basetypes.WithdrawState_Reviewing.String(),
		PlatformTransactionID: "00000000-0000-0000-0000-000000000000",
	}
)

func createWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithAccountID(&ret.AccountID, true),
		WithAddress(&ret.Address, true),
		WithAmount(&ret.Amount, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateWithdraw(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateWithdraw(t *testing.T) {
	ret.State = basetypes.WithdrawState_Rejected
	ret.StateStr = basetypes.WithdrawState_Rejected.String()
	ret.PlatformTransactionID = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithState(&ret.State, false),
		WithPlatformTransactionID(&ret.PlatformTransactionID, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateWithdraw(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetWithdraw(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getWithdraws(t *testing.T) {
	conds := &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		AccountID:  &commonpb.StringVal{Op: cruder.EQ, Value: ret.AccountID},
		State:      &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.State)},
		Amount:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.Amount},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetWithdraws(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteWithdraw(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteWithdraw(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetWithdraw(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestWithdraw(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createWithdraw", createWithdraw)
	t.Run("updateWithdraw", updateWithdraw)
	t.Run("getWithdraw", getWithdraw)
	t.Run("getWithdraws", getWithdraws)
	t.Run("deleteWithdraw", deleteWithdraw)
}
