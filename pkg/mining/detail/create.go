package detail

import (
	"context"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/detail"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/mining/detail"
	entmdetail "github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningdetail"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateDetail(ctx context.Context, in *npool.DetailReq) (*mgrpb.Detail, error) {
	now := uint32(time.Now().Unix())
	seconds := in.GetBenefitIntervalHours() * timedef.SecondsPerHour
	timestamp := now / seconds * seconds

	var info *mgrpb.Detail
	var info1 *ent.MiningDetail
	var err error

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info1, err = tx.
			MiningDetail.
			Query().
			Where(
				entmdetail.GoodID(uuid.MustParse(in.GetGoodID())),
				entmdetail.CoinTypeID(uuid.MustParse(in.GetCoinTypeID())),
				entmdetail.BenefitDate(timestamp),
			).
			Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info1 != nil {
			return nil
		}

		info1, err = tx.
			MiningDetail.
			Create().
			SetGoodID(uuid.MustParse(in.GetGoodID())).
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetAmount(decimal.RequireFromString(in.GetAmount())).
			SetBenefitDate(timestamp).
			Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	info = converter.Ent2Grpc(info1)

	return info, nil
}
