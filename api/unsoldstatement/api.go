package unsoldstatement

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/unsoldstatement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	unsoldstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	unsoldstatement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
