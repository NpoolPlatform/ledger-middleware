package detail

import (
	"context"

	// mgrpb "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/detail"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetDetailOnly(ctx context.Context, in *npool.GetDetailOnlyRequest) (*npool.GetDetailOnlyResponse, error) {
	return &npool.GetDetailOnlyResponse{}, status.Error(codes.Unimplemented, "not implemented")
}
