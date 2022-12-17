package detail

import (
	"context"

	mdetail "github.com/NpoolPlatform/ledger-middleware/pkg/mining/detail"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) CreateDetail(ctx context.Context, in *npool.CreateDetailRequest) (*npool.CreateDetailResponse, error) {
	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			return &npool.CreateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if _, err := uuid.Parse(in.GetInfo().GetGoodID()); err != nil {
		return &npool.CreateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		return &npool.CreateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	amount, err := decimal.NewFromString(in.GetInfo().GetAmount())
	if err != nil {
		return &npool.CreateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return &npool.CreateDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mdetail.CreateDetail(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailResponse{
		Info: info,
	}, nil
}
