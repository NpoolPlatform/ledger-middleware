package statement

import (
	"context"
	"encoding/json"
	"fmt"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/const"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	Reqs   []*crud.Req
	Conds  *crud.Conds
	Offset int32
	Limit  int32
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

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
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
			return nil
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

func WithIOType(_type *basetypes.IOType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		flag := false
		for ioType := range basetypes.IOType_value {
			if ioType == _type.String() && ioType != basetypes.IOType_DefaultType.String() {
				flag = true
			}
		}
		if !flag {
			return fmt.Errorf("invalid io type %v", *_type)
		}
		h.IOType = _type
		return nil
	}
}

func WithIOSubType(_type *basetypes.IOSubType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		flag := false
		for ioSubType := range basetypes.IOSubType_value {
			if ioSubType == _type.String() && ioSubType != basetypes.IOSubType_DefaultSubType.String() {
				flag = true
			}
		}
		if !flag {
			return fmt.Errorf("invalid io sub type %v", *_type)
		}
		h.IOSubType = _type
		return nil
	}
}

func WithAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
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

func WithFromCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FromCoinTypeID = &_id
		return nil
	}
}

func WithCoinUSDCurrency(currency *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if currency == nil {
			return nil
		}
		_currency, err := decimal.NewFromString(*currency)
		if err != nil {
			return err
		}
		h.CoinUSDCurrency = &_currency
		return nil
	}
}

func WithIOExtra(extra *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if extra == nil {
			return nil
		}
		if !json.Valid([]byte(*extra)) {
			return fmt.Errorf("io extra is invalid json str %v", *extra)
		}

		h.IOExtra = extra
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
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
		if conds.FromCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetFromCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.FromCoinTypeID = &cruder.Cond{
				Op:  conds.GetFromCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.IOType != nil {
			h.Conds.IOType = &cruder.Cond{
				Op:  conds.GetIOType().GetOp(),
				Val: conds.GetIOType().GetValue(),
			}
		}
		if conds.IOSubType != nil {
			h.Conds.IOSubType = &cruder.Cond{
				Op:  conds.GetIOSubType().GetOp(),
				Val: conds.GetIOSubType().GetValue(),
			}
		}
		if conds.IOExtra != nil {
			h.Conds.IOExtra = &cruder.Cond{
				Op:  conds.GetIOExtra().GetOp(),
				Val: conds.GetIOExtra().GetValue(),
			}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.StatementReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range reqs {
			_req := &crud.Req{}
			if req.AppID != nil {
				_id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &_id
			}
			if req.UserID != nil {
				_id, err := uuid.Parse(*req.UserID)
				if err != nil {
					return err
				}
				_req.UserID = &_id
			}
			if req.CoinTypeID != nil {
				_id, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &_id
			}
			if req.FromCoinTypeID != nil {
				_id, err := uuid.Parse(*req.FromCoinTypeID)
				if err != nil {
					return err
				}
				_req.FromCoinTypeID = &_id
			}
			if req.Amount != nil {
				amount, err := decimal.NewFromString(*req.Amount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("amount is less than 0 %v", *req.Amount)
				}
				_req.Amount = &amount
			}
			if req.CoinUSDCurrency != nil {
				currency, err := decimal.NewFromString(*req.CoinUSDCurrency)
				if err != nil {
					return err
				}
				if currency.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("coin usd currency is less than 0 %v", *req.CoinUSDCurrency)
				}
				_req.CoinUSDCurrency = &currency
			}

			if req.IOType == nil || req.IOSubType == nil {
				return fmt.Errorf("invalid io type or io subtype")
			}
			switch *req.IOType {
			case basetypes.IOType_Incoming:
				switch *req.IOSubType {
				case basetypes.IOSubType_Payment:
				case basetypes.IOSubType_MiningBenefit:
				case basetypes.IOSubType_Commission:
				case basetypes.IOSubType_TechniqueFeeCommission:
				case basetypes.IOSubType_Deposit:
				case basetypes.IOSubType_Transfer:
				case basetypes.IOSubType_OrderRevoke:
				default:
					return fmt.Errorf("io subtype not match io type, io subtype: %v, io type: %v", *req.IOSubType, *req.IOType)
				}
			case basetypes.IOType_Outcoming:
				switch *req.IOSubType {
				case basetypes.IOSubType_Payment:
				case basetypes.IOSubType_Withdrawal:
				case basetypes.IOSubType_Transfer:
				case basetypes.IOSubType_CommissionRevoke:
				default:
					return fmt.Errorf("io subtype not match io type, io subtype: %v, io type: %v", *req.IOSubType, *req.IOType)
				}
			default:
				return fmt.Errorf("invalid io type %v", *req.IOType)
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
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
