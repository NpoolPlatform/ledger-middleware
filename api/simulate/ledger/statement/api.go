package statement

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/simulate/ledger/statement"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	statement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	statement.RegisterMiddlewareServer(server, &Server{})
}
func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return statement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
