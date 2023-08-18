package profit

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entprofit "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/profit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Incoming   *decimal.Decimal
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.ProfitCreate, in *Req) (*ent.ProfitCreate, error) {
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

func UpdateSet(entity *ent.Profit, u *ent.ProfitUpdateOne, req *Req) (*ent.ProfitUpdateOne, error) {
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
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	if req.Incoming != nil {
        incoming = incoming.Add(entity.Incoming)
		u.SetIncoming(incoming)
	}
	return u, nil
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	CoinTypeID *cruder.Cond
	Incoming   *cruder.Cond
}

func SetQueryConds(q *ent.ProfitQuery, conds *Conds) (*ent.ProfitQuery, error) { //nolint
	q.Where(entprofit.DeletedAt(0))
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
			q.Where(entprofit.ID(id))
		default:
			return nil, fmt.Errorf("invalid id op field %v", conds.ID.Op)
		}
	}
	if conds.AppID != nil {
		appID, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entprofit.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app id op field %v", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		userID, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entprofit.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid user id op field %v", conds.UserID.Op)
		}
	}
	if conds.CoinTypeID != nil {
		coinTypeID, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type id")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entprofit.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.Incoming != nil {
		incoming, ok := conds.Incoming.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid incoming %v", conds.Incoming.Val)
		}
		switch conds.Incoming.Op {
		case cruder.LT:
			q.Where(entprofit.IncomingLT(incoming))
		case cruder.GT:
			q.Where(entprofit.IncomingGT(incoming))
		case cruder.EQ:
			q.Where(entprofit.IncomingEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid incoming op field %v", conds.Incoming.Op)
		}
	}

	return q, nil
}
