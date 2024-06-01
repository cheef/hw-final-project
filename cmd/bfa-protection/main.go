package main

import (
	"context"
	app "github.com/cheef/hw-final-project/internal/app"
	"github.com/cheef/hw-final-project/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	log := NewLogger()
	cfg, err := config.Load()

	if err != nil {
		log.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	application := app.New(ctx, log, cfg)

	if err := application.GRPCApp.Run(); err != nil {
		log.Error("failed to run GRPC", slog.String("error", err.Error()))
		os.Exit(1)
	}

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop

		if err := application.Storage.Stop(ctx); err != nil {
			log.Error("Error occurred during stopping storage", slog.String("error", err.Error()))
		}

		log.Info("Gracefully stopped")
		os.Exit(1)
	}()
}

func NewLogger() *slog.Logger {
	logOpts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logHandler := slog.NewJSONHandler(os.Stdout, logOpts)

	return slog.New(logHandler)
}
