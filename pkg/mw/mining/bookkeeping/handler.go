package bookkeeping

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	GoodID                    *uuid.UUID
	CoinTypeID                *uuid.UUID
	TotalAmount               *decimal.Decimal
	UnsoldAmount              *decimal.Decimal
	TechniqueServiceFeeAmount *decimal.Decimal
	BenefitDate               *uint32
	Offset                    int32
	Limit                     int32
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

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid good id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid coin type id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

//nolint
func WithTotalAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid total amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("total amount is less than equal 0 %v", *amount)
		}
		h.TotalAmount = &_amount
		return nil
	}
}

//nolint
func WithUnsoldAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid unsold amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("unsold amount is less than 0 %v", *amount)
		}
		h.UnsoldAmount = &_amount
		return nil
	}
}

//nolint
func WithTechniqueServiceFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid technique service fee amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("technique service fee amount is less than 0 %v", *amount)
		}
		h.TechniqueServiceFeeAmount = &_amount
		return nil
	}
}

func WithBenefitDate(date *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if date == nil {
			if must {
				return fmt.Errorf("invalid benefit date is must")
			}
			return nil
		}
		h.BenefitDate = date
		return nil
	}
}
