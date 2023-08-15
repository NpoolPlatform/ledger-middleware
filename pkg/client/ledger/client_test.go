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
	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"
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
	ret = npool.Ledger{
		ID:         uuid.NewString(),
		AppID:      uuid.NewString(),
		UserID:     uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		Incoming:   "0.000000000000000000",
		Outcoming:  "0.000000000000000000",
		Spendable:  "0.000000000000000000",
		Locked:     "0.000000000000000000",
	}
)

func createLedger(t *testing.T) {
	handler, err := ledger1.NewHandler(
		context.Background(),
		ledger1.WithID(&ret.ID),
		ledger1.WithAppID(&ret.AppID),
		ledger1.WithUserID(&ret.UserID),
		ledger1.WithCoinTypeID(&ret.CoinTypeID),
	)
	assert.Nil(t, err)

	info, err := handler.CreateLedger(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateLedger(t *testing.T) {
	ret.Incoming = "10.000000000000000000"
	ret.Outcoming = "1.000000000000000000"
	ret.Locked = "3.000000000000000000"
	ret.Spendable = "6.000000000000000000"

	handler, err := ledger1.NewHandler(
		context.Background(),
		ledger1.WithID(&ret.ID),
		ledger1.WithIncoming(&ret.Incoming),
		ledger1.WithOutcoming(&ret.Outcoming),
		ledger1.WithLocked(&ret.Locked),
		ledger1.WithSpendable(&ret.Spendable),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateLedger(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getLedger(t *testing.T) {
	info, err := GetLedger(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getLedgerOnly(t *testing.T) {
	info, err := GetLedgerOnly(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getLedgers(t *testing.T) {
	infos, _, err := GetLedgers(context.Background(), &npool.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &commonpb.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CoinTypeID: &commonpb.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 0)
	}
}

func deleteLedger(t *testing.T) {
	handler, err := ledger1.NewHandler(
		context.Background(),
		ledger1.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteLedger(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = handler.GetLedger(context.Background())
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

	t.Run("createLedger", createLedger)
	t.Run("updateLedger", updateLedger)
	t.Run("getLedger", getLedger)
	t.Run("getLedgerOnly", getLedgerOnly)
	t.Run("getLedgers", getLedgers)
	t.Run("deleteLedger", deleteLedger)

}
