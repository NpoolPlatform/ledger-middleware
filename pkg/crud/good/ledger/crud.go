package goodledger

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodledger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	GoodID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Amount     *decimal.Decimal
	ToPlatform *decimal.Decimal
	ToUser     *decimal.Decimal
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.GoodLedgerCreate, in *Req) *ent.GoodLedgerCreate {
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

	if in.Amount != nil {
		c.SetAmount(amount)
	}
	if in.ToPlatform != nil {
		c.SetToPlatform(toPlatform)
	}
	if in.ToUser != nil {
		c.SetToUser(toUser)
	}
	return c
}

func UpdateSet(u *ent.GoodLedgerUpdateOne, req *Req) *ent.GoodLedgerUpdateOne {
	amount := decimal.NewFromInt(0)
	if req.Amount != nil {
		amount = amount.Add(*req.Amount)
		u.SetAmount(amount)
	}
	toPlatform := decimal.NewFromInt(0)
	if req.ToPlatform != nil {
		toPlatform = toPlatform.Add(*req.ToPlatform)
		u.SetToPlatform(toPlatform)
	}
	toUser := decimal.NewFromInt(0)
	if req.ToUser != nil {
		toUser = toUser.Add(*req.ToUser)
		u.SetToUser(toUser)
	}

	return u
}

type Conds struct {
	ID         *cruder.Cond
	GoodID     *cruder.Cond
	CoinTypeID *cruder.Cond
	Amount     *cruder.Cond
	ToPlatform *cruder.Cond
	ToUser     *cruder.Cond
}

func SetQueryConds(q *ent.GoodLedgerQuery, conds *Conds) (*ent.GoodLedgerQuery, error) { //nolint
	q.Where(entgoodledger.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entgoodledger.ID(id))
		default:
			return nil, fmt.Errorf("invalid id op field %v", conds.ID.Op)
		}
	}
	if conds.GoodID != nil {
		goodID, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good id")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodledger.GoodID(goodID))
		default:
			return nil, fmt.Errorf("invalid good id op field %v", conds.GoodID.Op)
		}
	}
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodledger.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.LT:
			q.Where(entgoodledger.AmountLT(amount))
		case cruder.GT:
			q.Where(entgoodledger.AmountGT(amount))
		case cruder.EQ:
			q.Where(entgoodledger.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.ToPlatform != nil {
		toPlatform, ok := conds.ToPlatform.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid to platform %v", conds.ToPlatform.Val)
		}
		switch conds.ToPlatform.Op {
		case cruder.EQ:
			q.Where(entgoodledger.ToPlatform(toPlatform))
		default:
			return nil, fmt.Errorf("invalid to platform op field %v", conds.ToPlatform.Op)
		}
	}
	if conds.ToUser != nil {
		toUser, ok := conds.ToUser.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid to user %v", conds.ToUser.Val)
		}
		switch conds.ToUser.Op {
		case cruder.LT:
			q.Where(entgoodledger.AmountLT(toUser))
		case cruder.GT:
			q.Where(entgoodledger.AmountGT(toUser))
		case cruder.EQ:
			q.Where(entgoodledger.AmountEQ(toUser))
		default:
			return nil, fmt.Errorf("invalid to user op field %v", conds.ToUser.Op)
		}
	}
	return q, nil
}
