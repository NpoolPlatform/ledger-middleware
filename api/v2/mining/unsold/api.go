package unsold

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsold"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	unsold.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	unsold.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
