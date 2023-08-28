package goodledger

import (
    "fmt"

    "github.com/shopspring/decimal"

    "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
)

func CreateSetWithValidate(c *ent.GoodLedgerCreate, in *Req) (*ent.GoodLedgerCreate, error) {
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

    if amount.Cmp(
        toUser.Add(toPlatform),
    ) != 0 {
        return nil, fmt.Errorf("toPlatform (%v) + toUser (%v) != amount (%v)",
            toPlatform.String(), toUser.String(), amount.String())
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

    return c, nil
}

func UpdateSetWithValidate(entity *ent.GoodLedger, req *Req) (*ent.GoodLedgerUpdateOne, error) {
    amount := decimal.NewFromInt(0)
    if req.Amount != nil {
        amount = amount.Add(*req.Amount)
    }
    toPlatform := decimal.NewFromInt(0)
    if req.ToPlatform != nil {
        toPlatform = toPlatform.Add(*req.ToPlatform)
    }
    toUser := decimal.NewFromInt(0)
    if req.ToUser != nil {
        toUser = toUser.Add(*req.ToUser)
    }

    if amount.Add(entity.Amount).Cmp(
        toPlatform.Add(entity.ToPlatform).
            Add(toUser).
            Add(entity.ToUser),
    ) != 0 {
        return nil, fmt.Errorf("amount(%v + %v) != toPlatform(%v + %v) + toUser(%v + %v)",
            amount, entity.Amount, toPlatform, entity.ToPlatform, toUser, entity.ToUser,
        )
    }
    if amount.Add(entity.Amount).Cmp(decimal.NewFromInt(0)) < 0 {
        return nil, fmt.Errorf("amount less 0, %v + %v", amount.String(), entity.Amount)
    }
    if toPlatform.Add(entity.ToPlatform).Cmp(decimal.NewFromInt(0)) < 0 {
        return nil, fmt.Errorf("to platform less 0, %v + %v", toPlatform.String(), entity.ToPlatform)
    }
    if toUser.Add(entity.ToUser).Cmp(decimal.NewFromInt(0)) < 0 {
        return nil, fmt.Errorf("to user less %v + %v", toUser.String(), entity.ToUser)
    }
    stm := entity.Update()

    if req.Amount != nil {
        amount = amount.Add(entity.Amount)
        stm = stm.SetAmount(amount)
    }
    if req.ToPlatform != nil {
        toPlatform = toPlatform.Add(entity.ToPlatform)
        stm = stm.SetToPlatform(toPlatform)
    }
    if req.ToUser != nil {
        toUser = toUser.Add(entity.ToUser)
        stm = stm.SetToUser(toUser)
    }
    return stm, nil
}

