package statement

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createStatement(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Statement.Create(),
			&h.Req,
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}

//nolint
func (h *Handler) CreateStatement(ctx context.Context) (*npool.Statement, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid app id")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid user id")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid coin type id")
	}
	if h.IOType == nil {
		return nil, fmt.Errorf("invalid io type")
	}
	if h.IOSubType == nil {
		return nil, fmt.Errorf("invalid io sub type")
	}
	switch *h.IOType {
	case basetypes.IOType_Incoming:
		switch *h.IOSubType {
		case basetypes.IOSubType_Payment:
		case basetypes.IOSubType_MiningBenefit:
		case basetypes.IOSubType_Commission:
		case basetypes.IOSubType_TechniqueFeeCommission:
		case basetypes.IOSubType_Deposit:
		case basetypes.IOSubType_Transfer:
		case basetypes.IOSubType_OrderRevoke:
		default:
			return nil, fmt.Errorf("io subtype not match io type, io subtype: %v, io type: %v", *h.IOSubType, *h.IOType)
		}
	case basetypes.IOType_Outcoming:
		switch *h.IOSubType {
		case basetypes.IOSubType_Payment:
		case basetypes.IOSubType_Withdrawal:
		case basetypes.IOSubType_Transfer:
		case basetypes.IOSubType_CommissionRevoke:
		default:
			return nil, fmt.Errorf("io subtype not match io type, io subtype: %v, io type: %v", *h.IOSubType, *h.IOType)
		}
	default:
		return nil, fmt.Errorf("invalid io type %v", *h.IOType)
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.createStatement(ctx); err != nil {
		return nil, err
	}

	return h.GetStatement(ctx)
}
