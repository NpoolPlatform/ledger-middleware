package statement

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger"
	goodstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/statement"
	unsoldcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/unsold"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	entgoodstatement "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/goodstatement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkGoodStatementExist(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
    if req.GoodID == nil {
        return fmt.Errorf("invalid good id")
    }
    if req.CoinTypeID == nil {
        return fmt.Errorf("invalid coin type id")
    }
    if req.TotalAmount == nil {
        return fmt.Errorf("invalid total amount")
    }
    if req.UnsoldAmount == nil {
        return fmt.Errorf("invalid unsold amount")
    }
    if req.TechniqueServiceFeeAmount == nil {
        return fmt.Errorf("invalid technique service fee amount")
    }
    if req.BenefitDate == nil {
        return fmt.Errorf("invalid benefit date")
    }
    if req.ID == nil {
        exist, err := tx.
            GoodStatement.
            Query().
            Where(
            entgoodstatement.GoodID(*req.GoodID),
            entgoodstatement.CoinTypeID(*req.CoinTypeID),
            entgoodstatement.BenefitDate(*req.BenefitDate),
            ).
            Exist(ctx)
        if err != nil {
            return err
        }
        if exist {
            return fmt.Errorf("good statement already exist")
        }
    }
    return nil
}

//nolint
func (h *createHandler) tryCreateGoodStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateGoodLedgerStatement, *req.GoodID, *req.CoinTypeID, *req.BenefitDate)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	toPlatform := req.UnsoldAmount.Add(*req.TechniqueServiceFeeAmount)
	toUser := req.TotalAmount.Sub(toPlatform)
	if req.TotalAmount.Cmp(toPlatform.Add(toUser)) != 0 {
		return fmt.Errorf("TotalAmount(%v) != ToPlatform(%v) + ToUser(%v)", req.TotalAmount.String(), toPlatform.String(), toUser.String())
	}
	if _, err := goodstatementcrud.CreateSet(
		tx.GoodStatement.Create(),
		&goodstatementcrud.Req{
			ID:          req.ID,
			GoodID:      req.GoodID,
			CoinTypeID:  req.CoinTypeID,
			BenefitDate: req.BenefitDate,
			TotalAmount: req.TotalAmount,
			ToPlatform:  &toPlatform,
			ToUser:      &toUser,
			TechniqueServiceFeeAmount: req.TechniqueServiceFeeAmount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *createHandler) tryCreateUnsoldStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateGoodLedgerUnsoldStatement, *req.GoodID, *req.CoinTypeID, *req.BenefitDate)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	if _, err := unsoldcrud.CreateSet(
		tx.UnsoldStatement.Create(),
		&unsoldcrud.Req{
			GoodID:      req.GoodID,
			CoinTypeID:  req.CoinTypeID,
			Amount:      req.UnsoldAmount,
			BenefitDate: req.BenefitDate,
			StatementID: req.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateOrUpdateGoodLedger(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := goodledgercrud.SetQueryConds(
		tx.GoodLedger.Query(),
		&goodledgercrud.Conds{
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: *req.GoodID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	toPlatform := req.UnsoldAmount.Add(*req.TechniqueServiceFeeAmount)
	toUser := req.TotalAmount.Sub(toPlatform)
	if req.TotalAmount.Cmp(toPlatform.Add(toUser)) != 0 {
		return fmt.Errorf("TotalAmount(%v) != ToPlatform(%v) + ToUser(%v)", req.TotalAmount.String(), toPlatform.String(), toUser.String())
	}

	if info == nil {
		key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateGoodLedger, *req.GoodID, *req.CoinTypeID)
		if err := redis2.TryLock(key, 0); err != nil {
			return err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()

		stm, err := goodledgercrud.CreateSetWithValidate(
			tx.GoodLedger.Create(),
			&goodledgercrud.Req{
				GoodID:     req.GoodID,
				CoinTypeID: req.CoinTypeID,
				Amount:     req.TotalAmount,
				ToPlatform: &toPlatform,
				ToUser:     &toUser,
			},
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(ctx); err != nil {
			return err
		}
		return nil
	}

	stm1, err := goodledgercrud.UpdateSetWithValidate(
		info,
		&goodledgercrud.Req{
			Amount:     req.TotalAmount,
			ToPlatform: &toPlatform,
			ToUser:     &toUser,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm1.Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *Handler) CreateGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				if err := handler.checkGoodStatementExist(req, ctx, tx); err != nil {
					return err
				}
				id := uuid.New()
				if req.ID == nil {
					req.ID = &id
				}
				if err := handler.tryCreateGoodStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryCreateUnsoldStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryCreateOrUpdateGoodLedger(req, ctx, tx); err != nil {
					return err
				}
				ids = append(ids, *req.ID)
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &goodstatementcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetGoodStatements(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func (h *Handler) CreateGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	h.Reqs = append(h.Reqs, &h.Req)

	infos, err := h.CreateGoodStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos[0], nil
}
