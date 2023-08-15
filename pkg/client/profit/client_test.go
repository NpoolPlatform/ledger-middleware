package profit

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

	profit1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/profit"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"
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
	ret = npool.Profit{
		ID:         uuid.NewString(),
		AppID:      uuid.NewString(),
		UserID:     uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		Incoming:   "0.000000000000000000",
	}
)

func createProfit(t *testing.T) {
	info, err := CreateProfit(context.Background(), &npool.ProfitReq{
		ID:         &ret.ID,
		AppID:      &ret.AppID,
		UserID:     &ret.UserID,
		CoinTypeID: &ret.CoinTypeID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateProfit(t *testing.T) {
	ret.Incoming = "9999999.000000000000000000"

	handler, err := profit1.NewHandler(
		context.Background(),
		profit1.WithID(&ret.ID),
		profit1.WithIncoming(&ret.Incoming),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateProfit(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getProfit(t *testing.T) {
	info, err := GetProfit(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getProfits(t *testing.T) {
	infos, _, err := GetProfits(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteProfit(t *testing.T) {
	handler, err := profit1.NewHandler(
		context.Background(),
		profit1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteProfit(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetProfit(context.Background())
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

	t.Run("createProfit", createProfit)
	t.Run("updateProfit", updateProfit)
	t.Run("getProfit", getProfit)
	t.Run("getProfits", getProfits)
	t.Run("deleteProfit", deleteProfit)
}
