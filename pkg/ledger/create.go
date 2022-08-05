package ledger

import (
	"context"
	"fmt"

	detailmgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/detail"

	"github.com/shopspring/decimal"
)

func BookKeeping(ctx context.Context, in *detailmgrpb.DetailReq) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func UnlockBalance(
	ctx context.Context,
	appID, userID, coinTypeID string,
	ioType detailmgrpb.IOType,
	unlocked, outcoming decimal.Decimal,
	ioExtra string,
) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
