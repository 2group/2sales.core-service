package logging

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/2group/2sales.core-service/pkg/kafka"
)

// KafkaLogHandler is a custom slog handler that publishes logs to Kafka.
type KafkaLogHandler struct {
	publisher *kafka.KafkaPublisher
	topic     string
	minLevel  slog.Level
}

// NewKafkaLogHandler creates a new instance of KafkaLogHandler.
func NewKafkaLogHandler(publisher *kafka.KafkaPublisher, topic string, minLevel slog.Level) *KafkaLogHandler {
	return &KafkaLogHandler{
		publisher: publisher,
		topic:     topic,
		minLevel:  minLevel,
	}
}

// Enabled returns true if the log level is high enough.
func (h *KafkaLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

// Handle marshals the log record into JSON and publishes it to Kafka.
func (h *KafkaLogHandler) Handle(ctx context.Context, record slog.Record) error {
	// Here we include the basic fields.
	payload, err := json.Marshal(struct {
		Time    time.Time `json:"time"`
		Level   string    `json:"level"`
		Message string    `json:"message"`
	}{
		Time:    record.Time,
		Level:   record.Level.String(),
		Message: record.Message,
	})
	if err != nil {
		return err
	}

	return h.publisher.Publish(h.topic, payload)
}

// WithAttrs returns a new handler with additional attributes (not used in this example).
func (h *KafkaLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// WithGroup returns a new handler with the given group name (not used in this example).
func (h *KafkaLogHandler) WithGroup(name string) slog.Handler {
	return h
}

// MultiHandler forwards log records to multiple handlers.
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler creates a new MultiHandler.
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

// Enabled returns true if any underlying handler is enabled.
func (mh *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range mh.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle calls Handle on all underlying handlers.
func (mh *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range mh.handlers {
		if err := h.Handle(ctx, record); err != nil {
			return err
		}
	}
	return nil
}

// WithAttrs returns a new MultiHandler with each underlying handler updated.
func (mh *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(mh.handlers))
	for i, h := range mh.handlers {
		newHandlers[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: newHandlers}
}

// WithGroup returns a new MultiHandler with each underlying handler updated.
func (mh *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(mh.handlers))
	for i, h := range mh.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: newHandlers}
}
