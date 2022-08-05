//nolint:dupl
package ledger

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/profit"

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

	spendable := decimal.NewFromInt(0).Sub(outcoming)

	unlockedS := fmt.Sprintf("-%v", unlocked)
	outcomingS := outcoming.String()
	spendableS := spendable.String()

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

		ioType := detailmgrpb.IOType_Outcoming

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
