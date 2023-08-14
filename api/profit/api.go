package profit

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/profit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	profit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	profit.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
