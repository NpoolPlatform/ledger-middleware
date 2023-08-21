package profit

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/profit"

	"google.golang.org/grpc"
)

type Server struct {
	profit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	profit.RegisterMiddlewareServer(server, &Server{})
}
