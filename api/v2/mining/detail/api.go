package detail

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
