package statement

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/const"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/good/ledger/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	Rollback *bool
	Reqs     []*crud.Req
	Conds    *crud.Conds
	Limit    int32
	Offset   int32
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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
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
				return fmt.Errorf("invalid benefit date")
			}
			return nil
		}
		h.BenefitDate = date
		return nil
	}
}

func WithRollback(rollback *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if rollback == nil {
			if must {
				return fmt.Errorf("invalid rollback")
			}
			return nil
		}
		h.Rollback = rollback
		return nil
	}
}

//nolint
func WithReqs(reqs []*npool.GoodStatementReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range reqs {
			if must {
				if req.GoodID == nil {
					return fmt.Errorf("invalid good id")
				}
				if req.CoinTypeID == nil {
					return fmt.Errorf("invalid coin type id")
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
			}
			_req := &crud.Req{}
			if req.ID != nil {
				_id, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
				_req.ID = &_id
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
				amount, err := decimal.NewFromString(*req.UnsoldAmount)
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
				if *req.BenefitDate == 0 {
					return fmt.Errorf("invalid benefit date 0")
				}
				_req.BenefitDate = req.BenefitDate
			}
			if req.Rollback != nil {
				h.Rollback = req.Rollback
			}

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.Amount != nil {
			amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
			if err != nil {
				return err
			}
			h.Conds.Amount = &cruder.Cond{
				Op:  conds.GetAmount().GetOp(),
				Val: amount,
			}
		}
		if conds.BenefitDate != nil {
			h.Conds.BenefitDate = &cruder.Cond{
				Op:  conds.GetBenefitDate().GetOp(),
				Val: conds.GetBenefitDate().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
