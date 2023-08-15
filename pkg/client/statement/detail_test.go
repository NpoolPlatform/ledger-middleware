package statement

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	currency1 = decimal.RequireFromString("1.00045000000123012")

	entity = ent.Detail{
		ID:              uuid.New(),
		AppID:           uuid.New(),
		UserID:          uuid.New(),
		CoinTypeID:      uuid.New(),
		IoType:          npool.IOType_Incoming.String(),
		IoSubType:       npool.IOSubType_Payment.String(),
		Amount:          decimal.RequireFromString("9999999999999999999.999999999999999999"),
		FromCoinTypeID:  uuid.New(),
		CoinUsdCurrency: &currency1,
		IoExtra:         uuid.New().String(),
	}
)

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

}
