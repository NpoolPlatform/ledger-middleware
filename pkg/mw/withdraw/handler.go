package withdraw

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/const"
	crud "github.com/NpoolPlatform/ledger-middleware/pkg/crud/withdraw"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	FeeAmount *decimal.Decimal
	Rollback  *bool
	Conds     *crud.Conds
	Offset    int32
	Limit     int32
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid app id")
			}
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

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid user id")
			}
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

func WithAccountID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid account id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AccountID = &_id
		return nil
	}
}

func WithPlatformTransactionID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid platform transaction id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.PlatformTransactionID = &_id
		return nil
	}
}

func WithChainTransactionID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid chain transaction id")
			}
			return nil
		}
		h.ChainTransactionID = id
		return nil
	}
}

func WithAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
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

func WithAddress(address *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if address == nil {
			if must {
				return fmt.Errorf("invalid address")
			}
			return nil
		}
		h.Address = address
		return nil
	}
}

//nolint
func WithState(state *basetypes.WithdrawState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		switch *state {
		case basetypes.WithdrawState_Created:
		case basetypes.WithdrawState_Reviewing:
		case basetypes.WithdrawState_Approved:
		case basetypes.WithdrawState_Transferring:
		case basetypes.WithdrawState_PreRejected:
		case basetypes.WithdrawState_ReturnRejectedBalance:
		case basetypes.WithdrawState_Rejected:
		case basetypes.WithdrawState_PreFail:
		case basetypes.WithdrawState_ReturnFailBalance:
		case basetypes.WithdrawState_TransactionFail:
		case basetypes.WithdrawState_PreSuccessful:
		case basetypes.WithdrawState_SpendSuccessfulBalance:
		case basetypes.WithdrawState_Successful:
		default:
			return fmt.Errorf("invalid state")
		}
		h.State = state
		return nil
	}
}

func WithFeeAmount(feeAmount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if feeAmount == nil {
			if must {
				return fmt.Errorf("invalid feeamount")
			}
			return nil
		}
		_feeAmount, err := decimal.NewFromString(*feeAmount)
		if err != nil {
			return err
		}
		if _feeAmount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid feeamount")
		}
		h.FeeAmount = &_feeAmount
		return nil
	}
}

func WithReviewID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid review id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ReviewID = &_id
		return nil
	}
}

func WithRollback(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Rollback = b
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
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
		if conds.State != nil {
			state := conds.GetState().GetValue()
			h.Conds.State = &cruder.Cond{
				Op:  conds.GetState().GetOp(),
				Val: basetypes.WithdrawState(state),
			}
		}
		if conds.CreatedAt != nil {
			h.Conds.CreatedAt = &cruder.Cond{
				Op:  conds.GetCreatedAt().GetOp(),
				Val: conds.GetCreatedAt().GetValue(),
			}
		}
		if conds.AccountID != nil {
			id, err := uuid.Parse(conds.GetAccountID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AccountID = &cruder.Cond{
				Op:  conds.GetAccountID().GetOp(),
				Val: id,
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
