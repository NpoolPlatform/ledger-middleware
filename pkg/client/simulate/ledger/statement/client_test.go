package statement

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

	ledgercli "github.com/NpoolPlatform/ledger-middleware/pkg/client/simulate/ledger"
	profitcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/simulate/ledger/profit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger"
	profitpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/profit"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/statement"
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
	appID      = uuid.NewString()
	userID     = uuid.NewString()
	coinTypeID = uuid.NewString()

	deposit = npool.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "100",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_Deposit,
		IOSubTypeStr: basetypes.IOSubType_Deposit.String(),
		IOExtra:      fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	payment = npool.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "10",
		IOType:       basetypes.IOType_Outcoming,
		IOTypeStr:    basetypes.IOType_Outcoming.String(),
		IOSubType:    basetypes.IOSubType_Payment,
		IOSubTypeStr: basetypes.IOSubType_Payment.String(),
		IOExtra:      fmt.Sprintf(`{"PaymentID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	miningBenefit = npool.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "1",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_MiningBenefit,
		IOSubTypeStr: basetypes.IOSubType_MiningBenefit.String(),
		IOExtra:      fmt.Sprintf(`{"GoodID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	ledgerResult = ledgerpb.Ledger{
		EntID:      "",
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "101",
		Outcoming:  "10",
	}
)

func setup(t *testing.T) func(*testing.T) {
	deposits, err := CreateStatements(context.Background(), []*npool.StatementReq{{
		EntID:      &deposit.EntID,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &deposit.Amount,
		IOType:     &deposit.IOType,
		IOSubType:  &deposit.IOSubType,
		IOExtra:    &deposit.IOExtra,
	}})
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(deposits))
		deposit.CreatedAt = deposits[0].CreatedAt
		deposit.UpdatedAt = deposits[0].UpdatedAt
		deposit.ID = deposits[0].ID
		assert.Equal(t, &deposit, deposits[0])
	}

	payments, err := CreateStatements(context.Background(), []*npool.StatementReq{{
		EntID:      &payment.EntID,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &payment.Amount,
		IOType:     &payment.IOType,
		IOSubType:  &payment.IOSubType,
		IOExtra:    &payment.IOExtra,
	}})
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(payments))
		payment.CreatedAt = payments[0].CreatedAt
		payment.UpdatedAt = payments[0].UpdatedAt
		payment.ID = payments[0].ID
		assert.Equal(t, &payment, payments[0])
	}

	benefits, err := CreateStatements(context.Background(), []*npool.StatementReq{{
		EntID:      &miningBenefit.EntID,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &miningBenefit.Amount,
		IOType:     &miningBenefit.IOType,
		IOSubType:  &miningBenefit.IOSubType,
		IOExtra:    &miningBenefit.IOExtra,
	}})
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(benefits))
		miningBenefit.CreatedAt = benefits[0].CreatedAt
		miningBenefit.UpdatedAt = benefits[0].UpdatedAt
		miningBenefit.ID = benefits[0].ID
		assert.Equal(t, &miningBenefit, benefits[0])
	}
	return func(t *testing.T) {
		_, _ = DeleteStatement(context.Background(), &npool.StatementReq{EntID: &payment.EntID})
		_, _ = DeleteStatement(context.Background(), &npool.StatementReq{EntID: &miningBenefit.EntID})
	}
}

func compareLedger(t *testing.T) {
	info, err := ledgercli.GetLedgerOnly(context.Background(), &ledgerpb.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult.ID = info.ID
		ledgerResult.EntID = info.EntID
		ledgerResult.CreatedAt = info.CreatedAt
		ledgerResult.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult, info)
	}
}

var (
	profit = profitpb.Profit{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "1",
	}
)

func compareProfit(t *testing.T) {
	info, err := profitcli.GetProfitOnly(context.Background(), &profitpb.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		profit.ID = info.ID
		profit.EntID = info.EntID
		profit.CreatedAt = info.CreatedAt
		profit.UpdatedAt = info.UpdatedAt
	}

	info, err = profitcli.GetProfit(context.Background(), info.EntID)
	assert.Nil(t, err)
	assert.NotNil(t, info)

	infos, _, err := profitcli.GetProfits(context.Background(), &profitpb.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getStatement(t *testing.T) {
	info, err := GetStatement(context.Background(), deposit.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &deposit, info)
	}
}

//nolint
func getStatementOnly(t *testing.T) {
	info, err := GetStatementOnly(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOType)},
		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOSubType)},
		IOExtra:    &commonpb.StringVal{Op: cruder.LIKE, Value: deposit.IOExtra},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
	}
}

//nolint
func getStatements(t *testing.T) {
	infos, _, err := GetStatements(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOType)},
		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOSubType)},
		IOExtra:    &commonpb.StringVal{Op: cruder.LIKE, Value: deposit.IOExtra},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

var (
	ledgerResult2 = ledgerpb.Ledger{
		EntID:      "",
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "100",
		Outcoming:  "10",
	}
)

func rollbackStatements(t *testing.T) {
	infos, err := DeleteStatements(context.Background(), []*npool.StatementReq{{
		EntID:      &miningBenefit.EntID,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &miningBenefit.Amount,
		IOType:     &miningBenefit.IOType,
		IOSubType:  &miningBenefit.IOSubType,
		IOExtra:    &miningBenefit.IOExtra,
	}})
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(infos))
	assert.Equal(t, &miningBenefit, infos[0])

	ledgerResult2.ID = ledgerResult.ID
	ledgerResult2.EntID = ledgerResult.EntID
	info, err := ledgercli.GetLedger(context.Background(), ledgerResult.EntID)
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult2.CreatedAt = info.CreatedAt
		ledgerResult2.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult2, info)
	}
}

func compareProfit1(t *testing.T) {
	info, err := profitcli.GetProfit(context.Background(), profit.EntID)
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		assert.Equal(t, "0", info.Incoming)
	}
}

func tryGetStatement(t *testing.T) {
	info, err := GetStatement(context.Background(), profit.EntID)
	assert.Nil(t, err)
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

	teardown := setup(t)
	defer teardown(t)
	t.Run("compareLedger", compareLedger)
	t.Run("compareProfit", compareProfit)
	t.Run("getStatement", getStatement)
	t.Run("getStatementOnly", getStatementOnly)
	t.Run("getStatements", getStatements)
	t.Run("rollbackStatements", rollbackStatements)
	t.Run("compareProfit1", compareProfit1)
	t.Run("tryGetStatement", tryGetStatement)
}
