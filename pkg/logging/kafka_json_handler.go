package logging

import (
	"bytes"
	"context"
	"io"
	"log/slog"

	"github.com/2group/2sales.core-service/pkg/kafka"
)

// KafkaJSONHandler is a slog.Handler that uses the standard JSON formatting
// (which includes extra attributes) and then publishes the formatted output to Kafka.
type KafkaJSONHandler struct {
	publisher *kafka.KafkaPublisher
	topic     string
	opts      *slog.HandlerOptions
}

// NewKafkaJSONHandler creates a new KafkaJSONHandler.
func NewKafkaJSONHandler(publisher *kafka.KafkaPublisher, topic string, opts *slog.HandlerOptions) *KafkaJSONHandler {
	return &KafkaJSONHandler{
		publisher: publisher,
		topic:     topic,
		opts:      opts,
	}
}

// Enabled reports whether a log at the given level should be emitted.
// It compares the record level with the threshold provided in h.opts.Level.
func (h *KafkaJSONHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if h.opts.Level != nil {
		threshold := h.opts.Level.Level() // Convert the Leveler to a slog.Level.
		return level >= threshold
	}
	return true
}

// Handle formats the log record as JSON using a temporary JSON handler,
// then publishes that JSON to Kafka.
func (h *KafkaJSONHandler) Handle(ctx context.Context, record slog.Record) error {
	var buf bytes.Buffer
	// Create a temporary JSON handler that writes to buf.
	jsonHandler := slog.NewJSONHandler(&buf, h.opts)
	if err := jsonHandler.Handle(ctx, record); err != nil {
		return err
	}
	// Read the formatted JSON.
	data, err := io.ReadAll(&buf)
	if err != nil {
		return err
	}
	return h.publisher.Publish(h.topic, data)
}

// WithAttrs returns a new handler with additional attributes.
// For simplicity, we return the same handler.
func (h *KafkaJSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// WithGroup returns a new handler with the given group appended.
// For simplicity, we return the same handler.
func (h *KafkaJSONHandler) WithGroup(name string) slog.Handler {
	return h
}
