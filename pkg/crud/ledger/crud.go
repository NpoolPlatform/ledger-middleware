package ledger

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
	if in.Locked != nil {
		c.SetLocked(*in.Locked)
	}
	if in.Spendable != nil {
		c.SetSpendable(*in.Spendable)
	}
	return c
}

func UpdateSet(u *ent.LedgerUpdateOne, req *Req) *ent.LedgerUpdateOne {
	incoming := decimal.NewFromInt(0)
	if req.Incoming != nil {
		incoming = incoming.Add(*req.Incoming)
		u.SetIncoming(incoming)
	}

	locked := decimal.NewFromInt(0)
	if req.Locked != nil {
		locked = locked.Add(*req.Locked)
		u.SetLocked(locked)
	}

	outcoming := decimal.NewFromInt(0)
	if req.Outcoming != nil {
		outcoming = outcoming.Add(*req.Outcoming)
		u.SetOutcoming(outcoming)
	}

	spendable := decimal.NewFromInt(0)
	if req.Spendable != nil {
		spendable = spendable.Add(*req.Spendable)
		u.SetSpendable(spendable)
	}
	return u
}

func UpdateSetWithValidate(info *ent.Ledger, req *Req) (*ent.LedgerUpdateOne, error) {
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

	if incoming.Add(info.Incoming).
		Cmp(
			locked.Add(info.Locked).
				Add(outcoming).
				Add(info.Outcoming).
				Add(spendable).
				Add(info.Spendable),
		) != 0 {
		return nil, fmt.Errorf("outcoming (%v + %v) + locked (%v + %v) + spendable (%v + %v) != incoming (%v + %v)",
			outcoming, info.Outcoming, locked, info.Locked, spendable, info.Spendable, incoming, info.Incoming)
	}

	if locked.Add(info.Locked).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("locked (%v) + locked (%v) < 0", locked, info.Locked)
	}
	if incoming.Add(info.Incoming).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("incoming (%v) + incoming (%v) < 0", locked, info.Incoming)
	}
	if outcoming.Add(info.Outcoming).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("outcoming (%v) + outcoming (%v) < 0", locked, info.Outcoming)
	}
	if spendable.Add(info.Spendable).Cmp(decimal.NewFromInt(0)) < 0 {
		return nil, fmt.Errorf("spendable (%v) + spendable(%v) < 0", spendable, info.Spendable)
	}

	incoming = incoming.Add(info.Incoming)
	outcoming = outcoming.Add(info.Outcoming)
	locked = locked.Add(info.Locked)
	spendable = spendable.Add(info.Spendable)

	return UpdateSet(info.Update(), &Req{
		Incoming:  &incoming,
		Outcoming: &outcoming,
		Spendable: &spendable,
		Locked:    &locked,
	}), nil
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
			return nil, fmt.Errorf("invalid incoming %v", conds.Incoming.Val)
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
		case cruder.IN:
			q.Where(entledger.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid coin type ids op field %v", conds.CoinTypeIDs.Op)
		}
	}
	return q, nil
}
