package statement

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	profitcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/profit"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	if _, err := crud.CreateSet(
		tx.Statement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateOrAddLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := ledgercrud.SetQueryConds(tx.Ledger.Query(), &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: req.UserID},
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

	id := uuid.New()
	if req.ID == nil {
		req.ID = &id
	}

	incoming := decimal.NewFromInt(0)
	outcoming := decimal.NewFromInt(0)
	switch *req.IOType {
	case basetypes.IOType_Incoming:
		incoming = decimal.RequireFromString(req.Amount.String())
	case basetypes.IOType_Outcoming:
		outcoming = decimal.RequireFromString(req.Amount.String())
	default:
		return fmt.Errorf("invalid io type %v", *req.IOType)
	}
	spendable := incoming.Sub(outcoming)

	// create
	if info == nil {
		stm, err := ledgercrud.CreateSet(
			tx.Ledger.Create(),
			&ledgercrud.Req{
				ID:         req.ID,
				UserID:     req.UserID,
				CoinTypeID: h.CoinTypeID,
				Incoming:   &incoming,
				Outcoming:  &outcoming,
				Spendable:  &spendable,
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
	old, err := tx.Ledger.Get(ctx, info.ID)
	if err != nil {
		return err
	}
	if old == nil {
		return fmt.Errorf("ledger not exist, id %v", info.ID)
	}

	stm1, err := ledgercrud.UpdateSet(
		old,
		tx.Ledger.UpdateOneID(info.ID),
		&ledgercrud.Req{
			Incoming:  &incoming,
			Outcoming: &outcoming,
			Spendable: &spendable,
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

	h.Conds = &crud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		IOType:     &cruder.Cond{Op: cruder.EQ, Val: *h.IOType},
		IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *h.IOSubType},
		IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *h.IOExtra},
	}
	exist, err := h.ExistStatementConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		msg := fmt.Sprintf(
			"statement exist! AppID(%v), UserID(%v), CoinTypeID(%v), IOType(%v), IOSubType(%v), IOExtra(%v)",
			*h.AppID, *h.UserID, *h.CoinTypeID, *h.IOType, *h.IOSubType, *h.IOExtra,
		)
		logger.Sugar().Errorf(msg)
		return nil, fmt.Errorf("statement exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	h.Reqs = append(h.Reqs, &crud.Req{
		ID:         h.ID,
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     h.IOType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	})

	infos, err := h.CreateStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("to many records")
	}

	return infos[0], nil
}

func statementKey(in *crud.Req) string {
	extra := sha256.Sum256([]byte(*in.IOExtra))
	return fmt.Sprintf("%v:%v:%v:%v:%v:%v:%v:%v",
		commonpb.Prefix_PrefixCreateStatement,
		*in.AppID,
		*in.UserID,
		*in.CoinTypeID,
		in.IOType.String(),
		in.IOSubType.String(),
		*in.IOExtra,
		extra,
	)
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
			logger.Sugar().Errorf(msg)
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

				key := statementKey(req)
				if err := redis2.TryLock(key, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key)
				}()

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
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Get Statements
	return nil, nil
}
