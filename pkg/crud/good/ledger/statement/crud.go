package goodstatement

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID                        *uuid.UUID
	GoodID                    *uuid.UUID
	CoinTypeID                *uuid.UUID
	TotalAmount               *decimal.Decimal
	UnsoldAmount              *decimal.Decimal
	TechniqueServiceFeeAmount *decimal.Decimal
	BenefitDate               *uint32
	CreatedAt                 *uint32
	DeletedAt                 *uint32
}

func CreateSet(c *ent.GoodStatementCreate, in *Req) *ent.GoodStatementCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.GoodID != nil {
		c.SetGoodID(*in.GoodID)
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(*in.CoinTypeID)
	}
	if in.TotalAmount != nil {
		c.SetAmount(*in.TotalAmount)
	}
	if in.BenefitDate != nil {
		c.SetBenefitDate(*in.BenefitDate)
	}

	return c
}

func UpdateSet(u *ent.GoodStatementUpdateOne, req *Req) *ent.GoodStatementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	GoodID      *cruder.Cond
	CoinTypeID  *cruder.Cond
	Amount      *cruder.Cond
	BenefitDate *cruder.Cond
	IDs         *cruder.Cond
}

func SetQueryConds(q *ent.GoodStatementQuery, conds *Conds) (*ent.GoodStatementQuery, error) { //nolint
	q.Where(entgoodstatement.DeletedAt(0))
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
			q.Where(entgoodstatement.ID(id))
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
			q.Where(entgoodstatement.GoodID(goodID))
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
			q.Where(entgoodstatement.CoinTypeID(coinTypeID))
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
			q.Where(entgoodstatement.AmountLT(amount))
		case cruder.GT:
			q.Where(entgoodstatement.AmountGT(amount))
		case cruder.EQ:
			q.Where(entgoodstatement.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.BenefitDate != nil {
		benefitDate, ok := conds.BenefitDate.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid benefit date %v", conds.BenefitDate.Val)
		}
		switch conds.BenefitDate.Op {
		case cruder.LT:
			q.Where(entgoodstatement.BenefitDateLT(benefitDate))
		case cruder.GT:
			q.Where(entgoodstatement.BenefitDateGT(benefitDate))
		case cruder.EQ:
			q.Where(entgoodstatement.BenefitDateEQ(benefitDate))
		default:
			return nil, fmt.Errorf("invalid benefit date op field %v", conds.BenefitDate.Op)
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids %v", conds.IDs.Val)
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entgoodstatement.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid good statement op field %v", conds.IDs.Op)
		}
	}
	return q, nil
}
