package goodstatement

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodstatement"

	"google.golang.org/grpc"
)

type Server struct {
	goodstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodstatement.RegisterMiddlewareServer(server, &Server{})
}
