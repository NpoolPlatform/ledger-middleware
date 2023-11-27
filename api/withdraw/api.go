package withdraw

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	withdraw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	withdraw.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return withdraw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
