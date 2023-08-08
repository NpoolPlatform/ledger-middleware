package detail

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer"
	tracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer/detail"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/detail"

	"github.com/google/uuid"
)

func CreateSet(c *ent.DetailCreate, in *npool.DetailReq) (*ent.DetailCreate, error) {
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
	if in.IOType != nil {
		c.SetIoType(in.GetIOType().String())
	}
	if in.IOSubType != nil {
		c.SetIoSubType(in.GetIOSubType().String())
	}
	if in.Amount != nil {
		amount, err := decimal.NewFromString(in.GetAmount())
		if err != nil {
			return nil, err
		}
		c.SetAmount(amount)
	}
	if in.FromCoinTypeID != nil {
		c.SetFromCoinTypeID(uuid.MustParse(in.GetFromCoinTypeID()))
	}
	if in.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(in.GetCoinUSDCurrency())
		if err != nil {
			return nil, err
		}
		c.SetCoinUsdCurrency(currency)
	}
	if in.IOExtra != nil {
		c.SetIoExtra(in.GetIOExtra())
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}

	return c, nil
}

func Create(ctx context.Context, in *npool.DetailReq) (*ent.Detail, error) {
	var info *ent.Detail
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
		c, err := CreateSet(cli.Detail.Create(), in)
		if err != nil {
			return err
		}

		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.DetailReq) ([]*ent.Detail, error) {
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

	rows := []*ent.Detail{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.DetailCreate, len(in))
		for i, info := range in {
			bulk[i], err = CreateSet(tx.Detail.Create(), info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Detail.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Detail, error) {
	var info *ent.Detail
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
		info, err = cli.Detail.Query().Where(detail.ID(id)).Only(_ctx)
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

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.DetailQuery, error) { //nolint
	stm := cli.Detail.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.IOType != nil {
		switch conds.GetIOType().GetOp() {
		case cruder.EQ:
			stm.Where(detail.IoType(npool.IOType(conds.GetIOType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.IOSubType != nil {
		switch conds.GetIOSubType().GetOp() {
		case cruder.EQ:
			stm.Where(detail.IoSubType(npool.IOSubType(conds.GetIOSubType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Amount != nil {
		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAmount().GetOp() {
		case cruder.LT:
			stm.Where(detail.AmountLT(amount))
		case cruder.GT:
			stm.Where(detail.AmountGT(amount))
		case cruder.EQ:
			stm.Where(detail.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.FromCoinTypeID != nil {
		switch conds.GetFromCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.FromCoinTypeID(uuid.MustParse(conds.GetFromCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(conds.GetCoinUSDCurrency().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetCoinUSDCurrency().GetOp() {
		case cruder.LT:
			stm.Where(detail.CoinUsdCurrencyLT(currency))
		case cruder.GT:
			stm.Where(detail.CoinUsdCurrencyGT(currency))
		case cruder.EQ:
			stm.Where(detail.CoinUsdCurrencyEQ(currency))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.IOExtra != nil {
		switch conds.GetIOExtra().GetOp() {
		case cruder.LIKE:
			stm.Where(detail.IoExtraContains(conds.GetIOExtra().GetValue()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Detail, int, error) {
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

	rows := []*ent.Detail{}
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
			Order(ent.Desc(detail.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Detail, error) {
	var info *ent.Detail
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
		exist, err = cli.Detail.Query().Where(detail.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Detail, error) {
	var info *ent.Detail
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
		info, err = cli.Detail.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
