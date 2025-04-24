package main

import (
	"log/slog"
	"os"

	"github.com/2group/2sales.core-service/internal/app"
	"github.com/2group/2sales.core-service/internal/config"
	"github.com/2group/2sales.core-service/pkg/kafka"
	"github.com/2group/2sales.core-service/pkg/logging"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	kafkaPublisher, err := kafka.NewKafkaPublisher(cfg.KafkaBroker)
	if err != nil {
		slog.Error("failed to create Kafka publisher", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer kafkaPublisher.Close()
	log := logging.SetupLogger(kafkaPublisher, "core_service_logs", cfg.Env)

	log.Info("starting_application", "port", cfg.REST.Port)

	application := app.NewAPIServer(cfg, log)

	if err := application.Run(); err != nil {
		panic(err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log.With("service", "core_service")
}
