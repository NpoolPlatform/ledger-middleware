package general

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"

	converter "github.com/NpoolPlatform/ledger-middleware/pkg/converter/mining/general"
	entmgeneral "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/mininggeneral"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func CreateGeneral(ctx context.Context, in *mgrpb.GeneralReq) (*mgrpb.General, error) {
	var info *mgrpb.General
	var info1 *ent.MiningGeneral

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		_, err := tx.
			MiningGeneral.
			Query().
			Where(
				entmgeneral.GoodID(uuid.MustParse(in.GetGoodID())),
				entmgeneral.CoinTypeID(uuid.MustParse(in.GetCoinTypeID())),
			).
			Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}

		info1, err = tx.
			MiningGeneral.
			Create().
			SetGoodID(uuid.MustParse(in.GetGoodID())).
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetAmount(decimal.NewFromInt(0)).
			SetToPlatform(decimal.NewFromInt(0)).
			SetToUser(decimal.NewFromInt(0)).
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
