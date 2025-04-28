package logging

import (
	"bytes"
	"context"
	"log/slog"

	"github.com/2group/2sales.core-service/pkg/kafka"
)

// KafkaJSONHandler is a slog.Handler that uses the standard JSON formatting
// (which includes extra attributes) and then publishes the formatted output to Kafka.
type KafkaJSONHandler struct {
	publisher *kafka.KafkaPublisher
	topic     string
	opts      *slog.HandlerOptions

	// the bits we were dropping:
	attrs  []slog.Attr
	groups []string
}

func NewKafkaJSONHandler(pub *kafka.KafkaPublisher, topic string, opts *slog.HandlerOptions) *KafkaJSONHandler {
	return &KafkaJSONHandler{publisher: pub, topic: topic, opts: opts}
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

func (h *KafkaJSONHandler) Handle(ctx context.Context, rec slog.Record) error {
	var buf bytes.Buffer

	// 1) Start with a JSONHandler, but store it as the slog.Handler interface:
	var jsonH slog.Handler = slog.NewJSONHandler(&buf, h.opts)

	// 2) Propagate any static attrs you collected:
	if len(h.attrs) > 0 {
		jsonH = jsonH.WithAttrs(h.attrs)
	}
	// 3) Propagate any static groups:
	for _, g := range h.groups {
		jsonH = jsonH.WithGroup(g)
	}

	// 4) Finally format the record (this applies both the record's dynamic
	//    attrs AND your static attrs/groups):
	if err := jsonH.Handle(ctx, rec); err != nil {
		return err
	}

	// 5) Publish the raw bytes to Kafka:
	return h.publisher.Publish(h.topic, buf.Bytes())
}

// WithAttrs returns a new handler with additional attributes.
// For simplicity, we return the same handler.
func (h *KafkaJSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// append incoming attrs to our list
	return &KafkaJSONHandler{
		publisher: h.publisher,
		topic:     h.topic,
		opts:      h.opts,
		attrs:     append(h.attrs, attrs...),
		groups:    h.groups,
	}
}

// WithGroup returns a new handler with the given group appended.
// For simplicity, we return the same handler.
func (h *KafkaJSONHandler) WithGroup(name string) slog.Handler {
	return &KafkaJSONHandler{
		publisher: h.publisher,
		topic:     h.topic,
		opts:      h.opts,
		attrs:     h.attrs,
		groups:    append(h.groups, name),
	}
}
