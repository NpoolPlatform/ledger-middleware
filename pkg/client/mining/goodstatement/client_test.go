package goodstatement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodstatement"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodstatement"
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
	benefitDate = time.Now().Unix()
	ret         = npool.GoodStatement{
		ID:          uuid.NewString(),
		GoodID:      uuid.NewString(),
		CoinTypeID:  uuid.NewString(),
		Amount:      "9999999999.999999999999",
		BenefitDate: uint32(benefitDate),
	}
)

func getGoodStatement(t *testing.T) {
	handler, err := goodstatement1.NewHandler(
		context.Background(),
		goodstatement1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetGoodStatement(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getGoodStatementOnly(t *testing.T) {
	info, err := GetGoodStatementOnly(context.Background(), &npool.Conds{
		GoodID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	})
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func getGoodStatements(t *testing.T) {
	infos, _, err := GetGoodStatements(context.Background(), &npool.Conds{
		GoodID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteGoodStatement(t *testing.T) {
	handler, err := goodstatement1.NewHandler(
		context.Background(),
		goodstatement1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteGoodStatement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetGoodStatement(context.Background())
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

	t.Run("getGoodStatement", getGoodStatement)
	t.Run("getGoodStatementOnly", getGoodStatementOnly)
	t.Run("getGoodStatements", getGoodStatements)
	t.Run("deleteGoodStatement", deleteGoodStatement)
}
