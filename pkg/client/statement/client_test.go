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

	statement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/statement"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	ret = npool.Statement{
		ID:              uuid.NewString(),
		AppID:           uuid.NewString(),
		UserID:          uuid.NewString(),
		CoinTypeID:      uuid.NewString(),
		Amount:          "9999999999.999999999999",
		IOType:          basetypes.IOType_Incoming,
		IOTypeStr:       basetypes.IOType_Incoming.String(),
		IOSubType:       basetypes.IOSubType_Payment,
		IOSubTypeStr:    basetypes.IOSubType_Payment.String(),
		FromCoinTypeID:  uuid.New().String(),
		CoinUSDCurrency: "1.00045000000123012",
		IOExtra:         fmt.Sprintf(`{"OrderID": "%v", "PaymentID": "%v"}`, uuid.NewString(), uuid.NewString()),
	}
)

func createStatement(t *testing.T) {
	handler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithID(&ret.ID),
		statement1.WithAppID(&ret.AppID),
		statement1.WithUserID(&ret.UserID),
		statement1.WithCoinTypeID(&ret.CoinTypeID),
		statement1.WithAmount(&ret.Amount),
		statement1.WithIOType(&ret.IOType),
		statement1.WithIOSubType(&ret.IOSubType),
		statement1.WithFromCoinTypeID(&ret.FromCoinTypeID),
		statement1.WithCoinUSDCurrency(&ret.CoinUSDCurrency),
		statement1.WithIOExtra(&ret.IOExtra),
	)
	assert.Nil(t, err)

	info, err := handler.CreateStatement(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getStatement(t *testing.T) {
	info, err := GetStatement(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getStatementOnly(t *testing.T) {
	_, err := GetStatementOnly(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOType)},
		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOSubType)},
		Amount:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.Amount},
		IOExtra:    &commonpb.StringVal{Op: cruder.LIKE, Value: ret.IOExtra},
	})
	assert.Nil(t, err)
}

func getStatements(t *testing.T) {
	infos, _, err := GetStatements(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		IOType:     &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOType)},
		IOSubType:  &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(ret.IOSubType)},
		Amount:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.Amount},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteStatement(t *testing.T) {
	handler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteStatement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetStatement(context.Background())
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

	t.Run("createStatement", createStatement)
	t.Run("getStatement", getStatement)
	t.Run("getStatementOnly", getStatementOnly)
	t.Run("getStatements", getStatements)
	t.Run("deleteStatement", deleteStatement)
}
