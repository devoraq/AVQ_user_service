package grpc

import (
	"log/slog"
	"net"

	"github.com/DENFNC/awq_user_service/internal/app/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HandlerRegisterer interface {
	Register(grpc *grpc.Server)
}

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	addr       string
}

func NewApp(
	log *slog.Logger,
	addr string,
	reflect bool,
	handlers ...HandlerRegisterer,
) *App {
	const op = "grpc.NewApp"

	log = log.With("op", op)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.ValidateInterceptor(log),
			interceptor.PanicRecoveryInterceptor(log),
			interceptor.LoggerInterceptor(log),
		),
	)

	if reflect {
		log.Info(
			"Reflection enabled",
		)
		reflection.Register(grpcServer)
	}

	for _, handler := range handlers {
		handler.Register(grpcServer)
	}

	return &App{
		log:        log,
		grpcServer: grpcServer,
		addr:       addr,
	}
}

func (app *App) Start() {
	const op = "grpc.app.Start"

	log := app.log.With("op", op)

	if err := app.gRPCStart(); err != nil {
		log.Error(
			"Failed to start gRPC server",
			slog.String("err", err.Error()),
		)
		panic(err)
	}
}

func (app *App) Stop() {
	const op = "grpc.app.Stop"

	log := app.log.With("op", op)

	log.Info(
		"Stopping the server",
	)

	app.grpcServer.GracefulStop()
}

func (app *App) gRPCStart() error {
	const op = "grpc.app.gRPCStart"

	log := app.log.With("op", op)

	lis, err := net.Listen("tcp", app.addr)
	if err != nil {
		log.Error(
			"Failed to listen",
			slog.String("err", err.Error()),
		)
		return err
	}

	log.Info(
		"Starting gRPC server",
		slog.String("addr", lis.Addr().String()),
	)

	if err := app.grpcServer.Serve(lis); err != nil {
		log.Error(
			"Failed to gRPC serve",
			slog.String("err", err.Error()),
		)
		return err
	}

	return nil
}
