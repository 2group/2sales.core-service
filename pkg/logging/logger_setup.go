package logging

import (
	"log/slog"
	"os"

	"github.com/2group/2sales.core-service/pkg/kafka"
)

func SetupLogger(publisher *kafka.KafkaPublisher, topic, env string) {
	var level slog.Level
	switch env {
	case "local", "dev":
		level = slog.LevelDebug
	case "prod":
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}

	var handler slog.Handler
	if env == "local" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	} else {
		handler = NewKafkaJSONHandler(publisher, topic, &slog.HandlerOptions{Level: level})
	}

	// Attach static fields and set as default
	logger := slog.New(handler).With()
	slog.SetDefault(logger)
}
