// internal/utils/logging.go
package utils

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger")
	}
}

func GetLogger() *zap.Logger {
	return logger
}

func LogServiceError(service, action string, err error, fields ...zap.Field) {
	logFields := append([]zap.Field{
		zap.String("service", service),
		zap.String("act", action),
		zap.Error(err),
	}, fields...)
	logger.Error("error occurred", logFields...)
}

func LogServiceWarn(service, action, msg string, fields ...zap.Field) {
	logFields := append([]zap.Field{
		zap.String("service", service),
		zap.String("act", action),
	}, fields...)
	logger.Warn(msg, logFields...)
}

func LogServiceInfo(service, action, msg string, fields ...zap.Field) {
	logFields := append([]zap.Field{
		zap.String("service", service),
		zap.String("act", action),
	}, fields...)
	logger.Info(msg, logFields...)
}
