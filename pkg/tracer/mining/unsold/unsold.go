package unsold

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"
)

func trace(span trace1.Span, in *npool.UnsoldReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.String(fmt.Sprintf("Amount.%v", index), in.GetAmount()),
		attribute.Int64(fmt.Sprintf("BenefitDate.%v", index), int64(in.GetBenefitDate())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.UnsoldReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("Amount.Op", in.GetAmount().GetOp()),
		attribute.String("Amount.Value", in.GetAmount().GetValue()),
		attribute.String("BenefitDate.Op", in.GetBenefitDate().GetOp()),
		attribute.Int64("BenefitDate.Value", int64(in.GetBenefitDate().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.UnsoldReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
