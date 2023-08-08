package unsold

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"
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

var entity = ent.MiningUnsold{
	ID:          uuid.New(),
	GoodID:      uuid.New(),
	CoinTypeID:  uuid.New(),
	Amount:      decimal.RequireFromString("9999999999999999999.999999999999999999"),
	BenefitDate: uint32(time.Now().Unix()),
}

var (
	id          = entity.ID.String()
	goodID      = entity.GoodID.String()
	coinTypeID  = entity.CoinTypeID.String()
	amount      = entity.Amount.String()
	benefitDate = entity.BenefitDate

	req = npool.UnsoldReq{
		ID:          &id,
		GoodID:      &goodID,
		CoinTypeID:  &coinTypeID,
		Amount:      &amount,
		BenefitDate: &benefitDate,
	}
)

var info *ent.MiningUnsold

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.MiningUnsold{
		{
			ID:          uuid.New(),
			GoodID:      uuid.New(),
			CoinTypeID:  uuid.New(),
			Amount:      decimal.RequireFromString("10.00896"),
			BenefitDate: uint32(time.Now().Unix()),
		},
		{
			ID:          uuid.New(),
			GoodID:      uuid.New(),
			CoinTypeID:  uuid.New(),
			Amount:      decimal.RequireFromString("11.11111"),
			BenefitDate: uint32(time.Now().Unix()),
		},
	}

	reqs := []*npool.UnsoldReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_goodID := _entity.GoodID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_amount := _entity.Amount.String()
		_benefitDate := _entity.BenefitDate

		reqs = append(reqs, &npool.UnsoldReq{
			ID:          &_id,
			GoodID:      &_goodID,
			CoinTypeID:  &_coinTypeID,
			Amount:      &_amount,
			BenefitDate: &_benefitDate,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), entity.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestUnsold(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
