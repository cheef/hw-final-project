package bfaapp

import (
	"context"
	grpcapp "github.com/cheef/hw-final-project/internal/app/grpc"
	"github.com/cheef/hw-final-project/internal/config"
	"github.com/cheef/hw-final-project/internal/services"
	bfaservice "github.com/cheef/hw-final-project/internal/services/bfa_protection"
	sqlstorage "github.com/cheef/hw-final-project/internal/storage/sql"
	"log/slog"
)

type App struct {
	GRPCApp      *grpcapp.App
	LimitChecker *services.LimitChecker
	Storage      *sqlstorage.Storage
}

func New(ctx context.Context, log *slog.Logger, cfg *config.Config) *App {
	storage, _ := sqlstorage.NewStorage(ctx, cfg.Storage, log)
	limitChecker := services.NewLimitChecker(cfg, log)
	service := bfaservice.New(log, storage, limitChecker)
	grpcApp := grpcapp.New(log, service, cfg.GRPC.Port)

	return &App{
		GRPCApp:      grpcApp,
		LimitChecker: limitChecker,
		Storage:      storage,
	}
}
