package statement

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	GoodID                    *uuid.UUID
	CoinTypeID                *uuid.UUID
	TotalAmount               *decimal.Decimal
	UnsoldAmount              *decimal.Decimal
	TechniqueServiceFeeAmount *decimal.Decimal
	BenefitDate               *uint32
}

type Handler struct {
	*Req
	Reqs []*Req
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

//nolint
func WithReqs(reqs []*npool.GoodStatementsReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*Req{}
		for _, req := range reqs {
			_req := &Req{}
			if req.GoodID == nil {
				return fmt.Errorf("invalid good id ")
			}
			if req.CoinTypeID == nil {
				return fmt.Errorf("invalid coin type id ")
			}
			if req.TotalAmount == nil {
				return fmt.Errorf("invalid total amount")
			}
			if req.UnsoldAmount == nil {
				return fmt.Errorf("invalid unsold amount")
			}
			if req.TechniqueServiceFeeAmount == nil {
				return fmt.Errorf("invalid technique service fee amount")
			}
			if req.BenefitDate == nil {
				return fmt.Errorf("invalid benefit date")
			}
			if req.GoodID != nil {
				_id, err := uuid.Parse(*req.GoodID)
				if err != nil {
					return err
				}
				_req.GoodID = &_id
			}
			if req.CoinTypeID != nil {
				_id, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &_id
			}
			if req.TotalAmount != nil {
				amount, err := decimal.NewFromString(*req.TotalAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
					return fmt.Errorf("total amount is less than equal 0 %v", *req.TotalAmount)
				}
				_req.TotalAmount = &amount
			}
			if req.UnsoldAmount != nil {
				amount, err := decimal.NewFromString(*req.TotalAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("unsold amount is less than 0 %v", *req.UnsoldAmount)
				}
				_req.UnsoldAmount = &amount
			}
			if req.TechniqueServiceFeeAmount != nil {
				amount, err := decimal.NewFromString(*req.TechniqueServiceFeeAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("technique service fee amount is less than 0 %v", *req.TechniqueServiceFeeAmount)
				}
				_req.TechniqueServiceFeeAmount = &amount
			}

			if req.BenefitDate != nil {
				_req.BenefitDate = req.BenefitDate
			}

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
