package lock

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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	statementcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/statement"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	lockpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/lock"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
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

	deposit1 = npool.Statement{
		AppID:           appID,
		UserID:          userID,
		CoinTypeID:      coinTypeID,
		Amount:          "100",
		IOType:          basetypes.IOType_Incoming,
		IOTypeStr:       basetypes.IOType_Incoming.String(),
		IOSubType:       basetypes.IOSubType_Deposit,
		IOSubTypeStr:    basetypes.IOSubType_Deposit.String(),
		IOExtra:         fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString()),
		FromCoinTypeID:  "00000000-0000-0000-0000-000000000000",
		CoinUSDCurrency: "0",
	}

	deposit2 = npool.Statement{
		AppID:           appID,
		UserID:          userID,
		CoinTypeID:      coinTypeID,
		Amount:          "50",
		IOType:          basetypes.IOType_Incoming,
		IOTypeStr:       basetypes.IOType_Incoming.String(),
		IOSubType:       basetypes.IOSubType_Deposit,
		IOSubTypeStr:    basetypes.IOSubType_Deposit.String(),
		IOExtra:         fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString()),
		FromCoinTypeID:  "00000000-0000-0000-0000-000000000000",
		CoinUSDCurrency: "0",
	}

	req = lockpb.BalanceReq{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Amount:     "10",
	}

	ledgerResult1 = ledgerpb.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "150",
		Outcoming:  "0",
		Locked:     "10",
		Spendable:  "140",
	}

	ledgerResult2 = ledgerpb.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "150",
		Outcoming:  "0",
		Locked:     "0",
		Spendable:  "150",
	}
)

func createStatements(t *testing.T) {
	deposits, err := statementcli.CreateStatements(context.Background(), []*npool.StatementReq{
		{
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &deposit1.Amount,
			IOType:     &deposit1.IOType,
			IOSubType:  &deposit1.IOSubType,
			IOExtra:    &deposit1.IOExtra,
		}, {
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &deposit2.Amount,
			IOType:     &deposit2.IOType,
			IOSubType:  &deposit2.IOSubType,
			IOExtra:    &deposit2.IOExtra,
		},
		{
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &deposit2.Amount,
			IOType:     &deposit2.IOType,
			IOSubType:  &deposit2.IOSubType,
			IOExtra:    &deposit2.IOExtra,
		}})
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(deposits)) // the same batch of data cannot be written repeatedly.
	}
}

// func compareLedger(t *testing.T) {
// 	info, err := ledgercli.GetLedgerOnly(context.Background(), &ledgerpb.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
// 	})
// 	if assert.Nil(t, err) {
// 		assert.NotNil(t, info)
// 		ledgerResult.ID = info.ID
// 		ledgerResult.CreatedAt = info.CreatedAt
// 		ledgerResult.UpdatedAt = info.UpdatedAt
// 		assert.Equal(t, &ledgerResult, info)
// 	}
// }

// func compareProfit(t *testing.T) {
// 	info, err := profitcli.GetProfitOnly(context.Background(), &profitpb.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
// 	})
// 	if assert.Nil(t, err) {
// 		assert.NotNil(t, info)
// 		profit.ID = info.ID
// 		profit.CreatedAt = info.CreatedAt
// 		profit.UpdatedAt = info.UpdatedAt
// 	}

// 	info, err = profitcli.GetProfit(context.Background(), info.ID)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, info)

// 	infos, _, err := profitcli.GetProfits(context.Background(), &profitpb.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
// 	}, 0, 1)
// 	if assert.Nil(t, err) {
// 		assert.NotEqual(t, len(infos), 0)
// 	}
// }

// func getStatement(t *testing.T) {
// 	info, err := GetStatement(context.Background(), deposit.ID)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, &deposit, info)
// 	}
// }

// func getStatementOnly(t *testing.T) {
// 	info, err := GetStatementOnly(context.Background(), &npool.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
// 		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOType)},
// 		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOSubType)},
// 		IOExtra:    &commonpb.StringVal{Op: cruder.LIKE, Value: deposit.IOExtra},
// 	})
// 	if assert.Nil(t, err) {
// 		assert.NotNil(t, info)
// 	}
// }

// func getStatements(t *testing.T) {
// 	infos, _, err := GetStatements(context.Background(), &npool.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: userID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: coinTypeID},
// 		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOType)},
// 		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(deposit.IOSubType)},
// 		IOExtra:    &commonpb.StringVal{Op: cruder.LIKE, Value: deposit.IOExtra},
// 	}, 0, 1)
// 	if assert.Nil(t, err) {
// 		assert.NotEqual(t, len(infos), 0)
// 	}
// }

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createStatements", createStatements)
	// t.Run("compareLedger", compareLedger)
	// t.Run("compareProfit", compareProfit)
	// t.Run("getStatement", getStatement)
	// t.Run("getStatementOnly", getStatementOnly)
	// t.Run("getStatements", getStatements)
	// t.Run("unCreateStatements", unCreateStatements)
}
