package ledger

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/ledger-manager/pkg/db"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"

	detailpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/detail"
	generalpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/general"

	"github.com/google/uuid"
)

func GetIntervalGenerals(
	ctx context.Context, appID, userID string, start, end uint32, offset, limit int32,
) (
	infos []*generalpb.General, total uint32, err error,
) {
	details := []*ent.Detail{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Detail.
			Query().
			Where(
				detail.AppID(uuid.MustParse(appID)),
				detail.UserID(uuid.MustParse(userID)),
				detail.CreatedAtGT(start),
				detail.CreatedAtLT(end),
			)
		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		details, err = stm.
			Offset(int(offset)).
			Limit(int(limit)).
			All(ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}

	gMap := map[uuid.UUID]*generalpb.General{}
	for _, detail := range details {
		incoming := decimal.NewFromInt(0)
		outcoming := decimal.NewFromInt(0)

		switch detail.IoType {
		case detailpb.IOType_Incoming.String():
			incoming = incoming.Add(detail.Amount)
		case detailpb.IOType_Outcoming.String():
			outcoming = incoming.Add(detail.Amount)
		}

		general, ok := gMap[detail.CoinTypeID]
		if !ok {
			general = &generalpb.General{
				AppID:      appID,
				UserID:     userID,
				CoinTypeID: detail.CoinTypeID.String(),
				Incoming:   decimal.NewFromInt(0).String(),
				Outcoming:  decimal.NewFromInt(0).String(),
			}
		}

		general.Incoming = incoming.Add(decimal.RequireFromString(general.Incoming)).String()
		general.Outcoming = outcoming.Add(decimal.RequireFromString(general.Outcoming)).String()

		gMap[detail.CoinTypeID] = general
	}

	for _, general := range gMap {
		infos = append(infos, general)
	}

	return infos, total, nil
}
