package statement

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"

	"google.golang.org/grpc"
)

type Server struct {
	statement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	statement.RegisterMiddlewareServer(server, &Server{})
}
