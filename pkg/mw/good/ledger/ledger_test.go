package ledger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	goodstatement1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/statement"
	unsold1 "github.com/NpoolPlatform/ledger-middleware/pkg/mw/good/ledger/unsold"
	"github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodledgermwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger"
	goodstatementmwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	unsoldmwpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/unsold"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	id                        = uuid.NewString()
	unsoldStatementID         = uuid.NewString()
	goodID                    = uuid.NewString()
	coinTypeID                = uuid.NewString()
	totalAmount               = "400"
	unsoldAmount              = "100"
	techniqueServiceFeeAmount = "2"
	benefitDate               = uint32(time.Now().Unix())
)

func setup(t *testing.T) func(*testing.T) {
	reqs := []*goodstatementmwpb.GoodStatementReq{
		{
			ID:                        &id,
			UnsoldStatementID:         &unsoldStatementID,
			GoodID:                    &goodID,
			CoinTypeID:                &coinTypeID,
			TotalAmount:               &totalAmount,
			UnsoldAmount:              &unsoldAmount,
			TechniqueServiceFeeAmount: &techniqueServiceFeeAmount,
			BenefitDate:               &benefitDate,
		},
	}
	handler, err := goodstatement1.NewHandler(
		context.Background(),
		goodstatement1.WithReqs(reqs),
	)
	assert.Nil(t, err)

	infos, err := handler.CreateGoodStatements(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
		assert.Equal(t, id, infos[0].ID)
	}

	handler1, err := goodstatement1.NewHandler(context.Background())
	goodstatementID := uuid.MustParse(id)
	handler1.Req = &goodstatement1.Req{
		ID: &goodstatementID,
	}
	assert.Nil(t, err)

	return func(t *testing.T) {
		_, _ = handler1.DeleteGoodStatement(context.Background())
	}
}

func getUnsold(t *testing.T) {
	conds := &unsoldmwpb.Conds{
		ID:          &basetypes.StringVal{Op: cruder.EQ, Value: unsoldStatementID},
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: goodID},
		CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: coinTypeID},
		BenefitDate: &basetypes.Uint32Val{Op: cruder.EQ, Value: benefitDate},
	}
	handler, err := unsold1.NewHandler(
		context.Background(),
		unsold1.WithConds(conds),
	)
	assert.Nil(t, err)
	info, err := handler.GetUnsoldStatementOnly(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		assert.Equal(t, unsoldAmount, info.Amount)
	}
}

func getGoodLedger(t *testing.T) {
	conds := &goodledgermwpb.Conds{
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: goodID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: coinTypeID},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	info, err := handler.GetGoodLedgerOnly(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
		assert.Equal(t, "102", info.ToPlatform)
		assert.Equal(t, "298", info.ToUser)
	}
}

func TestGoodLedger(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getUnsold", getUnsold)
	t.Run("getGoodLedger", getGoodLedger)
}
