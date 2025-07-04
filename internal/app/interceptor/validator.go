package interceptor

import (
	"context"
	"log/slog"
	"reflect"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValidateInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	validator, err := protovalidate.New()
	if err != nil {
		log.Error(
			"failed to create protovalidate validator",
			slog.String("error", err.Error()),
		)
		panic("critical: cannot initialize protovalidate validator")
	}

	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		msg, ok := req.(interface{ ProtoReflect() protoreflect.Message })
		if !ok {
			log.Warn("request is not a protobuf message",
				slog.String("method", info.FullMethod),
				slog.Any("type", reflect.TypeOf(req)),
			)
			return nil, status.Error(codes.Internal, "invalid request type")
		}

		if err := validator.Validate(msg); err != nil {
			log.Warn("validation failed",
				slog.String("method", info.FullMethod),
				slog.String("error", err.Error()),
				slog.Any("request", req),
			)
			return nil, status.Error(codes.InvalidArgument, "validation failed")
		}

		return handler(ctx, req)
	}
}
