package statement

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/statement"
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

func CreateSet(c *ent.StatementCreate, in *Req) *ent.StatementCreate {
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

func UpdateSet(u *ent.StatementUpdateOne, req *Req) *ent.StatementUpdateOne {
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
	StartAt         *cruder.Cond
	EndAt           *cruder.Cond
	IDs             *cruder.Cond
}

func SetQueryConds(q *ent.StatementQuery, conds *Conds) (*ent.StatementQuery, error) { //nolint
	q.Where(entstatement.DeletedAt(0))
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
			q.Where(entstatement.ID(id))
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
			q.Where(entstatement.AppID(appID))
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
			q.Where(entstatement.UserID(userID))
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
			q.Where(entstatement.CoinTypeID(coinTypeID))
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
			q.Where(entstatement.IoType(ioType.String()))
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
			q.Where(entstatement.IoSubType(ioSubType.String()))
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
			q.Where(entstatement.AmountLT(amount))
		case cruder.GT:
			q.Where(entstatement.AmountGT(amount))
		case cruder.EQ:
			q.Where(entstatement.AmountEQ(amount))
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
			q.Where(entstatement.FromCoinTypeID(fromCoinTypeID))
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
			q.Where(entstatement.AmountLT(currency))
		case cruder.GT:
			q.Where(entstatement.AmountGT(currency))
		case cruder.EQ:
			q.Where(entstatement.AmountEQ(currency))
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
			q.Where(entstatement.IoExtraContains(extra))
		default:
			return nil, fmt.Errorf("invalid io extra op field %v", conds.IOExtra.Op)
		}
	}
	if conds.StartAt != nil {
		startAt, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid start  %v", conds.StartAt)
		}
		switch conds.StartAt.Op {
		case cruder.GT:
			q.Where(entstatement.CreatedAtGTE(startAt))
		case cruder.LT:
			q.Where(entstatement.CreatedAtLTE(startAt))
		default:
			return nil, fmt.Errorf("invalid start at op field %v", conds.StartAt.Op)
		}
	}
	if conds.EndAt != nil {
		endAT, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid end at  %v", conds.EndAt)
		}
		switch conds.EndAt.Op {
		case cruder.GT:
			q.Where(entstatement.CreatedAtGTE(endAT))
		case cruder.LT:
			q.Where(entstatement.CreatedAtLTE(endAT))
		default:
			return nil, fmt.Errorf("invalid end at op field %v", conds.EndAt.Op)
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids %v", conds.IDs.Val)
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entstatement.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid statement op field %v", conds.IDs.Op)
		}
	}
	return q, nil
}
