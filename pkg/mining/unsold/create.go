package unsold

import (
	"context"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/mining/unsold"
	entmunsold "github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningunsold"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateUnsold(ctx context.Context, in *npool.UnsoldReq) (*mgrpb.Unsold, error) {
	now := uint32(time.Now().Unix())
	seconds := in.GetBenefitIntervalHours() * timedef.SecondsPerHour
	timestamp := now / seconds * seconds

	var info *mgrpb.Unsold
	var info1 *ent.MiningUnsold

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		_, err := tx.
			MiningUnsold.
			Query().
			Where(
				entmunsold.GoodID(uuid.MustParse(in.GetGoodID())),
				entmunsold.CoinTypeID(uuid.MustParse(in.GetCoinTypeID())),
				entmunsold.BenefitDate(timestamp),
			).
			Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}

		info1, err = tx.
			MiningUnsold.
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
