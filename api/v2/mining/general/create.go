package general

import (
	"context"

	mgeneral "github.com/NpoolPlatform/ledger-middleware/pkg/mining/general"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CreateGeneral(ctx context.Context, in *npool.CreateGeneralRequest) (*npool.CreateGeneralResponse, error) {
	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			return &npool.CreateGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if _, err := uuid.Parse(in.GetInfo().GetGoodID()); err != nil {
		return &npool.CreateGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		return &npool.CreateGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mgeneral.CreateGeneral(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGeneralResponse{
		Info: info,
	}, nil
}
