package logging

import (
	"log/slog"
	"os"

	"github.com/2group/2sales.core-service/pkg/kafka"
)

func SetupLogger(publisher *kafka.KafkaPublisher, topic, env string) *slog.Logger {
	var level slog.Level
	switch env {
	case "local":
		level = slog.LevelDebug
	case "dev":
		level = slog.LevelDebug
	case "prod":
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}

	if env == "local" {
		// In local, log only to the console.
		consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		return slog.New(consoleHandler).With("service", "product_service")
	}

	// For non-local environments, use our Kafka JSON handler.
	opts := &slog.HandlerOptions{Level: level}
	kafkaJSONHandler := NewKafkaJSONHandler(publisher, topic, opts)
	return slog.New(kafkaJSONHandler).With("service", "product_service")
}
