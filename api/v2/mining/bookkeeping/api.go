package bookkeeping

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/bookkeeping"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	bookkeeping.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	bookkeeping.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
