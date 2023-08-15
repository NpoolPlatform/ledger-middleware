package unsoldstatement

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

	unsoldstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/unsoldstatement"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"
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
	ret         = npool.UnsoldStatement{
		ID:          uuid.NewString(),
		GoodID:      uuid.NewString(),
		CoinTypeID:  uuid.NewString(),
		Amount:      "9999999999.999999999999",
		BenefitDate: uint32(benefitDate),
	}
)

func createUnsoldStatement(t *testing.T) {
	handler, err := unsoldstatement1.NewHandler(
		context.Background(),
		unsoldstatement1.WithID(&ret.ID),
		unsoldstatement1.WithGoodID(&ret.GoodID),
		unsoldstatement1.WithCoinTypeID(&ret.CoinTypeID),
		unsoldstatement1.WithAmount(&ret.Amount),
		unsoldstatement1.WithBenefitDate(&ret.BenefitDate),
	)
	assert.Nil(t, err)

	info, err := handler.CreateUnsoldStatement(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getUnsoldStatement(t *testing.T) {
	handler, err := unsoldstatement1.NewHandler(
		context.Background(),
		unsoldstatement1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetUnsoldStatement(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getUnsoldStatementOnly(t *testing.T) {
	info, err := GetUnsoldStatementOnly(context.Background(), &npool.Conds{
		GoodID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	})
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func deleteUnsoldStatement(t *testing.T) {
	handler, err := unsoldstatement1.NewHandler(
		context.Background(),
		unsoldstatement1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteUnsoldStatement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetUnsoldStatement(context.Background())
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

	t.Run("createUnsoldStatement", createUnsoldStatement)
	t.Run("getUnsoldStatement", getUnsoldStatement)
	t.Run("getUnsoldStatementOnly", getUnsoldStatementOnly)
	t.Run("deleteUnsoldStatement", deleteUnsoldStatement)
}
