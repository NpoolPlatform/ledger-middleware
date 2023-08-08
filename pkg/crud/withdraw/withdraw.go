package withdraw

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer"
	tracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer/withdraw"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/withdraw"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/withdraw"

	"github.com/google/uuid"
)

func CreateSet(c *ent.WithdrawCreate, in *npool.WithdrawReq) *ent.WithdrawCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.CoinTypeID != nil {
		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
	}
	if in.AccountID != nil {
		c.SetAccountID(uuid.MustParse(in.GetAccountID()))
	}
	if in.Address != nil {
		c.SetAddress(in.GetAddress())
	}
	if in.PlatformTransactionID != nil {
		c.SetPlatformTransactionID(uuid.MustParse(in.GetPlatformTransactionID()))
	}

	c.SetAmount(decimal.RequireFromString(in.GetAmount()))
	c.SetState(npool.WithdrawState_Reviewing.String())

	return c
}

func Create(ctx context.Context, in *npool.WithdrawReq) (*ent.Withdraw, error) {
	var info *ent.Withdraw
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = CreateSet(cli.Withdraw.Create(), in).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.WithdrawReq) ([]*ent.Withdraw, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Withdraw{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.WithdrawCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Withdraw.Create(), info)
		}
		rows, err = tx.Withdraw.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.Withdraw, in *npool.WithdrawReq) (*ent.WithdrawUpdateOne, error) {
	stm := info.Update()

	if in.PlatformTransactionID != nil {
		stm.SetPlatformTransactionID(uuid.MustParse(in.GetPlatformTransactionID()))
	}
	if in.ChainTransactionID != nil {
		stm.SetChainTransactionID(in.GetChainTransactionID())
	}
	if in.State != nil {
		stm.SetState(in.GetState().String())
	}

	return stm, nil
}

func Update(ctx context.Context, in *npool.WithdrawReq) (*ent.Withdraw, error) {
	var info *ent.Withdraw
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Withdraw.Query().Where(withdraw.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query withdraw: %v", err)
		}

		stm, err := UpdateSet(info, in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update withdraw: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update withdraw: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Withdraw, error) {
	var info *ent.Withdraw
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Withdraw.Query().Where(withdraw.ID(id)).Only(_ctx)
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.WithdrawQuery, error) { //nolint
	stm := cli.Withdraw.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(withdraw.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(withdraw.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(withdraw.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(withdraw.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.AccountID != nil {
		switch conds.GetAccountID().GetOp() {
		case cruder.EQ:
			stm.Where(withdraw.AccountID(uuid.MustParse(conds.GetAccountID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(withdraw.State(npool.WithdrawState(conds.GetState().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.Amount != nil {
		incoming, err := decimal.NewFromString(conds.GetAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAmount().GetOp() {
		case cruder.LT:
			stm.Where(withdraw.AmountLT(incoming))
		case cruder.GT:
			stm.Where(withdraw.AmountGT(incoming))
		case cruder.EQ:
			stm.Where(withdraw.AmountEQ(incoming))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	if conds.CreatedAt != nil {
		switch conds.GetCreatedAt().GetOp() {
		case cruder.LT:
			stm.Where(withdraw.CreatedAtLT(conds.GetCreatedAt().GetValue()))
		case cruder.GT:
			stm.Where(withdraw.CreatedAtGT(conds.GetCreatedAt().GetValue()))
		case cruder.EQ:
			stm.Where(withdraw.CreatedAtEQ(conds.GetCreatedAt().GetValue()))
		default:
			return nil, fmt.Errorf("invalid withdraw field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Withdraw, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Withdraw{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Limit(limit).
			Order(ent.Desc(withdraw.FieldUpdatedAt)).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Withdraw, error) {
	var info *ent.Withdraw
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Withdraw.Query().Where(withdraw.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Withdraw, error) {
	var info *ent.Withdraw
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Withdraw.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
