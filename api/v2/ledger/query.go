package ledger

import (
	"context"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"

	curl "github.com/NpoolPlatform/ledger-manager/pkg/crud/general"

	converter "github.com/NpoolPlatform/ledger-manager/pkg/converter/general"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetGeneralOnly(ctx context.Context, in *npool.GetGeneralOnlyRequest) (*npool.GetGeneralOnlyResponse, error) {
	if in.Conds == nil {
		logger.Sugar().Errorw("GetGeneralOnly", "Conds", in.Conds)
		return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, "Conds is empty")
	}
	if in.Conds.ID != nil {
		if _, err := uuid.Parse(in.Conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.Conds.GetID().GetValue(), "error", err)
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, "Conds ID value is invalid")
		}
	}
	if in.Conds.AppID != nil {
		if _, err := uuid.Parse(in.Conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "AppID", in.Conds.GetAppID().GetValue(), "error", err)
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, "Conds AppID value is invalid")
		}
	}
	if in.Conds.CoinTypeID != nil {
		if _, err := uuid.Parse(in.Conds.GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("validate", "CoinTypeID", in.Conds.GetCoinTypeID().GetValue(), "error", err)
			return &npool.GetGeneralOnlyResponse{}, status.Error(codes.InvalidArgument, "Conds CoinTypeID value is invalid")
		}
	}
	info, err := curl.RowOnly(ctx, in.Conds)
	if err != nil {
		logger.Sugar().Errorw("GetGeneralOnly", "error", err)
		return &npool.GetGeneralOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGeneralOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
