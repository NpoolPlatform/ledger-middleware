package ledger

import (
	goodledger "github.com/NpoolPlatform/message/npool/ledger/mw/v2/good/ledger"

	"google.golang.org/grpc"
)

type Server struct {
	goodledger.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodledger.RegisterMiddlewareServer(server, &Server{})
}
