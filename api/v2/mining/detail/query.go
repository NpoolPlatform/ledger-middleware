package detail

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	mdetailmgrcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/mining/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) GetDetailOnly(ctx context.Context, in *npool.GetDetailOnlyRequest) (*npool.GetDetailOnlyResponse, error) {
	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return &npool.GetDetailOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return &npool.GetDetailOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			return &npool.GetDetailOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.Amount != nil {
		if _, err := decimal.NewFromString(conds.GetAmount().GetValue()); err != nil {
			return &npool.GetDetailOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := mdetailmgrcli.GetDetailOnly(ctx, conds)
	if err != nil {
		return &npool.GetDetailOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailOnlyResponse{
		Info: info,
	}, nil
}
func (s *Server) GetDetails(ctx context.Context, in *npool.GetDetailsRequest) (*npool.GetDetailsResponse, error) {
	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.Amount != nil {
		if _, err := decimal.NewFromString(conds.GetAmount().GetValue()); err != nil {
			return &npool.GetDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	infos, total, err := mdetailmgrcli.GetDetails(ctx, conds, int32(in.GetOffset()), int32(in.GetLimit()))
	if err != nil {
		return &npool.GetDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
