package general

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	general.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	general.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
