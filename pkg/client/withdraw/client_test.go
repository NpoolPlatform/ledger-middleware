package withdraw

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

	statementmwcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/ledger/statement"
	types "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	statementmwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
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
	ret = npool.Withdraw{
		ID:                    uuid.NewString(),
		AppID:                 uuid.NewString(),
		UserID:                uuid.NewString(),
		CoinTypeID:            uuid.NewString(),
		AccountID:             uuid.NewString(),
		Address:               uuid.NewString(),
		Amount:                "999.999999999",
		State:                 types.WithdrawState_Reviewing,
		StateStr:              types.WithdrawState_Reviewing.String(),
		PlatformTransactionID: "00000000-0000-0000-0000-000000000000",
	}
)

func createLedger(t *testing.T) {
	ioType := types.IOType_Incoming
	ioSubType := types.IOSubType_Deposit
	ioExtra := "{}"
	amount := "10000"

	info, err := statementmwcli.CreateStatement(
		context.Background(),
		&statementmwpb.StatementReq{
			AppID:      &ret.AppID,
			UserID:     &ret.UserID,
			CoinTypeID: &ret.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			IOExtra:    &ioExtra,
			Amount:     &amount,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func createWithdraw(t *testing.T) {
	info, err := CreateWithdraw(context.Background(), &npool.WithdrawReq{
		ID:         &ret.ID,
		AppID:      &ret.AppID,
		UserID:     &ret.UserID,
		CoinTypeID: &ret.CoinTypeID,
		AccountID:  &ret.AccountID,
		Address:    &ret.Address,
		Amount:     &ret.Amount,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateWithdraw(t *testing.T) {
	ret.State = types.WithdrawState_Rejected
	ret.StateStr = types.WithdrawState_Rejected.String()
	ret.PlatformTransactionID = uuid.NewString()

	info, err := UpdateWithdraw(context.Background(), &npool.WithdrawReq{
		ID:                    &ret.ID,
		PlatformTransactionID: &ret.PlatformTransactionID,
		State:                 &ret.State,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getWithdraw(t *testing.T) {
	info, err := GetWithdraw(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getWithdraws(t *testing.T) {
	infos, _, err := GetWithdraws(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		AccountID:  &commonpb.StringVal{Op: cruder.EQ, Value: ret.AccountID},
		State:      &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.State)},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteWithdraw(t *testing.T) {
	info, err := DeleteWithdraw(context.Background(), &npool.WithdrawReq{ID: &ret.ID})
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createLedger", createLedger)
	t.Run("createWithdraw", createWithdraw)
	t.Run("updateWithdraw", updateWithdraw)
	t.Run("getWithdraw", getWithdraw)
	t.Run("getWithdraws", getWithdraws)
	t.Run("deleteWithdraw", deleteWithdraw)
}
