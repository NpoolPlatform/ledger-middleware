package statement

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger"
	goodstatementcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/statement"
	unsoldcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/unsold"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) tryCreateOrUpdateGoodLedger(req *goodledgercrud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := goodledgercrud.SetQueryConds(tx.GoodLedger.Query(), &goodledgercrud.Conds{
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: req.GoodID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: req.CoinTypeID},
	})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return err
	}

	// create
	if info == nil {
		stm, err := goodledgercrud.CreateSet(
			tx.GoodLedger.Create(),
			&goodledgercrud.Req{
				GoodID:     req.GoodID,
				CoinTypeID: h.CoinTypeID,
				Amount:     req.Amount,
				ToPlatform: req.ToPlatform,
				ToUser:     req.ToUser,
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

func (h *createHandler) tryCreateGoodStatement(req *goodstatementcrud.Req, ctx context.Context, tx *ent.Tx) (*npool.GoodStatement, error) {
	info, err := goodstatementcrud.CreateSet(
		tx.GoodStatement.Create(),
		req,
	).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &npool.GoodStatement{
		ID:          info.ID.String(),
		GoodID:      info.GoodID.String(),
		CoinTypeID:  info.CoinTypeID.String(),
		BenefitDate: info.BenefitDate,
		CreatedAt:   info.CreatedAt,
		UpdatedAt:   info.UpdatedAt,
	}, nil
}

func (h *createHandler) tryCreateUnsoldStatement(req *unsoldcrud.Req, ctx context.Context, tx *ent.Tx) error {
	if _, err := unsoldcrud.CreateSet(
		tx.UnsoldStatement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
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
			if exist {
				msg := fmt.Sprintf(
					"good statement exist! GoodID(%v), CoinTypeID(%v),BenefitDate(%v)",
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

	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}
	infos := []*npool.GoodStatement{}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range reqs {
			_fn := func() error {
				goodStatementID := uuid.New()

				key := fmt.Sprintf("ledger-create-goodstatement:%v:%v:%v", *h.GoodID, *h.CoinTypeID, *h.BenefitDate)
				if err := redis2.TryLock(key, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key)
				}()

				info, err := handler.tryCreateGoodStatement(&goodstatementcrud.Req{
					ID:          &goodStatementID,
					GoodID:      req.GoodID,
					CoinTypeID:  req.CoinTypeID,
					BenefitDate: req.BenefitDate,
					Amount:      req.TotalAmount,
				}, ctx, tx)
				if err != nil {
					return err
				}
				infos = append(infos, info)

				if err := handler.tryCreateUnsoldStatement(&unsoldcrud.Req{
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

				if err := handler.tryCreateOrUpdateGoodLedger(&goodledgercrud.Req{
					GoodID:     req.GoodID,
					CoinTypeID: req.CoinTypeID,
					Amount:     req.TotalAmount,
					ToPlatform: &toPlatform,
					ToUser:     &toUser,
				}, ctx, tx); err != nil {
					return err
				}
				ids = append(ids, goodStatementID)
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

	return infos, nil
}
