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
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
	statementmwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
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
	lockID     = uuid.NewString()
	deposit1   = statementmwpb.Statement{
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

	deposit2 = statementmwpb.Statement{
		EntID:        uuid.NewString(),
		AppID:        appID,
		UserID:       userID,
		CoinTypeID:   coinTypeID,
		Amount:       "50",
		IOType:       basetypes.IOType_Incoming,
		IOTypeStr:    basetypes.IOType_Incoming.String(),
		IOSubType:    basetypes.IOSubType_Deposit,
		IOSubTypeStr: basetypes.IOSubType_Deposit.String(),
		IOExtra:      fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
)

func setup(t *testing.T) func(*testing.T) {
	deposits, err := statementcli.CreateStatements(context.Background(), []*statementmwpb.StatementReq{
		{
			EntID:      &deposit1.EntID,
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Amount:     &deposit1.Amount,
			IOType:     &deposit1.IOType,
			IOSubType:  &deposit1.IOSubType,
			IOExtra:    &deposit1.IOExtra,
		}, {
			EntID:      &deposit2.EntID,
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
	return func(t *testing.T) {
		_, _ = statementcli.DeleteStatement(context.Background(), &statementmwpb.StatementReq{EntID: &deposit2.EntID})
	}
}

var (
	locked        = "10"
	ioSubType     = basetypes.IOSubType_Withdrawal
	ioExtra       = fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString())
	ledgerResult1 = npool.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "150",
		Outcoming:  "0",
		Locked:     "10",
		Spendable:  "140",
	}

	ledgerResult2 = npool.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "150",
		Outcoming:  "0",
		Locked:     "0",
		Spendable:  "150",
	}

	spendResult = npool.Ledger{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Incoming:   "150",
		Outcoming:  "10",
		Locked:     "0",
		Spendable:  "140",
	}
)

func lockBalance(t *testing.T) {
	info, err := LockBalance(context.Background(), &npool.LockBalanceRequest{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Amount:     locked,
		LockID:     lockID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult1.ID = info.ID
		ledgerResult1.CreatedAt = info.CreatedAt
		ledgerResult1.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult1, info)
	}
}

func unlockBalance(t *testing.T) {
	info, err := UnlockBalance(context.Background(), &npool.UnlockBalanceRequest{
		LockID: lockID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		ledgerResult2.ID = info.ID
		ledgerResult2.CreatedAt = info.CreatedAt
		ledgerResult2.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ledgerResult2, info)
	}
}

var (
	statementID = uuid.NewString()
	lockID1     = uuid.NewString()
)

func spendBalance(t *testing.T) {
	// lock
	info, err := LockBalance(context.Background(), &npool.LockBalanceRequest{
		AppID:      appID,
		UserID:     userID,
		CoinTypeID: coinTypeID,
		Amount:     locked,
		LockID:     lockID1,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
	}

	// spend
	info1, err := SettleBalance(context.Background(), &npool.SettleBalanceRequest{
		LockID:      lockID1,
		IOSubType:   ioSubType,
		IOExtra:     ioExtra,
		StatementID: statementID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info1)
		spendResult.ID = info1.ID
		spendResult.CreatedAt = info1.CreatedAt
		spendResult.UpdatedAt = info1.UpdatedAt
		assert.Equal(t, &spendResult, info1)
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardowm := setup(t)
	defer teardowm(t)

	t.Run("lockBalance", lockBalance)
	t.Run("unlockBalance", unlockBalance)
	t.Run("spendBalance", spendBalance)
}
