package unsold

import (
	"context"

	munsold "github.com/NpoolPlatform/ledger-middleware/pkg/mining/unsold"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) CreateUnsold(ctx context.Context, in *npool.CreateUnsoldRequest) (*npool.CreateUnsoldResponse, error) {
	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			return &npool.CreateUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if _, err := uuid.Parse(in.GetInfo().GetGoodID()); err != nil {
		return &npool.CreateUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		return &npool.CreateUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	amount, err := decimal.NewFromString(in.GetInfo().GetAmount())
	if err != nil {
		return &npool.CreateUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return &npool.CreateUnsoldResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().GetBenefitIntervalHours() <= 0 {
		return &npool.CreateUnsoldResponse{}, status.Error(codes.InvalidArgument, "BenefitIntervalHours is invalid")
	}

	info, err := munsold.CreateUnsold(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateUnsoldResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUnsoldResponse{
		Info: info,
	}, nil
}
