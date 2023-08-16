package unsoldstatement

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/unsoldstatement"

	"google.golang.org/grpc"
)

type Server struct {
	unsoldstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	unsoldstatement.RegisterMiddlewareServer(server, &Server{})
}
