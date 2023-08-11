package statement

import (
	"context"
	"crypto/sha256"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	ledgercrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/ledger"
	profitcrud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/profit"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	ledgerhandler "github.com/NpoolPlatform/ledger-middleware/pkg/mw/ledger"
	profithandler "github.com/NpoolPlatform/ledger-middleware/pkg/mw/profit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type bookkeepingHandler struct {
	*Handler
	Unlocked  *decimal.Decimal
	Outcoming *decimal.Decimal
}

func statementKey(in *crud.Req) string {
	extra := sha256.Sum256([]byte(*in.IOExtra))
	return fmt.Sprintf("ledger-statement:%v:%v:%v:%v:%v:%v:%v",
		*in.AppID,
		*in.UserID,
		*in.CoinTypeID,
		in.IOType.String(),
		in.IOSubType.String(),
		*in.IOExtra,
		extra,
	)
}

func (h *bookkeepingHandler) tryCreateLedger(req *crud.Req, ctx context.Context, tx *ent.Tx) (string, error) {
	key := fmt.Sprintf("ledger-ledger:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)

	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ledger1 := &ledgerhandler.Handler{
		Req: ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
		},
	}
	ledger1.Conds = &ledgercrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
	}

	ledger, err := ledger1.GetLedgerOnly(ctx)
	if err != nil {
		return "", err
	}
	if ledger != nil {
		return ledger.ID, nil
	}

	info, err := ledger1.CreateLedger(ctx)
	if err != nil {
		return "", err
	}
	return info.ID, nil
}

func (h *bookkeepingHandler) tryCreateProfit(req *crud.Req, ctx context.Context, tx *ent.Tx) (string, error) {
	key := fmt.Sprintf("ledger-profit:%v:%v:%v", *h.AppID, *h.UserID, *h.CoinTypeID)

	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	profit1 := &profithandler.Handler{
		Req: profitcrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
		},
	}
	profit1.Conds = &profitcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
	}

	profit, err := profit1.GetProfitOnly(ctx)
	if err != nil {
		return "", err
	}
	if profit != nil {
		return profit.ID, nil
	}

	info, err := profit1.CreateProfit(ctx)
	if err != nil {
		return "", err
	}

	return info.ID, nil
}

type statementInfo struct {
	*crud.Req
	LedgerID string
	ProfitID string
}

func (h *bookkeepingHandler) tryBookKeepingV2(statements []statementInfo, ctx context.Context) error {
	// TODO: Remove duplicate record first

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, val := range statements {
			key := statementKey(val.Req)
			if err := redis2.TryLock(key, 0); err != nil {
				return err
			}
			defer func() {
				_ = redis2.Unlock(key)
			}()

			h.Conds = &crud.Conds{
				AppID:      &cruder.Cond{Op: cruder.EQ, Val: val.AppID},
				UserID:     &cruder.Cond{Op: cruder.EQ, Val: val.UserID},
				CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: val.CoinTypeID},
				IOType:     &cruder.Cond{Op: cruder.EQ, Val: val.IOType},
				IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: val.IOSubType},
				IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: val.IOExtra},
			}
			exist, err := h.ExistStatementConds(ctx)
			if err != nil {
				return err
			}
			// TODO: Return Or Continue
			if exist {
				continue
			}

			if _, err := crud.CreateSet(tx.Statement.Create(), val.Req).Save(ctx); err != nil {
				return err
			}

			incoming := decimal.NewFromInt(0)
			outcoming := decimal.NewFromInt(0)

			switch *val.IOType {
			case basetypes.IOType_Incoming:
				incoming = decimal.RequireFromString(val.Amount.String())
			case basetypes.IOType_Outcoming:
				outcoming = decimal.RequireFromString(val.Amount.String())
			default:
				return fmt.Errorf("invalid io type %v", *val.IOType)
			}

			spendable := incoming.Sub(outcoming)

			ledgerID, err := uuid.Parse(val.LedgerID)
			if err != nil {
				return err
			}
			ledger1 := &ledgerhandler.Handler{
				Req: ledgercrud.Req{
					ID:         &ledgerID,
					AppID:      val.AppID,
					UserID:     val.UserID,
					CoinTypeID: val.CoinTypeID,
					Incoming:   &incoming,
					Outcoming:  &outcoming,
					Spendable:  &spendable,
				},
			}
			if _, err := ledger1.UpdateLedger(ctx); err != nil {
				return err
			}

			profitAmount := decimal.NewFromInt(0)
			if *val.IOType == basetypes.IOType_Incoming {
				if *val.IOSubType == basetypes.IOSubType_MiningBenefit {
					profitAmount = incoming
				}
			}
			if profitAmount.Cmp(decimal.NewFromInt(0)) == 0 {
				return nil
			}

			profitID, err := uuid.Parse(val.ProfitID)
			if err != nil {
				return err
			}
			profit1 := &profithandler.Handler{
				Req: profitcrud.Req{
					ID:         &profitID,
					AppID:      val.AppID,
					UserID:     val.UserID,
					CoinTypeID: val.CoinTypeID,
					Incoming:   &profitAmount,
				},
			}
			if _, err := profit1.UpdateProfit(ctx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (h *Handler) BookKeepingV2(ctx context.Context) error {
	handler := &bookkeepingHandler{
		Handler: h,
	}

	statements := []statementInfo{}
	for _, req := range h.Reqs {
		err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			ledgerID, err := handler.tryCreateLedger(req, _ctx, tx)
			if err != nil {
				return err
			}
			profitID, err := handler.tryCreateProfit(req, _ctx, tx)
			if err != nil {
				return err
			}
			statements = append(statements, statementInfo{
				Req:      req,
				LedgerID: ledgerID,
				ProfitID: profitID,
			})
			return nil
		})
		if err != nil {
			return err
		}
	}
	return handler.tryBookKeepingV2(statements, ctx)
}

func (h *Handler) BookKeepingV2Out(ctx context.Context) error {

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, val := range h.Reqs {
			key := statementKey(val)
			if err := redis2.TryLock(key, 0); err != nil {
				return err
			}
			defer func() {
				_ = redis2.Unlock(key)
			}()

			// deal statement
			h.Conds = &crud.Conds{
				AppID:      &cruder.Cond{Op: cruder.EQ, Val: val.AppID},
				UserID:     &cruder.Cond{Op: cruder.EQ, Val: val.UserID},
				CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: val.CoinTypeID},
				IOType:     &cruder.Cond{Op: cruder.EQ, Val: val.IOType},
				IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: val.IOSubType},
				IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: val.IOExtra},
			}

			info, err := h.GetStatementOnly(ctx)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.ID)
				if err != nil {
					return err
				}
				h.ID = &id
				h.DeleteStatement(ctx)
			}

			// deal ledger
			incoming := decimal.NewFromInt(0)
			outcoming := decimal.NewFromInt(0)

			switch *val.IOType {
			case basetypes.IOType_Incoming:
				incoming = decimal.RequireFromString(val.Amount.String())
			case basetypes.IOType_Outcoming:
				outcoming = decimal.RequireFromString(val.Amount.String())
			default:
				return fmt.Errorf("invalid io type %v", *val.IOType)
			}

			spendable := incoming.Sub(outcoming)

			ledger1 := &ledgerhandler.Handler{
				Req: ledgercrud.Req{
					AppID:      val.AppID,
					UserID:     val.UserID,
					CoinTypeID: val.CoinTypeID,
				},
				Conds: &ledgercrud.Conds{
					AppID:      &cruder.Cond{Op: cruder.EQ, Val: val.AppID},
					UserID:     &cruder.Cond{Op: cruder.EQ, Val: val.UserID},
					CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: val.CoinTypeID},
				},
			}

			ledger, err := ledger1.GetLedgerOnly(ctx)
			if err != nil {
				return err
			}
			if ledger != nil {
				ledgerID, err := uuid.Parse(ledger.ID)
				if err != nil {
					return err
				}

				_incoming, err := decimal.NewFromString(fmt.Sprintf("-%v", incoming.String()))
				if err != nil {
					return err
				}
				_outcoming, err := decimal.NewFromString(fmt.Sprintf("-%v", outcoming.String()))
				if err != nil {
					return err
				}
				_spendable, err := decimal.NewFromString(fmt.Sprintf("-%v", spendable.String()))
				if err != nil {
					return err
				}

				ledger1 := &ledgerhandler.Handler{
					Req: ledgercrud.Req{
						ID:         &ledgerID,
						AppID:      val.AppID,
						UserID:     val.UserID,
						CoinTypeID: val.CoinTypeID,
						Incoming:   &_incoming,
						Outcoming:  &_outcoming,
						Spendable:  &_spendable,
					},
				}
				if _, err := ledger1.UpdateLedger(ctx); err != nil {
					return err
				}
			}

			if err != nil {
				return err
			}

			// deal profit
			profitAmount := decimal.NewFromInt(0)
			if *val.IOType == basetypes.IOType_Incoming {
				if *val.IOSubType == basetypes.IOSubType_MiningBenefit {
					profitAmount = incoming
				}
			}
			if profitAmount.Cmp(decimal.NewFromInt(0)) == 0 {
				return nil
			}

			profit1 := &profithandler.Handler{
				Req: profitcrud.Req{
					AppID:      val.AppID,
					UserID:     val.UserID,
					CoinTypeID: val.CoinTypeID,
				},
				Conds: &profitcrud.Conds{
					AppID:      &cruder.Cond{Op: cruder.EQ, Val: val.AppID},
					UserID:     &cruder.Cond{Op: cruder.EQ, Val: val.UserID},
					CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: val.CoinTypeID},
				},
			}
			profit, err := profit1.GetProfitOnly(ctx)
			if err != nil {
				return err
			}
			if profit != nil {
				profitID, err := uuid.Parse(profit.ID)
				if err != nil {
					return err
				}
				_profitAmount, err := decimal.NewFromString(fmt.Sprintf("-%v", profitAmount.String()))
				if err != nil {
					return err
				}
				profit1 = &profithandler.Handler{
					Req: profitcrud.Req{
						ID:         &profitID,
						AppID:      val.AppID,
						UserID:     val.UserID,
						CoinTypeID: val.CoinTypeID,
						Incoming:   &_profitAmount,
					},
				}
			}
			if _, err := profit1.UpdateProfit(ctx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (h *bookkeepingHandler) LockBalance(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid coin type id")
	}
	if h.Amount == nil {
		return fmt.Errorf("invalid amount in lock balance")
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		_ledgerID, err := h.tryCreateLedger(&h.Req, ctx, tx)
		if err != nil {
			return err
		}
		ledgerID, err := uuid.Parse(_ledgerID)
		if err != nil {
			return err
		}

		spendable, err := decimal.NewFromString(fmt.Sprintf("-%v", h.Amount.String()))
		if err != nil {
			return err
		}

		ledger1 := &ledgerhandler.Handler{
			Req: ledgercrud.Req{
				ID:         &ledgerID,
				AppID:      h.AppID,
				UserID:     h.UserID,
				CoinTypeID: h.CoinTypeID,
				Locked:     h.Amount,
				Spendable:  &spendable,
			},
		}
		if _, err := ledger1.UpdateLedger(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (h *bookkeepingHandler) UnlockBalance(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return fmt.Errorf("invalid coin type id")
	}
	if h.Unlocked == nil {
		return fmt.Errorf("invalid unlocked")
	}
	if h.Outcoming == nil {
		return fmt.Errorf("invalid outcoming")
	}
	if h.IOExtra == nil {
		return fmt.Errorf("invalid extra")
	}
	if h.Unlocked.Cmp(decimal.NewFromInt(0)) == 0 && h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
		return fmt.Errorf("nothing todo")
	}

	key := statementKey(&h.Req)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ioType := basetypes.IOType_Outcoming
	h.IOType = &ioType

	h.Conds = &crud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
		UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
		IOType:     &cruder.Cond{Op: cruder.EQ, Val: h.IOType},
		IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: h.IOSubType},
		IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
	}
	exist, err := h.ExistStatementConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("statement already exist, app id %v, user id %v, coin type id %v", *h.AppID, *h.UserID, *h.CoinTypeID)
	}

	spendable := h.Unlocked.Sub(*h.Outcoming)
	h.Amount = h.Outcoming
	_unlocked, err := decimal.NewFromString(fmt.Sprintf("-%v", *h.Unlocked))
	if err != nil {
		return err
	}

	ledger1 := &ledgerhandler.Handler{
		Req: ledgercrud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			Locked:     &_unlocked,
			Spendable:  &spendable,
			Outcoming:  h.Outcoming,
		},
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := h.tryCreateLedger(&h.Req, ctx, tx); err != nil {
			return err
		}

		if _, err := ledger1.UpdateLedger(ctx); err != nil {
			return err
		}

		if h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
			return nil
		}

		if _, err := h.CreateStatement(ctx); err != nil {
			return err
		}
		return nil
	})
}
