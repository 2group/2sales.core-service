package logging

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger
var slogLogger *SlogAdapter

func SetupLogger(env string) {
	switch strings.ToLower(env) {
	case "prod":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	var output io.Writer
	if env == "local" || env == "dev" {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		output = os.Stdout
	}

	logger = zerolog.New(output).With().Timestamp().Logger()
	slogLogger = NewSlogAdapter(logger)
}

func Logger() zerolog.Logger {
	return logger
}

func Slog() *SlogAdapter {
	return slogLogger
}

type SlogAdapter struct {
	z zerolog.Logger
}

func NewSlogAdapter(z zerolog.Logger) *SlogAdapter {
	return &SlogAdapter{z: z}
}

func (l *SlogAdapter) With(args ...any) *SlogAdapter {
	child := l.z.With()
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		child = child.Str(key, stringify(args[i+1]))
	}
	return &SlogAdapter{z: child.Logger()}
}

func (l *SlogAdapter) Info(msg string, args ...any) {
	l.write(l.z.Info(), msg, args...)
}

func (l *SlogAdapter) Debug(msg string, args ...any) {
	l.write(l.z.Debug(), msg, args...)
}

func (l *SlogAdapter) Error(msg string, args ...any) {
	l.write(l.z.Error(), msg, args...)
}

func (l *SlogAdapter) write(event *zerolog.Event, msg string, args ...any) {
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		event = event.Interface(key, args[i+1])
	}
	event.Msg(msg)
}

func stringify(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return "<non-string>"
}
