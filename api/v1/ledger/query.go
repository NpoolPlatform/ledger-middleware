//nolint:nolintlint,dupl
package ledger

import (
	"context"

	constant "github.com/NpoolPlatform/ledger-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ledger1 "github.com/NpoolPlatform/ledger-middleware/pkg/ledger"
	tracer "github.com/NpoolPlatform/ledger-middleware/pkg/tracer"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mw/v2/ledger"

	"github.com/google/uuid"
)

func (s *Server) GetIntervalGenerals(
	ctx context.Context, in *npool.GetIntervalGeneralsRequest,
) (
	*npool.GetIntervalGeneralsResponse, error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetIntervalGenerals")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetIntervalGenerals", "AppID", in.GetAppID(), "error", err)
		return &npool.GetIntervalGeneralsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetIntervalGenerals", "UserID", in.GetUserID(), "error", err)
		return &npool.GetIntervalGeneralsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = tracer.TraceStartEnd(span, in.GetStart(), in.GetEnd())
	span = tracer.TraceInvoker(span, "ledger", "ledger", "GetIntervalGenerals")

	infos, total, err := ledger1.GetIntervalGenerals(
		ctx,
		in.GetAppID(), in.GetUserID(),
		in.GetStart(), in.GetEnd(),
		in.GetOffset(), in.GetLimit(),
	)
	if err != nil {
		logger.Sugar().Errorw("GetIntervalGenerals", "error", err)
		return &npool.GetIntervalGeneralsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetIntervalGeneralsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetIntervalDetails(
	ctx context.Context, in *npool.GetIntervalDetailsRequest,
) (
	*npool.GetIntervalDetailsResponse, error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetIntervalDetails")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetIntervalDetails", "AppID", in.GetAppID(), "error", err)
		return &npool.GetIntervalDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetIntervalDetails", "UserID", in.GetUserID(), "error", err)
		return &npool.GetIntervalDetailsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = tracer.TraceStartEnd(span, in.GetStart(), in.GetEnd())
	span = tracer.TraceInvoker(span, "ledger", "ledger", "GetIntervalDetails")

	infos, total, err := ledger1.GetIntervalDetails(
		ctx,
		in.GetAppID(), in.GetUserID(),
		in.GetStart(), in.GetEnd(),
		in.GetOffset(), in.GetLimit(),
	)
	if err != nil {
		logger.Sugar().Errorw("GetIntervalDetails", "error", err)
		return &npool.GetIntervalDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetIntervalDetailsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetIntervalProfits(
	ctx context.Context, in *npool.GetIntervalProfitsRequest,
) (
	*npool.GetIntervalProfitsResponse, error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetIntervalProfits")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetIntervalProfits", "AppID", in.GetAppID(), "error", err)
		return &npool.GetIntervalProfitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetIntervalProfits", "UserID", in.GetUserID(), "error", err)
		return &npool.GetIntervalProfitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = tracer.TraceStartEnd(span, in.GetStart(), in.GetEnd())
	span = tracer.TraceInvoker(span, "ledger", "ledger", "GetIntervalProfits")

	infos, total, err := ledger1.GetIntervalProfits(
		ctx,
		in.GetAppID(), in.GetUserID(),
		in.GetStart(), in.GetEnd(),
		in.GetOffset(), in.GetLimit(),
	)
	if err != nil {
		logger.Sugar().Errorw("GetIntervalProfits", "error", err)
		return &npool.GetIntervalProfitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetIntervalProfitsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
