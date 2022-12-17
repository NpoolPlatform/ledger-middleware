package unsold

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"

	munsoldmgrcli "github.com/NpoolPlatform/ledger-manager/pkg/client/mining/unsold"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) GetUnsoldOnly(ctx context.Context, in *npool.GetUnsoldOnlyRequest) (*npool.GetUnsoldOnlyResponse, error) {
	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return &npool.GetUnsoldOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			return &npool.GetUnsoldOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			return &npool.GetUnsoldOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.Amount != nil {
		if _, err := decimal.NewFromString(conds.GetAmount().GetValue()); err != nil {
			return &npool.GetUnsoldOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := munsoldmgrcli.GetUnsoldOnly(ctx, conds)
	if err != nil {
		return &npool.GetUnsoldOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUnsoldOnlyResponse{
		Info: info,
	}, nil
}
