package lock

import (
	"context"
	"encoding/json"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Amount     *decimal.Decimal
	Unlocked   *decimal.Decimal
	Outcoming  *decimal.Decimal
	IOSubType  *basetypes.IOSubType
	IOExtra    *string
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return fmt.Errorf("invalid app id")
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return fmt.Errorf("invalid app id")
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return fmt.Errorf("invalid user id")
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return fmt.Errorf("coin type id")
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithIOSubType(_type *basetypes.IOSubType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return fmt.Errorf("invalid io sub type")
		}
		switch *_type {
		case basetypes.IOSubType_Withdrawal:
		case basetypes.IOSubType_Payment:
		default:
			return fmt.Errorf("invalid io sub type")
		}
		h.IOSubType = _type
		return nil
	}
}

func WithAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return fmt.Errorf("invalid amount")
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", *amount)
		}
		h.Amount = &_amount
		return nil
	}
}

func WithUnlocked(unlocked *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unlocked == nil {
			return fmt.Errorf("invalid unlock")
		}
		_unlocked, err := decimal.NewFromString(*unlocked)
		if err != nil {
			return err
		}
		if _unlocked.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", *unlocked)
		}
		h.Unlocked = &_unlocked
		return nil
	}
}

func WithOutcoming(outcoming *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if outcoming == nil {
			return fmt.Errorf("invalid outcoming")
		}
		_outcoming, err := decimal.NewFromString(*outcoming)
		if err != nil {
			return err
		}
		if _outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", *outcoming)
		}
		h.Outcoming = &_outcoming
		return nil
	}
}

func WithIOExtra(extra *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if extra == nil {
			return fmt.Errorf("invalid extra")
		}
		if !json.Valid([]byte(*extra)) {
			return fmt.Errorf("io extra is invalid json str %v", *extra)
		}

		h.IOExtra = extra
		return nil
	}
}
