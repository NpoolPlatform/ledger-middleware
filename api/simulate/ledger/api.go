package ledger

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledger.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	ledger.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return ledger.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
