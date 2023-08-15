package statement

import (
	"context"
	"fmt"

	"entgo.io/ent/entc/integration/edgefield/ent/info"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	profitcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/profit"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) tryCreateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if *req.IOSubType != basetypes.IOSubType_MiningBenefit {
		return nil
	}

	key := fmt.Sprintf("ledger-profit:%v:%v:%v", *req.AppID, *req.UserID, *req.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	if _, err := profitcrud.CreateSet(tx.Profit.Create(), &profitcrud.Req{
		AppID:      req.AppID,
		UserID:     req.UserID,
		CoinTypeID: req.CoinTypeID,
		Incoming:   req.Amount,
	}).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.CreateSet(tx.Statement.Create(), req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateOrAddLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) error {

	// tx.Ledger.Query().Where(entledger.)
	toPlatform := h.UnsoldAmount.Add(*h.TechniqueServiceFeeAmount)
	toUser := h.TotalAmount.Sub(toPlatform)

	// create
	if info == nil {
		if _, err := goodledgercrud.CreateSet(
			tx.GoodLedger.Create(),
			&goodledgercrud.Req{
				GoodID:     h.GoodID,
				CoinTypeID: h.CoinTypeID,
				ToPlatform: &toPlatform,
				ToUser:     &toUser,
				Amount:     h.TotalAmount,
			},
		).Save(ctx); err != nil {
			return nil
		}
		return nil
	}

	// update
	id, err := uuid.Parse(info.ID)
	if err != nil {
		return err
	}
	handler = &goodledger1.Handler{
		Req: goodledgercrud.Req{
			ID:         &id,
			ToPlatform: &toPlatform,
			ToUser:     &toUser,
			Amount:     h.TotalAmount,
		},
	}
	if _, err := handler.UpdateGoodLedger(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateStatement(ctx context.Context) (*npool.Statement, error) {
	switch *h.IOType {
	case basetypes.IOType_Incoming:
		switch *h.IOSubType {
		case basetypes.IOSubType_Payment:
		case basetypes.IOSubType_MiningBenefit:
		case basetypes.IOSubType_Commission:
		case basetypes.IOSubType_TechniqueFeeCommission:
		case basetypes.IOSubType_Deposit:
		case basetypes.IOSubType_Transfer:
		case basetypes.IOSubType_OrderRevoke:
		default:
			return nil, fmt.Errorf("io subtype not match io type, io subtype: %v, io type: %v", *h.IOSubType, *h.IOType)
		}
	case basetypes.IOType_Outcoming:
		switch *h.IOSubType {
		case basetypes.IOSubType_Payment:
		case basetypes.IOSubType_Withdrawal:
		case basetypes.IOSubType_Transfer:
		case basetypes.IOSubType_CommissionRevoke:
		default:
			return nil, fmt.Errorf("io subtype not match io type, io subtype: %v, io type: %v", *h.IOSubType, *h.IOType)
		}
	default:
		return nil, fmt.Errorf("invalid io type %v", *h.IOType)
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.createStatement(ctx); err != nil {
		return nil, err
	}

	return h.GetStatement(ctx)
}

func (h *Handler) CreateStatements(ctx context.Context) ([]*npool.Statement, error) {
	// Remove duplicate record first
	reqs := []*crud.Req{}
	for _, req := range h.Reqs {
		h.Conds = &crud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: *req.IOType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *req.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *req.IOExtra},
		}
		exist, err := h.ExistStatementConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			msg := fmt.Sprintf(
				"statement exist! AppID(%v), UserID(%v), CoinTypeID(%v), IOType(%v), IOSubType(%v), IOExtra(%v)",
				*req.AppID, *req.UserID, *req.CoinTypeID, *req.IOType, *req.IOSubType, *req.IOExtra,
			)
			logger.Sugar().Infof(msg)
			continue
		}
		reqs = append(reqs, req)
	}

	ids := []uuid.UUID{}
	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range reqs {
			_fn := func() error {
				id := uuid.New()
				if req.ID == nil {
					req.ID = &id
				}
				if err := handler.tryCreateStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryCreateProfit(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.tryCreateOrAddLedger(req, ctx, tx); err != nil {
					return err
				}

				ids = append(ids, id)
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}

		}

	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
