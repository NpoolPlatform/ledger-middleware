package general

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"

	mgeneralmgrcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/mining/general"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) GetGeneralOnly(ctx context.Context, in *npool.GetGeneralOnlyRequest) (*npool.GetGeneralOnlyResponse, error) {
	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.Amount != nil {
		if _, err := decimal.NewFromString(conds.GetAmount().GetValue()); err != nil {
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.ToPlatform != nil {
		if _, err := decimal.NewFromString(conds.GetToPlatform().GetValue()); err != nil {
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.ToUser != nil {
		if _, err := decimal.NewFromString(conds.GetToUser().GetValue()); err != nil {
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := mgeneralmgrcli.GetGeneralOnly(ctx, conds)
	if err != nil {
		return &npool.GetGeneralOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralOnlyResponse{
		Info: info,
	}, nil
}
