package main

import (
	"log/slog"
	"os"

	"github.com/2group/2sales.core-service/internal/app"
	"github.com/2group/2sales.core-service/internal/config"
	"github.com/2group/2sales.core-service/pkg/kafka"
	"github.com/2group/2sales.core-service/pkg/logging"
	"github.com/rs/zerolog/log"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	var kafkaPublisher *kafka.KafkaPublisher
	var err error
	if cfg.Env != "local" {
		kafkaPublisher, err = kafka.NewKafkaPublisher(cfg.KafkaBroker)
		defer kafkaPublisher.Close()
	}
	if err != nil {
		slog.Error("failed to create Kafka publisher", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logging.SetupLogger(cfg.Env)
	log.Info().Int("port", cfg.REST.Port).Msg("starting_application")
	log := setupLogger(cfg.Env)
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
