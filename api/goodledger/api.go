package goodledger

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/goodledger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodledger.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodledger.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
