package lock

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uuid.UUID
	Amount    *decimal.Decimal
	DeletedAt *uint32
}

func CreateSet(c *ent.LedgerLockCreate, in *Req) *ent.LedgerLockCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	return c
}

func UpdateSet(u *ent.LedgerLockUpdateOne, req *Req) *ent.LedgerLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
