package withdraw

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
)

func Ent2Grpc(row *ent.Withdraw) *npool.Withdraw {
	if row == nil {
		return nil
	}

	return &npool.Withdraw{
		ID:                    row.ID.String(),
		AppID:                 row.AppID.String(),
		UserID:                row.UserID.String(),
		CoinTypeID:            row.CoinTypeID.String(),
		AccountID:             row.AccountID.String(),
		Address:               row.Address,
		Amount:                row.Amount.String(),
		CreatedAt:             row.CreatedAt,
		PlatformTransactionID: row.PlatformTransactionID.String(),
		ChainTransactionID:    row.ChainTransactionID,
		State:                 npool.WithdrawState(npool.WithdrawState_value[row.State]),
	}
}

func Ent2GrpcMany(rows []*ent.Withdraw) []*npool.Withdraw {
	infos := []*npool.Withdraw{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
