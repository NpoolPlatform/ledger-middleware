package ledger

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
)

func CreateSetWithValidate(c *ent.LedgerCreate, in *Req) (*ent.LedgerCreate, error) {
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
	}
	locked := decimal.NewFromInt(0)
	if in.Locked != nil {
		locked = locked.Add(*in.Locked)
	}
	outcoming := decimal.NewFromInt(0)
	if in.Outcoming != nil {
		outcoming = outcoming.Add(*in.Outcoming)
	}
	spendable := decimal.NewFromInt(0)
	if in.Spendable != nil {
		spendable = spendable.Add(*in.Spendable)
	}

	if incoming.Cmp(
		outcoming.Add(locked).
			Add(spendable),
	) != 0 {
		return nil, fmt.Errorf("outcoming (%v) + locked (%v) + spendable (%v) != incoming (%v)",
			outcoming, locked, spendable, incoming)
	}

	if in.Incoming != nil {
		c.SetIncoming(incoming)
	}
	if in.Outcoming != nil {
		c.SetOutcoming(outcoming)
	}
	if in.Locked != nil {
		c.SetLocked(locked)
	}
	if in.Spendable != nil {
		c.SetSpendable(spendable)
	}
	return c, nil
}

func UpdateSetWithValidate(entity *ent.Ledger, req *Req) (*ent.LedgerUpdateOne, error) {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
	}
	locked := decimal.NewFromInt(0)
	if req.Locked != nil {
		locked = locked.Add(*req.Locked)
	}
	outcoming := decimal.NewFromInt(0)
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
	}
	spendable := decimal.NewFromInt(0)
	if req.Spendable != nil {
		spendable = spendable.Add(*req.Spendable)
	}

	if incoming.Add(entity.Incoming).
		Cmp(
			locked.Add(entity.Locked).
				Add(outcoming).
				Add(entity.Outcoming).
				Add(spendable).
				Add(entity.Spendable),
		) != 0 {
		return nil, fmt.Errorf("outcoming (%v + %v) + locked (%v + %v) + spendable (%v + %v) != incoming (%v + %v)",
			outcoming, entity.Outcoming, locked, entity.Locked, spendable, entity.Spendable, incoming, entity.Incoming)
	}

	if locked.Add(entity.Locked).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("locked (%v) + locked (%v) < 0", locked, entity.Locked)
	}
	if incoming.Add(entity.Incoming).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("incoming (%v) + incoming (%v) < 0", locked, entity.Incoming)
	}
	if outcoming.Add(entity.Outcoming).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("outcoming (%v) + outcoming (%v) < 0", locked, entity.Outcoming)
	}
	if spendable.Add(entity.Spendable).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("spendable (%v) + spendable(%v) < 0", spendable, entity.Spendable)
	}

	stm := entity.Update()

	if req.Incoming != nil {
		incoming = incoming.Add(entity.Incoming)
		stm = stm.SetIncoming(incoming)
	}
	if req.Outcoming != nil {
		outcoming = outcoming.Add(entity.Outcoming)
		stm = stm.SetOutcoming(outcoming)
	}
	if req.Locked != nil {
		locked = locked.Add(entity.Locked)
		stm = stm.SetLocked(locked)
	}
	if req.Spendable != nil {
		spendable = spendable.Add(entity.Spendable)
		stm = stm.SetSpendable(spendable)
	}
	return stm, nil
}
