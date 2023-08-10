package detail

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
// 	commontracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer"
// 	tracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer/mining/detail"
// 	"github.com/shopspring/decimal"
// 	"go.opentelemetry.io/otel"
// 	"go.opentelemetry.io/otel/codes"

// 	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
// 	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
// 	detail "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/miningdetail"
// 	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
// 	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

// 	"github.com/google/uuid"
// )

// func CreateSet(c *ent.MiningDetailCreate, in *npool.DetailReq) (*ent.MiningDetailCreate, error) {
// 	if in.ID != nil {
// 		c.SetID(uuid.MustParse(in.GetID()))
// 	}
// 	if in.GoodID != nil {
// 		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
// 	}
// 	if in.CoinTypeID != nil {
// 		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
// 	}
// 	if in.Amount != nil {
// 		amount, err := decimal.NewFromString(in.GetAmount())
// 		if err != nil {
// 			return nil, err
// 		}
// 		c.SetAmount(amount)
// 	}
// 	if in.BenefitDate != nil {
// 		c.SetBenefitDate(in.GetBenefitDate())
// 	}
// 	if in.CreatedAt != nil {
// 		c.SetCreatedAt(in.GetCreatedAt())
// 	}
// 	return c, nil
// }

// func Create(ctx context.Context, in *npool.DetailReq) (*ent.MiningDetail, error) {
// 	var info *ent.MiningDetail
// 	var err error

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = tracer.Trace(span, in)

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		c, err := CreateSet(cli.MiningDetail.Create(), in)
// 		if err != nil {
// 			return err
// 		}

// 		info, err = c.Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return info, nil
// }

// func CreateBulk(ctx context.Context, in []*npool.DetailReq) ([]*ent.MiningDetail, error) {
// 	var err error

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = tracer.TraceMany(span, in)

// 	rows := []*ent.MiningDetail{}
// 	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
// 		bulk := make([]*ent.MiningDetailCreate, len(in))
// 		for i, info := range in {
// 			bulk[i], err = CreateSet(tx.MiningDetail.Create(), info)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		rows, err = tx.MiningDetail.CreateBulk(bulk...).Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }

// func Row(ctx context.Context, id uuid.UUID) (*ent.MiningDetail, error) {
// 	var info *ent.MiningDetail
// 	var err error

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = commontracer.TraceID(span, id.String())

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		info, err = cli.MiningDetail.Query().Where(detail.ID(id)).Only(_ctx)
// 		if ent.IsNotFound(err) {
// 			return nil
// 		}
// 		return err
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return info, nil
// }

// func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.MiningDetailQuery, error) { //nolint
// 	stm := cli.MiningDetail.Query()
// 	if conds.ID != nil {
// 		switch conds.GetID().GetOp() {
// 		case cruder.EQ:
// 			stm.Where(detail.ID(uuid.MustParse(conds.GetID().GetValue())))
// 		default:
// 			return nil, fmt.Errorf("invalid detail field")
// 		}
// 	}
// 	if conds.GoodID != nil {
// 		switch conds.GetGoodID().GetOp() {
// 		case cruder.EQ:
// 			stm.Where(detail.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
// 		default:
// 			return nil, fmt.Errorf("invalid detail field")
// 		}
// 	}
// 	if conds.CoinTypeID != nil {
// 		switch conds.GetCoinTypeID().GetOp() {
// 		case cruder.EQ:
// 			stm.Where(detail.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
// 		default:
// 			return nil, fmt.Errorf("invalid detail field")
// 		}
// 	}
// 	if conds.Amount != nil {
// 		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
// 		if err != nil {
// 			return nil, err
// 		}
// 		switch conds.GetAmount().GetOp() {
// 		case cruder.LT:
// 			stm.Where(detail.AmountLT(amount))
// 		case cruder.GT:
// 			stm.Where(detail.AmountGT(amount))
// 		case cruder.EQ:
// 			stm.Where(detail.AmountEQ(amount))
// 		default:
// 			return nil, fmt.Errorf("invalid detail field")
// 		}
// 	}
// 	if conds.BenefitDate != nil {
// 		switch conds.GetBenefitDate().GetOp() {
// 		case cruder.LT:
// 			stm.Where(detail.BenefitDateLT(conds.GetBenefitDate().GetValue()))
// 		case cruder.GT:
// 			stm.Where(detail.BenefitDateGT(conds.GetBenefitDate().GetValue()))
// 		case cruder.EQ:
// 			stm.Where(detail.BenefitDateEQ(conds.GetBenefitDate().GetValue()))
// 		default:
// 			return nil, fmt.Errorf("invalid detail field")
// 		}
// 	}
// 	return stm, nil
// }

// func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.MiningDetail, int, error) {
// 	var err error

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = tracer.TraceConds(span, conds)
// 	span = commontracer.TraceOffsetLimit(span, offset, limit)

// 	rows := []*ent.MiningDetail{}
// 	var total int
// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		stm, err := setQueryConds(conds, cli)
// 		if err != nil {
// 			return err
// 		}

// 		total, err = stm.Count(_ctx)
// 		if err != nil {
// 			return err
// 		}

// 		rows, err = stm.
// 			Offset(offset).
// 			Order(ent.Desc(detail.FieldUpdatedAt)).
// 			Limit(limit).
// 			All(_ctx)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	return rows, total, nil
// }

// func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.MiningDetail, error) {
// 	var info *ent.MiningDetail
// 	var err error

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = tracer.TraceConds(span, conds)

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		stm, err := setQueryConds(conds, cli)
// 		if err != nil {
// 			return err
// 		}

// 		info, err = stm.Only(_ctx)
// 		if err != nil {
// 			if ent.IsNotFound(err) {
// 				return nil
// 			}
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return info, nil
// }

// func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
// 	var err error
// 	var total int

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = tracer.TraceConds(span, conds)

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		stm, err := setQueryConds(conds, cli)
// 		if err != nil {
// 			return err
// 		}

// 		total, err = stm.Count(_ctx)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return 0, err
// 	}

// 	return uint32(total), nil
// }

// func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
// 	var err error
// 	exist := false

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = commontracer.TraceID(span, id.String())

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		exist, err = cli.MiningDetail.Query().Where(detail.ID(id)).Exist(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return false, err
// 	}

// 	return exist, nil
// }

// func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
// 	var err error
// 	exist := false

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = tracer.TraceConds(span, conds)

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		stm, err := setQueryConds(conds, cli)
// 		if err != nil {
// 			return err
// 		}

// 		exist, err = stm.Exist(_ctx)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return false, err
// 	}

// 	return exist, nil
// }

// func Delete(ctx context.Context, id uuid.UUID) (*ent.MiningDetail, error) {
// 	var info *ent.MiningDetail
// 	var err error

// 	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
// 	defer span.End()

// 	defer func() {
// 		if err != nil {
// 			span.SetStatus(codes.Error, "db operation fail")
// 			span.RecordError(err)
// 		}
// 	}()

// 	span = commontracer.TraceID(span, id.String())

// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		info, err = cli.MiningDetail.UpdateOneID(id).
// 			SetDeletedAt(uint32(time.Now().Unix())).
// 			Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return info, nil
// }
