package unsold

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"
)

func Ent2Grpc(row *ent.MiningUnsold) *npool.Unsold {
	if row == nil {
		return nil
	}

	return &npool.Unsold{
		ID:          row.ID.String(),
		GoodID:      row.GoodID.String(),
		CoinTypeID:  row.CoinTypeID.String(),
		Amount:      row.Amount.String(),
		BenefitDate: row.BenefitDate,
		CreatedAt:   row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.MiningUnsold) []*npool.Unsold {
	infos := []*npool.Unsold{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
