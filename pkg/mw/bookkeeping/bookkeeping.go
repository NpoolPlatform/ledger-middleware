package bookkeeping

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
	statementhandler "github.com/NpoolPlatform/ledger-middleware/pkg/mw/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type bookkeepingHandler struct {
	*Handler
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

//nolint
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

// nolint
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

//nolint
func (h *bookkeepingHandler) tryBookKeeping(statements []statementInfo, ctx context.Context) error {
	// TODO: Remove duplicate record first

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, val := range statements {
			key := statementKey(val.Req)
			if err := redis2.TryLock(key, 0); err != nil {
				return err
			}
			//TODO: defer in for loop
			defer func() {
				_ = redis2.Unlock(key)
			}()

			handler := &statementhandler.Handler{
				Conds: &crud.Conds{
					AppID:      &cruder.Cond{Op: cruder.EQ, Val: val.AppID},
					UserID:     &cruder.Cond{Op: cruder.EQ, Val: val.UserID},
					CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: val.CoinTypeID},
					IOType:     &cruder.Cond{Op: cruder.EQ, Val: val.IOType},
					IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: val.IOSubType},
					IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: val.IOExtra},
				},
			}
			exist, err := handler.ExistStatementConds(ctx)
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

func (h *Handler) BookKeeping(ctx context.Context) error {
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
	return handler.tryBookKeeping(statements, ctx)
}

//nolint
func (h *Handler) BookKeepingOut(ctx context.Context) error {

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
			handler := &statementhandler.Handler{
				Conds: &crud.Conds{
					AppID:      &cruder.Cond{Op: cruder.EQ, Val: val.AppID},
					UserID:     &cruder.Cond{Op: cruder.EQ, Val: val.UserID},
					CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: val.CoinTypeID},
					IOType:     &cruder.Cond{Op: cruder.EQ, Val: val.IOType},
					IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: val.IOSubType},
					IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: val.IOExtra},
				},
			}

			info, err := handler.GetStatementOnly(ctx)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.ID)
				if err != nil {
					return err
				}
				h.ID = &id
				if _, err := handler.DeleteStatement(ctx); err != nil {
					return err
				}
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

func (h *Handler) LockBalanceOut(ctx context.Context) error {
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
		ledger1 := &ledgerhandler.Handler{
			Req: ledgercrud.Req{
				AppID:      h.AppID,
				UserID:     h.UserID,
				CoinTypeID: h.CoinTypeID,
			},
			Conds: &ledgercrud.Conds{
				AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
				UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
				CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
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

			locked, err := decimal.NewFromString(fmt.Sprintf("-%v", h.Amount.String()))
			if err != nil {
				return err
			}
			ledger1 := &ledgerhandler.Handler{
				Req: ledgercrud.Req{
					ID:         &ledgerID,
					AppID:      h.AppID,
					UserID:     h.UserID,
					CoinTypeID: h.CoinTypeID,
					Locked:     &locked,
					Spendable:  h.Amount,
				},
			}
			if _, err := ledger1.UpdateLedger(ctx); err != nil {
				return err
			}
		}

		return nil
	})
}

func (h *Handler) LockBalance(ctx context.Context) error {
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

	handler := &bookkeepingHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		_ledgerID, err := handler.tryCreateLedger(&crud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			IOType:     h.IOType,
			IOSubType:  h.IOSubType,
			IOExtra:    h.IOExtra,
		}, ctx, tx)
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

//nolint
func (h *Handler) UnlockBalanceOut(ctx context.Context) error {
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

	ioType := basetypes.IOType_Outcoming
	h.IOType = &ioType

	key := statementKey(&crud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     h.IOType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	})
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler := &statementhandler.Handler{
		Conds: &crud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: h.IOType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: h.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
		},
	}
	statement, err := handler.GetStatementOnly(ctx)
	if err != nil {
		return err
	}

	if statement != nil {
		statementID, err := uuid.Parse(statement.ID)
		if err != nil {
			return err
		}
		h.ID = &statementID

		//TODO:
		spendable := h.Unlocked.Sub(*h.Outcoming)

		_outcoming, err := decimal.NewFromString(fmt.Sprintf("-%v", *h.Outcoming))
		if err != nil {
			return err
		}

		return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
			ledger1 := &ledgerhandler.Handler{
				Req: ledgercrud.Req{
					AppID:      h.AppID,
					UserID:     h.UserID,
					CoinTypeID: h.CoinTypeID,
					Locked:     h.Unlocked,
					Spendable:  &spendable,
					Outcoming:  &_outcoming,
				},
			}
			if _, err := ledger1.UpdateLedger(ctx); err != nil {
				return err
			}

			if h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
				return nil
			}

			if _, err := handler.DeleteStatement(ctx); err != nil {
				return err
			}

			return nil
		})
	}
	return nil
}

//nolint
func (h *Handler) UnlockBalance(ctx context.Context) error {
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

	key := statementKey(&crud.Req{
		AppID:      h.AppID,
		UserID:     h.UserID,
		CoinTypeID: h.CoinTypeID,
		IOType:     h.IOType,
		IOSubType:  h.IOSubType,
		IOExtra:    h.IOExtra,
	})
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	ioType := basetypes.IOType_Outcoming
	h.IOType = &ioType

	handler := &statementhandler.Handler{
		Conds: &crud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: h.CoinTypeID},
			IOType:     &cruder.Cond{Op: cruder.EQ, Val: h.IOType},
			IOSubType:  &cruder.Cond{Op: cruder.EQ, Val: h.IOSubType},
			IOExtra:    &cruder.Cond{Op: cruder.LIKE, Val: h.IOExtra},
		},
	}
	exist, err := handler.ExistStatementConds(ctx)
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

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		bookkeeping1 := &bookkeepingHandler{
			Handler: h,
		}
		if _, err := bookkeeping1.tryCreateLedger(&crud.Req{
			AppID:      h.AppID,
			UserID:     h.UserID,
			CoinTypeID: h.CoinTypeID,
			IOType:     h.IOType,
			IOSubType:  h.IOSubType,
			IOExtra:    h.IOExtra,
		}, ctx, tx); err != nil {
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
		if _, err := ledger1.UpdateLedger(ctx); err != nil {
			return err
		}

		if h.Outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
			return nil
		}

		if _, err := handler.CreateStatement(ctx); err != nil {
			return err
		}
		return nil
	})
}
