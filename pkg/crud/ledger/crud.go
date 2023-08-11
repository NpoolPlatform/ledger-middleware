package general

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entledger "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/ledger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Incoming   *decimal.Decimal
	Outcoming  *decimal.Decimal
	Locked     *decimal.Decimal
	Spendable  *decimal.Decimal
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.LedgerCreate, in *Req) *ent.LedgerCreate {
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
	if in.Incoming != nil {
		c.SetIncoming(*in.Incoming)
	}
	if in.Outcoming != nil {
		c.SetOutcoming(*in.Outcoming)
	}
	if in.Spendable != nil {
		c.SetSpendable(*in.Spendable)
	}
	if in.Locked != nil {
		c.SetLocked(*in.Locked)
	}
	return c
}

func UpdateSet(entity *ent.Ledger, u *ent.LedgerUpdateOne, req *Req) (*ent.LedgerUpdateOne, error) {
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
	if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("incoming (%v) < 0", incoming)
	}
	if outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("outcoming (%v) < 0", outcoming)
	}
	if spendable.Add(entity.Spendable).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("spendable (%v) + spendable(%v) < 0", spendable, entity.Spendable)
	}

	if req.Incoming != nil {
		incoming = incoming.Add(entity.Incoming)
		u.SetIncoming(incoming)
	}
	if req.Outcoming != nil {
		outcoming = outcoming.Add(entity.Outcoming)
		u.SetOutcoming(outcoming)
	}
	if req.Locked != nil {
		locked = locked.Add(entity.Locked)
		u.SetLocked(locked)
	}
	if req.Spendable != nil {
		spendable = spendable.Add(entity.Spendable)
		u.SetSpendable(spendable)
	}
	return u, nil
}

type Conds struct {
	ID          *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	Incoming    *cruder.Cond
	Outcoming   *cruder.Cond
	Spendable   *cruder.Cond
	Locked      *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

func SetQueryConds(q *ent.LedgerQuery, conds *Conds) (*ent.LedgerQuery, error) { //nolint
	q.Where(entledger.DeletedAt(0))
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
			q.Where(entledger.ID(id))
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
			q.Where(entledger.AppID(appID))
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
			q.Where(entledger.UserID(userID))
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
			q.Where(entledger.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.Incoming != nil {
		incoming, ok := conds.Incoming.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid io type %v", conds.Incoming.Val)
		}
		switch conds.Incoming.Op {
		case cruder.EQ:
			q.Where(entledger.Incoming(incoming))
		default:
			return nil, fmt.Errorf("invalid incoming op field %v", conds.Incoming.Op)
		}
	}
	if conds.Outcoming != nil {
		outcoming, ok := conds.Outcoming.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid outcoming %v", conds.Outcoming.Val)
		}
		switch conds.Outcoming.Op {
		case cruder.EQ:
			q.Where(entledger.Outcoming(outcoming))
		default:
			return nil, fmt.Errorf("invalid outcoming op field %v", conds.Outcoming.Op)
		}
	}
	if conds.Spendable != nil {
		spendable, ok := conds.Spendable.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid spendable %v", conds.Spendable.Val)
		}
		switch conds.Spendable.Op {
		case cruder.LT:
			q.Where(entledger.SpendableLT(spendable))
		case cruder.GT:
			q.Where(entledger.SpendableGT(spendable))
		case cruder.EQ:
			q.Where(entledger.SpendableEQ(spendable))
		default:
			return nil, fmt.Errorf("invalid spendable op field %v", conds.Spendable.Op)
		}
	}
	if conds.Locked != nil {
		locked, ok := conds.Locked.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid locked %v", conds.Locked.Val)
		}
		switch conds.Locked.Op {
		case cruder.EQ:
			q.Where(entledger.Locked(locked))
		default:
			return nil, fmt.Errorf("invalid locked op field %v", conds.Locked.Op)
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid coin type ids %v", conds.CoinTypeIDs.Val)
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.LIKE:
			q.Where(entledger.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid coin type ids op field %v", conds.CoinTypeIDs.Op)
		}
	}
	return q, nil
}
