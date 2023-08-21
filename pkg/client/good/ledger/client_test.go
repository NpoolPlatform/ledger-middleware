package goodledger

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"bou.ke/monkey"
// 	"github.com/NpoolPlatform/go-service-framework/pkg/config"
// 	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
// 	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
// 	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"

// 	goodledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/mining/goodledger"
// 	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
// 	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func init() {
// 	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
// 		return
// 	}
// 	if err := testinit.Init(); err != nil {
// 		fmt.Printf("cannot init test stub: %v\n", err)
// 	}
// }

// var (
// 	ret = npool.GoodLedger{
// 		ID:         uuid.NewString(),
// 		GoodID:     uuid.NewString(),
// 		CoinTypeID: uuid.NewString(),
// 		Amount:     "0.000000000000000000",
// 		ToPlatform: "0.000000000000000000",
// 		ToUser:     "0.000000000000000000",
// 	}
// )

// func updateGoodLedger(t *testing.T) {
// 	ret.Amount = "10.000000000000000000"
// 	ret.ToPlatform = "3.000000000000000000"
// 	ret.ToUser = "7.000000000000000000"

// 	handler, err := goodledger1.NewHandler(
// 		context.Background(),
// 		goodledger1.WithID(&ret.ID),
// 		goodledger1.WithAmount(&ret.Amount),
// 		goodledger1.WithToPlatform(&ret.ToPlatform),
// 		goodledger1.WithToUser(&ret.ToUser),
// 	)
// 	assert.Nil(t, err)

// 	info, err := handler.UpdateGoodLedger(context.Background())
// 	if assert.Nil(t, err) {
// 		ret.UpdatedAt = info.UpdatedAt
// 		assert.Equal(t, &ret, info)
// 	}
// }

// func getGoodLedger(t *testing.T) {
// 	handler, err := goodledger1.NewHandler(
// 		context.Background(),
// 		goodledger1.WithID(&ret.ID),
// 	)
// 	assert.Nil(t, err)

// 	info, err := handler.GetGoodLedger(context.Background())
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, &ret, info)
// 	}
// }

// func getGoodLedgerOnly(t *testing.T) {
// 	info, err := GetGoodLedgerOnly(context.Background(), &npool.Conds{
// 		GoodID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.GoodID},
// 		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
// 	})
// 	assert.Nil(t, err)
// 	assert.NotNil(t, info)
// }

// func deleteGoodLedger(t *testing.T) {
// 	handler, err := goodledger1.NewHandler(
// 		context.Background(),
// 		goodledger1.WithID(&ret.ID),
// 	)
// 	assert.Nil(t, err)

// 	info, err := handler.DeleteGoodLedger(context.Background())
// 	assert.Nil(t, err)
// 	assert.NotNil(t, info)

// 	info, err = handler.GetGoodLedger(context.Background())
// 	assert.Nil(t, err)
// 	assert.Nil(t, info)
// }

// func TestClient(t *testing.T) {
// 	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
// 		return
// 	}

// 	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

// 	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
// 		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	})

// 	t.Run("updateGoodLedger", updateGoodLedger)
// 	t.Run("getGoodLedger", getGoodLedger)
// 	t.Run("getGoodLedgerOnly", getGoodLedgerOnly)
// 	t.Run("deleteGoodLedger", deleteGoodLedger)
// }
