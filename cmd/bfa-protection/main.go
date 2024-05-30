package main

import (
	app "github.com/cheef/hw-final-project/internal/app"
	"github.com/cheef/hw-final-project/internal/config"
	"log/slog"
	"os"
)

func main() {
	logOpts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logHandler := slog.NewJSONHandler(os.Stdout, logOpts)
	log := slog.New(logHandler)
	cfg, err := config.Load()

	if err != nil {
		log.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	application := app.New(log, cfg)

	if err := application.GRPCApp.Run(); err != nil {
		log.Error("failed to run GRPC", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
