package log

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

// ZapDevLogger creates a zap dev logger.
func ZapDevLogger() (logr.Logger, error) {
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		return logr.Logger{}, fmt.Errorf("create zap logger instance: %w", err)
	}

	return zapr.NewLogger(zapLog), nil
}
