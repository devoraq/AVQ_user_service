package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		timeNow := time.Now()

		resp, err = handler(ctx, req)

		var code codes.Code
		if err != nil {
			if s, ok := status.FromError(err); ok {
				code = s.Code()
			} else {
				code = codes.Unknown
			}
		} else {
			code = codes.OK
		}

		log.Info(
			"Request",
			slog.String("method", info.FullMethod),
			slog.String("status", code.String()),
			slog.Int("code", int(code)),
			slog.Duration("time", time.Since(timeNow)),
		)

		return resp, err
	}
}
