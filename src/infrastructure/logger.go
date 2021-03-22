package infrastructure

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	logLevelDebug = "debug"
	logLevelInfo  = "info"
	logLevelWarn  = "warn"
	logLevelErr   = "error"
)

func NewLogger(logLevel string) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()

	switch logLevel {
	case logLevelDebug:
		cfg.Level.SetLevel(zap.DebugLevel)
	case logLevelInfo:
		cfg.Level.SetLevel(zap.InfoLevel)
	case logLevelWarn:
		cfg.Level.SetLevel(zap.WarnLevel)
	case logLevelErr:
		cfg.Level.SetLevel(zap.ErrorLevel)
	default:
		return nil, errors.Errorf("unexpected_log_level: `%s`", logLevel)
	}

	cfg.DisableStacktrace = true
	logger, err := cfg.Build(zap.AddStacktrace(zap.PanicLevel))
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
