package general

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
// 	"github.com/shopspring/decimal"

// 	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

// 	testinit "github.com/NpoolPlatform/ledger-middleware/pkg/testinit"
// 	valuedef "github.com/NpoolPlatform/message/npool"
// 	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"
// 	"github.com/google/uuid"

// 	"github.com/stretchr/testify/assert"
// )

// func init() {
// 	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
// 		return
// 	}
// 	if err := testinit.Init(); err != nil {
// 		fmt.Printf("cannot init test stub: %v\n", err)
// 	}
// }

// var entity = ent.MiningGeneral{
// 	ID:         uuid.New(),
// 	GoodID:     uuid.New(),
// 	CoinTypeID: uuid.New(),
// 	Amount:     decimal.NewFromInt(0),
// 	ToPlatform: decimal.NewFromInt(0),
// 	ToUser:     decimal.NewFromInt(0),
// }

// var (
// 	id         = entity.ID.String()
// 	appID      = entity.GoodID.String()
// 	coinTypeID = entity.CoinTypeID.String()
// 	amount     = entity.Amount.String()
// 	toPlatform = entity.ToPlatform.String()
// 	toUser     = entity.ToUser.String()

// 	req = npool.GeneralReq{
// 		ID:         &id,
// 		GoodID:     &appID,
// 		CoinTypeID: &coinTypeID,
// 		Amount:     &amount,
// 		ToPlatform: &toPlatform,
// 		ToUser:     &toUser,
// 	}
// )

// var info *ent.MiningGeneral

// func create(t *testing.T) {
// 	var err error
// 	info, err = Create(context.Background(), &req)
// 	if assert.Nil(t, err) {
// 		entity.UpdatedAt = info.UpdatedAt
// 		entity.CreatedAt = info.CreatedAt
// 		assert.Equal(t, info.String(), entity.String())
// 	}
// }

// func createBulk(t *testing.T) {
// 	entities := []*ent.MiningGeneral{
// 		{
// 			ID:         uuid.New(),
// 			GoodID:     uuid.New(),
// 			CoinTypeID: uuid.New(),
// 			Amount:     decimal.NewFromInt(0),
// 			ToPlatform: decimal.NewFromInt(0),
// 			ToUser:     decimal.NewFromInt(0),
// 		},
// 		{
// 			ID:         uuid.New(),
// 			GoodID:     uuid.New(),
// 			CoinTypeID: uuid.New(),
// 			Amount:     decimal.NewFromInt(0),
// 			ToPlatform: decimal.NewFromInt(0),
// 			ToUser:     decimal.NewFromInt(0),
// 		},
// 	}

// 	reqs := []*npool.GeneralReq{}
// 	for _, _entity := range entities {
// 		_id := _entity.ID.String()
// 		_goodID := _entity.GoodID.String()
// 		_coinTypeID := _entity.CoinTypeID.String()
// 		_amount := _entity.Amount.String()
// 		_toPlatform := _entity.ToPlatform.String()
// 		_toUser := _entity.ToUser.String()

// 		reqs = append(reqs, &npool.GeneralReq{
// 			ID:         &_id,
// 			GoodID:     &_goodID,
// 			CoinTypeID: &_coinTypeID,
// 			Amount:     &_amount,
// 			ToPlatform: &_toPlatform,
// 			ToUser:     &_toUser,
// 		})
// 	}
// 	infos, err := CreateBulk(context.Background(), reqs)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, len(infos), 2)
// 	}
// }

// func add(t *testing.T) {
// 	amount = "30"
// 	toPlatform = "10"
// 	toUser = "20"

// 	req.Amount = &amount
// 	req.ToPlatform = &toPlatform
// 	req.ToUser = &toUser

// 	entity.Amount, _ = decimal.NewFromString(amount)
// 	entity.ToPlatform, _ = decimal.NewFromString(toPlatform)
// 	entity.ToUser, _ = decimal.NewFromString(toUser)

// 	info, err := AddFields(context.Background(), &req)
// 	if assert.Nil(t, err) {
// 		entity.UpdatedAt = info.UpdatedAt
// 		assert.Equal(t, info.String(), entity.String())
// 	}
// }

// func row(t *testing.T) {
// 	var err error
// 	info, err = Row(context.Background(), info.ID)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, info.String(), entity.String())
// 	}
// }

// func rows(t *testing.T) {
// 	infos, total, err := Rows(context.Background(),
// 		&npool.Conds{
// 			ID: &valuedef.StringVal{
// 				Value: id,
// 				Op:    cruder.EQ,
// 			},
// 		}, 0, 0)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, total, 1)
// 		assert.Equal(t, infos[0].String(), entity.String())
// 	}
// }

// func rowOnly(t *testing.T) {
// 	var err error
// 	info, err = RowOnly(context.Background(),
// 		&npool.Conds{
// 			ID: &valuedef.StringVal{
// 				Value: id,
// 				Op:    cruder.EQ,
// 			},
// 		})
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, info.String(), entity.String())
// 	}
// }

// func count(t *testing.T) {
// 	count, err := Count(context.Background(),
// 		&npool.Conds{
// 			ID: &valuedef.StringVal{
// 				Value: id,
// 				Op:    cruder.EQ,
// 			},
// 		},
// 	)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, count, uint32(1))
// 	}
// }

// func exist(t *testing.T) {
// 	exist, err := Exist(context.Background(), entity.ID)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, exist, true)
// 	}
// }

// func existConds(t *testing.T) {
// 	exist, err := ExistConds(context.Background(),
// 		&npool.Conds{
// 			ID: &valuedef.StringVal{
// 				Value: id,
// 				Op:    cruder.EQ,
// 			},
// 		},
// 	)
// 	if assert.Nil(t, err) {
// 		assert.Equal(t, exist, true)
// 	}
// }

// func deleteA(t *testing.T) {
// 	info, err := Delete(context.Background(), entity.ID)
// 	if assert.Nil(t, err) {
// 		entity.DeletedAt = info.DeletedAt
// 		assert.Equal(t, info.String(), entity.String())
// 	}
// }

// func TestGeneral(t *testing.T) {
// 	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
// 		return
// 	}
// 	t.Run("create", create)
// 	t.Run("createBulk", createBulk)
// 	t.Run("add", add)
// 	t.Run("row", row)
// 	t.Run("rows", rows)
// 	t.Run("rowOnly", rowOnly)
// 	t.Run("exist", exist)
// 	t.Run("existConds", existConds)
// 	t.Run("count", count)
// 	t.Run("delete", deleteA)
// }