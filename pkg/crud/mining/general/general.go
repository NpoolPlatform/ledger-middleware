package general

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"
// 	commontracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer"
// 	tracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer/mining/general"
// 	"github.com/shopspring/decimal"
// 	"go.opentelemetry.io/otel"
// 	"go.opentelemetry.io/otel/codes"

// 	"github.com/NpoolPlatform/ledger-middleware/pkg/db"
// 	"github.com/NpoolPlatform/ledger-middleware/pkg/db/ent"
// 	general "github.com/NpoolPlatform/ledger-middleware/pkg/db/ent/mininggeneral"
// 	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
// 	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"

// 	"github.com/google/uuid"
// )

// func CreateSet(c *ent.MiningGeneralCreate, in *npool.GeneralReq) (*ent.MiningGeneralCreate, error) {
// 	if in.ID != nil {
// 		c.SetID(uuid.MustParse(in.GetID()))
// 	}
// 	if in.GoodID != nil {
// 		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
// 	}
// 	if in.CoinTypeID != nil {
// 		c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
// 	}

// 	c.SetAmount(decimal.NewFromInt(0))
// 	c.SetToPlatform(decimal.NewFromInt(0))
// 	c.SetToUser(decimal.NewFromInt(0))
// 	return c, nil
// }

// func Create(ctx context.Context, in *npool.GeneralReq) (*ent.MiningGeneral, error) {
// 	var info *ent.MiningGeneral
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
// 		c, err := CreateSet(cli.MiningGeneral.Create(), in)
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

// func CreateBulk(ctx context.Context, in []*npool.GeneralReq) ([]*ent.MiningGeneral, error) {
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

// 	rows := []*ent.MiningGeneral{}
// 	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
// 		bulk := make([]*ent.MiningGeneralCreate, len(in))
// 		for i, info := range in {
// 			bulk[i] = tx.MiningGeneral.Create()
// 			if info.ID != nil {
// 				bulk[i].SetID(uuid.MustParse(info.GetID()))
// 			}
// 			if info.GoodID != nil {
// 				bulk[i].SetGoodID(uuid.MustParse(info.GetGoodID()))
// 			}
// 			if info.CoinTypeID != nil {
// 				bulk[i].SetCoinTypeID(uuid.MustParse(info.GetCoinTypeID()))
// 			}
// 			bulk[i].SetAmount(decimal.NewFromInt(0))
// 			bulk[i].SetToPlatform(decimal.NewFromInt(0))
// 			bulk[i].SetToUser(decimal.NewFromInt(0))
// 		}
// 		rows, err = tx.MiningGeneral.CreateBulk(bulk...).Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }

// func AddFieldsSet(info *ent.MiningGeneral, in *npool.GeneralReq) (*ent.MiningGeneralUpdateOne, error) {
// 	var err error

// 	amount := decimal.NewFromInt(0)
// 	if in.Amount != nil {
// 		amount, err = decimal.NewFromString(in.GetAmount())
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	toPlatform := decimal.NewFromInt(0)
// 	if in.ToPlatform != nil {
// 		toPlatform, err = decimal.NewFromString(in.GetToPlatform())
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	toUser := decimal.NewFromInt(0)
// 	if in.ToUser != nil {
// 		toUser, err = decimal.NewFromString(in.GetToUser())
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if amount.Cmp(toPlatform.Add(toUser)) < 0 {
// 		return nil, fmt.Errorf("amount < toPlatform + toUser")
// 	}
// 	if amount.Cmp(decimal.NewFromInt(0)) < 0 {
// 		return nil, fmt.Errorf("amount < 0")
// 	}
// 	if toPlatform.Cmp(decimal.NewFromInt(0)) < 0 {
// 		return nil, fmt.Errorf("toPlatform < 0")
// 	}
// 	if toUser.Cmp(decimal.NewFromInt(0)) < 0 {
// 		return nil, fmt.Errorf("toUser < 0")
// 	}

// 	stm := info.Update()

// 	if in.Amount != nil {
// 		stm = stm.AddAmount(amount)
// 	}
// 	if in.ToUser != nil {
// 		stm = stm.AddToUser(toUser)
// 	}
// 	if in.ToPlatform != nil {
// 		stm = stm.AddToPlatform(toPlatform)
// 	}

// 	return stm, nil
// }

// func AddFields(ctx context.Context, in *npool.GeneralReq) (*ent.MiningGeneral, error) {
// 	var info *ent.MiningGeneral
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

// 	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
// 		info, err = tx.MiningGeneral.Query().Where(general.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail query general: %v", err)
// 		}

// 		stm, err := AddFieldsSet(info, in)
// 		if err != nil {
// 			return err
// 		}

// 		info, err = stm.Save(_ctx)
// 		if err != nil {
// 			return fmt.Errorf("fail update general: %v", err)
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("fail update general: %v", err)
// 	}

// 	return info, nil
// }

// func Row(ctx context.Context, id uuid.UUID) (*ent.MiningGeneral, error) {
// 	var info *ent.MiningGeneral
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
// 		info, err = cli.MiningGeneral.Query().Where(general.ID(id)).Only(_ctx)
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

// func SetQueryConds(conds *npool.Conds, stm *ent.MiningGeneralQuery) (*ent.MiningGeneralQuery, error) { //nolint
// 	if conds.ID != nil {
// 		switch conds.GetID().GetOp() {
// 		case cruder.EQ:
// 			stm.Where(general.ID(uuid.MustParse(conds.GetID().GetValue())))
// 		default:
// 			return nil, fmt.Errorf("invalid general field")
// 		}
// 	}
// 	if conds.GoodID != nil {
// 		switch conds.GetGoodID().GetOp() {
// 		case cruder.EQ:
// 			stm.Where(general.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
// 		default:
// 			return nil, fmt.Errorf("invalid general field")
// 		}
// 	}
// 	if conds.CoinTypeID != nil {
// 		switch conds.GetCoinTypeID().GetOp() {
// 		case cruder.EQ:
// 			stm.Where(general.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
// 		default:
// 			return nil, fmt.Errorf("invalid general field")
// 		}
// 	}
// 	if conds.Amount != nil {
// 		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
// 		if err != nil {
// 			return nil, err
// 		}
// 		switch conds.GetAmount().GetOp() {
// 		case cruder.LT:
// 			stm.Where(general.AmountLT(amount))
// 		case cruder.GT:
// 			stm.Where(general.AmountGT(amount))
// 		case cruder.EQ:
// 			stm.Where(general.AmountEQ(amount))
// 		default:
// 			return nil, fmt.Errorf("invalid general field")
// 		}
// 	}
// 	if conds.ToPlatform != nil {
// 		toPlatform, err := decimal.NewFromString(conds.GetToPlatform().GetValue())
// 		if err != nil {
// 			return nil, err
// 		}
// 		switch conds.GetToPlatform().GetOp() {
// 		case cruder.LT:
// 			stm.Where(general.ToPlatformLT(toPlatform))
// 		case cruder.GT:
// 			stm.Where(general.ToPlatformGT(toPlatform))
// 		case cruder.EQ:
// 			stm.Where(general.ToPlatformEQ(toPlatform))
// 		default:
// 			return nil, fmt.Errorf("invalid general field")
// 		}
// 	}
// 	if conds.ToUser != nil {
// 		toUser, err := decimal.NewFromString(conds.GetToUser().GetValue())
// 		if err != nil {
// 			return nil, err
// 		}
// 		switch conds.GetToUser().GetOp() {
// 		case cruder.LT:
// 			stm.Where(general.ToUserLT(toUser))
// 		case cruder.GT:
// 			stm.Where(general.ToUserGT(toUser))
// 		case cruder.EQ:
// 			stm.Where(general.ToUserEQ(toUser))
// 		default:
// 			return nil, fmt.Errorf("invalid general field")
// 		}
// 	}
// 	return stm, nil
// }

// func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.MiningGeneral, int, error) {
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

// 	rows := []*ent.MiningGeneral{}
// 	var total int
// 	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
// 		stm, err := SetQueryConds(conds, cli.MiningGeneral.Query())
// 		if err != nil {
// 			return err
// 		}

// 		total, err = stm.Count(_ctx)
// 		if err != nil {
// 			return err
// 		}

// 		rows, err = stm.
// 			Offset(offset).
// 			Order(ent.Desc(general.FieldUpdatedAt)).
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

// func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.MiningGeneral, error) {
// 	var info *ent.MiningGeneral
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
// 		stm, err := SetQueryConds(conds, cli.MiningGeneral.Query())
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
// 		stm, err := SetQueryConds(conds, cli.MiningGeneral.Query())
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
// 		exist, err = cli.MiningGeneral.Query().Where(general.ID(id)).Exist(_ctx)
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
// 		stm, err := SetQueryConds(conds, cli.MiningGeneral.Query())
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

// func Delete(ctx context.Context, id uuid.UUID) (*ent.MiningGeneral, error) {
// 	var info *ent.MiningGeneral
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
// 		info, err = cli.MiningGeneral.UpdateOneID(id).
// 			SetDeletedAt(uint32(time.Now().Unix())).
// 			Save(_ctx)
// 		return err
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return info, nil
// }
