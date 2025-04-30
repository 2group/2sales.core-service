package logging

import (
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var (
	rootLogger zerolog.Logger
	once       sync.Once
)

// SetupLogger initializes the global logger once based on the environment.
// env: "prod" -> InfoLevel, others ("dev","local",etc.) -> DebugLevel.
func SetupLogger(env string) {
	once.Do(func() {
		// 1) Set global level
		level := zerolog.DebugLevel
		if strings.EqualFold(env, "prod") {
			level = zerolog.InfoLevel
		}
		zerolog.SetGlobalLevel(level)

		// 2) Choose output writer
		var output io.Writer
		if strings.EqualFold(env, "dev") || strings.EqualFold(env, "local") {
			output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		} else {
			output = os.Stdout
		}

		// 3) Build root logger with timestamp and caller info
		rootLogger = zerolog.New(output).
			With().
			Timestamp().
			Caller().
			Logger()
	})
}

// L returns the global logger instance. Call SetupLogger before using.
func L() *zerolog.Logger {
	return &rootLogger
}

// FromContext returns a logger enriched with "correlation_id" from ctx if present.
// If no correlation ID is found, returns the root logger.
