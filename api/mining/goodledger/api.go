package goodledger

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/mining/goodledger"

	"google.golang.org/grpc"
)

type Server struct {
	goodledger.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodledger.RegisterMiddlewareServer(server, &Server{})
}
