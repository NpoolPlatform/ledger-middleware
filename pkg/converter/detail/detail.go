package detail

import (
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/detail"
	"github.com/shopspring/decimal"
)

func Ent2Grpc(row *ent.Detail) *npool.Detail {
	if row == nil {
		return nil
	}

	info := &npool.Detail{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		IOType:          npool.IOType(npool.IOType_value[row.IoType]),
		IOSubType:       npool.IOSubType(npool.IOSubType_value[row.IoSubType]),
		Amount:          row.Amount.String(),
		FromCoinTypeID:  row.FromCoinTypeID.String(),
		CoinUSDCurrency: decimal.NewFromInt(0).String(),
		IOExtra:         row.IoExtra,
		CreatedAt:       row.CreatedAt,
	}

	if row.CoinUsdCurrency != nil {
		info.CoinUSDCurrency = row.CoinUsdCurrency.String()
	}

	return info
}

func Ent2GrpcMany(rows []*ent.Detail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
