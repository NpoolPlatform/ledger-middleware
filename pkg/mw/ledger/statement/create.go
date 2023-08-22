package statement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	profitcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/profit"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	commonpb "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) tryCreateOrUpdateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if *req.IOSubType != basetypes.IOSubType_MiningBenefit {
		return nil
	}

	stm, err := profitcrud.SetQueryConds(
		tx.Profit.Query(),
		&profitcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	// create
	if info == nil {
		key := fmt.Sprintf("%v:%v:%v:%v", commonpb.Prefix_PrefixCreateLedgerProfit, *req.AppID, *req.UserID, *req.CoinTypeID)
		if err := redis2.TryLock(key, 0); err != nil {
			return err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()
		stm, err := profitcrud.CreateSet(tx.Profit.Create(), &profitcrud.Req{
			AppID:      req.AppID,
			UserID:     req.UserID,
			CoinTypeID: req.CoinTypeID,
			Incoming:   req.Amount,
		})
		if err != nil {
			return err
		}
		if _, err := stm.Save(ctx); err != nil {
			return err
		}
		return nil
	}

	// update
	old, err := tx.Profit.Get(ctx, info.ID)
	if err != nil {
		return err
	}

	stm1, err := profitcrud.UpdateSet(
		old,
		tx.Profit.UpdateOneID(info.ID),
		&profitcrud.Req{
			AppID:      req.AppID,
			UserID:     req.UserID,
			CoinTypeID: req.CoinTypeID,
			Incoming:   req.Amount,
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
func (h *createHandler) tryCreateStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := crud.SetQueryConds(
		tx.Statement.Query(),
		&crud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: *req.IOType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: *req.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: *req.IOExtra},
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
		return fmt.Errorf("statement already exist")
	}

	key := fmt.Sprintf("%v:%v:%v:%v",
		commonpb.Prefix_PrefixCreateLedgerStatement,
		*req.AppID,
		*req.UserID,
		*req.CoinTypeID,
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	if _, err := crud.CreateSet(
		tx.Statement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) tryCreateOrUpdateLedger(req *ledgercrud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := ledgercrud.SetQueryConds(tx.Ledger.Query(), &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
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

	// create
	if info == nil {
		key := fmt.Sprintf("%v:%v:%v:%v",
			commonpb.Prefix_PrefixCreateLedger,
			*req.AppID,
			*req.UserID,
			*req.CoinTypeID,
		)
		if err := redis2.TryLock(key, 0); err != nil {
			return err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()

		stm, err := ledgercrud.CreateSet(
			tx.Ledger.Create(),
			req,
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
			Incoming:  req.Incoming,
			Outcoming: req.Outcoming,
			Spendable: req.Spendable,
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
		return nil, fmt.Errorf("statement already exist")
	}

	h.Reqs = []*crud.Req{&h.Req}
	infos, err := h.CreateStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	return infos[0], nil
}

//nolint
func (h *Handler) CreateStatements(ctx context.Context) ([]*npool.Statement, error) {
	// to ensure the accuracy of the ledger, the same batch of data cannot be written repeatedly.
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

				if err := handler.tryCreateStatement(req, ctx, tx); err != nil {
					return err
				}

				if err := handler.tryCreateOrUpdateProfit(req, ctx, tx); err != nil {
					return err
				}

				if h.ChangeLedger != nil && !*h.ChangeLedger { // just create statement, do not update ledger
					return nil
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
				locked := decimal.NewFromInt(0)

				if err := handler.tryCreateOrUpdateLedger(&ledgercrud.Req{
					AppID:      req.AppID,
					UserID:     req.UserID,
					CoinTypeID: req.CoinTypeID,
					Incoming:   &incoming,
					Outcoming:  &outcoming,
					Spendable:  &spendable,
					Locked:     &locked,
				}, ctx, tx); err != nil {
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

	h.Conds = &crud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (h *createHandler) tryDeleteStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if req.ID == nil {
		return fmt.Errorf("invalid statement id")
	}
	now := uint32(time.Now().Unix())
	if _, err := crud.UpdateSet(
		tx.Statement.UpdateOneID(*req.ID),
		&crud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint
func (h *Handler) UnCreateStatements(ctx context.Context) ([]*npool.Statement, error) {
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return nil, fmt.Errorf("invalid id in uncrate statements")
		}
		ids = append(ids, *req.ID)
	}

	h.Conds = &crud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}

	statementMap := map[string]*npool.Statement{}
	for _, info := range infos {
		statementMap[info.ID] = info
	}

	reqs := []*crud.Req{}
	for _, req := range h.Reqs {
		_, ok := statementMap[req.ID.String()]
		if !ok {
			logger.Sugar().Errorf("statement not exist, id %v", req.ID)
			continue
		}
		reqs = append(reqs, req)
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range reqs {
			_fn := func() error {
				key := fmt.Sprintf("ledger-uncreate-statement:%v:%v:%v:%v", *req.AppID, *req.UserID, *req.CoinTypeID, *req.IOExtra)
				if err := redis2.TryLock(key, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key)
				}()

				if err := handler.tryDeleteStatement(req, ctx, tx); err != nil {
					return err
				}

				profitAmount := decimal.RequireFromString(fmt.Sprintf("-%v", req.Amount.String()))
				if err := handler.tryCreateOrUpdateProfit(&crud.Req{
					AppID:      req.AppID,
					UserID:     req.UserID,
					CoinTypeID: req.CoinTypeID,
					Amount:     &profitAmount,
					IOSubType:  req.IOSubType,
				}, ctx, tx); err != nil {
					return err
				}

				incoming := decimal.NewFromInt(0)
				outcoming := decimal.NewFromInt(0)
				switch *req.IOType {
				case basetypes.IOType_Incoming:
					incoming = decimal.RequireFromString(fmt.Sprintf("-%v", req.Amount.String()))
				case basetypes.IOType_Outcoming:
					outcoming = decimal.RequireFromString(fmt.Sprintf("-%v", req.Amount.String()))
				default:
					return fmt.Errorf("invalid io type %v", *req.IOType)
				}
				spendable := incoming.Sub(outcoming)

				if err := handler.tryCreateOrUpdateLedger(&ledgercrud.Req{
					AppID:      req.AppID,
					UserID:     req.UserID,
					CoinTypeID: req.CoinTypeID,
					Incoming:   &incoming,
					Outcoming:  &outcoming,
					Spendable:  &spendable,
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

	return infos, nil
}
