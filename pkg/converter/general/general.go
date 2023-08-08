package general

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/general"
)

func Ent2Grpc(row *ent.General) *npool.General {
	if row == nil {
		return nil
	}

	return &npool.General{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Incoming:   row.Incoming.String(),
		Locked:     row.Locked.String(),
		Outcoming:  row.Outcoming.String(),
		Spendable:  row.Spendable.String(),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.General) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
