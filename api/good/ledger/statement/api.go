package statement

import (
	"context"

	goodstatement "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type Server struct {
	goodstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodstatement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return goodstatement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
