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

func (h *createHandler) createGoodStatement(ctx context.Context, tx *ent.Tx, req *goodstatementcrud.Req) error {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateGoodLedgerStatement, *req.GoodID, *req.CoinTypeID, *req.BenefitDate)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	exist, err := tx.
		GoodStatement.
		Query().
		Where(
			entgoodstatement.GoodID(*req.GoodID),
			entgoodstatement.CoinTypeID(*req.CoinTypeID),
			entgoodstatement.BenefitDate(*req.BenefitDate),
			entgoodstatement.DeletedAt(0),
		).
		Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("good statement already exist")
	}

	toPlatform := req.UnsoldAmount.Add(*req.TechniqueServiceFeeAmount)
	toUser := req.TotalAmount.Sub(toPlatform)
	if req.TotalAmount.Cmp(toPlatform.Add(toUser)) != 0 {
		return fmt.Errorf(
			"TotalAmount(%v) != ToPlatform(%v) + ToUser(%v)",
			req.TotalAmount.String(),
			toPlatform.String(),
			toUser.String(),
		)
	}
	if _, err := goodstatementcrud.CreateSet(
		tx.GoodStatement.Create(),
		&goodstatementcrud.Req{
			EntID:                     req.EntID,
			GoodID:                    req.GoodID,
			CoinTypeID:                req.CoinTypeID,
			BenefitDate:               req.BenefitDate,
			TotalAmount:               req.TotalAmount,
			ToPlatform:                &toPlatform,
			ToUser:                    &toUser,
			TechniqueServiceFeeAmount: req.TechniqueServiceFeeAmount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createUnsoldStatement(ctx context.Context, tx *ent.Tx, req *goodstatementcrud.Req) error {
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
			StatementID: req.EntID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createOrUpdateGoodLedger(ctx context.Context, tx *ent.Tx, req *goodstatementcrud.Req) error {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateGoodLedger, *req.GoodID, *req.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()
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

	if req.TotalAmount.Cmp(req.UnsoldAmount.Add(*req.TechniqueServiceFeeAmount)) < 0 {
		return fmt.Errorf(
			"unsold amount(%v) + techniqueservicefeeamount(%v) < total amount(%v)",
			req.TotalAmount,
			req.UnsoldAmount,
			req.TechniqueServiceFeeAmount,
		)
	}
	toPlatform := req.UnsoldAmount.Add(*req.TechniqueServiceFeeAmount)
	toUser := req.TotalAmount.Sub(toPlatform)
	if req.TotalAmount.Cmp(toPlatform.Add(toUser)) != 0 {
		return fmt.Errorf(
			"TotalAmount(%v) != ToPlatform(%v) + ToUser(%v)",
			req.TotalAmount,
			toPlatform,
			toUser,
		)
	}

	if info == nil {
		if _, err := goodledgercrud.CreateSet(
			tx.GoodLedger.Create(),
			&goodledgercrud.Req{
				GoodID:     req.GoodID,
				CoinTypeID: req.CoinTypeID,
				Amount:     req.TotalAmount,
				ToPlatform: &toPlatform,
				ToUser:     &toUser,
			},
		).Save(ctx); err != nil {
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

func (h *Handler) CreateGoodStatements(ctx context.Context) ([]*npool.GoodStatement, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				id := uuid.New()
				if req.EntID == nil {
					req.EntID = &id
				}
				if err := handler.createGoodStatement(ctx, tx, req); err != nil {
					return err
				}
				if err := handler.createUnsoldStatement(ctx, tx, req); err != nil {
					return err
				}
				if err := handler.createOrUpdateGoodLedger(ctx, tx, req); err != nil {
					return err
				}
				ids = append(ids, *req.EntID)
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
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
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
	h.Reqs = []*goodstatementcrud.Req{&h.Req}
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
