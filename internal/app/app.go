package bfaapp

import (
	grpcapp "github.com/cheef/hw-final-project/internal/app/grpc"
	"github.com/cheef/hw-final-project/internal/config"
	"github.com/cheef/hw-final-project/internal/services"
	bfaservice "github.com/cheef/hw-final-project/internal/services/bfa_protection"
	"log/slog"
)

type App struct {
	GRPCApp      *grpcapp.App
	LimitChecker *services.LimitChecker
}

func New(log *slog.Logger, cfg *config.Config) *App {
	limitChecker := services.NewLimitChecker(cfg, log)
	service := bfaservice.New(log, limitChecker)
	grpcApp := grpcapp.New(log, service, cfg.GRPC.Port)

	return &App{
		GRPCApp:      grpcApp,
		LimitChecker: limitChecker,
	}
}
