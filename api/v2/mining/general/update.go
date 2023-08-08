package general

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"

	mgeneralmgrcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/mining/general"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *Server) AddGeneral(ctx context.Context, in *npool.AddGeneralRequest) (*npool.AddGeneralResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.AddGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Amount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetAmount()); err != nil {
			return &npool.AddGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetInfo().ToPlatform != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetToPlatform()); err != nil {
			return &npool.AddGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetInfo().ToUser != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetToUser()); err != nil {
			return &npool.AddGeneralResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := mgeneralmgrcli.AddGeneral(ctx, in.GetInfo())
	if err != nil {
		return &npool.AddGeneralResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AddGeneralResponse{
		Info: info,
	}, nil
}
