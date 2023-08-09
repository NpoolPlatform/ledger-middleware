package goodledger

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entdetail "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID              *uuid.UUID
	AppID           *uuid.UUID
	UserID          *uuid.UUID
	CoinTypeID      *uuid.UUID
	IOType          *basetypes.IOType
	IOSubType       *basetypes.IOSubType
	Amount          *decimal.Decimal
	FromCoinTypeID  *uuid.UUID
	CoinUSDCurrency *decimal.Decimal
	IOExtra         *string
	CreatedAt       *uint32
	DeletedAt       *uint32
}

func CreateSet(c *ent.DetailCreate, in *Req) *ent.DetailCreate {
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
	if in.IOType != nil {
		c.SetIoType(in.IOType.String())
	}
	if in.IOSubType != nil {
		c.SetIoSubType(in.IOSubType.String())
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	if in.FromCoinTypeID != nil {
		c.SetFromCoinTypeID(*in.FromCoinTypeID)
	}
	if in.CoinUSDCurrency != nil {
		c.SetCoinUsdCurrency(*in.CoinUSDCurrency)
	}
	if in.IOExtra != nil {
		c.SetIoExtra(*in.IOExtra)
	}
	return c
}

func UpdateSet(u *ent.DetailUpdateOne, req *Req) *ent.DetailUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID              *cruder.Cond
	AppID           *cruder.Cond
	UserID          *cruder.Cond
	CoinTypeID      *cruder.Cond
	IOType          *cruder.Cond
	IOSubType       *cruder.Cond
	Amount          *cruder.Cond
	FromCoinTypeID  *cruder.Cond
	CoinUSDCurrency *cruder.Cond
	IOExtra         *cruder.Cond
}

func SetQueryConds(q *ent.DetailQuery, conds *Conds) (*ent.DetailQuery, error) { //nolint
	q.Where(entdetail.DeletedAt(0))
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
			q.Where(entdetail.ID(id))
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
			q.Where(entdetail.AppID(appID))
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
			q.Where(entdetail.UserID(userID))
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
			q.Where(entdetail.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.IOType != nil {
		ioType, ok := conds.IOType.Val.(basetypes.IOType)
		if !ok {
			return nil, fmt.Errorf("invalid io type %v", conds.IOType.Val)
		}
		switch conds.IOType.Op {
		case cruder.EQ:
			q.Where(entdetail.IoType(ioType.String()))
		default:
			return nil, fmt.Errorf("invalid io type op field %v", conds.IOType.Op)
		}
	}
	if conds.IOSubType != nil {
		ioSubType, ok := conds.IOSubType.Val.(basetypes.IOSubType)
		if !ok {
			return nil, fmt.Errorf("invalid io type %v", conds.IOSubType.Val)
		}
		switch conds.IOSubType.Op {
		case cruder.EQ:
			q.Where(entdetail.IoSubType(ioSubType.String()))
		default:
			return nil, fmt.Errorf("invalid io sub type op field %v", conds.IOSubType.Op)
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.LT:
			q.Where(entdetail.AmountLT(amount))
		case cruder.GT:
			q.Where(entdetail.AmountGT(amount))
		case cruder.EQ:
			q.Where(entdetail.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.FromCoinTypeID != nil {
		fromCoinTypeID, ok := conds.FromCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid from coin type id %v", conds.FromCoinTypeID.Val)
		}
		switch conds.FromCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entdetail.FromCoinTypeID(fromCoinTypeID))
		default:
			return nil, fmt.Errorf("invalid from coin type id op field %v", conds.FromCoinTypeID.Op)
		}
	}
	if conds.CoinUSDCurrency != nil {
		currency, ok := conds.CoinUSDCurrency.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid coin usd currency %v", conds.CoinUSDCurrency.Val)
		}
		switch conds.CoinUSDCurrency.Op {
		case cruder.LT:
			q.Where(entdetail.AmountLT(currency))
		case cruder.GT:
			q.Where(entdetail.AmountGT(currency))
		case cruder.EQ:
			q.Where(entdetail.AmountEQ(currency))
		default:
			return nil, fmt.Errorf("invalid coin usd currency op field %v", conds.CoinUSDCurrency.Op)
		}
	}
	if conds.IOExtra != nil {
		extra, ok := conds.IOExtra.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid io extra %v", conds.IOExtra.Val)
		}
		switch conds.IOExtra.Op {
		case cruder.LIKE:
			q.Where(entdetail.IoExtraContains(extra))
		default:
			return nil, fmt.Errorf("invalid io extra op field %v", conds.IOExtra.Op)
		}
	}
	return q, nil
}
