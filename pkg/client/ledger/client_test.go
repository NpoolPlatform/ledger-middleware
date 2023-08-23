package ledger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	statementcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	ledgerpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	locked    = "10"
	ioSubType = basetypes.IOSubType_Withdrawal
	ioExtra   = fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString())
	req       = ledgerpb.LedgerReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		IOSubType:  &ioSubType,
		IOExtra:    &ioExtra,
		Spendable:  &locked,
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

	spendResult = ledgerpb.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "150",
		Outcoming:  "10",
		Locked:     "0",
		Spendable:  "140",
	}
)

func insertSameDataTwice(t *testing.T) {
	_, err := statementcli.CreateStatements(context.Background(), []*npool.StatementReq{
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
		},
	})
	assert.NotNil(t, err) // the same batch of data cannot be written repeatedly.
}

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
	})
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(deposits))
	}
}

func lockBalance(t *testing.T) {
	info, err := SubBalance(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult1.ID = info.ID
		ledgerResult1.CreatedAt = info.CreatedAt
		ledgerResult1.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult1, info)
	}
}

func unlockBalance(t *testing.T) {
	info, err := AddBalance(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult2.ID = info.ID
		ledgerResult2.CreatedAt = info.CreatedAt
		ledgerResult2.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult2, info)
	}
}

var (
	spendable = "10"
	spendReq  = &ledgerpb.LedgerReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		IOSubType:  &ioSubType,
		IOExtra:    &ioExtra,
		Locked:     &spendable,
	}
)

func spendBalance(t *testing.T) {
	// lock
	info, err := SubBalance(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
	}

	// spend
	info1, err := SubBalance(context.Background(), spendReq)
	if assert.Nil(t, err) {
		assert.NotNil(t, info1)
		assert.Equal(t, &spendResult, info1)
	}
}

func unspendBalance(t *testing.T) {
	// unspend
	info, err := AddBalance(context.Background(), spendReq)
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		assert.Equal(t, &ledgerResult1, info)
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

	t.Run("insertSameDataTwice", insertSameDataTwice)
	t.Run("createStatements", createStatements)
	t.Run("lockBalance", lockBalance)
	t.Run("unlockBalance", unlockBalance)
	t.Run("spendBalance", spendBalance)
	t.Run("unspendBalance", unspendBalance)
}
