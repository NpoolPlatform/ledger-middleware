package profit

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
)

func CreateSetWithValidate(c *ent.ProfitCreate, in *Req) (*ent.ProfitCreate, error) {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.UserID != nil {
		c.SetUserID(*in.UserID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
	}

	incoming := decimal.NewFromInt(0)
	if in.Incoming != nil {
		incoming = incoming.Add(*in.Incoming)
		if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return nil, fmt.Errorf("profit incoming less than 0 %v", incoming.String())
		}
		c.SetIncoming(incoming)
	}
	return c, nil
}

func UpdateSetWithValidate(entity *ent.Profit, req *Req) (*ent.ProfitUpdateOne, error) {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
	}
	if incoming.Add(entity.Incoming).
		Cmp(
			decimal.NewFromInt(0),
		) < 0 {
		return nil, fmt.Errorf("incoming (%v) + entity.incoming (%v) < 0",
			incoming, entity.Incoming)
	}

	stm := entity.Update()

	if req.Incoming != nil {
		incoming = incoming.Add(entity.Incoming)
		stm = stm.SetIncoming(incoming)
	}
	if req.DeletedAt != nil {
		stm = stm.SetDeletedAt(*req.DeletedAt)
	}
	return stm, nil
}
