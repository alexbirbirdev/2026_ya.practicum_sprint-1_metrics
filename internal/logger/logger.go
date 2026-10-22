package logger

import (
	"log"

	"go.uber.org/zap"
)

var AppLogger *zap.Logger = zap.NewNop()

func InitAppLogger(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		log.Fatal("problem in parsing logger level: %s", err)
	}

	cfg := zap.NewDevelopmentConfig()

	cfg.Level = lvl

	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatal("problem with building zap logger: %s", err)
	}
	AppLogger = zapLogger
	return nil
}
