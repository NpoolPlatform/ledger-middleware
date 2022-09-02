package api

import (
	"context"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/api/v1/ledger"
	ledger2 "github.com/NpoolPlatform/ledger-middleware/api/v2/ledger"
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
	ledger2.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledger.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
