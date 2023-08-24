package goodledger

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
)

func CreateSetWithValidate(c *ent.GoodLedgerCreate, in *Req) (*ent.GoodLedgerCreate, error) {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.GoodID != nil {
		c.SetGoodID(*in.GoodID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
	}

	amount := decimal.NewFromInt(0)
	if in.Amount != nil {
		amount = amount.Add(*in.Amount)
	}
	toPlatform := decimal.NewFromInt(0)
	if in.ToPlatform != nil {
		toPlatform = toPlatform.Add(*in.ToPlatform)
	}
	toUser := decimal.NewFromInt(0)
	if in.ToUser != nil {
		toUser = toUser.Add(*in.ToUser)
	}

	if amount.Cmp(
		toUser.Add(toPlatform),
	) != 0 {
		return nil, fmt.Errorf("toPlatform (%v) + toUser (%v) != amount (%v)",
			toPlatform.String(), toUser.String(), amount.String())
	}

	if in.Amount != nil {
		c.SetAmount(amount)
	}
	if in.ToPlatform != nil {
		c.SetToPlatform(toPlatform)
	}
	if in.ToUser != nil {
		c.SetToUser(toUser)
	}

	return c, nil
}

func UpdateSetWithValidate(entity *ent.GoodLedger, req *Req) (*ent.GoodLedgerUpdateOne, error) {
	amount := decimal.NewFromInt(0)
	if req.Amount != nil {
		amount = amount.Add(*req.Amount)
	}
	toPlatform := decimal.NewFromInt(0)
	if req.ToPlatform != nil {
		toPlatform = toPlatform.Add(*req.ToPlatform)
	}
	toUser := decimal.NewFromInt(0)
	if req.ToUser != nil {
		toUser = toUser.Add(*req.ToUser)
	}

	if amount.Cmp(toPlatform.Add(toUser)) < 0 {
		return nil, fmt.Errorf("amount %v < toplatform %v + touser %v", amount.String(), toPlatform.String(), toUser.String())
	}
	if amount.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("amount less 0 %v", amount.String())
	}
	if toPlatform.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("toplatform less 0 %v", toPlatform.String())
	}
	if toUser.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("touser less 0 %v", toUser.String())
	}

	stm := entity.Update()

	if req.Amount != nil {
		amount = amount.Add(entity.Amount)
		stm = stm.SetAmount(amount)
	}
	if req.ToPlatform != nil {
		toPlatform = toPlatform.Add(entity.ToPlatform)
		stm = stm.SetToPlatform(toPlatform)
	}
	if req.ToUser != nil {
		toUser = toUser.Add(entity.ToUser)
		stm = stm.SetToUser(toUser)
	}
	return stm, nil
}
