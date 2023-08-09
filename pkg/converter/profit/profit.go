package profit

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"
)

func Ent2Grpc(row *ent.Profit) *npool.Profit {
	if row == nil {
		return nil
	}

	return &npool.Profit{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Incoming:   row.Incoming.String(),
	}
}

func Ent2GrpcMany(rows []*ent.Profit) []*npool.Profit {
	infos := []*npool.Profit{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
