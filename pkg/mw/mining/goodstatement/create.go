package goodstatement

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodledger"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/mining/goodstatement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodstatement"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *Handler) CreateGoodStatement(ctx context.Context) (*npool.GoodStatement, error) {
	h.Conds = &crud.Conds{
		GoodID:      &cruder.Cond{Op: cruder.EQ, Val: h.GoodID},
		CoinTypeID:  &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		BenefitDate: &cruder.Cond{Op: cruder.EQ, Val: h.BenefitDate},
	}

	exist, err := h.ExistGoodStatementConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("statement exist, goodid(%v), cointypeid(%v), benefitdate(%v)", *h.GoodID, *h.CoinTypeID, *h.BenefitDate)
	}

	// id := uuid.New()
	// if h.ID == nil {
	// 	h.ID = &id
	// }

	// if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
	// 	if _, err := crud.CreateSet(
	// 		cli.GoodStatement.Create(),
	// 		&h.Req,
	// 	).Save(_ctx); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }); err != nil {
	// 	return nil, err
	// }

	return h.GetGoodStatement(ctx)
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

func (h *createHandler) tryCreateGoodStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.CreateSet(
		tx.GoodStatement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateUnsoldStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.CreateSet(
		tx.GoodStatement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	reqs := []*crud.Req{}

	ids := []uuid.UUID{}
	handler := &createHandler{
		Handler: h,
	}

	db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range reqs {
			_fn := func() error {
				id := uuid.New()
				if req.ID == nil {
					req.ID = &id
				}

				key := fmt.Sprintf("ledger-create-goodstatement:%v:%v:%v", *h.GoodID, *h.CoinTypeID, *h.BenefitDate)
				if err := redis2.TryLock(key, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key)
				}()

				if err := handler.tryCreateGoodStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryCreateUnsoldStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryCreateOrUpdateGoodLedger(&goodledgercrud.Req{
					GoodID: req.GoodID,
					CoinTypeID: req.CoinTypeID,
					Amount: req.Amount,
					ToPlatform: req.,
				}, ctx, tx); err != nil {
					return err
				}
				return nil
			}

			if err := _fn(); err != nil {
				return err
			}
			return nil
		}

		return nil
	})
}
