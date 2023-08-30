package statement

import (
	"context"
	"fmt"

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
	ids []uuid.UUID
}

//nolint
func (h *createHandler) validate() error {
	for _, req := range h.Reqs {
		if req.AppID == nil {
			return fmt.Errorf("invalid app id")
		}
		if req.UserID == nil {
			return fmt.Errorf("invalid user id")
		}
		if req.CoinTypeID == nil {
			return fmt.Errorf("invalid coin type id")
		}
		if req.Amount == nil {
			return fmt.Errorf("invalid amount")
		}
		if req.IOExtra == nil {
			return fmt.Errorf("invalid io extra")
		}
		if req.IOType == nil {
			return fmt.Errorf("invalid io type")
		}
		if req.IOSubType == nil {
			return fmt.Errorf("invalid io sub type")
		}
		switch *req.IOType {
		case basetypes.IOType_Incoming:
			switch *req.IOSubType {
			case basetypes.IOSubType_Payment:
			case basetypes.IOSubType_MiningBenefit:
			case basetypes.IOSubType_Commission:
			case basetypes.IOSubType_TechniqueFeeCommission:
			case basetypes.IOSubType_Deposit:
			case basetypes.IOSubType_Transfer:
			case basetypes.IOSubType_OrderRevoke:
			default:
				return fmt.Errorf("io subtype not match io type")
			}
		case basetypes.IOType_Outcoming:
			switch *req.IOSubType {
			case basetypes.IOSubType_Payment:
			case basetypes.IOSubType_Withdrawal:
			case basetypes.IOSubType_Transfer:
			case basetypes.IOSubType_CommissionRevoke:
			default:
				return fmt.Errorf("io subtype not match io type")
			}
		default:
			return fmt.Errorf("invalid io type %v", *req.IOType)
		}
	}
	return nil
}

func (h *createHandler) createOrUpdateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
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
		stm, err := profitcrud.CreateSetWithValidate(
			tx.Profit.Create(),
			&profitcrud.Req{
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

	stm1, err := profitcrud.UpdateSetWithValidate(
		old,
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

func (h *createHandler) createStatement(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	if req.ID == nil {
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

	id := uuid.New()
	if req.ID == nil {
		req.ID = &id
	}
	h.ids = append(h.ids, *req.ID)

	if _, err := crud.CreateSet(
		tx.Statement.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createOrUpdateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) error {
	stm, err := ledgercrud.SetQueryConds(
		tx.Ledger.Query(),
		&ledgercrud.Conds{
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

		if _, err := ledgercrud.CreateSet(
			tx.Ledger.Create(),
			&ledgercrud.Req{
				AppID:      req.AppID,
				UserID:     req.UserID,
				CoinTypeID: req.CoinTypeID,
				Incoming:   &incoming,
				Outcoming:  &outcoming,
				Locked:     &locked,
				Spendable:  &spendable,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	stm1, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming:  &incoming,
			Outcoming: &outcoming,
			Spendable: &spendable,
			Locked:    &locked,
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

func (h *Handler) CreateStatements(ctx context.Context) ([]*npool.Statement, error) {
	handler := &createHandler{
		Handler: h,
	}
	handler.ids = []uuid.UUID{}
	if err := handler.validate(); err != nil {
		return nil, err
	}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			_fn := func() error {
				if err := handler.createStatement(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.createOrUpdateProfit(req, ctx, tx); err != nil {
					return err
				}
				if err := handler.createOrUpdateLedger(req, ctx, tx); err != nil {
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

	h.Conds = &crud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: handler.ids},
	}
	h.Offset = 0
	h.Limit = int32(len(handler.ids))

	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (h *Handler) CreateStatement(ctx context.Context) (*npool.Statement, error) {
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
