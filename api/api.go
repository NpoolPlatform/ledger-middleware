package api

import (
	"context"

	bookkeeping "github.com/NpoolPlatform/ledger-middleware/api/bookkeeping"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/api/ledger"
	goodledger "github.com/NpoolPlatform/ledger-middleware/api/mining/goodledger"
	goodstatement "github.com/NpoolPlatform/ledger-middleware/api/mining/goodstatement"
	unsold "github.com/NpoolPlatform/ledger-middleware/api/mining/unsoldstatement"
	profit "github.com/NpoolPlatform/ledger-middleware/api/profit"
	statement "github.com/NpoolPlatform/ledger-middleware/api/statement"
	withdraw "github.com/NpoolPlatform/ledger-middleware/api/withdraw"

	ledger "github.com/NpoolPlatform/message/npool/ledger/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledger.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	ledger.RegisterMiddlewareServer(server, &Server{})
	ledger1.Register(server)
	goodledger.Register(server)
	statement.Register(server)
	profit.Register(server)
	bookkeeping.Register(server)
	withdraw.Register(server)
	goodstatement.Register(server)
	unsold.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledger.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
