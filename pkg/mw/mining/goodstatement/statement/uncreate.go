package statement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodledger"
	goodstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodstatement"
	unsoldcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/unsoldstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodstatement"
	"github.com/shopspring/decimal"
)

type unCreateHandler struct {
	*Handler
}

func (h *unCreateHandler) tryUpdateGoodLedger(req *goodledgercrud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := goodledgercrud.SetQueryConds(tx.GoodLedger.Query(), &goodledgercrud.Conds{
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: req.GoodID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: req.CoinTypeID},
	})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("good ledger not exist")
		}
		return err
	}

	// update
	old, err := tx.GoodLedger.Get(ctx, info.ID)
	if err != nil {
		return err
	}
	if old == nil {
		return fmt.Errorf("ledger not exist, id %v", info.ID)
	}

	stm1, err := goodledgercrud.UpdateSet(
		old,
		tx.GoodLedger.UpdateOneID(info.ID),
		&goodledgercrud.Req{
			Amount:     req.Amount,
			ToPlatform: req.ToPlatform,
			ToUser:     req.ToUser,
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
func (h *unCreateHandler) tryDeleteGoodStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := goodstatementcrud.SetQueryConds(
		tx.GoodStatement.Query(),
		&goodstatementcrud.Conds{
			GoodID:      &cruder.Cond{Op: cruder.EQ, Val: *req.GoodID},
			CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
			BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: *req.BenefitDate},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := goodstatementcrud.UpdateSet(
		tx.GoodStatement.UpdateOneID(info.ID),
		&goodstatementcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *unCreateHandler) tryDeleteUnsoldStatement(req *unsoldcrud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := unsoldcrud.SetQueryConds(
		tx.UnsoldStatement.Query(),
		&unsoldcrud.Conds{
			GoodID:      &cruder.Cond{Op: cruder.EQ, Val: *req.GoodID},
			CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
			BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: *req.BenefitDate},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := unsoldcrud.UpdateSet(
		tx.UnsoldStatement.UpdateOneID(info.ID),
		&unsoldcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UnCreateGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	reqs := []*Req{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			stm, err := goodstatementcrud.SetQueryConds(
				tx.GoodStatement.Query(),
				&goodstatementcrud.Conds{
					GoodID:      &cruder.Cond{Op: cruder.EQ, Val: *req.GoodID},
					CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
					BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: *req.BenefitDate},
				},
			)
			if err != nil {
				return err
			}
			exist, err := stm.Exist(ctx)
			if err != nil {
				return err
			}
			if !exist {
				msg := fmt.Sprintf(
					"good statement not exist! GoodID(%v), CoinTypeID(%v),BenefitDate(%v)",
					*req.GoodID, *req.CoinTypeID, *req.BenefitDate,
				)
				logger.Sugar().Errorf(msg)
				continue
			}
			reqs = append(reqs, req)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	handler := &unCreateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range reqs {
			_fn := func() error {
				key := fmt.Sprintf("ledger-delete-goodstatement:%v:%v:%v", *h.GoodID, *h.CoinTypeID, *h.BenefitDate)
				if err := redis2.TryLock(key, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key)
				}()

				if err := handler.tryDeleteGoodStatement(&goodstatementcrud.Req{
					GoodID:      req.GoodID,
					CoinTypeID:  req.CoinTypeID,
					BenefitDate: req.BenefitDate,
					Amount:      req.TotalAmount,
				}, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryDeleteUnsoldStatement(&unsoldcrud.Req{
					GoodID:      req.GoodID,
					CoinTypeID:  req.CoinTypeID,
					Amount:      req.UnsoldAmount,
					BenefitDate: req.BenefitDate,
				}, ctx, tx); err != nil {
					return err
				}

				toPlatform := h.UnsoldAmount.Add(*h.TechniqueServiceFeeAmount)
				toUser := h.TotalAmount.Sub(toPlatform)
				if h.TotalAmount.Cmp(toPlatform.Add(toUser)) != 0 {
					return fmt.Errorf("TotalAmount(%v) != ToPlatform(%v) + ToUser(%v)", h.TotalAmount.String(), toPlatform.String(), toUser.String())
				}

				_amount := decimal.RequireFromString(fmt.Sprintf("-%v", req.TotalAmount.String()))
				_toUser := decimal.RequireFromString(fmt.Sprintf("-%v", toUser.String()))
				_toPlatform := decimal.RequireFromString(fmt.Sprintf("-%v", toPlatform.String()))
				if err := handler.tryUpdateGoodLedger(&goodledgercrud.Req{
					GoodID:     req.GoodID,
					CoinTypeID: req.CoinTypeID,
					Amount:     &_amount,
					ToPlatform: &_toPlatform,
					ToUser:     &_toUser,
				}, ctx, tx); err != nil {
					return err
				}
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

	return nil, nil
}
