package lock

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger/lock"

	"google.golang.org/grpc"
)

type Server struct {
	lock.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	lock.RegisterMiddlewareServer(server, &Server{})
}
