package ledger

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v1/ledger"

	"github.com/NpoolPlatform/ledger-manager/api/detail"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if err := detail.
}
