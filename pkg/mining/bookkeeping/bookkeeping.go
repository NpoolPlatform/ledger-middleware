package bookkeeping

import (
	"context"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"

	detailmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/detail"
	generalmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/general"
	unsoldmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"

	detailcrud "github.com/NpoolPlatform/ledger-manager/pkg/crud/mining/detail"
	generalcrud "github.com/NpoolPlatform/ledger-manager/pkg/crud/mining/general"
	unsoldcrud "github.com/NpoolPlatform/ledger-manager/pkg/crud/mining/unsold"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/shopspring/decimal"
)

func BookKeeping(
	ctx context.Context,
	goodID, coinTypeID string,
	total, unsold, techniqueServiceFee decimal.Decimal,
	benefitDate uint32,
) error {
	totalS := total.String()
	unsoldS := unsold.String()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		stm1, err := detailcrud.CreateSet(
			tx.MiningDetail.Create(),
			&detailmgrpb.DetailReq{
				GoodID:      &goodID,
				CoinTypeID:  &coinTypeID,
				Amount:      &totalS,
				BenefitDate: &benefitDate,
			})
		if err != nil {
			return err
		}

		_, err = stm1.Save(_ctx)
		if err != nil {
			return err
		}

		stm2, err := generalcrud.SetQueryConds(&generalmgrpb.Conds{
			GoodID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: goodID,
			},
			CoinTypeID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: coinTypeID,
			},
		}, tx.MiningGeneral.Query())
		if err != nil {
			return err
		}

		g, err := stm2.ForUpdate().Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if g == nil {
			c, err := generalcrud.CreateSet(
				tx.MiningGeneral.Create(),
				&generalmgrpb.GeneralReq{
					GoodID:     &goodID,
					CoinTypeID: &coinTypeID,
				})
			if err != nil {
				return err
			}

			g, err = c.Save(_ctx)
			if err != nil {
				return err
			}
		}

		toPlatform := unsold.Add(techniqueServiceFee)
		toUserS := total.Sub(toPlatform).String()
		toPlatformS := toPlatform.String()

		stm3, err := generalcrud.AddFieldsSet(g, &generalmgrpb.GeneralReq{
			Amount:     &totalS,
			ToPlatform: &toPlatformS,
			ToUser:     &toUserS,
		})
		if err != nil {
			return err
		}

		_, err = stm3.Save(_ctx)
		if err != nil {
			return err
		}

		stm4, err := unsoldcrud.CreateSet(
			tx.MiningUnsold.Create(),
			&unsoldmgrpb.UnsoldReq{
				GoodID:      &goodID,
				CoinTypeID:  &coinTypeID,
				Amount:      &unsoldS,
				BenefitDate: &benefitDate,
			})
		if err != nil {
			return err
		}

		_, err = stm4.Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
