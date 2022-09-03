//nolint:dupl
package ledger

import (
	"context"
	"crypto/sha256"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/profit"

	detailcli "github.com/NpoolPlatform/ledger-manager/pkg/client/detail"
	detailcrud "github.com/NpoolPlatform/ledger-manager/pkg/crud/detail"
	detailmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/detail"

	generalcli "github.com/NpoolPlatform/ledger-manager/pkg/client/general"
	generalcrud "github.com/NpoolPlatform/ledger-manager/pkg/crud/general"
	generalmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/general"

	profitcli "github.com/NpoolPlatform/ledger-manager/pkg/client/profit"
	profitcrud "github.com/NpoolPlatform/ledger-manager/pkg/crud/profit"
	profitmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/profit"

	commonpb "github.com/NpoolPlatform/message/npool"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	errno "github.com/NpoolPlatform/ledger-middleware/pkg/errno"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func TryCreateProfit(ctx context.Context, appID, userID, coinTypeID string) (string, error) {
	key := fmt.Sprintf("ledger-profit:%v:%v:%v", appID, userID, coinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	profit1, err := profitcli.GetProfitOnly(ctx, &profitmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: coinTypeID,
		},
	})
	if err != nil {
		return "", err
	}
	if profit1 != nil {
		return profit1.ID, nil
	}

	profit1, err = profitcli.CreateProfit(ctx, &profitmgrpb.ProfitReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
	})
	if err != nil {
		return "", err
	}

	return profit1.ID, nil
}

func TryCreateGeneral(ctx context.Context, appID, userID, coinTypeID string) (string, error) {
	key := fmt.Sprintf("ledger-general:%v:%v:%v", appID, userID, coinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	general1, err := generalcli.GetGeneralOnly(ctx, &generalmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: coinTypeID,
		},
	})
	if err != nil {
		return "", err
	}
	if general1 != nil {
		return general1.ID, nil
	}

	general1, err = generalcli.CreateGeneral(ctx, &generalmgrpb.GeneralReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
	})
	if err != nil {
		return "", err
	}

	return general1.ID, nil
}

func detailKey(in *detailmgrpb.DetailReq) string {
	extra := sha256.Sum256([]byte(in.GetIOExtra()))
	return fmt.Sprintf("ledger-detail:%v:%v:%v:%v:%v:%v:%v",
		in.GetAppID(),
		in.GetUserID(),
		in.GetCoinTypeID(),
		in.GetIOType(),
		in.GetIOSubType(),
		in.GetIOExtra(),
		extra,
	)
}

func BookKeeping(ctx context.Context, in *detailmgrpb.DetailReq) error { //nolint
	val, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		return err
	}
	if val.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("invalid amount")
	}

	var generalID string
	var profitID string

	if generalID, err = TryCreateGeneral(
		ctx,
		in.GetAppID(), in.GetUserID(), in.GetCoinTypeID(),
	); err != nil {
		return err
	}

	if profitID, err = TryCreateProfit(
		ctx,
		in.GetAppID(), in.GetUserID(), in.GetCoinTypeID(),
	); err != nil {
		return err
	}

	key := detailKey(in)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	conds := &detailmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetUserID(),
		},
		IOType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(in.GetIOType()),
		},
		IOSubType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(in.GetIOSubType()),
		},
		IOExtra: &commonpb.StringVal{
			Op:    cruder.LIKE,
			Value: in.GetIOExtra(),
		},
	}

	// For commission, we just ignore coin type ID here
	if in.GetIOSubType() != detailmgrpb.IOSubType_Commission {
		conds.CoinTypeID = &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetCoinTypeID(),
		}
	}

	exist, err := detailcli.ExistDetailConds(ctx, conds)
	if err != nil {
		return err
	}
	if exist {
		return errno.ErrAlreadyExists
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c1, err := detailcrud.CreateSet(tx.Detail.Create(), in)
		if err != nil {
			return err
		}

		_, err = c1.Save(ctx)
		if err != nil {
			return err
		}

		incomingD := decimal.NewFromInt(0)
		outcomingD := decimal.NewFromInt(0)

		switch in.GetIOType() {
		case detailmgrpb.IOType_Incoming:
			incomingD = decimal.RequireFromString(in.GetAmount())
		case detailmgrpb.IOType_Outcoming:
			outcomingD = decimal.RequireFromString(in.GetAmount())
		default:
			return fmt.Errorf("invalid iotype")
		}

		spendableD := incomingD.Sub(outcomingD)

		incoming := incomingD.String()
		outcoming := outcomingD.String()
		spendable := spendableD.String()

		info, err := tx.
			General.
			Query().
			Where(
				general.ID(uuid.MustParse(generalID)),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		c2, err := generalcrud.UpdateSet(info, &generalmgrpb.GeneralReq{
			AppID:      in.AppID,
			UserID:     in.UserID,
			CoinTypeID: in.CoinTypeID,
			Incoming:   &incoming,
			Outcoming:  &outcoming,
			Spendable:  &spendable,
		})
		if err != nil {
			return err
		}

		_, err = c2.Save(ctx)
		if err != nil {
			return err
		}

		profitAmountD := decimal.NewFromInt(0)

		if in.GetIOType() == detailmgrpb.IOType_Incoming {
			if in.GetIOSubType() == detailmgrpb.IOSubType_MiningBenefit {
				profitAmountD = incomingD
			}
		}
		if profitAmountD.Cmp(decimal.NewFromInt(0)) == 0 {
			return nil
		}

		profitAmount := profitAmountD.String()

		info1, err := tx.
			Profit.
			Query().
			Where(
				profit.ID(uuid.MustParse(profitID)),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		c3, err := profitcrud.UpdateSet(info1, &profitmgrpb.ProfitReq{
			AppID:      in.AppID,
			UserID:     in.UserID,
			CoinTypeID: in.CoinTypeID,
			Incoming:   &profitAmount,
		})
		if err != nil {
			return err
		}

		_, err = c3.Save(ctx)
		return err
	})
}

type detailInfo struct {
	Detail    *detailmgrpb.DetailReq
	GeneralID string
	ProfitID  string
}

//nolint:funlen,gocyclo
func BookKeepingV2(ctx context.Context, infos []*detailmgrpb.DetailReq) error {
	detailInfos := []detailInfo{}

	for _, in := range infos {
		val, err := decimal.NewFromString(in.GetAmount())
		if err != nil {
			return err
		}
		if val.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid amount")
		}

		var generalID string
		var profitID string

		if generalID, err = TryCreateGeneral(
			ctx,
			in.GetAppID(), in.GetUserID(), in.GetCoinTypeID(),
		); err != nil {
			return err
		}

		if profitID, err = TryCreateProfit(
			ctx,
			in.GetAppID(), in.GetUserID(), in.GetCoinTypeID(),
		); err != nil {
			return err
		}
		detailInfos = append(detailInfos, detailInfo{
			Detail:    in,
			GeneralID: generalID,
			ProfitID:  profitID,
		})
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, val := range detailInfos {
			conds := &detailmgrpb.Conds{
				AppID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.Detail.GetAppID(),
				},
				UserID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.Detail.GetUserID(),
				},
				IOType: &commonpb.Int32Val{
					Op:    cruder.EQ,
					Value: int32(val.Detail.GetIOType()),
				},
				IOSubType: &commonpb.Int32Val{
					Op:    cruder.EQ,
					Value: int32(val.Detail.GetIOSubType()),
				},
				IOExtra: &commonpb.StringVal{
					Op:    cruder.LIKE,
					Value: val.Detail.GetIOExtra(),
				},
			}

			// For commission, we just ignore coin type ID here
			if val.Detail.GetIOSubType() != detailmgrpb.IOSubType_Commission {
				conds.CoinTypeID = &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.Detail.GetCoinTypeID(),
				}
			}

			exist, err := detailcli.ExistDetailConds(ctx, conds)
			if err != nil {
				return err
			}
			if exist {
				return errno.ErrAlreadyExists
			}

			c1, err := detailcrud.CreateSet(tx.Detail.Create(), val.Detail)
			if err != nil {
				return err
			}

			_, err = c1.Save(ctx)
			if err != nil {
				return err
			}

			incomingD := decimal.NewFromInt(0)
			outcomingD := decimal.NewFromInt(0)

			switch val.Detail.GetIOType() {
			case detailmgrpb.IOType_Incoming:
				incomingD = decimal.RequireFromString(val.Detail.GetAmount())
			case detailmgrpb.IOType_Outcoming:
				outcomingD = decimal.RequireFromString(val.Detail.GetAmount())
			default:
				return fmt.Errorf("invalid iotype")
			}

			spendableD := incomingD.Sub(outcomingD)

			incoming := incomingD.String()
			outcoming := outcomingD.String()
			spendable := spendableD.String()

			info, err := tx.
				General.
				Query().
				Where(
					general.ID(uuid.MustParse(val.GeneralID)),
				).
				ForUpdate().
				Only(ctx)
			if err != nil {
				return err
			}

			c2, err := generalcrud.UpdateSet(info, &generalmgrpb.GeneralReq{
				AppID:      val.Detail.AppID,
				UserID:     val.Detail.UserID,
				CoinTypeID: val.Detail.CoinTypeID,
				Incoming:   &incoming,
				Outcoming:  &outcoming,
				Spendable:  &spendable,
			})
			if err != nil {
				return err
			}

			_, err = c2.Save(ctx)
			if err != nil {
				return err
			}

			profitAmountD := decimal.NewFromInt(0)

			if val.Detail.GetIOType() == detailmgrpb.IOType_Incoming {
				if val.Detail.GetIOSubType() == detailmgrpb.IOSubType_MiningBenefit {
					profitAmountD = incomingD
				}
			}
			if profitAmountD.Cmp(decimal.NewFromInt(0)) == 0 {
				continue
			}

			profitAmount := profitAmountD.String()

			info1, err := tx.
				Profit.
				Query().
				Where(
					profit.ID(uuid.MustParse(val.ProfitID)),
				).
				ForUpdate().
				Only(ctx)
			if err != nil {
				return err
			}

			c3, err := profitcrud.UpdateSet(info1, &profitmgrpb.ProfitReq{
				AppID:      val.Detail.AppID,
				UserID:     val.Detail.UserID,
				CoinTypeID: val.Detail.CoinTypeID,
				Incoming:   &profitAmount,
			})
			if err != nil {
				return err
			}

			_, err = c3.Save(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// nolint
func UnlockBalance(
	ctx context.Context,
	appID, userID, coinTypeID string,
	ioSubType detailmgrpb.IOSubType,
	unlocked, outcoming decimal.Decimal,
	ioExtra string,
) error {
	var generalID string
	var err error

	if generalID, err = TryCreateGeneral(
		ctx, appID, userID, coinTypeID,
	); err != nil {
		return err
	}

	if unlocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid unlocked")
	}
	if outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid outcoming")
	}
	if unlocked.Cmp(decimal.NewFromInt(0)) == 0 && outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
		return fmt.Errorf("nothing todo")
	}

	spendable := unlocked.Sub(outcoming)

	unlockedS := fmt.Sprintf("-%v", unlocked)
	outcomingS := outcoming.String()
	spendableS := spendable.String()

	ioType := detailmgrpb.IOType_Outcoming

	detailReq := &detailmgrpb.DetailReq{
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		IOType:     &ioType,
		IOSubType:  &ioSubType,
		Amount:     &outcomingS,
		IOExtra:    &ioExtra,
	}

	key := detailKey(detailReq)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	exist, err := detailcli.ExistDetailConds(ctx, &detailmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: detailReq.GetAppID(),
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: detailReq.GetUserID(),
		},
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: detailReq.GetCoinTypeID(),
		},
		IOType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(detailReq.GetIOType()),
		},
		IOSubType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(detailReq.GetIOSubType()),
		},
		IOExtra: &commonpb.StringVal{
			Op:    cruder.LIKE,
			Value: detailReq.GetIOExtra(),
		},
	})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("already exist")
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			General.
			Query().
			Where(
				general.ID(uuid.MustParse(generalID)),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		c1, err := generalcrud.UpdateSet(info, &generalmgrpb.GeneralReq{
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			Outcoming:  &outcomingS,
			Spendable:  &spendableS,
			Locked:     &unlockedS,
		})
		if err != nil {
			return err
		}

		_, err = c1.Save(ctx)
		if err != nil {
			return err
		}

		if outcoming.Cmp(decimal.NewFromInt(0)) == 0 {
			return nil
		}

		c2, err := detailcrud.CreateSet(tx.Detail.Create(), &detailmgrpb.DetailReq{
			AppID:      &appID,
			UserID:     &userID,
			CoinTypeID: &coinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			Amount:     &outcomingS,
			IOExtra:    &ioExtra,
		})
		if err != nil {
			return err
		}

		_, err = c2.Save(ctx)
		return err
	})
}

func LockBalance(
	ctx context.Context,
	appID, userID, coinTypeID string,
	amount decimal.Decimal,
) error {
	var generalID string
	var err error

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("invalid amount")
	}

	if generalID, err = TryCreateGeneral(
		ctx, appID, userID, coinTypeID,
	); err != nil {
		return err
	}

	locked := amount.String()
	spendable := fmt.Sprintf("-%v", amount)

	_, err = generalcli.AddGeneral(ctx, &generalmgrpb.GeneralReq{
		ID:         &generalID,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		Locked:     &locked,
		Spendable:  &spendable,
	})
	return err
}
