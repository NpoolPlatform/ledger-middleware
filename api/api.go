package api

import (
	"context"

	goodledger "github.com/NpoolPlatform/ledger-middleware/api/good/ledger"
	goodstatement "github.com/NpoolPlatform/ledger-middleware/api/good/ledger/statement"
	ledger1 "github.com/NpoolPlatform/ledger-middleware/api/ledger"
	profit "github.com/NpoolPlatform/ledger-middleware/api/ledger/profit"
	statement "github.com/NpoolPlatform/ledger-middleware/api/ledger/statement"
	withdraw "github.com/NpoolPlatform/ledger-middleware/api/withdraw"
	couponwithdraw "github.com/NpoolPlatform/ledger-middleware/api/withdraw/coupon"

	ledger "github.com/NpoolPlatform/message/npool/ledger/mw/v2"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledger.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	ledger.RegisterMiddlewareServer(server, &Server{})
	ledger1.Register(server)
	goodledger.Register(server)
	statement.Register(server)
	profit.Register(server)
	withdraw.Register(server)
	couponwithdraw.Register(server)
	goodstatement.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledger.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := ledger1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := statement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := profit.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := withdraw.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := goodstatement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
