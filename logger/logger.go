package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	zapLogger *zap.Logger
	sugar     *zap.SugaredLogger
}

func New(logLevel string, production bool) (*Logger, error) {
	var cfg zap.Config
	if production {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	level, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		return nil, err
	}
	cfg.Level = level

	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{
		zapLogger: zapLogger,
		sugar:     zapLogger.Sugar(),
	}, nil
}
