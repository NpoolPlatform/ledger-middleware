package bookkeeping

import (
	"context"

	bookkeeping1 "github.com/NpoolPlatform/ledger-middleware/pkg/mining/bookkeeping"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/bookkeeping"

	mdetailmgrcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/mining/detail"
	mdetailmgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if _, err := uuid.Parse(in.GetGoodID()); err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	totalAmount, err := decimal.NewFromString(in.GetTotalAmount())
	if err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if totalAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, "TotalAmount is invalid")
	}
	unsoldAmount, err := decimal.NewFromString(in.GetUnsoldAmount())
	if err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if unsoldAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, "UnsoldAmount is invalid")
	}
	feeAmount, err := decimal.NewFromString(in.GetTechniqueServiceFeeAmount())
	if err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if feeAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, "FeeAmount is invalid")
	}
	if in.GetBenefitDate() == 0 {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, "BenefitDate is invalid")
	}

	detail, err := mdetailmgrcli.GetDetailOnly(ctx, &mdetailmgrpb.Conds{
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
		BenefitDate: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: in.GetBenefitDate(),
		},
	})
	if err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if detail != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, "Benefit exist")
	}

	err = bookkeeping1.BookKeeping(
		ctx,
		in.GetGoodID(),
		in.GetCoinTypeID(),
		totalAmount,
		unsoldAmount,
		feeAmount,
		in.GetBenefitDate(),
	)
	if err != nil {
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}
