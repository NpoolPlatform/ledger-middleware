package withdraw

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entwithdraw "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
)

type Req struct {
	ID                    *uuid.UUID
	AppID                 *uuid.UUID
	UserID                *uuid.UUID
	CoinTypeID            *uuid.UUID
	AccountID             *uuid.UUID
	Address               *string
	Amount                *decimal.Decimal
	PlatformTransactionID *uuid.UUID
	ChainTransactionID    *string
	State                 *basetypes.WithdrawState
	CreatedAt             *uint32
	DeletedAt             *uint32
}

func CreateSet(c *ent.WithdrawCreate, in *Req) *ent.WithdrawCreate {
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
	if in.AccountID != nil {
		c.SetAccountID(*in.AccountID)
	}
	if in.Address != nil {
		c.SetAddress(*in.Address)
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	if in.PlatformTransactionID != nil {
		c.SetPlatformTransactionID(*in.PlatformTransactionID)
	}

	c.SetState(basetypes.WithdrawState_Reviewing.String())
	return c
}

func UpdateSet(u *ent.WithdrawUpdateOne, req *Req) *ent.WithdrawUpdateOne {
	if req.PlatformTransactionID != nil {
		u.SetPlatformTransactionID(*req.PlatformTransactionID)
	}
	if req.ChainTransactionID != nil {
		u.SetChainTransactionID(*req.ChainTransactionID)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	CoinTypeID *cruder.Cond
	AccountID  *cruder.Cond
	Address    *cruder.Cond
	State      *cruder.Cond
	Amount     *cruder.Cond
}

func SetQueryConds(q *ent.WithdrawQuery, conds *Conds) (*ent.WithdrawQuery, error) { //nolint
	q.Where(entwithdraw.DeletedAt(0))
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
			q.Where(entwithdraw.ID(id))
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
			q.Where(entwithdraw.AppID(appID))
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
			q.Where(entwithdraw.UserID(userID))
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
			q.Where(entwithdraw.CoinTypeID(coinTypeID))
		default:
			return nil, fmt.Errorf("invalid coin type id op field %v", conds.CoinTypeID.Op)
		}
	}
	if conds.AccountID != nil {
		accountID, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid account id %v", conds.AccountID.Val)
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entwithdraw.AccountID(accountID))
		default:
			return nil, fmt.Errorf("invalid account id op field %v", conds.AccountID.Op)
		}
	}
	if conds.Address != nil {
		address, ok := conds.Address.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid address %v", conds.Address.Val)
		}
		switch conds.Address.Op {
		case cruder.EQ:
			q.Where(entwithdraw.Address(address))
		default:
			return nil, fmt.Errorf("invalid address op field %v", conds.Address.Op)
		}
	}

	if conds.Amount != nil {
		Amount, ok := conds.Amount.Val.(decimal.Decimal)
		if !ok {
			return nil, fmt.Errorf("invalid amount %v", conds.Amount.Val)
		}
		switch conds.Amount.Op {
		case cruder.EQ:
			q.Where(entwithdraw.Amount(Amount))
		default:
			return nil, fmt.Errorf("invalid amount op field %v", conds.Amount.Op)
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.WithdrawState)
		if !ok {
			return nil, fmt.Errorf("invalid state %v", conds.State.Val)
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entwithdraw.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid state op field %v", conds.State.Op)
		}
	}
	return q, nil
}
