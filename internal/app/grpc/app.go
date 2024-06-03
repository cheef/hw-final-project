package grpcapp

import (
	"fmt"
	bfagrpc "github.com/cheef/hw-final-project/internal/grpc/bfa_protection"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
	log        *slog.Logger
}

func New(log *slog.Logger, service bfagrpc.BFAProtection, port int) *App {
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))

	bfagrpc.Register(gRPCServer, service)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (app *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))

	if err != nil {
		return fmt.Errorf("%s: %w", "grpcapp.Run", err)
	}

	app.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := app.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", "grpcapp.Run", err)
	}

	return nil
}
