package general

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"
)

func Ent2Grpc(row *ent.MiningGeneral) *npool.General {
	if row == nil {
		return nil
	}

	return &npool.General{
		ID:         row.ID.String(),
		GoodID:     row.GoodID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Amount:     row.Amount.String(),
		ToPlatform: row.ToPlatform.String(),
		ToUser:     row.ToUser.String(),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.MiningGeneral) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
