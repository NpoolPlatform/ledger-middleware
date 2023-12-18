package coupon

import (
	"github.com/NpoolPlatform/message/npool/ledger/mw/v2/withdraw/coupon"
	"google.golang.org/grpc"
)

type Server struct {
	coupon.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coupon.RegisterMiddlewareServer(server, &Server{})
}
