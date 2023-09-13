package withdraw

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw"

	"google.golang.org/grpc"
)

type Server struct {
	withdraw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	withdraw.RegisterMiddlewareServer(server, &Server{})
}
