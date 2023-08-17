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
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
		ID:              "",
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
	payment = npool.Statement{
		ID:              "",
		AppID:           appID,
		UserID:          userID,
		CoinTypeID:      coinTypeID,
		Amount:          "10",
		IOType:          basetypes.IOType_Outcoming,
		IOTypeStr:       basetypes.IOType_Outcoming.String(),
		IOSubType:       basetypes.IOSubType_Payment,
		IOSubTypeStr:    basetypes.IOSubType_Payment.String(),
		IOExtra:         fmt.Sprintf(`{"PaymentID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
		FromCoinTypeID:  "00000000-0000-0000-0000-000000000000",
		CoinUSDCurrency: "0",
	}
	miningBenefit = npool.Statement{
		ID:              "",
		AppID:           appID,
		UserID:          userID,
		CoinTypeID:      coinTypeID,
		Amount:          "1",
		IOType:          basetypes.IOType_Incoming,
		IOTypeStr:       basetypes.IOType_Incoming.String(),
		IOSubType:       basetypes.IOSubType_MiningBenefit,
		IOSubTypeStr:    basetypes.IOSubType_MiningBenefit.String(),
		IOExtra:         fmt.Sprintf(`{"GoodID": "%v", "OrderID": "%v"}`, uuid.NewString(), uuid.NewString()),
		FromCoinTypeID:  "00000000-0000-0000-0000-000000000000",
		CoinUSDCurrency: "0",
	}
	ledgerResult = ledgerpb.Ledger{
		ID:         "",
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "101",
		Outcoming:  "10",
		Locked:     "0",
		Spendable:  "91",
	}
)

func createStatements(t *testing.T) {
	deposits, err := CreateStatements(context.Background(), []*npool.StatementReq{{
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
		ledgerResult.ID = info.ID
		assert.Equal(t, &ledgerResult, info)
	}
}

func getStatement(t *testing.T) {
	info, err := GetStatement(context.Background(), deposit.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &deposit, info)
	}
}

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
		ID:         "",
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "101",
		Outcoming:  "0",
		Locked:     "0",
		Spendable:  "101",
	}
)

func unCreateStatements(t *testing.T) {
	_, err := UnCreateStatements(context.Background(), []*npool.StatementReq{{
		ID:         &payment.ID,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Amount:     &payment.Amount,
		IOType:     &payment.IOType,
		IOSubType:  &payment.IOSubType,
		IOExtra:    &payment.IOExtra,
	}})
	assert.Nil(t, err)

	ledgerResult2.ID = ledgerResult.ID
	handler, err := ledger1.NewHandler(
		context.Background(),
		ledger1.WithID(&ledgerResult.ID),
	)

	assert.Nil(t, err)
	info, err := handler.GetLedger(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ledgerResult2, info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createStatements", createStatements)
	t.Run("getStatement", getStatement)
	t.Run("getStatementOnly", getStatementOnly)
	t.Run("getStatements", getStatements)
	t.Run("unCreateStatements", unCreateStatements)
}
