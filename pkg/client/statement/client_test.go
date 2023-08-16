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

	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
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

	deposit = npool.Statement{
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "100.000000000000000000",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_Deposit,
		IOSubTypeStr: basetypes.IOSubType_Deposit.String(),
		IOExtra:      fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	payment = npool.Statement{
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "10.000000000000000000",
		IOType:       basetypes.IOType_Outcoming,
		IOTypeStr:    basetypes.IOType_Outcoming.String(),
		IOSubType:    basetypes.IOSubType_Payment,
		IOSubTypeStr: basetypes.IOSubType_Payment.String(),
		IOExtra:      fmt.Sprintf(`{"PaymentID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	miningBenefit = npool.Statement{
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "1.000000000000000000",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_MiningBenefit,
		IOSubTypeStr: basetypes.IOSubType_MiningBenefit.String(),
		IOExtra:      fmt.Sprintf(`{"GoodID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
	ledgerResult = ledgerpb.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "101.000000000000000000",
		Outcoming:  "10.000000000000000000",
		Locked:     "0.000000000000000000",
		Spendable:  "91.000000000000000000",
	}
)

func createStatements(t *testing.T) {
	reqs := []*npool.StatementReq{}
	reqs = append(reqs, &npool.StatementReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &deposit.Amount,
		IOType:     &deposit.IOType,
		IOSubType:  &deposit.IOSubType,
		IOExtra:    &deposit.IOExtra,
	})

	reqs = append(reqs, &npool.StatementReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &payment.Amount,
		IOType:     &payment.IOType,
		IOSubType:  &payment.IOSubType,
		IOExtra:    &payment.IOExtra,
	})

	reqs = append(reqs, &npool.StatementReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &miningBenefit.Amount,
		IOType:     &miningBenefit.IOType,
		IOSubType:  &miningBenefit.IOSubType,
		IOExtra:    &miningBenefit.IOExtra,
	})

	infos, err := CreateStatements(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, 3, len(infos))
	}

	handler, err := ledger1.NewHandler(
		context.Background(),
		ledger1.WithAppID(&appID),
		ledger1.WithUserID(&userID),
		ledger1.WithCoinTypeID(&coinTypeID),
	)
	assert.Nil(t, err)

	handler.Conds = &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *handler.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *handler.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *handler.CoinTypeID},
	}

	info, err := handler.GetLedgerOnly(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ledgerResult, info)
	}
}

// func getStatement(t *testing.T) {
// 	info, err := GetStatement(context.Background(), ret.ID)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, &ret, info)
// 	}
// }

// func getStatementOnly(t *testing.T) {
// 	_, err := GetStatementOnly(context.Background(), &npool.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
// 		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOType)},
// 		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOSubType)},
// 		Amount:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.Amount},
// 		IOExtra:    &commonpb.StringVal{Op: cruder.LIKE, Value: ret.IOExtra},
// 	})
// 	assert.Nil(t, err)
// }

// func getStatements(t *testing.T) {
// 	infos, _, err := GetStatements(context.Background(), &npool.Conds{
// 		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
// 		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
// 		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOType)},
// 		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOSubType)},
// 		Amount:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.Amount},
// 	}, 0, 1)
// 	if assert.Nil(t, err) {
// 		assert.NotEqual(t, len(infos), 0)
// 	}
// }

// func deleteStatement(t *testing.T) {
// 	handler, err := statement1.NewHandler(
// 		context.Background(),
// 		statement1.WithID(&ret.ID),
// 	)
// 	assert.Nil(t, err)

// 	info, err := handler.DeleteStatement(context.Background())
// 	assert.Nil(t, err)
// 	assert.NotNil(t, info)

// 	info, err = handler.GetStatement(context.Background())
// 	assert.Nil(t, err)
// 	assert.Nil(t, info)
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
	// t.Run("getStatement", getStatement)
	// t.Run("getStatementOnly", getStatementOnly)
	// t.Run("getStatements", getStatements)
	// t.Run("deleteStatement", deleteStatement)
}
