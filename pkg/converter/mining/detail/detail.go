package detail

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"
)

func Ent2Grpc(row *ent.MiningDetail) *npool.Detail {
	if row == nil {
		return nil
	}

	return &npool.Detail{
		ID:          row.ID.String(),
		GoodID:      row.GoodID.String(),
		CoinTypeID:  row.CoinTypeID.String(),
		Amount:      row.Amount.String(),
		BenefitDate: row.BenefitDate,
		CreatedAt:   row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.MiningDetail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
