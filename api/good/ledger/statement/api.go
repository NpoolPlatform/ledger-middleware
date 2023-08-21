package statement

import (
	goodstatement "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger/statement"

	"google.golang.org/grpc"
)

type Server struct {
	goodstatement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodstatement.RegisterMiddlewareServer(server, &Server{})
}
