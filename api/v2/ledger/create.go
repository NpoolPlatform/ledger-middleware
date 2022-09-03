package ledger

import (
	"context"
	"errors"

	"github.com/NpoolPlatform/ledger-manager/api/detail"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/ledger"

	errno "github.com/NpoolPlatform/ledger-middleware/pkg/errno"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if len(in.GetInfos()) == 0 {
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, "params is empty")
	}

	for _, val := range in.GetInfos() {
		if err := detail.Validate(val); err != nil {
			logger.Sugar().Errorw("BookKeeping", "error", err)
			return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if err := ledger1.BookKeepingV2(ctx, in.GetInfos()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		if errors.Is(err, errno.ErrAlreadyExists) {
			return &npool.BookKeepingResponse{}, status.Error(codes.AlreadyExists, err.Error())
		}
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}
