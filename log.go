package log

import (
	"context"
	"io"
	"os"

	"github.com/go-logr/logr"
)

type ctxValueType string

const (
	logKey ctxValueType = "_logger"
)

// From extracts a logger from a context. If one does not exist, a new zap dev logger is created. This will panic if
// unable to create the zap dev logger.
func From(ctx context.Context) logr.Logger {
	if ctx == nil {
		return newLogger()
	}

	if logger, ok := ctx.Value(logKey).(logr.Logger); ok {
		return logger
	}

	return newLogger()
}

func newLogger() logr.Logger {
	l, err := ZapDevLogger()
	if err != nil {
		panic(err)
	}

	return l
}

// LoggerOption is an option for configuring the logger.
type LoggerOption func(config *LoggerConfig)

// LoggerOutput sets the output location for the logger.
func LoggerOutput(w io.Writer) LoggerOption {
	if w == nil {
		panic("logger output cannot be nil")
	}

	return func(config *LoggerConfig) {
		config.out = w
	}
}

// WithLogger creates a context with an existing logger.
func WithLogger(ctx context.Context, logger logr.Logger) context.Context {
	return context.WithValue(ctx, logKey, logger)
}

// LoggerConfig is logger configuration.NewLogger
type LoggerConfig struct {
	out io.Writer
}

func newLoggerConfig() *LoggerConfig {
	config := &LoggerConfig{
		out: os.Stderr,
	}

	return config
}
